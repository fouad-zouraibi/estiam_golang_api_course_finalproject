package repos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phramos07/finalproject/types"
)

type UserRepository interface {
	GetUser(context.Context, string) (*types.User, error)
	CreateUser(context.Context, *types.User) error
}

type userRepositoryImpl struct {
	dbConn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) UserRepository {
	return &userRepositoryImpl{
		dbConn: conn,
	}
}

/*REMOVE THIS METHOD*/
/*REMOVE THIS METHOD*/
const SQL_GET_USER = `
		select 
			u.id,
			u.username,
			u.pass
		from
			"user" as u
		where u.id = $1;`

func (repo *userRepositoryImpl) GetUser(c context.Context, userId string) (*types.User, error) {
	rows, err := repo.dbConn.Query(c, SQL_GET_USER, userId)
	if err != nil {
		return nil, fmt.Errorf("error during query to get user: %v", err)
	}

	if rows.Next() {
		user := &types.User{}
		err = rows.Scan(
			&user.Id,
			&user.Password,
			&user.Username,
		)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, nil
}

/*IMPLEMENT THIS METHOD*/
const SQL_INSERT_USER = `
	INSERT INTO "user" (id, username, pass) VALUES ($1, $2, $3);`

func (repo *userRepositoryImpl) CreateUser(c context.Context, user *types.User) error {
	_, err := repo.dbConn.Exec(c, SQL_INSERT_USER, user.Id, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("error during query to create user: %v", err)
	}
	return nil
}
