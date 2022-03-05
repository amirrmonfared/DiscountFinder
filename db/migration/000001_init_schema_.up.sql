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

CREATE TABLE "on_sale" (
  "id" bigserial NOT NULL,
  "link" varchar PRIMARY KEY NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "first" ("link");

CREATE INDEX ON "second" ("link");

CREATE INDEX ON "on_sale" ("link");
