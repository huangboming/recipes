# Simple recipes project

Learn from "Building Distributed Application in Gin" by Mohamed Labouardy

## Run the project

Prerequisites:

- Download Docker
- Download MongoDB
- Download Node.js
- Download Go

Run the project with the following instructions:

1. Download the project
2. Run MongoDB and Redis in Docker: `docker run -d --name mongodb -v mongodb-data:/data/db -e ?MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:latest` and `docker run -d --name redis -p 6379:6379 redis:latest`
3. Open terminal, go to `YOUR_GO_PATH/recipes-api/cmd`, run `MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo X_API_KEY=apikey go run main.go`
4. Open terminal, go to `YOUR_GO_PATH/recipes-web`, run `npm install` and `npm run build`
