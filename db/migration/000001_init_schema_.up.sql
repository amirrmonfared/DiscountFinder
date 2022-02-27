CREATE TABLE "first" (
  "id" bigserial NOT NULL,
  "link" varchar PRIMARY KEY NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "second" (
  "id" bigserial NOT NULL,
  "link" varchar PRIMARY KEY NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "onSale" (
  "id" bigserial NOT NULL,
  "link" varchar PRIMARY KEY NOT NULL,
  "price" bigint NOT NULL,
  "salePer" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "first" ("link");

CREATE INDEX ON "second" ("link");

CREATE INDEX ON "onSale" ("link");
