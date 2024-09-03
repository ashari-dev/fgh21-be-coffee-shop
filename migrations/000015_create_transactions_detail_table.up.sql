CREATE TABLE "transaction_details" (
    "id" SERIAL PRIMARY KEY,
    "quantity" int,
    "product_id" INT REFERENCES "products"("id"),
    "variant_id" INT REFERENCES "product_variants"("id"),
    "product_size_id" INT REFERENCES "product_sizes"("id")
)