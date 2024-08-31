CREATE TABLE "category_products" (
    "id" SERIAL PRIMARY KEY,
    "category_id" INT REFERENCES "categories"("id"),
    "product_id" INT REFERENCES "products"("id")
)