package main

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type vote struct {
// 	PollId  string   `json:"pollId"`
// 	Options []string `json:"optionIds"`
// }

// type Option struct {
// 	ID          primitive.ObjectID `bson:"_id" json:"id"`
// 	Description string             `json:"description"`
// 	Count       int64              `json:"count"`
// }

// type Poll struct {
// 	ID        primitive.ObjectID `bson:"_id" json:"id"`
// 	Title     string             `json:"title"`
// 	Options   []Option           `json:"options"`
// 	CreatedAt time.Time          `json:"createdAt"`
// 	IsActive  bool               `json:"isActive"`
// }

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectToMongo(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	failOnError(err, "failed to connect to mongo")
	return client
}

func countVote(v vote) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connectToMongo(ctx)
	defer client.Disconnect(ctx)

	collection := client.Database("poll").Collection("polls")
	var poll *Poll
	id, _ := primitive.ObjectIDFromHex(v.PollId)
	collection.FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne()).Decode(&poll)

	log.Println(poll)
	for _, voteOption := range v.Options {
		for _, pollOption := range poll.Options {
			if strings.Compare(voteOption, pollOption.ID.Hex()) == 0 {
				pollOption.Count++
				log.Println(poll)
			}
		}
		// result := collection.UpdateByID(
		// 	ctx,
		// 	bson.M{"_id", id},
		// 	bson.D{
		// 		{"$set", bson.D{{"count"}}}
		// }  )
	}

	//_, err := collection.UpdateOne(context)
	//failOnError(err, "Failed to save to mongo")
	//return err
}

func main() {

}
