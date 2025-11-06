-- +goose Up
-- +goose StatementBegin

CREATE TABLE "role_permission"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "role_id" uuid NULL REFERENCES "roles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "permission_id" uuid NULL REFERENCES "permissions" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE ("role_id", "permission_id")
);
CREATE TRIGGER "updated_at_role_permission" BEFORE UPDATE ON "role_permission" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_role_permission" AS TABLE "role_permission" WITH NO DATA;
CREATE TRIGGER "soft_delete_role_permission" BEFORE DELETE ON "role_permission" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_role_permission";
DROP TABLE "role_permission";

-- +goose StatementEnd
