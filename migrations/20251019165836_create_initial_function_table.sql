-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION trigger_set_timestamp() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW(); RETURN NEW; END; $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_soft_delete()
RETURNS TRIGGER AS $$
DECLARE 
  deleted_table_name text := CONCAT('deleted_', TG_TABLE_NAME::regclass::text);
  queryInsert text;
BEGIN
  queryInsert := format('INSERT INTO %I SELECT * FROM %I WHERE id = %L', lower(deleted_table_name), TG_TABLE_NAME::regclass::text, OLD.id);
  EXECUTE queryInsert;
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE "general_settings" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "key" character varying NOT NULL UNIQUE,
  "value" text NOT NULL DEFAULT '',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL
);
CREATE TRIGGER "updated_at_general_settings" BEFORE UPDATE ON "general_settings" FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE RULE "protect_default_general_settings_delete" AS
ON DELETE TO "general_settings"
DO INSTEAD NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "general_settings";

DROP FUNCTION IF EXISTS trigger_soft_delete();
DROP FUNCTION IF EXISTS trigger_set_timestamp();
DROP EXTENSION IF EXISTS "uuid-ossp";

-- +goose StatementEnd
