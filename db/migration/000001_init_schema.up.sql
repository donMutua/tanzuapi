CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial,
  "event_name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "about" text
);

CREATE TABLE "tickets" (
  "id" SERIAL PRIMARY KEY,
  "event_id" bigserial,
  "price" bigint,
  "start_date" date,
  "end_date" date,
  "mode" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "speakers" (
  "id" SERIAL PRIMARY KEY,
  "event_id" bigserial,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "bio" text NOT NULL
);

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "events" ("event_name");

CREATE INDEX ON "tickets" ("start_date");

CREATE INDEX ON "tickets" ("end_date");

CREATE INDEX ON "tickets" ("start_date", "end_date");

CREATE INDEX ON "speakers" ("first_name");

CREATE INDEX ON "speakers" ("last_name");

CREATE INDEX ON "speakers" ("event_id");

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "speakers" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");
