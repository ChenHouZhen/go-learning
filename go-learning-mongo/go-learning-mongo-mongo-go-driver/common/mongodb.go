package common

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Mongo_URL_Prefix string = "mongodb://"
)

type MongoDBConfig struct {
	PoolLimit     int
	SyncTimeout   time.Duration
	SocketTimeout time.Duration
	DialTimeout   time.Duration
	Host          string
	User          string
	Password      string
	Database      string
}

func getMongoUrl(objConf MongoDBConfig) string{
	var szUrl string = Mongo_URL_Prefix + objConf.User + ":" + objConf.Password + "@" +
		objConf.Host + "/" + objConf.Database
	return szUrl
}

func NewMongoClient(objConf MongoDBConfig) (*mongo.Client, error) {

	var ptrMgoClient *mongo.Client
	var anyErr error

	szUrl := getMongoUrl(objConf)
	//_ptrLog.Infof("url: %v", szUrl)

	ptrMgoClient, anyErr = mongo.NewClient(options.Client().ApplyURI(szUrl),
		options.Client().SetMaxPoolSize(uint64(objConf.PoolLimit)),
		options.Client().SetSocketTimeout(objConf.SocketTimeout*time.Second))
	if anyErr != nil {
		//_ptrLog.Errorf("mongo.NewClient error: %v", anyErr)
		return ptrMgoClient, anyErr
	}
	objCtx, funCancel := context.WithTimeout(context.Background(), objConf.DialTimeout*time.Second)
	defer funCancel()

	anyErr = ptrMgoClient.Connect(objCtx)
	if anyErr != nil {
		//_ptrLog.Errorf("Connect error: %v", anyErr)
		return ptrMgoClient, anyErr
	}

	//_ptrLog.Infof("mongo.NewClient: %v", szUrl)
	return ptrMgoClient, anyErr
}

func NewMongoClientByUrl(szMongoUrl string) (*mongo.Client, error) {

	var ptrMgoClient *mongo.Client
	var anyErr error

	ptrMgoClient, anyErr = mongo.NewClient(options.Client().ApplyURI(szMongoUrl))
	if anyErr != nil {
		//_ptrLog.Errorf("mongo.NewClient error: %v", anyErr)
		return ptrMgoClient, anyErr
	}
	objCtx, funCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer funCancel()

	anyErr = ptrMgoClient.Connect(objCtx)
	if anyErr != nil {
		//_ptrLog.Errorf("Connect error: %v", anyErr)
		return ptrMgoClient, anyErr
	}

	//_ptrLog.Infof("mongo.NewClient: %v", szMongoUrl)
	return ptrMgoClient, anyErr
}
