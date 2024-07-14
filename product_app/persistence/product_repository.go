package persistence

import (
	"context"
	"golangtour/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(store string) []domain.Product
	AddProduct(product domain.Product) error
	GetProductById(id int64) (domain.Product, error)
	DeleteProductById(id int64) error
	UpdateProduct(product domain.Product) error
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

	return extractProductRows(productRows)
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()
	query := `SELECT * FROM products where store = $1`
	productRows, err := productRepository.dbPool.Query(ctx, query, storeName)

	if err != nil {
		log.Error("Unable to get products:", err)
		return []domain.Product{}
	}

	return extractProductRows(productRows)
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()
	query := `INSERT INTO products (name, price, discount, store) VALUES($1, $2, $3, $4)`
	addNewProduct, err := productRepository.dbPool.Exec(ctx, query, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Unable to add product:", err)
		return err
	}
	log.Info("Product added successfully", addNewProduct)

	return nil
}

func extractProductRows(productRows pgx.Rows) []domain.Product {
	products := []domain.Product{}
	var product domain.Product

	for productRows.Next() {
		err := productRows.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
		if err != nil {
			log.Error("Unable to get product:", err)
			return []domain.Product{}
		}

		products = append(products, domain.Product{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Discount: product.Discount,
			Store:    product.Store,
		})
	}
	return products
}

func (productRepository *ProductRepository) GetProductById(id int64) (domain.Product, error) {
	ctx := context.Background()
	query := `SELECT * FROM products where id = $1`
	productRow := productRepository.dbPool.QueryRow(ctx, query, id)

	var product domain.Product
	err := productRow.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
	if err != nil {
		log.Error("Unable to get product:", err)
		return domain.Product{}, err
	}

	return product, nil
}

func (productRepository *ProductRepository) DeleteProductById(id int64) error {
	ctx := context.Background()

	_, getError := productRepository.GetProductById(id)
	if getError != nil {
		log.Error("Unable to get product:", getError)
		return getError
	}

	query := `DELETE FROM products where id = $1`
	_, deleteError := productRepository.dbPool.Exec(ctx, query, id)
	if deleteError != nil {
		log.Error("Unable to delete product:", deleteError)
		return deleteError
	}

	return nil
}

func (productRepository *ProductRepository) UpdateProduct(product domain.Product) error {
	ctx := context.Background()

	_, getError := productRepository.GetProductById(product.Id)
	if getError != nil {
		log.Error("Unable to get product:", getError)
		return getError
	}

	query := `UPDATE products SET name = $1, price = $2, discount = $3, store = $4 WHERE id = $5`
	_, updateError := productRepository.dbPool.Exec(ctx, query, product.Name, product.Price, product.Discount, product.Store, product.Id)
	if updateError != nil {
		log.Error("Unable to update product:", updateError)
		return updateError
	}

	return nil
}
