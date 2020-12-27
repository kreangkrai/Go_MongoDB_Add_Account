package Controller

import (
	"context"
	"log"
	"time"

	"github.com/kriangkrai/Mee/MongoDB/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadAccount(name string) []Models.AccountModel {

	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{"name": name}
	SingleResult, errFind := collection.Find(ctx, filter)
	if errFind != nil {
		panic(errFind)
	}
	cancelFindOne()

	datas := []Models.AccountModel{}
	defer SingleResult.Close(ctx)
	for SingleResult.Next(ctx) {
		var episode Models.AccountModel
		if err := SingleResult.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, episode)
	}

	return datas
}

func InsertAccount(account Models.AccountModel) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, Models.AccountModel{ID: primitive.NewObjectID(), Name: account.Name, Password: account.Password})

	if err != nil {
		return nil, err
	}
	return res, nil
}
