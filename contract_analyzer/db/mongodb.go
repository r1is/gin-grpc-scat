package db

import (
	"context"
	"fmt"
	"log"
	"pkg/apis"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

var client *qmgo.QmgoClient

// 初始化
func InitMongoDB() {
	ctx := context.Background()
	dsn := fmt.Sprintf("mongodb://%v:27017", "mongodb")
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: dsn, Database: "SCAT", Coll: "VulDetection"})

	if err != nil {
		panic(err)
	}
	client = cli
}

func GetClient() *qmgo.QmgoClient {
	return client
}

func InsertData(data interface{}) error {
	_, err := client.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

// 通过requestId查询数据库
func QueryDataByReqId(requestId string) (interface{}, error) {
	query := bson.M{"requestid": requestId}
	result := apis.AnalyzeSourceCodeResp{}
	err := client.Find(context.TODO(), query).One(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 退出后关闭
func CloseDB() {
	if client != nil {
		if err := client.Close(context.TODO()); err != nil {
			log.Fatalf("Failed to close MongoDB connection: %v", err)
		}
	}
}
