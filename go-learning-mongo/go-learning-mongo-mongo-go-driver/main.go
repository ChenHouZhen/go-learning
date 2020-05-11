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
			fmt.Printf("InsertUser err! %+v", err)
		}else {
			fmt.Println("InsertUser ok!")
		}
	}()
	createTime := time.Now().Format("2006-01-02 15:04:05")

	objUser := base_common.SysUser{
		UserId:     userId,
		Username:   username,
		Password:   password,
		Salt:       salt,
		Email:     	email,
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


func ListAllUser() (arySysUser []*base_common.SysUser, err error) {
	defer func() {
		if err!= nil {
			fmt.Printf("ListAllUser err! %+v", err)
		}else {
			fmt.Println("ListAllUser ok!")
		}
	}()
	arySysUser = make([]*base_common.SysUser,0)
	var objCtx context.Context = context.Background()
	ptrCollection := GetDBInstance().Database("my").Collection("sys_user")
	//fmt.Printf("database:%+v, collection:%+v, filter:%+v",
	//	ptrCollection.Database(), ptrCollection.Name())
	objFilter := bson.M{
	}

	ptrCursor, err :=ptrCollection.Find(objCtx, objFilter)

	if err != nil {
		return
	}
	defer ptrCursor.Close(objCtx)

	for ptrCursor.Next(objCtx) {
		var objUser base_common.SysUser
		err = ptrCursor.Decode(&objUser)
		if err != nil {
			fmt.Printf("err ：%+v\n", err)
			continue
		}
		arySysUser = append(arySysUser, &objUser)
	}

	return arySysUser, nil
}

func QueryUserByName(szName string) (objUser *base_common.SysUser, err error) {
	defer func() {
		if err!= nil {
			fmt.Printf("QueryUserByName err! %+v", err)
		}else {
			fmt.Println("QueryUserByName ok!")
		}
	}()
	// 默认值，不然当查询不到值得时候，会返回 nil
	objUser = &base_common.SysUser{}
	var objCtx context.Context = context.Background()
	ptrCollection := GetDBInstance().Database("my").Collection("sys_user")
	objFilter := bson.M{
		"username" : szName,
	}

	ptrSingleResult :=ptrCollection.FindOne(objCtx, objFilter)
	err = ptrSingleResult.Decode(objUser)

	// 当查询不存在的数据时， Decode 就会报错，mongo: no documents in result
	// 但是用 Cursor 循环（查询多条）的时候不会出现这个异常， SingleResult 才出现
	if err == mongo.ErrNoDocuments {
		err = nil
	}

	if err!=nil {
		return
	}
	return objUser, nil
}

func ListUserByStatus(nStatus int) (arySysUser []*base_common.SysUser, err error) {
	defer func() {
		if err!= nil {
			fmt.Printf("ListUserByStatus err! %+v", err)
		}else {
			fmt.Println("ListUserByStatus ok!")
		}
	}()
	arySysUser = make([]*base_common.SysUser,0)
	var objCtx context.Context = context.Background()
	ptrCollection := GetDBInstance().Database("my").Collection("sys_user")
	objFilter := bson.M{
		"status" :nStatus,
	}

	ptrCursor, err :=ptrCollection.Find(objCtx, objFilter)
	if err != nil {
		return
	}
	defer ptrCursor.Close(objCtx)

	for ptrCursor.Next(objCtx) {
		var objUser base_common.SysUser
		err = ptrCursor.Decode(&objUser)
		if err != nil {
			fmt.Printf("err ：%+v\n", err)
			continue
		}
		arySysUser = append(arySysUser, &objUser)
	}

	return arySysUser, nil
}


func PageUser(skip, limit int64) (arySysUser []*base_common.SysUser, nCount int64, err error) {
	defer func() {
		if err!= nil {
			fmt.Printf("PageUser err! %+v", err)
		}else {
			fmt.Println("PageUser ok!")
		}
	}()
	arySysUser = make([]*base_common.SysUser,0)
	var objCtx context.Context = context.Background()
	ptrCollection := GetDBInstance().Database("my").Collection("sys_user")
	objFilter := bson.M{
	}

	ptrFindOptions := options.Find()
	ptrFindOptions.SetSkip(skip)
	ptrFindOptions.SetLimit(limit)

	ptrCursor, err :=ptrCollection.Find(objCtx, objFilter, ptrFindOptions)
	if err != nil {
		return
	}
	defer ptrCursor.Close(objCtx)

	for ptrCursor.Next(objCtx) {
		var objUser base_common.SysUser
		err = ptrCursor.Decode(&objUser)
		if err != nil {
			fmt.Printf("err ：%+v\n", err)
			continue
		}
		arySysUser = append(arySysUser, &objUser)
	}

	// 查总数

	nCount, err = ptrCollection.CountDocuments(objCtx, objFilter)
	if err != nil {
		return
	}

	return arySysUser, nCount, nil
}

func main() {
	var err error

	defer func() {
		if err!= nil {
			fmt.Printf("main err! %+v",err)
		}else {
			fmt.Println("main ok!")
		}
	}()

	fmt.Println("---------------------------------- 插入 -----------------------------")
	err = InsertUser(1, "root", "123456", "123456", "843207296@qq.com", "110")
	err = InsertUser(2, "chenhz", "123456", "123456", "chenhouzhen@qq.com", "119")
	err = InsertUser(3, "admin", "admin", "123456", "admin@qq.com", "120")
	err = InsertUser(4, "xiaoming", "xiaoming", "123456", "xiaoming@qq.com", "199")


	fmt.Println("-------------------------- 全部查询 -----------------------------")
	arySysUser, err :=  ListAllUser()
	if err!= nil {
		return
	}
	for _, obj := range arySysUser{
		fmt.Println(fmt.Sprintf("%+v", obj))
	}

	fmt.Println("-------------------------- 根据名称查询 -----------------------------")

	objUser,err := QueryUserByName("测试2")
	if err!= nil {
		return
	}
	if objUser == nil {
		 fmt.Println("user is nil")
	}

	fmt.Printf("%+v\n",objUser)
	//----------------
	fmt.Println("-------------------------- 根据状态查询 -----------------------------")

	aryStatusUser, err := ListUserByStatus(1)
	if err!= nil {
		return
	}
	for _, obj := range aryStatusUser{
		fmt.Println(fmt.Sprintf("%+v", obj))
	}
	//----------------
	fmt.Println("-------------------------- 分页查询 -----------------------------")

	aryPageUser,nCount, err := PageUser(1,2)
	if err!= nil {
		return
	}
	for _, obj := range aryPageUser{
		fmt.Println(fmt.Sprintf("%+v", obj))
	}
	fmt.Printf("nCount:%+v\n",nCount)

}