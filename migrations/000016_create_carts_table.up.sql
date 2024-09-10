CREATE TABLE "carts" (
    "id" SERIAL PRIMARY KEY,
    "transaction_detail_id" INT REFERENCES "transaction_details"("id"),
    "quantity" int,
    "variant_id" INT REFERENCES "product_variants"("id"),
    "sizes_id" INT REFERENCES "product_sizes"("id"),
    "product_id" INT REFERENCES "products"("id"),
    "user_id" INT REFERENCES "users"("id")
);