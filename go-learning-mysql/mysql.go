package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const(
	USER_NAME = "root"
	PASS_WORD = "123456"
	NETWORK = "tcp"
	SERVER  = "127.0.0.1"
	PORT = 3306
	DATEBASE = "test"
)

type Test struct {
	id	int64	`json:"id"`
	a	int64	`json:"a"`
	b	int64	`json:"b"`
} 


var ptrDB *sql.DB

func init() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USER_NAME, PASS_WORD, NETWORK, SERVER, PORT, DATEBASE)
	fmt.Printf("connet sql:%s\n",conn)
	DB, anyErr := sql.Open("mysql", conn)
	if anyErr != nil {
		fmt.Printf("connection to mysql failed : %v\n",anyErr)
	}
	//最大连接周期，超时的连接就close
	DB.SetConnMaxLifetime(100 * time.Second)
	// 最大连接数
	DB.SetMaxOpenConns(100)
	ptrDB = DB
}

func Insert(DB *sql.DB)  {
	result, err := DB.Exec("insert into test(a,b) values(?,?)", 2,4)
	if err != nil {
		fmt.Printf("Insert data failed, err :%v\n",err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Sprintf("Get insert id failed ,err:%v\n",err)
	}
	fmt.Printf("Insert data id: %d\n",lastInsertId)

	// 影响的行数
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get rowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Printf("Affected rows:%d\n",rowsaffected)
}

func QueryById(DB *sql.DB, id int64)  {
	test := new(Test)
	row := DB.QueryRow("select id, a, b from test where id = ?",id)

	if err := row.Scan(&test.id, &test.a, &test.b); err != nil {
		fmt.Printf("scan failed,err:%v",err)
		return
	}

	fmt.Printf("single row data:%+v",*test)
}


func main() {
	//Insert(ptrDB)
	QueryById(ptrDB,1)
}