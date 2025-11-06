-- +goose Up
-- +goose StatementBegin

CREATE TABLE "user_photos"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "image" character varying(100) NOT NULL,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_user_photos" BEFORE UPDATE ON "user_photos" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_user_photos" AS TABLE "user_photos" WITH NO DATA;
CREATE TRIGGER "soft_delete_user_photos" BEFORE DELETE ON "user_photos" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_user_photos";
DROP TABLE "user_photos";

-- +goose StatementEnd
