CREATE TABLE "profile" (
    "id" SERIAL PRIMARY KEY,
    "full_name" VARCHAR(50),
    "phone_number" VARCHAR(50),
    "address" TEXT,
    "image" VARCHAR(255),
    "user_id"int REFERENCES "users"("id")
)
