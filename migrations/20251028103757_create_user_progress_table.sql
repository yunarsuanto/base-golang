-- +goose Up
-- +goose StatementBegin

CREATE TABLE "user_progress"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "lesson_id" uuid NULL REFERENCES "lessons" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "score" float NOT NULL DEFAULT 0,
  "starting_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "completed_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "attempt_count" integer NOT NULL DEFAULT 0,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_user_progress" BEFORE UPDATE ON "user_progress" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_user_progress" AS TABLE "user_progress" WITH NO DATA;
CREATE TRIGGER "soft_delete_user_progress" BEFORE DELETE ON "user_progress" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_user_progress";
DROP TABLE "user_progress";

-- +goose StatementEnd
