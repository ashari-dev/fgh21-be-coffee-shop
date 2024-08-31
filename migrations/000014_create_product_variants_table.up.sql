CREATE TABLE "product_variants" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50),
    "add_price" int,
    "stock" int,
    "product_id" INT REFERENCES "products"("id")
)