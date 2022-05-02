# api-dispatch
A proof-of-concept dispatch service.

**The problem**: Given a known list of drivers and their geographic whereabouts,
and given a known location for a trip pickup, how do we select the nearest 
drivers quickly and efficiently?

## Dependencies
Built on:
* [go-envconfig](https://github.com/sethvargo/go-envconfig) for env configuration
* [fx](https://github.com/uber-go/fx) for dependency injection
* [postgres](https://www.postgresql.org/) for SQL DB
* [migrate](https://github.com/golang-migrate/migrate) for DB migrations
* [sqlboiler](https://github.com/volatiletech/sqlboiler) for schema-generated, strongly-typed ORM
* [materialize](https://materialize.com/) for fast views on "who's nearby" query
* [h3](https://h3geo.org/), a hexagonal hierarchical geospatial indexing system

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
Using [grpcurl](https://github.com/fullstorydev/grpcurl):
```bash
(
cat << EOF
{
  "locations": [
    {
      "driver_id": "greenpoint",
      "timestamp": "2022-05-02T03:45:11Z",
      "lat_lng": {"latitude": 40.7302797, "longitude": -73.9487438}
    },
    {
      "driver_id": "wburg",
      "timestamp": "2022-05-02T03:45:11Z",
      "lat_lng": {"latitude": 40.7082168, "longitude": -73.95753}
    }
  ]
}
EOF
) | grpcurl -plaintext -d @ localhost:8080 coop.drivers.dispatch.v1beta1.DispatchService/Ingest
```

### Getting the nearest drivers
We provide a [gRPC endpoint](idl/coop/drivers/dispatch/v1beta1/api.proto) that 
allows clients to get the nearest drivers to a specific point:
```protobuf
service DispatchService {
  // Returns driver IDs of nearest drivers
  rpc Dispatch(DispatchRequest) returns (DispatchResponse) {}
}

message DispatchRequest {
  LatLng point = 1;
}

message DispatchResponse {
  repeated string driver_ids = 1;
}
```

#### Request
Here we're requesting a pickup at `(40.7110694,-73.9514453)`, an Argentinian 
restaurant in Williamsburg.

You can use the CLI:
```bash
go run cmd/dispatch/dispatch.go dispatch --latitude 40.7110694 --longitude -73.9514453
```

You can also use grpcurl:
```bash
(
cat << EOF
{
  "location": {
    "latitude": 40.7110694,
    "longitude": -73.9514453
  }
}
EOF
) | grpcurl -plaintext -d @ localhost:8080 coop.drivers.dispatch.v1beta1.DispatchService/Dispatch
```

#### Response
Because the pickup is in Williamsburg, a driver in that neighborhood appears
ranked above a driver in Greenpoint (north of Williamsburg).
```json
{
  "results": [
    {
      "driverId": "wburg",
      "resolution": 8
    },
    {
      "driverId": "greenpoint",
      "resolution": 7
    }
  ]
}
```

### Materialized View
[Materialize](https://materialize.com/) will power a query that retrieves the 
nearest drivers to a given point.