package database

import (
	"context"
	"fmt"
	"go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

func (db *DB) SaveCreator(inputCreator *model.NewCreator) *model.Creator {
	collection := db.client.Database("graphql_art").Collection("creators")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, inputCreator)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Creator{
		ID:   res.InsertedID.(primitive.ObjectID).Hex(),
		Name: inputCreator.Name,
		Age:  inputCreator.Age,
		//Arts: 		inputCreator.Arts,
	}
}

func (db *DB) SaveArt(inputArt *model.NewArt) *model.Art {
	collection := db.client.Database("graphql_art").Collection("arts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, inputArt)
	if err != nil {
		log.Fatal(err)
	}

	creator := db.FindCreatorByID(inputArt.CreatorID)
	return &model.Art{
		ID:      res.InsertedID.(primitive.ObjectID).Hex(),
		Name:    inputArt.Name,
		Type:    inputArt.Type,
		Creator: creator,
	}
}

func (db *DB) FindCreatorByID(ID string) *model.Creator {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("graphql_art").Collection("creators")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	creator := model.Creator{}
	err = res.Decode(&creator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(creator)
	return &creator
}

func (db *DB) FindArtByID(ID string) *model.Art {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("graphql_art").Collection("arts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	//fmt.Println(res.DecodeBytes())

	artDecoded := model.ArtDecoded{}
	err = res.Decode(&artDecoded)
	if err != nil {
		log.Fatal(err)
	}

	creator := db.FindCreatorByID(artDecoded.CreatorId)
	art := model.Art{artDecoded.ID, artDecoded.Name, artDecoded.Type, creator}

	return &art
}

func (db *DB) FindAllCreators() []*model.Creator {
	collection := db.client.Database("graphql_art").Collection("creators")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var creators []*model.Creator
	for cur.Next(ctx) {
		var creator *model.Creator
		err := cur.Decode(&creator)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(creator.ID)
		creators = append(creators, creator)
	}
	return creators
}

func (db *DB) FindAllArts() []*model.Art {
	collection := db.client.Database("graphql_art").Collection("arts")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var arts []*model.Art
	for cur.Next(ctx) {
		artDecoded := model.ArtDecoded{}
		err := cur.Decode(&artDecoded)
		if err != nil {
			log.Fatal(err)
		}

		creator := db.FindCreatorByID(artDecoded.CreatorId)
		art := model.Art{artDecoded.ID, artDecoded.Name, artDecoded.Type, creator}

		arts = append(arts, &art)
	}
	return arts
}
