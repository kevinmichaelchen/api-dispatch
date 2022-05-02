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
  "results": [
    {
      "driverId": "GPT-Kana-Hashi",
      "resolution": 9
    },
    {
      "driverId": "GPT-Good-Room",
      "resolution": 9
    },
    {
      "driverId": "GPT-Christinas",
      "resolution": 9
    },
    {
      "driverId": "GPT-Pelicana-Chicken",
      "resolution": 9
    },
    {
      "driverId": "GPT-St-Vitus",
      "resolution": 8
    },
    {
      "driverId": "GPT-Lobster-Joint",
      "resolution": 8
    },
    {
      "driverId": "GPT-Le-Fanfare",
      "resolution": 8
    },
    {
      "driverId": "WBG-Chimu-Bistro",
      "resolution": 8
    },
    {
      "driverId": "GPT-Sweetleaf-Coffee-Roasters",
      "resolution": 8
    },
    {
      "driverId": "WBG-Bernies",
      "resolution": 8
    },
    {
      "driverId": "WBG-Llama-Inn",
      "resolution": 8
    },
    {
      "driverId": "GPT-Beer-Ale",
      "resolution": 8
    },
    {
      "driverId": "GPT-Wenwen",
      "resolution": 8
    },
    {
      "driverId": "GPT-Esme",
      "resolution": 8
    },
    {
      "driverId": "GPT-El-Born",
      "resolution": 8
    },
    {
      "driverId": "WBG-Birria-Landia",
      "resolution": 7
    }
  ]
}
```

### Materialized View
[Materialize](https://materialize.com/) will power a query that retrieves the 
nearest drivers to a given point.