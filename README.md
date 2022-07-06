# Majoo Goalng Bootcamp

## Run project

 - Go to root folder
 - Run command
    ```
    go run ./main.go
    ```
- Server will run in `localhost:8000`


## Detail project's folder
- api : routing for endpoint
    - admin_book    : endpoint for book's service that need authorization
    - book  : endpoint for book's service that open for public
    - user  : endpoint for user's service

- database : setup database connection, model, and seeder
    - database  : setup database connection, call migration, and seeder
    - seeder    : database's seeder

- dto   : database object (model)
    - Book : Book's model
    - Model : Template model (equals to gorm.Model. Modify in naming for response)
    - User : User's model

- repository : all logic for CRUD
    - repo_book : book's repository
    - user_book : user's repository

- security  : service for security
    - JWTAuth   : JWTAuth logic

- server    : routing to each services
    - routes

- usecase   : all controller logic
    - ucase_book    : book's controler
    - ucase_user    : user's controler

- util  : utilities. Used as helper
    - config    : load data from config.env
    - encryption    : logic for hashing data

## Testing Hit Endpoint
- Use Postman
- List endpoint (HOST: localhost:8000):
```
    Public
    - GET {HOST}/v1/book        (Fetch all book)
    - GET {HOST}/v1/book/:ID    (Fetch book by id)
    - POST {HOST}/v1/login      (user login)
        -> form body: email, password. Check seeder data in database/seeder.go

    Only Authorized User
    - POST {HOST}/v1/book   (Insert new book)
        -> form / raw body: name, author, description, publisher, isbn. Example request:
            {
            "name": "Hunger Games",
            "author": "Suzanne Collins",
            "description": "The Hunger Games is a 2008 dystopian novel by the American writer Suzanne Collins.",
            "publisher": "Scholastic Press",
            "isbn": "978-0-439-02352-8"
            }
    - PUT {HOST}/v1/book    (Update book)
        -> form / raw body: id, name, author, description, publisher, isbn. Example request:
            {
                "id": 5,
                "name": "Hunger Games",
                "author": "Suzanne Collins",
                "description": "The Hunger Games is a 2008 dystopian novel by the American writer Suzanne Collins.",
                "publisher": "Scholastic Press",
                "isbn": "978-0-439-02352-8"
            }
    - DELETE {HOST}/v1/book/:ID (Delete book)
```