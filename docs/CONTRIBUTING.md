## Getting started
```bash
docker-compose up
API_KEY=GOOGLE_MAPS_KEY go run main.go
```

## Lines of code
```bash
tokei -e models -e idl
```

## Dependencies
We use the following Go dependencies:

* [go-envconfig](https://github.com/sethvargo/go-envconfig) for env configuration
* [fx](https://github.com/uber-go/fx) for dependency injection
* [cobra](https://github.com/spf13/cobra) for CLI
* [postgres](https://www.postgresql.org/) for SQL DB
* [migrate](https://github.com/golang-migrate/migrate) for DB migrations
* [sqlboiler](https://github.com/volatiletech/sqlboiler) for schema-generated, strongly-typed ORM
* [xid](https://github.com/rs/xid) for random ID generation
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) for validation

## Protobufs
To regenerate protobufs, run 
```bash
make gen-proto
```

## Schema

The gRPC handler will use [H3](https://h3geo.org/) to calculate information
about the provided geographic coordinates. We'll persist which hexagonal cells
the driver is currently at various resolutions, as well as any neighboring
cells. (See the SQL in the [schema](../schema) directory).

### Running migrations

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

### Generating SQLBoiler code

We use [sqlboiler](https://github.com/volatiletech/sqlboiler) to auto-generate
a strongly-typed ORM by pointing it at our current schema.

```bash
# Install sqlboiler
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# Generate code
make gen-models
```

## Seeding driver locations

```bash
go run cmd/dispatch/*.go ingest trips --file seed-trips.json
go run cmd/dispatch/*.go ingest drivers --file seed-drivers.json
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

## Connecting to postgres

```bash
psql postgres://postgres:postgres@localhost:5432/dispatch

dispatch=# select driver_id from driver_location ;
           driver_id           
-------------------------------
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
(16 rows)
```