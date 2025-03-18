CREATE TYPE "attendance_type_enum" AS ENUM ('in', 'out');

CREATE TABLE IF NOT EXISTS "attendances" (
    "id" SERIAL PRIMARY KEY,
    "created_at" TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "type" attendance_type_enum NOT NULL,
        "user_id" VARCHAR NOT NULL,
        "longitude" FLOAT NOT NULL,
        "latitude" FLOAT NOT NULL,
        "note" TEXT,
        CONSTRAINT fk_attendances_user_id FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);