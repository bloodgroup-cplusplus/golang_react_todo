package middleware

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")
	clientOptions := options.Client().ApplyURL(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

}

func loadTheEnv() {
	err := godotenv.Load("env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func CreateTask() {

}

func TaskComplete() {

}

func UndoTask() {

}

func DeleteTask() {

}

func deleteAllTasks() {

}
