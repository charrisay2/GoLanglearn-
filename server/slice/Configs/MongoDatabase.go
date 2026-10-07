package configs

import (
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func Connect() error {
	// mongoURI := os.Getenv("MONGO_URI")
	// dbName := os.Getenv("DB_NAME")
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb+srv://2431540075_db_user:PcjkPfxTw13deoEf@cluster0.erdxaxu.mongodb.net"))
	if err != nil {
		return err
	}
	DB = client.Database("users")
	return nil
}
