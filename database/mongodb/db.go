package mongodb

import (
	"Amooz/pkg/common"
	"Amooz/pkg/config"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDb struct {
	mongoDb *mongo.Database
}

func NewMongoDb(cfg config.MongodbConfig) (*MongoDb, error) {
	ConnetionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Pass, cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(ConnetionString).SetMaxPoolSize(uint64(cfg.MaxPoolSize))
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("Connect Error ", err)
	}

	// Check the connection
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Info("Ping Error ", err)
		return nil, err
	}

	existDb(client, cfg)
	//seedData(database)
	log.Info("mongoClient connected")
	return &MongoDb{mongoDb: client.Database(cfg.DbName)}, nil
}

func seedData(database *mongo.Database) {
	//database.Collection(models_grpc.Collection_ByBit_MGIIL).
	//	InsertOne(context.TODO(), models_grpc.ByBitMarketGetInstrumentsInfoLinear{
	//
	//		baseCoin:    "BTC-2",
	//		Status:    "Status-2",
	//		CreatedAt: time.Now(),
	//	})

}
func existDb(client *mongo.Client, cfg config.MongodbConfig) {

	var d *mongo.Database
	filter := bson.D{{"name", bson.D{{"$eq", cfg.DbName}}}}
	names, err := client.ListDatabaseNames(context.Background(), filter)
	if err != nil {
		log.Info("ListDatabaseNames Error ", err)
	}
	if len(names) == 0 {
		d = client.Database(cfg.DbName)
		filter := bson.D{{}}
		collections, err := d.ListCollections(context.Background(), filter)
		if err != nil {
			log.Info("ListCollections Error ", err)
		}
		var colls []string

		if err = collections.All(context.TODO(), &colls); err != nil {
			log.Info("ListCollections All Error ", err)
		}
		if len(colls) == 0 {
			d.CreateCollection(context.Background(), common.Collection_CryptoCurrency)
			d.CreateCollection(context.Background(), common.Collection_FearGreed)
		}
	}
	if d == nil {
		d = client.Database(cfg.DbName)
	}
	//return d
}
func (m *MongoDb) MongoConn() *mongo.Database {
	return m.mongoDb
}
