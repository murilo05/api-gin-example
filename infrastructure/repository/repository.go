package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.engdb.com.br/apigin/domain/entities"
	"gitlab.engdb.com.br/apigin/interfaces"
)

type mysqlRepo struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewMysqlUserRepository(db *sql.DB) interfaces.UserRepo {
	return &mysqlRepo{
		DB: db,
	}
}

func (m *mysqlRepo) GetUser(ctx context.Context) ([]entities.User, int, error) {

	query := "SELECT * FROM users"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	users := []entities.User{}

	for rows.Next() {
		var user entities.User

		err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Age)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil, 500, err
		}

		users = append(users, user)
	}
	if len(users) == 0 {
		return users, 204, nil
	}

	return users, 200, nil
}

func (m *mysqlRepo) GetUserById(ctx context.Context, userId string) (*entities.User, int, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE ID = %s", userId)

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	user := entities.User{}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Age)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil, 500, err
		}
	} else {
		return nil, 204, nil
	}

	return &user, 200, nil
}

func (m *mysqlRepo) CreateUser(ctx context.Context, body entities.User) (int, error) {

	query := fmt.Sprintf("INSERT INTO users (name,last_name,age) VALUES ('%s','%s','%v')", body.Name, body.LastName, body.Age)
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	rows.Close()

	return 200, nil
}

func (m *mysqlRepo) DeleteUser(ctx context.Context, userId string) (int, error) {

	query := fmt.Sprintf("DELETE from users WHERE id = %s", userId)

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	rows.Close()

	return 200, nil
}

func (m *mysqlRepo) EditUser(ctx context.Context, body entities.User, userId string) (int, error) {

	query := fmt.Sprintf("UPDATE users SET name = '%s', last_name = '%s', age = %v WHERE id = %s", body.Name, body.LastName, body.Age, userId)

	fmt.Println(query, "QUERY")

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	rows.Close()

	return 200, nil
}
