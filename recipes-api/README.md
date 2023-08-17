# Simple recipe API

Learn from "Building Distributed Application in Gin" by Mohamed Labouardy

Done:

- RESTfulAPI with Gin
- Database(in Docker): MongoDB, Redis(for cache)
- Authorization with API key(disabled with comments)


Example: `MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo X_API_KEY=apikey go run main.go`

## APIs

- `GET /recipes`: fetch all recipes in the database.
- `POST /recipes`: add a recipe to the database.
- `PUT /recipes/{id}`: update a recipe in the database.
- `DELETE /recipes/{id}`: delete a recipe in the database
- `GET /recipes/search?tag={tag}`: search recipes by tags
