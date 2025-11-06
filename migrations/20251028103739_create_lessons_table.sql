-- +goose Up
-- +goose StatementBegin


CREATE TABLE "category_lessons"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" character varying(100) NOT NULL UNIQUE,
  "description" TEXT,
  "category_lesson_id" uuid NULL REFERENCES "category_lessons" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "media" character varying(100) NOT NULL,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_category_lessons" BEFORE UPDATE ON "category_lessons" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_category_lessons" AS TABLE "category_lessons" WITH NO DATA;
CREATE TRIGGER "soft_delete_category_lessons" BEFORE DELETE ON "category_lessons" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

---------------------------------------------

CREATE TABLE "lessons"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" character varying(100) NOT NULL,
  "description" TEXT,
  "category_lesson_id" uuid NULL REFERENCES "category_lessons" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "media" character varying(100) NOT NULL,
  "level" smallint NOT NULL DEFAULT 1,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL,
  UNIQUE(title, level)
);
CREATE TRIGGER "updated_at_lessons" BEFORE UPDATE ON "lessons" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_lessons" AS TABLE "lessons" WITH NO DATA;
CREATE TRIGGER "soft_delete_lessons" BEFORE DELETE ON "lessons" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_lessons";
DROP TABLE "lessons";

DROP TABLE "deleted_category_lessons";
DROP TABLE "category_lessons";

-- +goose StatementEnd
