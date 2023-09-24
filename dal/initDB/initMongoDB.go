package initDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MONGODB_DB *mongo.Client

func InitMongoDB() {
	var err error
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	MONGODB_DB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = MONGODB_DB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
