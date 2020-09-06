CREATE TABLE "todo" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "completed" boolean NOT NULL DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "todo" ("id");