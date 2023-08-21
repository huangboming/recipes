// TODO: get env variable from .env file: https://github.com/joho/godotenv
// MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo REDIS_URL=127.0.0.1:6379 CORS_ORIGIN=http://localhost:3000 X_API_KEY=apikey go run main.go
// MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo JWT_SECRET=eUbP9shywUygMx7u go run main.go

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"recipes-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// global variables

var recipesHandler *handlers.RecipesHandler

func init() {

	// read data from JSON (if needed)
	// recipes = make([]Recipe, 0)
	// file, _ := os.ReadFile("../internal/models/data/recipes.json")
	// _ = json.Unmarshal([]byte(file), &recipes)

	// connect to mongodb
	ctx := context.Background()
	client, _ := mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping(ctx)
	fmt.Println(status)

	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)

	// insert JSON data to database
	// var listOfRecipes []interface{}
	// for _, recipe := range recipes {
	// 	listOfRecipes = append(listOfRecipes, recipe)
	// }
	// collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	// insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Inserted recipes: ", len(insertManyResult.InsertedIDs))
}

func main() {
	r := gin.Default()
	// CORS
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{os.Getenv("CORS_ORIGIN")},
	// 	AllowMethods:     []string{"GET", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "X-API-KEY"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	r.Use(cors.Default())
	r.GET("/recipes", recipesHandler.ListRecipesHandler)
	r.POST("/recipes", recipesHandler.NewRecipeHandler)
	r.PUT("/recipes/:id", recipesHandler.UpdateRecipesHandler)
	r.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	r.GET("/recipes/search", recipesHandler.SearchRecipesHanlder)

	// add authorization with API key minddleware
	// authorized := r.Group("/")
	// authorized.Use(handlers.AuthMiddleware())
	// // APIs
	// authorized.GET("/recipes", recipesHandler.ListRecipesHandler)
	// authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
	// authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipesHandler)
	// authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	// authorized.GET("/recipes/search", recipesHandler.SearchRecipesHanlder)

	r.Run()
}
