package infrastructure

import (
	"context"
	"golangtour/common/postresql"
	"golangtour/domain"
	"golangtour/persistence"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postresql.GetConnectionPool(
		ctx,
		postresql.Config{
			Host:                  "localhost",
			Port:                  "6432",
			Username:              "postgres",
			Password:              "postgres",
			DbName:                "productapp",
			MaxConnections:        "10",
			MaxConnectionIdleTime: "30s",
		})

	productRepository = persistence.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool) // Clear test data
}

func TestGetAllProductsByStore(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}

	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool) // Clear test data
}

func TestAddProduct(t *testing.T) {
	expectedProduct := domain.Product{
		Id:       1,
		Name:     "Kupa",
		Price:    50.0,
		Discount: 0.0,
		Store:    "ÜZÜM MARKET",
	}

	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    50.0,
		Discount: 0.0,
		Store:    "ÜZÜM MARKET",
	}

	t.Run("AddProduct", func(t *testing.T) {
		err := productRepository.AddProduct(newProduct)
		assert.Nil(t, err)

		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProduct, actualProducts[0])
	})
	clear(ctx, dbPool) // Clear test data
}

func TestGetProductById(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	expectedProduct := domain.Product{
		Id:       1,
		Name:     "AirFryer",
		Price:    3000.0,
		Discount: 22.0,
		Store:    "ABC TECH",
	}

	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, err := productRepository.GetProductById(1)
		assert.Nil(t, err)
		assert.Equal(t, expectedProduct, actualProduct)
	})

	clear(ctx, dbPool) // Clear test data
}

func TestDeleteProductById(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	t.Run("DeleteProductById", func(t *testing.T) {
		err := productRepository.DeleteProductById(1)
		assert.Nil(t, err)

		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
	})

	clear(ctx, dbPool) // Clear test data
}

func TestUpdateProduct(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	expectedProduct := domain.Product{
		Id:       1,
		Name:     "Kulaklık",
		Price:    100.0,
		Discount: 1.0,
		Store:    "ABC TECH",
	}

	updatedProduct := domain.Product{
		Id:       1,
		Name:     "Kulaklık",
		Price:    100.0,
		Discount: 1.0,
		Store:    "ABC TECH",
	}

	t.Run("UpdateProduct", func(t *testing.T) {
		err := productRepository.UpdateProduct(updatedProduct)
		assert.Nil(t, err)

		actualProduct, err := productRepository.GetProductById(1)
		assert.Nil(t, err)
		assert.Equal(t, expectedProduct, actualProduct)
	})

	clear(ctx, dbPool) // Clear test data
}
