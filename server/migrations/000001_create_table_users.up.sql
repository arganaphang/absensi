CREATE TYPE "user_role_enum" AS ENUM ('admin', 'staff');

CREATE TABLE IF NOT EXISTS "users" (
    "id" VARCHAR PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "fullname" VARCHAR(255) NOT NULL,
    "birthdate" DATE NOT NULL,
    "position" VARCHAR(255) NOT NULL,
    "password" TEXT NOT NULL,
    "phone" VARCHAR(20) NOT NULL,
    "address" TEXT,
    "role" user_role_enum NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);