/*
Package ranking provides functions to sort a list of search results (either
trips or drivers, depending on the query).

The assumption is that when results are queried from the datastore, they will
already have a rough geospatial ordering. In our system, we're using a hexagonal
grid to provide an initial ordering. (Other systems might do something similar
by using as-the-crow-flies distance and ranking results based on concentric
rings around an origin point.)

In any case, we need to sort results even further, taking into account
additional factors, such as the duration it takes the driver to get to the trip.

We'll also want to consider additional factors:
* trip payment
* trip start time (compared to the current time)
* expected trip duration
* driver seniority
* driver eligibility

TODO thinking out loud......
These fields may or may not be stored in this service's data model, so a gRPC
call may be need to be made to enrich our models with this information at
ranking time. In fact, I would be a proponent of keeping this service's data
models simple and only based on geospatial information; if the array of ranking
variables grows increasingly complex, no schema changes will need to be made.
That said, one might distinguish between a filter-variable (such as a driver's
eligibility to claim a certain kind of trip) versus a ranking-variable (like
expected trip payment), with the former being something this service accounts
for in its schema.
*/
package ranking
