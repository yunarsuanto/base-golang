-- +goose Up
-- +goose StatementBegin

CREATE TABLE "user_profiles"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "national_id" character varying(30) NOT NULL UNIQUE,
  "fullname" character varying(100) NOT NULL DEFAULT '',
  "email" character varying(100) NOT NULL DEFAULT '',
  "phone" character varying(30) NOT NULL DEFAULT '',
  "address" TEXT NOT NULL DEFAULT '',
  "postal_code" character varying(15) NOT NULL DEFAULT '',
  "age" smallint NOT NULL DEFAULT 0,
  "latitude" DOUBLE PRECISION NOT NULL DEFAULT 0,
  "longitude" DOUBLE PRECISION NOT NULL DEFAULT 0,
  "profile_image" character varying(100) NOT NULL DEFAULT '',
  "national_id_image" character varying(100) NOT NULL DEFAULT '',
  "guardian_name" character varying(100) NOT NULL DEFAULT '',
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_user_profiles" BEFORE UPDATE ON "user_profiles" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_user_profiles" AS TABLE "user_profiles" WITH NO DATA;
CREATE TRIGGER "soft_delete_user_profiles" BEFORE DELETE ON "user_profiles" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_user_profiles";
DROP TABLE "user_profiles";

-- +goose StatementEnd
