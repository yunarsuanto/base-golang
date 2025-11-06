-- +goose Up
-- +goose StatementBegin

CREATE TABLE "users"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" character varying(100) NOT NULL UNIQUE,
  "password" character varying NOT NULL,
  "is_active" boolean DEFAULT false,
  "provider_id" TEXT UNIQUE,
  "provider" VARCHAR(20) DEFAULT 'local',
  "token_verification" VARCHAR(255),
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_users" BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_users" AS TABLE "users" WITH NO DATA;
CREATE TRIGGER "soft_delete_users" BEFORE DELETE ON "users" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_users";
DROP TABLE "users";

-- +goose StatementEnd
