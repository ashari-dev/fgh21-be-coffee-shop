CREATE TABLE "products" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(50),
    "description" TEXT,
    "price" INT,
    "user_id" INT REFERENCES "users"("id")
)