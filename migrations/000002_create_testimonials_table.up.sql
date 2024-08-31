CREATE TABLE "testimonials" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50),
    "profession" VARCHAR(50),
    "comment" TEXT,
    "rating" INT,
    "image" VARCHAR(255)
)