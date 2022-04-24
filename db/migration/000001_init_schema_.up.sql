CREATE TABLE "first" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "brand" varchar NOT NULL,
  "link" varchar UNIQUE NOT NULL,
  "price" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "on_sale" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "brand" varchar NOT NULL,
  "link" varchar UNIQUE NOT NULL,
  "price" varchar NOT NULL,
  "prv_price" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "first" ("link");

CREATE INDEX ON "on_sale" ("link");
