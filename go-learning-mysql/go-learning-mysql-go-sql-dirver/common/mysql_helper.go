package common

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type MySQLDb struct {
	DB *sql.DB
}

type MySQLTx struct {
	Tx *sql.Tx
}

func Open(szDataSourceName string) (*MySQLDb, error) {
	ptrDB, anyErr := sql.Open("mysql", szDataSourceName)
	if anyErr != nil {
		return nil, anyErr
	}

	anyErr = ptrDB.Ping()
	if anyErr != nil {
		return nil, anyErr
	}
	ptrMySQLDB := &MySQLDb{
		DB: ptrDB,
	}
	return ptrMySQLDB, nil
}

func (ptrDb *MySQLDb) SetMaxOpenConns(n int) {
	ptrDb.DB.SetMaxOpenConns(n)
}
func (ptrDb *MySQLDb) SetMaxIdleConns(n int) {
	ptrDb.DB.SetMaxIdleConns(n)
}
func (ptrDb *MySQLDb) SetConnMaxLifetime(d time.Duration) {
	ptrDb.DB.SetConnMaxLifetime(d)
}

func (ptrDb *MySQLDb) QueryRowsCallback(query string, fnRowsScan func(ptrRows *sql.Rows) (bBreak bool), args ...interface{}) error {
	ptrRows, anyErr := (ptrDb.DB).Query(query, args...)
	if ptrRows != nil {
		defer ptrRows.Close()
	}
	if anyErr != nil {
		return anyErr
	}
	for ptrRows.Next() {
		if fnRowsScan(ptrRows) {
			break
		}
	}
	return anyErr
}

func (ptrDb *MySQLDb) QueryRowsCallbackWithParam(query string, fnRowsScan func(ptrRows *sql.Rows, aryParam interface{}) (bBreak bool), aryParam interface{}, args ...interface{}) error {
	ptrRows, anyErr := (ptrDb.DB).Query(query, args...)
	if ptrRows != nil {
		defer ptrRows.Close()
	}
	if anyErr != nil {
		return anyErr
	}
	for ptrRows.Next() {
		if fnRowsScan(ptrRows, aryParam) {
			break
		}
	}
	return anyErr
}

// 不推荐使用这个接口,建议使用QueryRowsCallback接口
func (ptrDb *MySQLDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r, err := (ptrDb.DB).Query(query, args...)
	if err != nil {
	}
	return r, err
}


func (ptrDb *MySQLDb) QueryRow(query string, args ...interface{}) *sql.Row {
	return (ptrDb.DB).QueryRow(query, args...)
}

func (ptrDb *MySQLDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	r, err := (ptrDb.DB).Exec(query, args...)
	//var nRow, nInsID int64
	//var anyRowErr, anyInsErr error
	if r != nil {
		//nRow, anyRowErr = r.RowsAffected()
		//nInsID, anyInsErr = r.LastInsertId()
	}
	//utils.Logging2CloudInfof("SQL:%s, args:%v, row:%v, rowErr:%v, insID:%v, insErr:%v", query, args, nRow, anyRowErr, nInsID, anyInsErr)
	if err != nil {
		//utils.Logging2CloudWarningf("db mysql query error:%v, sql:%s, args:%v", err, query, args)
	}
	return r, err
}

func (ptrDb *MySQLDb) Begin() (*MySQLTx, error) {
	ptrTx, anyErr := ptrDb.DB.Begin()
	return &MySQLTx{
		Tx: ptrTx,
	}, anyErr
}

func (ptrTX *MySQLTx) Rollback() error {
	return ptrTX.Tx.Rollback()
}

func (ptrTX *MySQLTx) Commit() error {
	return ptrTX.Tx.Commit()
}

func (ptrTX *MySQLTx) QueryRowsCallback(query string, fnRowsScan func(ptrRows *sql.Rows) (bBreak bool), args ...interface{}) error {
	ptrRows, anyErr := (ptrTX.Tx).Query(query, args...)
	if ptrRows != nil {
		defer ptrRows.Close()
	}
	if anyErr != nil {
		return anyErr
	}
	for ptrRows.Next() {
		if fnRowsScan(ptrRows) {
			break
		}
	}
	return anyErr
}

// 不推荐使用这个接口,建议使用QueryRowsCallback接口
func (ptrTX *MySQLTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r, err := (ptrTX.Tx).Query(query, args...)
	if err != nil {
	}
	return r, err
}

func (ptrTX *MySQLTx) QueryRowsCallbackWithParam(query string, fnRowsScan func(ptrRows *sql.Rows, aryParam interface{}) (bBreak bool), aryParam interface{}, args ...interface{}) error {
	ptrRows, anyErr := (ptrTX.Tx).Query(query, args...)
	if ptrRows != nil {
		defer ptrRows.Close()
	}
	if anyErr != nil {
		//utils.Logging2CloudWarningf("db mysql query err:%v, sql:%s, args:%v", anyErr, query, args)
		return anyErr
	}
	for ptrRows.Next() {
		if fnRowsScan(ptrRows, aryParam) {
			break
		}
	}
	return anyErr
}


func (ptrTX *MySQLTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	r, err := (ptrTX.Tx).Exec(query, args...)
	//var nRow, nInsID int64
	//var anyRowErr, anyInsErr error
	if r != nil {
		//nRow, anyRowErr = r.RowsAffected()
		//nInsID, anyInsErr = r.LastInsertId()
	}
	if err != nil {
		//utils.Logging2CloudWarningf("db mysql query error:%v, sql:%s, args:%v", err, query, args)
	}
	return r, err
}

func IsDupKey(anyErr error) bool {
	if objDriverErr, bOK := anyErr.(*mysql.MySQLError); bOK {
		if objDriverErr.Number == 1062 { // just one code(more err code not handle)
			// Handle the duplicate entry
			return true
		}
	}
	return false
}
