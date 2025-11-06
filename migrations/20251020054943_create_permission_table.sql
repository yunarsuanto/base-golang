-- +goose Up
-- +goose StatementBegin

CREATE TABLE "permissions"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL UNIQUE,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_permissions" BEFORE UPDATE ON "permissions" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_permissions" AS TABLE "permissions" WITH NO DATA;
CREATE TRIGGER "soft_delete_permissions" BEFORE DELETE ON "permissions" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_permissions";
DROP TABLE "permissions";

-- +goose StatementEnd
