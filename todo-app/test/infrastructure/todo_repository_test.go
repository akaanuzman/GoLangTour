package infrastructure

import (
	"context"
	"os"
	"testing"
	"time"
	"todoapp/common/postresql"
	"todoapp/domains"
	"todoapp/persistence"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var todoRepository persistence.ITodoRepository
var dbPool *pgxpool.Pool
var ctx context.Context

var expectedTodos = []domains.Todo{
	{
		Id:          1,
		Title:       "Buy Milk",
		Description: "Buy 1lt milk from the grocery store",
		CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsDone:      false,
		DueDate:     nil,
	},
	{
		Id:          2,
		Title:       "Buy Bread",
		Description: "Buy 1 loaf of bread from the bakery",
		CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsDone:      false,
		DueDate:     nil,
	},
	{
		Id:          3,
		Title:       "Buy Eggs",
		Description: "Buy 1 dozen eggs from the grocery store",
		CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsDone:      false,
		DueDate:     nil,
	},
	{
		Id:          4,
		Title:       "Buy Butter",
		Description: "Buy 1 stick of butter from the grocery store",
		CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsDone:      false,
		DueDate:     nil,
	},
}

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postresql.GetConnectionPool(
		ctx,
		postresql.PostreSQLConfig{
			Host:                  "localhost",
			Port:                  "6432",
			Username:              "postgres",
			Password:              "postgres",
			DbName:                "todoapp",
			MaxConnections:        "10",
			MaxConnectionIdleTime: "30s",
		})

	todoRepository = persistence.NewTodoRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllTodos(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	t.Run("GetAllTodos", func(t *testing.T) {
		actualTodos := todoRepository.GetAllTodos()
		assert.Equal(t, 4, len(actualTodos))
		for i, actual := range actualTodos {
			assert.Equal(t, expectedTodos[i].Title, actual.Title)
			assert.Equal(t, expectedTodos[i].Description, actual.Description)
			assert.Equal(t, expectedTodos[i].IsDone, actual.IsDone)
			assert.True(t, expectedTodos[i].CreatedAt.Equal(actual.CreatedAt))
			assert.True(t, expectedTodos[i].UpdatedAt.Equal(actual.UpdatedAt))
		}
	})

	clear(ctx, dbPool) // Clear test data
}

func TestGetTodoById(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	expectedTodo := domains.Todo{
		Id:          1,
		Title:       "Buy Milk",
		Description: "Buy 1lt milk from the grocery store",
		CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsDone:      false,
		DueDate:     nil,
	}

	t.Run("GetTodoById", func(t *testing.T) {
		actualTodo, err := todoRepository.GetTodoById(1)
		assert.Nil(t, err)
		assert.Equal(t, expectedTodo.Id, actualTodo.Id)
		assert.Equal(t, expectedTodo.Title, actualTodo.Title)
		assert.Equal(t, expectedTodo.Description, actualTodo.Description)
		assert.Equal(t, expectedTodo.IsDone, actualTodo.IsDone)
		assert.True(t, expectedTodo.CreatedAt.Equal(actualTodo.CreatedAt))
		assert.True(t, expectedTodo.UpdatedAt.Equal(actualTodo.UpdatedAt))
	})

	clear(ctx, dbPool) // Clear test data
}

func TestGetUndoneTodos(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	t.Run("GetUndoneTodos", func(t *testing.T) {
		actualTodos := todoRepository.GetDoneOrUndoneTodos(false)
		assert.Equal(t, 4, len(actualTodos))
		for i, actual := range actualTodos {
			assert.Equal(t, expectedTodos[i].Title, actual.Title)
			assert.Equal(t, expectedTodos[i].Description, actual.Description)
			assert.Equal(t, expectedTodos[i].IsDone, actual.IsDone)
			assert.True(t, expectedTodos[i].CreatedAt.Equal(actual.CreatedAt))
			assert.True(t, expectedTodos[i].UpdatedAt.Equal(actual.UpdatedAt))
		}
	})
}

func TestGetDoneTodos(t *testing.T) {
	ctx = context.Background()

	setup(ctx, dbPool) // Setup test data

	t.Run("GetDoneTodos", func(t *testing.T) {
		actualTodos := todoRepository.GetDoneOrUndoneTodos(true)
		assert.Equal(t, 0, len(actualTodos))
	})

	clear(ctx, dbPool) // Clear test data
}
