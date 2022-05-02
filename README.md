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

### Materialized View
[Materialize](https://materialize.com/) will power a query that retrieves the 
nearest drivers to a given point.