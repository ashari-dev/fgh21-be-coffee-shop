CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "email" VARCHAR(50) UNIQUE,
    "password" VARCHAR(255),
    "role_id" INT REFERENCES "roles"("id")
)