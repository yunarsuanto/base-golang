-- +goose Up
-- +goose StatementBegin

CREATE TABLE "users"
(
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL,
  "username" character varying(100) NOT NULL UNIQUE,
  "password" character varying NOT NULL,
  "created_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "updated_by" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_users" BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "deleted_users" AS TABLE "users" WITH NO DATA;
CREATE TRIGGER "soft_delete_users" BEFORE DELETE ON "users" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();
------------------------------------------
INSERT INTO "users" ("name", "username", "password") VALUES
('super_admin',	'superadmin',	'$2a$10$7dBcJCNxbl12LXBhiBtJsO/cgQ6IEn.qdej8kJneSoqggEbE/YKNK');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_users";
DROP TABLE "users";

-- +goose StatementEnd
