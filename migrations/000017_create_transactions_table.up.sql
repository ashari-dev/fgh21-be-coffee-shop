CREATE TABLE "transactions" (
    "id" SERIAL PRIMARY KEY,
    "no_order" int,
    "add_full_name" VARCHAR(50),
    "add_email" VARCHAR(50),
    "add_address" TEXT,
    "payment" VARCHAR(50),
    "user_id" INT REFERENCES "users"("id"),
    "transaction_detail_id" INT REFERENCES "transaction_details"("id"),
    "order_type_id" INT REFERENCES "order_types"("id"),
    "transaction_status_id" INT REFERENCES "transaction_status"("id")
)