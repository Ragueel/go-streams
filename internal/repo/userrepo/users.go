package userrepo

import (
	"context"
	goqu "github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"go-streams/internal/domain/dbmodel"
)

type UsersRepository struct {
	connection *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UsersRepository {
	return &UsersRepository{conn}
}

func (repo *UsersRepository) GetUser(ctx context.Context, id uuid.UUID) <-chan *dbmodel.User {
	res := make(chan *dbmodel.User)
	go func() {
		defer close(res)
		query := goqu.Select().From("users").Where()
		repo.connection.Exec(query.ToSQL())
		res <- nil
	}()

	return res
}

func (repo UsersRepository) CreateUser() {
}

func (repo UsersRepository) SaveUser(user *dbmodel.User) {

}

func runQuery(ctx context.Context, connection *pgx.Conn, query string) <-chan interface{} {
	result := make(chan interface{})
	go func() {
		rows, err := connection.QueryEx(ctx, query, nil)
		defer rows.Close()
		if err != nil {
			result <- nil
		}
		allRows, _ := rows.Values()
		for _, item := range allRows {
			result <- item
		}
	}()

	return result
}
