package controllers

import (
	"blog/src/models"
	"blog/src/services"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController struct {
	postService services.IPostService
}

func NewPostController(postService services.IPostService) *PostController {
	return &PostController{postService: postService}
}

func (controller *PostController) CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user ID"})
	}

	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	post.Id = primitive.NewObjectID()
	post.AuthorID = authorID
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if err := controller.postService.CreatePost(&post); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}

func (controller *PostController) UpdatePost(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user ID"})
	}

	post.Id = id
	post.AuthorID = authorID
	post.UpdatedAt = time.Now()

	if err := controller.postService.UpdatePost(&post); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}

func (controller *PostController) DeletePost(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	if err := controller.postService.DeletePost(id.Hex()); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Post deleted"})
}

func (controller *PostController) GetAllPosts(c echo.Context) error {
	posts, err := controller.postService.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, posts)
}

func (controller *PostController) GetPostById(c echo.Context) error {
	id := c.Param("id")
	post, err := controller.postService.GetPostById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}
