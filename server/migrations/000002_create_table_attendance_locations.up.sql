CREATE TABLE IF NOT EXISTS "attendance_locations" (
    "id" SERIAL PRIMARY KEY,
    "longitude" FLOAT NOT NULL,
    "latitude" FLOAT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);