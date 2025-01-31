package database

import (
	
	"log"
	"os"
	// "time"
	// "context"
	"fmt"

	// add for postgresql
	"database/sql"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// add
	_ "github.com/lib/pq" // PostgreSQL driver

)

//PostgreSQL Connection
func ConnectToDB() *sql.DB {
	// load the .env file
	if err := godotenv.Load(); err != nil{
		log.Fatal("Error loading .env file")
	}

	// Get DB Credentials from environment variable
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	// Construct the connection string
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DBUser, DBPassword, DBName, DBHost, DBPort,
	)

	//Open the DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal("Failed to connect to the database: ", err)
	}

	// ping the db to verify the connection
	err = db.Ping()
	if err != nil{
		log.Fatal("Failed to ping the database: ", err)
	}

	fmt.Println("Connected to PostgreSQL")
	return db
}





//MongoDB Connection
// func DBinstance() *mongo.Client{
// 	// Load .env file	
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	DBUserName := os.Getenv("DB_USERNAME")
// 	DBPassword := os.Getenv("DB_PASSWORD")

// 	if DBUserName == "" || DBPassword == "" {
// 		log.Fatal("Database credentials are missing. Set DB_USERNAME and DB_PASSWORD in environment variables.")
// 	}

// 	MongoDbURI := fmt.Sprintf(
// 		"mongodb+srv://%s:%s@restaurant-management-g.gbdz0.mongodb.net/?retryWrites=true&w=majority&appName=Restaurant-Management-Golang", 
// 		DBUserName, DBPassword,
// 	)  // you can use the online link as well
// 	// fmt.Println(MongoDb) // remove in production

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDbURI))
// 	if err != nil{
// 		log.Fatal(err)
// 	}


// 	fmt.Println("Connected to mongodb")
// 	return client
// }

// var Client *mongo.Client = DBinstance()

// func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
// 	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

// 	return collection
// }