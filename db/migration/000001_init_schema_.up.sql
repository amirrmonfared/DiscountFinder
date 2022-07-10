CREATE TABLE "products" (
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
  "previous_price" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "products" ("link");

CREATE INDEX ON "on_sale" ("link");
