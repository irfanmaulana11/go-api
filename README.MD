# Interview Assesment 
#### by Muhammad Irfan Maulana

### Introduction
- This project use GIN as the framework. ([GIN](https://github.com/gin-gonic/gin)), PostgreSQL, Gorm and Goose
- Unit test on db repositories
- Auto db migrations

### Installation
1. Run `go mod tidy`
2. Rename `.env.example` file to `.env`
3. Make sure database config in `.env` file
4. Create Database with name `api_go`

### Postman Collection
you can see documentation and run collections via [postman collections here](https://documenter.getpostman.com/view/20624986/2s9YR6ZDih) or import manuallly fie `API GO.postman_collection.json`

### How to Run
1. run
    ```
    docker-compose up
    ```
2. stop
    ```
    docker-compose down
    ```
3. if you need to re-build docker / update docker container (if you have any changes of code), please run 
    ```
    docker-compose up --build
    ```
