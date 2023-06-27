# fiber-example
* install packages with go install
* move `env_sample` to `.env` and set credentials
* create database `CREATE DATABASE IF NOT EXISTS database_name` where database_name was previous set in your `env` file
* run migrations `go run main.go --migrate=yes`
* run proyect `go run main.go`


# Endpoints
* sign up and payload
```
curl --location 'localhost:8000/api/v1/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "juan",
  "password": "user123",
  "email": "juan@domain.com"
}'
```

* sign in and payload
```
curl --location 'localhost:8000/api/v1/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "juan@domain.com",
    "password": "user123"
}'
```

* authenticate route
```
curl --location 'localhost:8000/api/v1/current_user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ4ODQ5NzQsIm5hbWUiOiJqYWltZUBkb21haW4uY29tIn0.Vsv4HeoeRmb5ajBeEy1W7wEyOoTQJAmAF9UflaFJhuM'
```
