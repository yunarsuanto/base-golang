-- +goose Up
-- +goose StatementBegin

CREATE TABLE "user_role"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "role_id" uuid NULL REFERENCES "roles" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "is_active" boolean DEFAULT false,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE ("user_id", "role_id")
);
CREATE TRIGGER "updated_at_user_role" BEFORE UPDATE ON "user_role" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_user_role" AS TABLE "user_role" WITH NO DATA;
CREATE TRIGGER "soft_delete_user_role" BEFORE DELETE ON "user_role" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_user_role";
DROP TABLE "user_role";

-- +goose StatementEnd
