CREATE TABLE "public"."trip"
(
    "id"               text        NOT NULL,
    "created_at"       timestamptz NOT NULL DEFAULT now(),
    "scheduled_for"    timestamptz NOT NULL,
    "latitude"         float       NOT NULL,
    "longitude"        float       NOT NULL,
    "expected_pay"     float       NOT NULL,
    "r7_cell"          text,
    "r8_cell"          text,
    "r9_cell"          text,
    "r10_cell"         text,
    "r7_k1_neighbors"  text[],
    "r8_k1_neighbors"  text[],
    "r8_k2_neighbors"  text[],
    "r9_k1_neighbors"  text[],
    "r9_k2_neighbors"  text[],
    "r10_k1_neighbors" text[],
    "r10_k2_neighbors" text[],
    PRIMARY KEY ("id")
);

CREATE INDEX idx_trip_r7_k1_neighbors
    ON public.trip USING btree (r7_k1_neighbors);

CREATE INDEX idx_trip_r8_k1_neighbors
    ON public.trip USING btree (r8_k1_neighbors);
CREATE INDEX idx_trip_r8_k2_neighbors
    ON public.trip USING btree (r8_k2_neighbors);

CREATE INDEX idx_trip_r9_k1_neighbors
    ON public.trip USING btree (r9_k1_neighbors);
CREATE INDEX idx_trip_r9_k2_neighbors
    ON public.trip USING btree (r9_k2_neighbors);

CREATE INDEX idx_trip_r10_k1_neighbors
    ON public.trip USING btree (r10_k1_neighbors);
CREATE INDEX idx_trip_r10_k2_neighbors
    ON public.trip USING btree (r10_k2_neighbors);