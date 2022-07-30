# product-stock-web-service

## Description
This is an example of implementation of Product Stock Webservice in Go (Golang) projects.

This Project was developed with Clean Architecture in mind
 * Independent of Frameworks. 
 * Testable.
 * Independent of UI. 
 * Independent of Database.
 * Independent of any external agency.

This project has  4 Domain layer :
 * Domain Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

### How To Run This Project
> Make Sure you have run the /migrations/create_product_tables.sql in your mysql

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make test
```

#### Run the Applications
Here is the steps to run the application

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/NagarjunNagesh/product-stock-web-service.git

#move to project
$ cd product-stock-web-service

# Build the docker image first
$ docker-compose up --build

# Run the application
$ go run ./cmd/api/main.go

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:8080/products

# Stop
$ make stop
```
