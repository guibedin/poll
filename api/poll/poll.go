package poll

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Vote struct {
	OptionIDs []string `json:"options"`
}

type Option struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Description string             `json:"description"`
	Count       int64              `json:"count"`
}

type Poll struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `json:"title"`
	Options   []Option           `json:"options"`
	CreatedAt time.Time          `json:"createdAt"`
	IsActive  bool               `json:"isActive"`
}

func connectToMongo(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	failOnError(err, "failed to connect to mongo")
	return client
}

func New(title string, opts []Option) (*Poll, error) {
	// Create options with ID
	var options []Option
	for _, opt := range opts {
		options = append(options, Option{
			ID:          primitive.NewObjectID(),
			Description: opt.Description,
			Count:       0,
		})
	}
	// Create poll object
	poll := &Poll{
		ID:        primitive.NewObjectID(),
		Title:     title,
		Options:   options,
		CreatedAt: time.Now(),
		IsActive:  true,
	}

	// Save poll to DB
	err := poll.Save()

	// Return poll
	return poll, err
}

func (p *Poll) Save() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connectToMongo(ctx)
	defer client.Disconnect(ctx)

	collection := client.Database("poll").Collection("polls")
	_, err := collection.InsertOne(ctx, p)
	failOnError(err, "Failed to save to mongo")
	return err
}

func Get(pollId string) *Poll {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connectToMongo(ctx)
	defer client.Disconnect(ctx)

	collection := client.Database("poll").Collection("polls")
	var result *Poll
	id, _ := primitive.ObjectIDFromHex(pollId)
	err := collection.FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne(),
	).Decode(&result)
	failOnError(err, "Failed to FindOne")

	return result
}
