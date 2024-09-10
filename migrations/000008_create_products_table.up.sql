CREATE TABLE "products" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(50),
    "description" TEXT,
    "price" INT,
    "stock" INT,
    "user_id" INT REFERENCES "users"("id")
)