# Golang API example with GIN

This is a repository with examples of HTTP requests with GIN and mySQL

## Requirements

**Go 1.16 >**
```
$ Get installer in: https://go.dev/learn/
```

**MySQL Database**
```
$ Get installers in:
- XAMPP: https://www.apachefriends.org/pt_br/index.html
- HeidiSQL: https://www.heidisql.com/download.php
```

## Getting Started



All environment variables are in `env.yaml`, configure them according to your database.

All executables are defined inside `main.go`.

```
$ go run main.go
```

The server should start at: [`http:localhost:8080`](http:localhost:8080)

## Glossary

1. Database: SQL Commands to create your db.
2. Domain: Entities, model and usecase
3. Infrastructure: Basic configuration, Database connection, handler and repository
4. Interfaces: All project interfaces
5. Utils: Packages that are used throughout the project

## Architecture Goals

1. Follow the concepts of the clean architecture, with extensive use of interfaces.


## Available Endpoint

### `POST /createUser`

- Create a  new user.

```
curl --location --request POST 'http://localhost:8080/createUser' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "string",
    "lastName": "string",
    "age": 5
}'
```

### `GET /getUsers`

- Get all users

```
curl --location --request GET 'http://localhost:8080/getUsers' \
```


### `GET /getUserById/:id`

- Get user by ID


```
curl --location --request GET 'http://localhost:8080/getUserById/10'
```

### `DELETE /:id`

- Delete an user

```
curl --location --request DELETE 'http://localhost:8080/10'
```

### `PUT /:id`

- Update an existing user

```
curl --location --request PUT 'http://localhost:8080/11' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "String",
    "lastName": "String",
    "age": 10
}'

```
