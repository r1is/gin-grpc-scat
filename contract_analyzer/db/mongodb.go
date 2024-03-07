package db

import (
	"context"
	"fmt"
	"pkg/apis"

	"pkg/log"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

var client *qmgo.QmgoClient

// 初始化
func InitMongoDB() {
	ctx := context.TODO()
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

func UpdateStatusByReqId(reqId string, data interface{}) error {
	ctx := context.TODO()
	filter := bson.M{"requestid": reqId}
	update := bson.M{"$set": bson.M{"resp": data}}
	err := client.Collection.UpdateOne(ctx, filter, update) // Updated code
	if err != nil {
		return err
	}

	return nil
}

func UpdateVluResByReqId(reqId string, data interface{}) error {
	ctx := context.TODO()
	filter := bson.M{"requestid": reqId}
	update := bson.M{"$set": bson.M{"result": data}}
	err := client.Collection.UpdateOne(ctx, filter, update) // Updated code
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

// 判断RequestId是否存在，存在返回true，不存在返回false
func IsExistRequestId(requestId string) bool {
	ctx := context.Background()
	query := bson.M{"requestid": requestId}

	count, err := client.Find(ctx, query).Count()
	if err != nil {
		log.LogError(fmt.Sprintf("Failed to count: %v", err))
	}

	log.LogInfo(fmt.Sprintf("RequestId: %v, count: %v", requestId, count))

	if count > 0 {
		return true
	} else {
		return false
	}
}

// 退出后关闭
func CloseDB() {
	if client != nil {
		if err := client.Close(context.TODO()); err != nil {
			log.LogError(fmt.Sprintf("Failed to close MongoDB connection: %v", err))
		}
	}
}
