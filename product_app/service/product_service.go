package service

import (
	"errors"
	"golangtour/domain"
	"golangtour/persistence"
	"golangtour/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	GetAll() []domain.Product
	GetAllByStore(store string) []domain.Product
	Update(product domain.Product) error
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) Add(productCreate model.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}

	product := domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	}

	return productService.productRepository.AddProduct(product)
}

func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteProductById(productId)
}

func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetProductById(productId)
}

func (productService *ProductService) GetAll() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

func (productService *ProductService) GetAllByStore(store string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(store)
}

func (productService *ProductService) Update(product domain.Product) error {
	return productService.productRepository.UpdateProduct(product)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70 {
		return errors.New("Discount cannot be greater more than 70%")
	}
	return nil
}
