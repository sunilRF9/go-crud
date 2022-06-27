package routes
import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)
func main() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
	    ApplyURI("mongodb+srv://sunil:<PASSWORD>@cluster0.bdg5i.mongodb.net/?retryWrites=true&w=majority").
	    SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
	    log.Fatal(err)
	}
}
