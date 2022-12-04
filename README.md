# api-dispatch

[![Lines Of Code](https://aschey.tech/tokei/github/kevinmichaelchen/api-dispatch?category=code&style=for-the-badge)](https://github.com/kevinmichaelchen/api-dispatch)

## Table of Contents

1. [Introduction](#introduction)
   1. [The problem](#the-problem)
   2. [The solution](#the-solution)
2. [Project structure](#project-structure)
3. [How does it work](#how-does-it-work)
   1. [Geospatial Indexing](#geospatial-indexing)
   2. [Ingestion](#ingestion)
   3. [Querying](#querying)
4. [Getting started](#getting-started)

## Introduction
A proof-of-concept dispatch service.

### The problem
Given a known list of drivers and their geographic whereabouts,
and given a known location for a trip pickup, how do we select the nearest
drivers? Conversely, how do we select the best trips for drivers?

### The solution
A combination of
[h3](https://h3geo.org/) (a hexagonal hierarchical geospatial indexing system)
and distance matrix APIs, such as
[Google Maps](https://developers.google.com/maps/documentation/distance-matrix/distance-matrix)
or
[Open Source Routing Machine](http://project-osrm.org/) (OSRM).

## Project structure

| Directory                                      | Description                                          |
|------------------------------------------------|------------------------------------------------------|
| [`./cmd`](./cmd)                               | CLI for making gRPC requests                         |
| [`./idl`](./idl/coop/drivers/dispatch/v1beta1) | Protobufs (Interface Definition Language)            |
| [`./internal/app`](./internal/app)             | App dependency injection / initialization            |
| [`./internal/idl`](./internal/idl)             | Auto-generated protobufs                             |
| [`./internal/models`](./internal/models)       | Auto-generated ORM / models                          |
| [`./internal/service`](./internal/service)     | Service layer / Business logic                       |
| [`./schema`](./schema)                         | SQL migration scripts                                |
| [`./pkg/grpc`](./pkg/grpc)                     | gRPC interceptors                                    |
| [`./pkg/maps`](./pkg/maps)                     | Geo-related logic (e.g., geocoding, distance matrix) |

## How does it work

### Ingestion
See [`./docs/ingestion.md`](./docs/ingestion.md).

### Geospatial Indexing

H3 is covered in more detail in [`./docs/h3.md`](./docs/h3.md).

Briefly, H3 tessellatees the world into hexagons at various resolutions.
We rank by distance using the concept of
[k-rings](https://h3geo.org/docs/api/traversal/#kring),
which are akin to [concentric circles](https://en.wikipedia.org/wiki/Concentric_objects).

### Querying

There are two sorting/ranking passes.
The first pass is for general reachability.
The second pass is to precisely rank results based on distance/duration.

The first sorts by k-ring (concentric circle). Results that are in tighter concentric circles, at finer (more zoomed in) resolutions, will rank higher.

The second pass sorts based on the duration it would take the driver to travel
to the pickup point. For this, we use a Distance Matrix request.

A third ranking phase could potentially take place downstream, where other non-geographical factors are considered,
such as a driver's rating or seniority, or a trip's expected payment, start time, expected duration, etc.

## Getting started

First, spin everything up:
```bash
# Step 1: Start containers in detached (background) mode
docker-compose up -d

# Step 2: Create the database schema
make migrate-up

# Step 3: Start app (you can optionally specify a API_KEY env var for Google Maps)
go run main.go

# Step 4: Seed trips
go run cmd/dispatch/*.go ingest trips --file seed-trips.json

# Step 5: Seed drivers
go run cmd/dispatch/*.go ingest drivers --file seed-drivers.json
```

Finally, hit the API (using [HTTPie](https://httpie.io/))
```bash
http POST \
  http://localhost:8081/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers \
    limit=2 \
    pickup_location:='{"latitude": 40.73010864595388, "longitude": -73.95094555260256}'
```

Here we're requesting a pickup at [Key Food Supermarkets](https://goo.gl/maps/xUnzhGm2h1Hpcx6q7)
(Greenpoint).

Notice how drivers in neighboring hex cells at higher (finer) resolutions appear
above those who neighbor the pickup location in lower (coarser) resolutions.

(See these results in [Google Maps](https://www.google.com/maps/dir/Key+Food+Supermarkets/Christina's/Lobster+Joint/Pelicana+Chicken/Wenwen,+Manhattan+Avenue,+Brooklyn,+NY/Esme,+Manhattan+Avenue,+Brooklyn,+NY/Sweetleaf+Coffee+Roasters/Good+Room/KanaHashi/El+Born,+Manhattan+Avenue,+Brooklyn,+NY/@40.7298698,-73.9590101,16z/data=!3m2!4b1!5s0x89c25946ba8690d1:0x75343887f28c8143!4m62!4m61!1m5!1m1!1s0x89c2594776e1a533:0x6e12c8b9202752d8!2m2!1d-73.9509796!2d40.7299092!1m5!1m1!1s0x89c25940d3ef382f:0x694b17e017a97e4d!2m2!1d-73.953922!2d40.7291077!1m5!1m1!1s0x89c2593e9472f533:0x50e900372535289c!2m2!1d-73.9552774!2d40.7354038!1m5!1m1!1s0x89c2593f58cb8ce7:0xa1a951245a59d32d!2m2!1d-73.9547015!2d40.7314082!1m5!1m1!1s0x89c2595f20f6ca99:0xee8ef144f9bfebaf!2m2!1d-73.9550488!2d40.7339961!1m5!1m1!1s0x89c2593f3d6abf81:0x8d77248129051b4c!2m2!1d-73.9549203!2d40.7332267!1m5!1m1!1s0x89c2593ee8a352d7:0x5bc1971dd74cfcd7!2m2!1d-73.9553735!2d40.7345144!1m5!1m1!1s0x89c25946b09ba761:0x90af8ca50b67a075!2m2!1d-73.9529121!2d40.7269376!1m5!1m1!1s0x89c2593f3f8c4f21:0x543eba709caaa83!2m2!1d-73.9548886!2d40.7326259!1m5!1m1!1s0x89c259444f48269b:0x668274ceb7e6b645!2m2!1d-73.9512545!2d40.7243531!3e0))

```json
{
   "results": [
      {
         "driver": {
            "driverId": "GPT-Christinas",
            "mostRecentHeartbeat": "0001-01-01T00:00:00Z",
            "currentLocation": {
               "latitude": 40.729116923462385,
               "longitude": -73.95392251222499
            }
         },
         "distanceMeters": 471,
         "duration": "227s",
         "location": {
            "latitude": 40.729116923462385,
            "longitude": -73.95392251222499
         },
         "address": "853 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 10,
         "kValue": 2
      },
      {
         "driver": {
            "driverId": "GPT-Pelicana-Chicken",
            "mostRecentHeartbeat": "0001-01-01T00:00:00Z",
            "currentLocation": {
               "latitude": 40.731416756150395,
               "longitude": -73.95470131162523
            }
         },
         "distanceMeters": 706,
         "duration": "292s",
         "location": {
            "latitude": 40.731416756150395,
            "longitude": -73.95470131162523
         },
         "address": "941 Manhattan Ave, Brooklyn, NY 11222, USA",
         "resolution": 9,
         "kValue": 2
      }
   ],
   "pickupAddress": "216 Greenpoint Ave, Brooklyn, NY 11222, USA"
}
```
