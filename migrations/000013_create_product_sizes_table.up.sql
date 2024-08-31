CREATE TABLE "product_sizes" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50),
    "add_price" int,
    "product_id" INT REFERENCES "products"("id")
)