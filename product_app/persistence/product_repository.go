package persistence

import (
	"context"
	"golangtour/domain"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")

	if err != nil {
		log.Error("Unable to get products:", err)
		return []domain.Product{}
	}

	products := []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		err = productRows.Scan(&id, &name, &price, &discount, &store)
		if err != nil {
			log.Error("Unable to get product:", err)
			return []domain.Product{}
		}

		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})

	}
	return products
}
