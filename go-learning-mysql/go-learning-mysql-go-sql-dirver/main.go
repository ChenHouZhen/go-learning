package main

import (
	"database/sql"
	"fmt"
	base_common "go-learning/go-learning-common"
	"go-learning/go-learning-mysql/go-learning-mysql-go-sql-dirver/common"
	"time"
)

var _ptrMysql *common.MySQLDb

func init() {

	ptrMysql, anyErr := common.Open("root:123456@tcp(127.0.0.1:3306)/robot")
	ptrMysql.SetMaxOpenConns(500)
	ptrMysql.SetMaxIdleConns(200)
	ptrMysql.SetConnMaxLifetime(time.Second * 300)
	if anyErr != nil {
		panic(anyErr)
	}

	_ptrMysql = ptrMysql

	fmt.Println("mysql init...")
}

func GetDBInstance() *common.MySQLDb {
	return _ptrMysql
}


func ListAllUser() (arySysUser []*base_common.SysUser, err error) {
	arySysUser = make([]*base_common.SysUser,0)
	err = GetDBInstance().QueryRowsCallback("select * from sys_user",
		func(ptrRows *sql.Rows) (bBreak bool) {
			var objUser base_common.SysUser
			if err := ptrRows.Scan(&objUser.UserId, &objUser.Username, &objUser.Password, &objUser.Salt,
				&objUser.Email, &objUser.Mobile, &objUser.Status, &objUser.DeptId, &objUser.CreateTime, &objUser.Avatar); err != nil {
				return
			}
			arySysUser = append(arySysUser, &objUser)
			return
		})
	return
}

func QueryUserByName(szName string) (objUser *base_common.SysUser, err error) {
	objUser = &base_common.SysUser{}
	row := GetDBInstance().QueryRow("select * from sys_user where username = ?", szName)
	err = row.Scan(&objUser.UserId, &objUser.Username, &objUser.Password, &objUser.Salt,
		&objUser.Email, &objUser.Mobile, &objUser.Status, &objUser.DeptId, &objUser.CreateTime, &objUser.Avatar)
	return
}

func ListUserByStatus(nStatus int) (arySysUser []*base_common.SysUser, err error) {
	arySysUser = make([]*base_common.SysUser, 0)
	err = GetDBInstance().QueryRowsCallback("select * from sys_user where status = ?",
		func(ptrRows *sql.Rows) (bBreak bool) {
			var objUser base_common.SysUser
			if err := ptrRows.Scan(&objUser.UserId, &objUser.Username, &objUser.Password, &objUser.Salt,
				&objUser.Email, &objUser.Mobile, &objUser.Status, &objUser.DeptId, &objUser.CreateTime, &objUser.Avatar); err != nil {
				fmt.Printf("rows Scan err: %+v, nStatus:%+v\n", err,nStatus)
				return
			}
			arySysUser = append(arySysUser, &objUser)
			return
		},
		nStatus)
	return
}



func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("err！ %+v",err)
		}else {
			fmt.Printf("ok!")
		}
	}()

	fmt.Println("-------------------------- 全部查询 -----------------------------")
	arySysUser, err :=  ListAllUser()
	if err!= nil {
		err = fmt.Errorf("method: %s err :%+v","ListAllUser", err)
		return
	}
	for _, obj := range arySysUser{
		fmt.Println(fmt.Sprintf("%+v", obj))
	}

	// -----
	fmt.Println("-------------------------- 根据名称查询 -----------------------------")

	objUser,err := QueryUserByName("测试2")
	if err!= nil {
		// Scan 当 数据库存储得值为NULL, 这样无法转成 string 或者 int 类型， Scan 就会报错
		// Scan error on column index 4, name "email": converting NULL to string is unsupported

		// 查询字段和 接收得字段长度不匹配得时候 , Scan 就会报错
		// sql: expected 10 destination arguments in Scan, not 9

		// 当查询不存在的数据时， Scan 就会报错，但是用 Rows 循环（查询多条）的时候不会出现这个异常， Row 才出现
		// no rows in result set
		if err == sql.ErrNoRows {
			fmt.Println("数据不存在")
			err = nil
		}else {
			err = fmt.Errorf("method: %s err : %+v","QueryUserByName",  err)
			return
		}
	}
	fmt.Printf("%+v\n",objUser)

	//----------------
	fmt.Println("-------------------------- 根据状态查询 -----------------------------")

	aryStatusUser, err := ListUserByStatus(1)
	if err!= nil {
		err = fmt.Errorf("method: %s err: %+v","ListUserByStatus", err)
		return
	}
	for _, obj := range aryStatusUser{
		fmt.Println(fmt.Sprintf("%+v", obj))
	}
}