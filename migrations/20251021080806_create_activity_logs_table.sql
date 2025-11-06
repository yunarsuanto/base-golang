-- +goose Up
-- +goose StatementBegin

CREATE TABLE "activity_logs"
(
    "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" uuid NULL REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    "host" character varying NOT NULL,
    "path" character varying NOT NULL,
    "body" TEXT NULL,
    "status_code" SMALLINT NOT NULL,
    "error_message" character varying NOT NULL,
    "ip_address" character varying(20) NOT NULL,
    "user_agent" character varying NOT NULL,
    "memory_usage" NUMERIC(15,12) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "deleted_activity_logs" AS TABLE "activity_logs" WITH NO DATA;
CREATE TRIGGER "soft_delete_activity_logs" BEFORE DELETE ON "activity_logs" FOR EACH ROW EXECUTE PROCEDURE trigger_soft_delete();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

---------------------------------------------

DROP TABLE "deleted_activity_logs";
DROP TABLE "activity_logs";

-- +goose StatementEnd
