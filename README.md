# api-dispatch

A proof-of-concept dispatch service.

**The problem**: Given a known list of drivers and their geographic whereabouts,
and given a known location for a trip pickup, how do we select the nearest
drivers quickly and efficiently?

## Dependencies

Built on:

* [go-envconfig](https://github.com/sethvargo/go-envconfig) for env configuration
* [fx](https://github.com/uber-go/fx) for dependency injection
* [cobra](https://github.com/spf13/cobra) for CLI
* [postgres](https://www.postgresql.org/) for SQL DB
* [migrate](https://github.com/golang-migrate/migrate) for DB migrations
* [sqlboiler](https://github.com/volatiletech/sqlboiler) for schema-generated, strongly-typed ORM
* [h3](https://h3geo.org/), a hexagonal hierarchical geospatial indexing system
* [materialize](https://materialize.com/) — can we use parameterized queries? :question:

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

### Schema

The gRPC handler will use [H3](https://h3geo.org/) to calculate information
about the provided geographic coordinates. We'll persist which hexagonal cells
the driver is currently at various resolutions, as well as any neighboring
cells. (See the SQL in the [schema](./schema) directory).

#### Running migrations

To run database migrations:

```bash
# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create a migration script
migrate create -dir ./schema -ext sql init

# Run all migrations
migrate -path ./schema -database postgres://postgres:postgres@localhost:5432/dispatch\?sslmode=disable up

# Undo migrations
migrate -path ./schema -database postgres://postgres:postgres@localhost:5432/dispatch\?sslmode=disable down
```

#### Generating SQLBoiler code

We use [sqlboiler](https://github.com/volatiletech/sqlboiler) to auto-generate
a strongly-typed ORM by pointing it at our current schema.

```bash
# Install sqlboiler
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# Generate code
make gen-models
```

### Seeding driver locations

Using [`seed.json`](./seed.json):

```bash
go run cmd/dispatch/dispatch.go ingest --file seed.json
```

Or using [grpcurl](https://github.com/fullstorydev/grpcurl):

```bash
(
cat << EOF
{
  "locations": [
    {
      "driver_id": "greenpoint",
      "timestamp": "2022-05-02T03:45:11Z",
      "lat_lng": {"latitude": 40.7302797, "longitude": -73.9487438}
    }
  ]
}
EOF
) | grpcurl -plaintext -d @ localhost:8080 coop.drivers.dispatch.v1beta1.DispatchService/Ingest
```

### Connecting to postgres

```bash
psql postgres://postgres:postgres@localhost:5432/dispatch

dispatch=# select driver_id from driver_location ;
           driver_id           
-------------------------------
 greenpoint
 wburg
 GPT-Beer-Ale
 GPT-St-Vitus
 GPT-Le-Fanfare
 GPT-Lobster-Joint
 GPT-Sweetleaf-Coffee-Roasters
 GPT-Wenwen
 GPT-Esme
 GPT-Kana-Hashi
 GPT-Pelicana-Chicken
 GPT-Christinas
 GPT-Good-Room
 GPT-El-Born
 WBG-Bernies
 WBG-Llama-Inn
 WBG-Chimu-Bistro
 WBG-Birria-Landia
(18 rows)
```

### Getting the nearest drivers

#### Request

Here we're requesting a pickup at [Key Food Supermarkets](https://goo.gl/maps/xUnzhGm2h1Hpcx6q7)
(Greenpoint).

You can use the CLI:

```bash
go run cmd/dispatch/dispatch.go dispatch --latitude 40.73010864595388 --longitude -73.95094555260256
```

You can also use grpcurl:

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

```json
{
  "results":  [
    {
      "driverId":  "GPT-Christinas",
      "driverLocation":  {
        "latitude":  40.729212580192396,
        "longitude":  -73.95367193640175
      },
      "resolution":  10,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Lobster-Joint",
      "driverLocation":  {
        "latitude":  40.73570796316877,
        "longitude":  -73.95471296204035
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Pelicana-Chicken",
      "driverLocation":  {
        "latitude":  40.73145299880781,
        "longitude":  -73.95453125752603
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Wenwen",
      "driverLocation":  {
        "latitude":  40.7340725164784,
        "longitude":  -73.95478344708282
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Esme",
      "driverLocation":  {
        "latitude":  40.73341597708627,
        "longitude":  -73.95453615584985
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Sweetleaf-Coffee-Roasters",
      "driverLocation":  {
        "latitude":  40.73470993736533,
        "longitude":  -73.95447884704288
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Good-Room",
      "driverLocation":  {
        "latitude":  40.727136833286565,
        "longitude":  -73.95260914370321
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Kana-Hashi",
      "driverLocation":  {
        "latitude":  40.73267839045372,
        "longitude":  -73.95498100296005
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "GPT-El-Born",
      "driverLocation":  {
        "latitude":  40.72458479176929,
        "longitude":  -73.95095436146362
      },
      "resolution":  9,
      "kValue":  2
    },
    {
      "driverId":  "WBG-Bernies",
      "driverLocation":  {
        "latitude":  40.722135621331645,
        "longitude":  -73.95011426935746
      },
      "resolution":  8,
      "kValue":  2
    },
    {
      "driverId":  "WBG-Chimu-Bistro",
      "driverLocation":  {
        "latitude":  40.715273706507176,
        "longitude":  -73.95167992349239
      },
      "resolution":  8,
      "kValue":  2
    },
    {
      "driverId":  "WBG-Birria-Landia",
      "driverLocation":  {
        "latitude":  40.71454465722973,
        "longitude":  -73.95241070138934
      },
      "resolution":  8,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Beer-Ale",
      "driverLocation":  {
        "latitude":  40.737689977054565,
        "longitude":  -73.95528488057654
      },
      "resolution":  8,
      "kValue":  2
    },
    {
      "driverId":  "GPT-St-Vitus",
      "driverLocation":  {
        "latitude":  40.737116268235525,
        "longitude":  -73.95425922312981
      },
      "resolution":  8,
      "kValue":  2
    },
    {
      "driverId":  "GPT-Le-Fanfare",
      "driverLocation":  {
        "latitude":  40.73640659952994,
        "longitude":  -73.95616178430338
      },
      "resolution":  8,
      "kValue":  2
    },
    {
      "driverId":  "WBG-Llama-Inn",
      "driverLocation":  {
        "latitude":  40.717046310599365,
        "longitude":  -73.94969794585268
      },
      "resolution":  8,
      "kValue":  2
    }
  ]
}
```

### H3 resolutions
H3 supports [multiple resolutions](https://h3geo.org/docs/core-library/restable):

<img src="./docs/hex.png" />
<img src="./docs/hex-annotated.png" />

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

### Materialized View

[Materialize](https://materialize.com/) will power a query that retrieves the
nearest drivers to a given point.