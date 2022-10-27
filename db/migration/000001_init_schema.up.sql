CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "event_name" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "about" text,
  "cost" decimal,
  "start_time" timestamptz,
  "end_time" timestamptz
);

CREATE TABLE "tickets" (
  "id" SERIAL PRIMARY KEY,
  "event_id" bigserial,
  "attendee" varchar,
  "payment_status" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "payment_history" (
  "id" SERIAL PRIMARY KEY,
  "ticket_id" bigserial,
  "status" varchar,
  "transaction_id" varchar,
  "transaction_type" varchar,
  "recipient" varchar,
  "parent_transaction" varchar,
  "email" varchar,
  "recieved" timestamptz
);

CREATE TABLE "speakers" (
  "id" SERIAL PRIMARY KEY,
  "event_id" bigserial,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "bio" text NOT NULL
);

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "events" ("event_name");

CREATE INDEX ON "events" ("start_time");

CREATE INDEX ON "events" ("end_time");

CREATE INDEX ON "events" ("start_time", "end_time");

CREATE INDEX ON "tickets" ("attendee");

CREATE INDEX ON "speakers" ("first_name");

CREATE INDEX ON "speakers" ("last_name");

CREATE INDEX ON "speakers" ("event_id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("attendee") REFERENCES "users" ("username");

ALTER TABLE "payment_history" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id");

ALTER TABLE "speakers" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");
