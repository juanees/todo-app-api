# Todo App API

Simple TODO Backend app built using

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) 

![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

## Run Locally

Clone the project

```bash
  git clone https://github.com/juanees/todo-app-api.git
```

Go to the project directory

```bash
  cd todo-app-api
```

If you are using Windows you [must install mingw](https://github.com/mattn/go-sqlite3/issues/467)

```bash
   choco install mingw
```

Start the server

```bash
  go run .\cmd\main.go
```

## API Reference

#### Get all todos

```http
  GET /api/v1/todos
```

#### Get todo

```http
  GET /api/v1/todos/:code
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of todo to fetch |

```bash
   curl --location --request GET 'localhost:8080/api/v1/todos/228fd7d1-cec0-4486-ae0e-4db638d302db'
```

#### Add todo

```http
  POST /api/v1/todos
```

```bash
   curl --location --request POST 'localhost:8080/api/v1/todos' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Test todo",
    "description": "test description",
    "completed": false
}'
```

#### Update todo

```http
  PATCH /api/v1/todos/:code
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of todo to update |


```bash
   curl --location --request PATCH 'localhost:8080/api/v1/todos/228fd7d1-cec0-4486-ae0e-4db638d302db' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Changed name",
    "description": "New description",
    "completed": false
}'
```

#### Delete todo

```http
  DELETE /api/v1/todos/:code
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of todo to delete |


```bash
   curl --location --request DELETE 'localhost:8080/api/v1/todos/228fd7d1-cec0-4486-ae0e-4db638d302db'
```
## Roadmap

- Unit Tests

- IOC or Dependency Injection

