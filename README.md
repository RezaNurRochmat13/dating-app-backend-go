# Dating Apps

Dating apps built using Golang.

## Directory sturcture

Directory structure of the project.

- config => Contains configuration database, libraries etc.
- controller => Contains presentation logic.
- entities => Contains core of database models.
- middleware => Contains custom middlewares.
- repository => Contains database operations.
- service => Contains bussiness logic application.
- utils => Contains utility functions.
- tests => Contains unit test files.

## Technology

Technology used when built application.

- Golang => Main programming language.
- PostgreSQL => Main database.
- GORM => Main Object Relational Mapper.

## How to run service

Guide how to run application.

- Clone this repository
- Setup the databases on config/database.go
- Run command `go run .`
- Access the endpoint

## API Documentation

API Documentation on these dating apps.

#### Auth

Endpoint of Auth Login

-- Login
```
- URL : /api/auth/login
- Method : POST
- Request Body : JSON
{
  "email": "xxx@xxx.com",
  "password": "xxx"
}
- Response : 200 OK if valid user and generated token
```

-- Register
```
- URL : /api/auth/register
- Method : POST
- Request Body : JSON
{
  "username": "xxx",  
  "email": "xxx@xxx.com",
  "password": "xxx"
}
- Response : 201 OK if valid user
```

#### Like User

Endpoint of liked user

-- Like
```
- URL : /api/users/:id/like
- Method : POST
- Request Body : JSON
{
  "liker_id": "xxx"
}
- Response : 200 OK if valid payload
```

-- All Users
```
- URL : /api/users/
- Method : GET
- Request Body : -
- Response : 200 OK
```

## Author

- Reza Nur Rochmat <rezanurrochmat@gmail.com>