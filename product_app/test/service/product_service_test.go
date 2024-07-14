package service

import (
	"golangtour/domain"
	"golangtour/service"
	"golangtour/service/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "ABC TECH",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)

	productService = service.NewProductService(fakeProductRepository)
	os.Exit(m.Run())
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("Should return all products", func(t *testing.T) {
		acutalProducts := productService.GetAll()
		assert.Equal(t, 2, len(acutalProducts))
	})
}

func Test_WhenNoValidationErrorOccured_ShouldAddProduct(t *testing.T) {
	t.Run("Should add product", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "Laptop",
			Price:    5000.0,
			Discount: 0.0,
			Store:    "ABC TECH",
		})
		acutalProducts := productService.GetAll()
		assert.Equal(t, 3, len(acutalProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Laptop",
			Price:    5000.0,
			Discount: 0.0,
			Store:    "ABC TECH",
		}, acutalProducts[len(acutalProducts)-1])
	})
}

func Test_WhenDiscountIsHigherThen70_ShouldNotAddProduct(t *testing.T) {
	t.Run("WhenDiscountIsHigherThen70_ShouldNotAddProduct", func(t *testing.T) {
		err := productService.Add(model.ProductCreate{
			Name:     "Laptop",
			Price:    5000.0,
			Discount: 71.0,
			Store:    "ABC TECH",
		})

		actualProducts := productService.GetAll()
		assert.Equal(t, 2, len(actualProducts))
		assert.Equal(t, "Discount cannot be greater more than 70%", err.Error())
	})
}
