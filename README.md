# api-dispatch

## Table of Contents

1. [Introduction](#introduction)
   1. [The problem](#the-problem)
   1. [The solution](#the-solution)
   1. [Not considered yet](#not-considered-yet)
1. [Project structure](#project-structure)
1. [How does it work](#how-does-it-work)
   1. [Ingestion](#ingestion)
      1. [Geospatial Indexing](#geospatial-indexing)
   1. [Getting the nearest trips or drivers](#getting-the-nearest-trips-or-drivers)
      1. [Request](#request)
      2. [Response](#response)

## Introduction
A proof-of-concept dispatch service.

### The problem
Given a known list of drivers and their geographic whereabouts,
and given a known location for a trip pickup, how do we select the nearest
drivers? Conversely, how do we select the best trips for drivers?

### The solution
Combine [Google Maps](https://developers.google.com/maps/documentation/distance-matrix/distance-matrix)
and [h3](https://h3geo.org/) (a hexagonal hierarchical geospatial indexing system).

### Not considered yet
1. Driver eligibility for trip

## Project structure

| Directory                                    | Description                               |
|----------------------------------------------|-------------------------------------------|
| [`./cmd`](./cmd)                             | CLI for making gRPC requests              |
| [`./idl`](./idl)                             | Protobufs (Interface Definition Language) |
| [`./internal/app`](./internal/app)           | App dependency injection / initialization |
| [`./internal/distance`](./internal/distance) | Google Maps Distance Matrix logic         |
| [`./internal/idl`](./internal/idl)           | Auto-generated protobufs                  |
| [`./internal/models`](./internal/models)     | Auto-generated ORM / models               |
| [`./internal/service`](./internal/service)   | Service layer / Business logic            |
| [`./schema`](./schema)                       | SQL migration scripts                     |

## How does it work

### Ingestion

We expose two [gRPC endpoints](idl/coop/drivers/dispatch/v1beta1/api.proto) to
ingest:
1. batches of driver location pings
2. batches of trips

A driver location ping takes the form:
```
(time, driver_id, lat, lng)
```

A trip takes the form:
```
(id, scheduled_for, expected_pay, lat, lng)
```

**Note about scalability**: Postgres might not be the right tool for this kind
of time-series data, so maybe we can assume that somewhere downstream these
locations are persisted to something like
[RedisTimeSeries](https://redis.io/docs/stack/timeseries/) and periodically
compacted into Postgres.

#### Geospatial Indexing
See [`./docs/h3.md`](./docs/h3.md).

### Getting the nearest trips or drivers

Resolution and k-ring value are used for the first round of filtering. If 
someone isn't in the 1-ring at resolution 7 (pretty zoomed out), they won't 
appear in search results. After that, we sort by Google Maps-estimated duration
(time to get to pickup location).

#### Request

Here we're requesting a pickup at [Key Food Supermarkets](https://goo.gl/maps/xUnzhGm2h1Hpcx6q7)
(Greenpoint).

You can use the CLI:

```bash
# Get nearest drivers to a specified pickup (passenger) location
go run cmd/dispatch/*.go nearest drivers --latitude 40.73010864595388 --longitude -73.95094555260256

# Get nearest trips to a specified driver location
go run cmd/dispatch/*.go nearest trips --latitude 40.73010864595388 --longitude -73.95094555260256
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
         "distanceMeters": 471,
         "duration": "227s",
         "driverLocation": {
            "latitude": 40.729116923462385,
            "longitude": -73.95392251222499
         },
         "address": "853 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 10,
         "kValue": 2
      },
      {
         "driverId": "GPT-Pelicana-Chicken",
         "distanceMeters": 706,
         "duration": "285s",
         "driverLocation": {
            "latitude": 40.731416756150395,
            "longitude": -73.95470131162523
         },
         "address": "941 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-El-Born",
         "distanceMeters": 806,
         "duration": "301s",
         "driverLocation": {
            "latitude": 40.72436580353396,
            "longitude": -73.95124766347774
         },
         "address": "651 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-Good-Room",
         "distanceMeters": 768,
         "duration": "317s",
         "driverLocation": {
            "latitude": 40.726944958544514,
            "longitude": -73.95291323476157
         },
         "address": "98 Meserole Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-Kana-Hashi",
         "distanceMeters": 843,
         "duration": "320s",
         "driverLocation": {
            "latitude": 40.732637499546215,
            "longitude": -73.95488693544799
         },
         "address": "981 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-Esme",
         "distanceMeters": 911,
         "duration": "337s",
         "driverLocation": {
            "latitude": 40.733235412885314,
            "longitude": -73.95491763917049
         },
         "address": "999 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-Wenwen",
         "distanceMeters": 999,
         "duration": "340s",
         "driverLocation": {
            "latitude": 40.7340094734467,
            "longitude": -73.95504667163209
         },
         "address": "1029 Manhattan Ave # 1, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-Lobster-Joint",
         "distanceMeters": 1132,
         "duration": "342s",
         "driverLocation": {
            "latitude": 40.73541316092886,
            "longitude": -73.95527626749518
         },
         "address": "1073 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "GPT-Le-Fanfare",
         "distanceMeters": 1223,
         "duration": "368s",
         "driverLocation": {
            "latitude": 40.73622468817931,
            "longitude": -73.95551737528102
         },
         "address": "1103 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      },
      {
         "driverId": "WBG-Bernies",
         "distanceMeters": 1193,
         "duration": "374s",
         "driverLocation": {
            "latitude": 40.721899506693795,
            "longitude": -73.95055397598824
         },
         "address": "332 Driggs Ave, Brooklyn, NY 11222, USA",
         "resolution": 8,
         "kValue": 2
      }
   ],
   "pickupAddress": "216 Greenpoint Ave, Brooklyn, NY 11222, USA"
}
```
