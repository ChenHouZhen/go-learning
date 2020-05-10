package main

import (
	"context"
	"fmt"
	base_common "go-learning/go-learning-common"
	"go-learning/go-learning-mongo/go-learning-mongo-mongo-go-driver/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ptrMongoClient *mongo.Client

func init() {
	ptrClient, anyErr := common.NewMongoClientByUrl(common.Mongo_URL_Prefix+"127.0.0.1")
	if anyErr != nil {
		fmt.Printf("err: %+v",anyErr)
		return
	}
	_ptrMongoClient = ptrClient
}

func GetDBInstance() *mongo.Client {
	return _ptrMongoClient
}

func InsertUser(userId int64, username,password,salt, email,mobile string) (err error){
	defer func() {
		if err != nil {
			fmt.Printf("err! %+v", err)
		}else {
			fmt.Println("ok!")
		}
	}()
	createTime := time.Now().Format("2006-01-02 15:04:05")

	objUser := base_common.SysUser{
		UserId:     userId,
		Username:   username,
		Password:   password,
		Salt:       salt,
		Email:     email,
		Mobile:     mobile,
		Status:     1,
		DeptId:     1,
		CreateTime: createTime,
		Avatar:     "",
	}

	bUpsert := true
	opts := options.ReplaceOptions{
		Upsert:                   &bUpsert,
	}

	objFilter := bson.M{
		"userId" : userId,
	}

	ptrCollection := GetDBInstance().Database("my").Collection("sys_user")

	_, err = ptrCollection.ReplaceOne(context.Background(), objFilter, objUser, &opts)
	if err!= nil {
		err = fmt.Errorf("obj:%+v err:%+v", objUser, err)
		return err
	}

	return nil
}

func main() {
	var err error

	defer func() {
		if err!= nil {
			fmt.Printf("err! %+v",err)
		}else {
			fmt.Println("ok!")
		}
	}()

	fmt.Println("---------------------------------- 插入 -----------------------------")
	err = InsertUser(1, "root", "123456", "123456", "843207296@qq.com", "110")
}