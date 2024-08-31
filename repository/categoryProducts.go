package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllcategoryproduct(data models.CategoryProduct) ([]models.CategoryProduct, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM category_products`

	rows, err := db.Query(context.Background(), sql)

	if err != nil {
		return []models.CategoryProduct{}, err
	}

	categoryproducts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.CategoryProduct])

	if err != nil {
		return []models.CategoryProduct{}, err
	}

	return categoryproducts, err
}
func FindCategoryProductByCategoryId(id int) ([]models.CategoryProduct, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM "category_products" where "category_id"=$1;`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return []models.CategoryProduct{}, err
	}

	categoryproducts, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.CategoryProduct])

	if err != nil {
		return []models.CategoryProduct{}, err
	}

	return categoryproducts, err

}
