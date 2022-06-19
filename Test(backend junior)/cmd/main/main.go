package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/NetworkPy/TestTaskJuniorBackDev/internal/config"
	"github.com/NetworkPy/TestTaskJuniorBackDev/internal/tokens"
	"github.com/gin-gonic/gin"
	_ "github.com/openshift/osin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	pathConf string
)

func init() {
	flag.StringVar(&pathConf, pathConf, "config/conf.yaml", "path to configuration file")
	flag.Parse()
}

func main() {

	newConfig, err := config.NewConfig(pathConf)

	if err != nil {
		log.Fatalln(err)
	}

	db, err := mongo.NewClient(options.Client().ApplyURI(newConfig.MongoUrl))

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Connect(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	token := tokens.NewToken(db, newConfig.SecretKey)

	router := gin.Default()

	handler := tokens.NewHandler(token)

	handler.Register(router)

	start(router, newConfig)
}

func start(router *gin.Engine, config *config.Config) {

	fmt.Printf("Start server at %s", config.Port)

	log.Fatalln(router.Run(fmt.Sprintf("%s", config.Port)))
}
