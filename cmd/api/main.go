package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"product-stock-web-service/domain"
	"product-stock-web-service/product/middleware"
	"product-stock-web-service/product/repository/mysql"
	"product-stock-web-service/product/usecase"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e, ar := initializeMySql()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	productUseCase := usecase.NewProductUsecase(ar, timeoutContext)
	middleware.NewProductHandler(e, productUseCase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func initializeMySql() (*echo.Echo, domain.ProductRepository) {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	/*defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()*/

	e := echo.New()
	ar := mysql.NewProductRepository(dbConn)
	return e, ar
}
