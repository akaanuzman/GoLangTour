package controller

import (
	"golangtour/domain"
	"golangtour/service"
	"net/http"
	"strconv"

	"golangtour/controller/request"
	"golangtour/controller/response"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (controller *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", controller.GetById)
	e.GET("/api/v1/products", controller.GetAll)
	e.POST("/api/v1/products", controller.Add)
	e.DELETE("/api/v1/products/:id", controller.DeleteById)
}

func (controller *ProductController) GetById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	product, err := controller.productService.GetById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponse(product))
}

func (controller *ProductController) GetAll(c echo.Context) error {
	store := c.QueryParam("store")
	if len(store) == 0 {
		allProducts := controller.productService.GetAll()
		return c.JSON(http.StatusOK, response.ToResponseList(allProducts))
	}
	productsWithGivenStore := controller.productService.GetAllByStore(store)
	return c.JSON(http.StatusOK, response.ToResponseList(productsWithGivenStore))
}

func (controller *ProductController) Add(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	bindErr := c.Bind(&addProductRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	err := controller.productService.Add(addProductRequest.ToModel())

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, addProductRequest)
}

func (controller *ProductController) Update(c echo.Context) error {
	var updateProductRequest request.AddProductRequest
	bindErr := c.Bind(&updateProductRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	model := updateProductRequest.ToModel()
	err := controller.productService.Update(domain.Product{
		Name:     model.Name,
		Price:    model.Price,
		Discount: model.Discount,
		Store:    model.Store,
	})

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, updateProductRequest)
}

func (controller *ProductController) DeleteById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	err := controller.productService.DeleteById(int64(productId))
	if err != nil {
		println(err.Error())
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}
