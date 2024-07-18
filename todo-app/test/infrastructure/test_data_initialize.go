package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_TODOS = `INSERT INTO todos (title, description, created_at, updated_at)
VALUES('Buy Milk', 'Buy 1lt milk from the grocery store', '2023-01-01 00:00:00', '2023-01-01 00:00:00'),
('Buy Bread', 'Buy 1 loaf of bread from the bakery', '2023-01-01 00:00:00', '2023-01-01 00:00:00'),
('Buy Eggs', 'Buy 1 dozen eggs from the grocery store', '2023-01-01 00:00:00', '2023-01-01 00:00:00'),
('Buy Butter', 'Buy 1 stick of butter from the grocery store', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertTodosResult, insertTodosErr := dbPool.Exec(ctx, INSERT_TODOS)
	if insertTodosErr != nil {
		log.Error(insertTodosErr)
	} else {
		log.Info(fmt.Sprintf("Todos data created with %d rows", insertTodosResult.RowsAffected()))
	}
}
