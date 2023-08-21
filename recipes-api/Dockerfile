FROM golang:latest AS build_env

WORKDIR /go/src/app/recipes-api

# install dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# build app
RUN CGO_ENABLED=0 GOOS=linux go build -o main /go/src/app/recipes-api/cmd

# run app
EXPOSE 8080
CMD ["/go/src/app/recipes-api/main"]

