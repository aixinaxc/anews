package xrom_mysql

import (
	"fmt"
	"ginp/base"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DataEngine *xorm.Engine

func InitXorm()  {
	DataEngine,err := xorm.NewEngine("mysql",base.MysqlUrl)
	if err != nil {
		fmt.Println("mysql打开失败",err)
	}
	DataEngine.SetConnMaxLifetime(base.MysqlMaxLifetime)
	DataEngine.SetMaxOpenConns(base.MysqlMaxOpenConn)
	DataEngine.SetMaxIdleConns(base.MysqlMaxIdleConn)
	DataEngine.ShowSQL(true)
	err = DataEngine.Ping()
	if err != nil {
		fmt.Println("mysql没有ping成功",err)
	}
}

func InsertXORMMsg(t interface{}) bool {
	//engine := Client()
	if DataEngine == nil {
		fmt.Println("mysql连接失败")
		return false
	}

	_,err := DataEngine.Insert(t)
	if err != nil {
		fmt.Println("数据插入失败",err)
		return false
	}
	return true
}

func FindXROMMsg(sql string) []map[string][]byte  {
	//Enging := Client()
	if DataEngine == nil {
		fmt.Println("mysql连接失败")
		return nil
	}
	results,err := DataEngine.Query(sql)
	if err != nil {
		fmt.Println("查询数据失败")
		return nil
	}
	return results
}

