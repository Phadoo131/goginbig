CREATE TABLE "instore" (
  "book" varchar PRIMARY KEY,
  "owner" varchar NOT NULL,
  "bookcount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "book" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "returnandborrow" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "book" varchar NOT NULL,
  "bookcount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "returnandborrow" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "returnandborrow" ADD FOREIGN KEY ("book") REFERENCES "instore" ("book");

ALTER TABLE "returnandborrow" ADD FOREIGN KEY ("bookcount") REFERENCES "instore" ("bookcount");

ALTER TABLE "accounts" ADD FOREIGN KEY ("book") REFERENCES "instore" ("book");

