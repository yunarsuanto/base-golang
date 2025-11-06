-- +goose Up
-- +goose StatementBegin

CREATE TABLE "lesson_items"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "lesson_id" uuid NULL REFERENCES "lessons" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "content" TEXT NOT NULL DEFAULT '',
  "order" smallint NOT NULL DEFAULT 0,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_lesson_items" BEFORE UPDATE ON "lesson_items" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lesson_items" AS TABLE "lesson_items" WITH NO DATA;
CREATE TRIGGER "soft_delete_lesson_items" BEFORE DELETE ON "lesson_items" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_lesson_items";
DROP TABLE "lesson_items";

-- +goose StatementEnd
