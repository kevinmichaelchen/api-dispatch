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
compacted into Postgres using a RedisGears Python function.