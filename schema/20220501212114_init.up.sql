CREATE TABLE "public"."driver_location"
(
    "id"               text        NOT NULL,
    "created_at"       timestamptz NOT NULL DEFAULT now(),
    "driver_id"        text        NOT NULL,
    "latitude"         float       NOT NULL,
    "longitude"        float       NOT NULL,
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

CREATE INDEX idx_driver_location_r7_k1_neighbors
    ON public.driver_location USING btree (r7_k1_neighbors);

CREATE INDEX idx_driver_location_r8_k1_neighbors
    ON public.driver_location USING btree (r8_k1_neighbors);
CREATE INDEX idx_driver_location_r8_k2_neighbors
    ON public.driver_location USING btree (r8_k2_neighbors);

CREATE INDEX idx_driver_location_r9_k1_neighbors
    ON public.driver_location USING btree (r9_k1_neighbors);
CREATE INDEX idx_driver_location_r9_k2_neighbors
    ON public.driver_location USING btree (r9_k2_neighbors);

CREATE INDEX idx_driver_location_r10_k1_neighbors
    ON public.driver_location USING btree (r10_k1_neighbors);
CREATE INDEX idx_driver_location_r10_k2_neighbors
    ON public.driver_location USING btree (r10_k2_neighbors);