CREATE TABLE "product_images" (
    "id" SERIAL PRIMARY KEY,
    "image" VARCHAR(255),
    "product_id" int REFERENCES "products"("id")
)