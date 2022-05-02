# api-dispatch

A proof-of-concept dispatch service.

**The problem**: Given a known list of drivers and their geographic whereabouts,
and given a known location for a trip pickup, how do we select the nearest
drivers quickly and efficiently? Conversely, how do we select the best trips for
drivers?

**The solution**: Combine [Google Maps](https://developers.google.com/maps/documentation/distance-matrix/distance-matrix)
and [h3](https://h3geo.org/) (a hexagonal hierarchical geospatial indexing system).

**Not considered yet**:
1. Driver eligibility for trip
2. Trip payment
3. Time until trip start time

## Project structure

| Directory                                  | Description                               |
|--------------------------------------------|-------------------------------------------|
| [`./cmd`](./cmd)                           | CLI for making gRPC requests              |
| [`./idl`](./idl)                           | Protobufs (Interface Definition Language) |
| [`./internal/app`](./internal/app)         | App dependency injection / initialization |
| [`./internal/idl`](./internal/idl)         | Auto-generated protobufs                  |
| [`./internal/models`](./internal/models)   | Auto-generated ORM / models               |
| [`./internal/service`](./internal/service) | Service layer / Business logic            |
| [`./schema`](./schema)                     | SQL migration scripts                     |

## How does it work?

### Location Ingestion

We expose a [gRPC endpoint](idl/coop/drivers/dispatch/v1beta1/api.proto) to
ingest location pings in batch. Pings take the form of:

```
(time, driver_id, lat, lng)
```

**Note about scalability**: Postgres might not be the right tool for this kind
of time-series data, so maybe we can assume that somewhere downstream these
locations are persisted to something like
[RedisTimeSeries](https://redis.io/docs/stack/timeseries/) and periodically
compacted into Postgres.

### H3 Geospatial Indexing
As location pings are ingested, we use the H3 library to figure out which hex 
cells the driver is currently in at various resolutions, as well as k-rings
which are essentially sets of *1*st-degree, *2*nd-degree, or _k_-degree 
neighbors.

#### H3 resolutions
H3 supports [multiple resolutions](https://h3geo.org/docs/core-library/restable):

<div style="display: flex; justify-content: space-between;">
<img src="./docs/hex.png" style="margin-right:15px;" />
<img src="./docs/hex-annotated.png" />
  </div>

Each finer-resolution cell is 7 times smaller than its coarser parent.

Brooklyn is 250 km<sup>2</sup> (one cell at Resolution 5)...

Williamsburg is 5 km<sup>2</sup> (one cell at Resolution 7)...

| Resolution | Avg Hex Area               | Avg Hex Edge Length (km) | Number of unique indexes |
|------------|----------------------------|--------------------------|--------------------------|
| 5          | 252.9 km<sup>2</sup>       | 8.5 km                   | 2,016,842                |
| 6          | 36.13 km<sup>2</sup>       | 3.2 km                   | 14,117,882               |
| 7          | 5.16 km<sup>2</sup>        | 1.2 km                   | 98,825,162               |
| 8          | 737327.6 m<sup>2</sup>     | 461 m                    | 691,776,122              |
| 9          | 105332.5 m<sup>2</sup>     | 174 m                    | 4,842,432,842            |
| 10         | 15047.5 m<sup>2</sup>      | 65 m                     | 33,897,029,882           |
| 11         | 2149.6 m<sup>2</sup>       | 24 m                     | 237,279,209,162          |

### Getting the nearest drivers

#### Request

Here we're requesting a pickup at [Key Food Supermarkets](https://goo.gl/maps/xUnzhGm2h1Hpcx6q7)
(Greenpoint).

You can use the CLI:

```bash
go run cmd/dispatch/dispatch.go dispatch --latitude 40.73010864595388 --longitude -73.95094555260256
```

Or [grpcurl](https://github.com/fullstorydev/grpcurl):
```bash
(
cat << EOF
{
  "location": {
    "latitude": 40.73010864595388,
    "longitude": -73.95094555260256
  }
}
EOF
) | grpcurl -plaintext -d @ localhost:8080 coop.drivers.dispatch.v1beta1.DispatchService/Dispatch
```

#### Response

Because the pickup is in [Key Food Supermarkets](https://goo.gl/maps/xUnzhGm2h1Hpcx6q7)
(Greenpoint), drivers in that neighborhood appear above others.

Notice how drivers in neighboring hex cells at higher (finer) resolutions appear
above those who neighbor the pickup location in lower (coarser) resolutions.

(See these results in [Google Maps](https://www.google.com/maps/dir/Key+Food+Supermarkets/Christina's/Lobster+Joint/Pelicana+Chicken/Wenwen,+Manhattan+Avenue,+Brooklyn,+NY/Esme,+Manhattan+Avenue,+Brooklyn,+NY/Sweetleaf+Coffee+Roasters/Good+Room/KanaHashi/El+Born,+Manhattan+Avenue,+Brooklyn,+NY/@40.7298698,-73.9590101,16z/data=!3m2!4b1!5s0x89c25946ba8690d1:0x75343887f28c8143!4m62!4m61!1m5!1m1!1s0x89c2594776e1a533:0x6e12c8b9202752d8!2m2!1d-73.9509796!2d40.7299092!1m5!1m1!1s0x89c25940d3ef382f:0x694b17e017a97e4d!2m2!1d-73.953922!2d40.7291077!1m5!1m1!1s0x89c2593e9472f533:0x50e900372535289c!2m2!1d-73.9552774!2d40.7354038!1m5!1m1!1s0x89c2593f58cb8ce7:0xa1a951245a59d32d!2m2!1d-73.9547015!2d40.7314082!1m5!1m1!1s0x89c2595f20f6ca99:0xee8ef144f9bfebaf!2m2!1d-73.9550488!2d40.7339961!1m5!1m1!1s0x89c2593f3d6abf81:0x8d77248129051b4c!2m2!1d-73.9549203!2d40.7332267!1m5!1m1!1s0x89c2593ee8a352d7:0x5bc1971dd74cfcd7!2m2!1d-73.9553735!2d40.7345144!1m5!1m1!1s0x89c25946b09ba761:0x90af8ca50b67a075!2m2!1d-73.9529121!2d40.7269376!1m5!1m1!1s0x89c2593f3f8c4f21:0x543eba709caaa83!2m2!1d-73.9548886!2d40.7326259!1m5!1m1!1s0x89c259444f48269b:0x668274ceb7e6b645!2m2!1d-73.9512545!2d40.7243531!3e0))

```json
{
  "results": [
    {
      "driverId": "GPT-Christinas",
      "distanceMeters": 484,
      "duration": "230s",
      "driverLocation": {
        "latitude": 40.729212580192396,
        "longitude": -73.95367193640175
      },
      "resolution": 10,
      "kValue": 2
    },
    {
      "driverId": "GPT-Wenwen",
      "distanceMeters": 1004,
      "duration": "341s",
      "driverLocation": {
        "latitude": 40.7340725164784,
        "longitude": -73.95478344708282
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-Esme",
      "distanceMeters": 923,
      "duration": "339s",
      "driverLocation": {
        "latitude": 40.73341597708627,
        "longitude": -73.95453615584985
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-El-Born",
      "distanceMeters": 790,
      "duration": "297s",
      "driverLocation": {
        "latitude": 40.72458479176929,
        "longitude": -73.95095436146362
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-Good-Room",
      "distanceMeters": 481,
      "duration": "231s",
      "driverLocation": {
        "latitude": 40.727136833286565,
        "longitude": -73.95260914370321
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-Kana-Hashi",
      "distanceMeters": 831,
      "duration": "317s",
      "driverLocation": {
        "latitude": 40.73267839045372,
        "longitude": -73.95498100296005
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-Lobster-Joint",
      "distanceMeters": 1154,
      "duration": "347s",
      "driverLocation": {
        "latitude": 40.73570796316877,
        "longitude": -73.95471296204035
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-Sweetleaf-Coffee-Roasters",
      "distanceMeters": 1229,
      "duration": "370s",
      "driverLocation": {
        "latitude": 40.73470993736533,
        "longitude": -73.95447884704288
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "GPT-Pelicana-Chicken",
      "distanceMeters": 708,
      "duration": "286s",
      "driverLocation": {
        "latitude": 40.73145299880781,
        "longitude": -73.95453125752603
      },
      "resolution": 9,
      "kValue": 2
    },
    {
      "driverId": "WBG-Llama-Inn",
      "distanceMeters": 1919,
      "duration": "463s",
      "driverLocation": {
        "latitude": 40.717046310599365,
        "longitude": -73.94969794585268
      },
      "resolution": 8,
      "kValue": 2
    }
  ]
}
```
