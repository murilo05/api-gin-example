package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.engdb.com.br/apigin/domain/entities"
)

func GetUser(ctx context.Context) ([]entities.User, int, error) {
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"apigin")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM users")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	users := []entities.User{}
	for results.Next() {
		var user entities.User

		err = results.Scan(&user.Id, &user.Name, &user.LastName, &user.Age)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil, 500, err
		}

		// append the product into products array
		users = append(users, user)
	}
	if len(users) == 0 {
		return users, 204, nil
	}

	return users, 200, nil
}

func GetUserById(ctx context.Context, userId string) (*entities.User, int, error) {
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"apigin")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM users WHERE ID = ?", userId)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 500, err
	}

	user := entities.User{}

	if results.Next() {
		err = results.Scan(&user.Id, &user.Name, &user.LastName, &user.Age)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil, 500, err
		}
	} else {
		return nil, 204, nil
	}

	return &user, 200, nil
}

func CreateUser(ctx context.Context, body entities.User) (int, error) {
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"apigin")
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (name,last_name,age) VALUES (?,?,?)", body.Name, body.LastName, body.Age)

	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	defer insert.Close()

	return 200, nil
}

func DeleteUser(ctx context.Context, userId string) (int, error) {
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"apigin")
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	defer db.Close()

	delete, err := db.Query("DELETE from users WHERE id = ?", userId)

	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	defer delete.Close()

	return 200, nil
}

func EditUser(ctx context.Context, body entities.User, userId string) (int, error) {
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"apigin")
	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	defer db.Close()

	delete, err := db.Query("UPDATE users SET name = ?, last_name = ?, age = ? WHERE id = ?", body.Name, body.LastName, body.Age, userId)

	if err != nil {
		fmt.Println("Err", err.Error())
		return 500, err
	}

	defer delete.Close()

	return 200, nil
}
