package main

import (
	"os"
	// "github.com/gin-gonic/gin"
	"github.com/JJFelix/restaurant_management/database"
	"github.com/JJFelix/restaurant_management/internal/database"
	// routes "github.com/JJFelix/restaurant_management/routes"
	// middleware "github.com/JJFelix/restaurant_management/middleware"
	// "go.mongodb.org/mongo-driver/mongo" // change to postgres
	"database/sql"

)


// mongodb
// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// postgresql
var db *sql.DB
var queries *database.Queries

func main(){
	port := os.Getenv("PORT")

	if port ==""{
		port = "8000"
	}
	
	// initialize PostgreSQL db connection
	dbConn = database.ConnectToDB()
	defer dbConn.Close()

	// initialize the queries
	queries = db.New(dbConn)


	// using gin router
	// router := gin.New()
	// router.SetTrustedProxies([]string{})
	// router.Use(gin.Logger())

	// routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	// routes.FoodRoutes(router)
	// routes.MenuRoutes(router)
	// routes.TableRoutes(router)
	// routes.OrderRoutes(router)
	// routes.OrderItemRoutes(router)
	// routes.InvoiceRoutes(router)

	// router.Run(":" + port)


}