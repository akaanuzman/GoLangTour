package service

import (
	"golangtour/domain"
	"golangtour/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{products: initialProducts}
}

func (productRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return productRepository.products
}

func (productRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	return productRepository.products
}

func (productRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	productRepository.products = append(productRepository.products, domain.Product{
		Id:       int64(len(productRepository.products) + 1),
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

func (productRepository *FakeProductRepository) GetProductById(id int64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (productRepository *FakeProductRepository) DeleteProductById(id int64) error {
	return nil
}

func (productRepository *FakeProductRepository) UpdateProduct(product domain.Product) error {
	return nil
}
