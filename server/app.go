package server

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	nounHttp "github.com/PhilWhittingham/vocab-helper-de/noun/delivery/http"
	nounMongo "github.com/PhilWhittingham/vocab-helper-de/noun/repository/mongo"
	"github.com/PhilWhittingham/vocab-helper-de/noun/service"
)

func initDB() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("MONGO_URI")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("MONGO_DB_NAME"))
}

func Run(port string) {
	router := gin.Default()

	db := initDB()
	nr := nounMongo.NewNounRepository(db, viper.GetString("MONGO_COLLECTION_NAME_NOUNS"))
	s := service.NewNounService(nr)

	group := router.Group("/api")

	nounHttp.RegisterHTTPEndpoints(group, s)

	router.Run("localhost:" + port)
}
