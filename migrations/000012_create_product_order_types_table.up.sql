CREATE TABLE "product_order_types" (
    "id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "products"("id")
    "order_type_id" INT REFERENCES "order_types"("id")
)