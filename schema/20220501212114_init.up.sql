CREATE TABLE "public"."driver_location"
(
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
    "r9_k1_neighbors"  text[],
    "r10_k1_neighbors" text[],
    PRIMARY KEY ("created_at", "driver_id")
);