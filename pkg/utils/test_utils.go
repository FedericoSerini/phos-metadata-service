package utils

import (
	"context"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI    string
	mongoClient *mongo.Client
)

func GetMongoClient() *mongo.Client {
	return mongoClient
}

// Setup MongoDB Test Container
func setupTestContainer() string {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6.0",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
	}

	mongoC, err := testcontainers.GenericContainer(context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		log.Fatal(err)
	}

	// Get Container Host & Port
	host, err := mongoC.Host(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	port, err := mongoC.MappedPort(context.Background(), "27017")
	if err != nil {
		log.Fatal(err)
	}

	mongoURI := "mongodb://" + host + ":" + port.Port()
	return mongoURI
}

// Connect to the test MongoDB instance
func SetupTestDB() {
	mongoURI = setupTestContainer()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client
}

// Tear down MongoDB connection
func TeardownTestDB() {
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
