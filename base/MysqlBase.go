package base

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"graduation_project_socket/base/inter"
)

type MysqlBase struct {
	db_name string
	user_name string
	password string
	port string
	host string
}

type Db struct {
	DbInter *inter.DbInter
	MysqlBase MysqlBase
	Conn *sql.DB
}


func (db *Db) getCon(){
	var dbConf MysqlBase = MysqlBase{
		"mysql","root","guoyuzhao123","3306","localhost",
	}
	var twoParam = dbConf.user_name + ":" + dbConf.password + "@tcp("+dbConf.host+":"+dbConf.port+")/graduation_project?charset=utf8"
	db.Conn, _ = sql.Open(dbConf.db_name, twoParam)
}

func (db Db) Select(sql string)  {

}

// 更新操作
func (db *Db) Up(sql string,data ...interface{}) int64 {
	db.getCon()
	defer db.Conn.Close()
	var num int64
	stmt, err := db.Conn.Prepare(sql)
	if err != nil {
		fmt.Print(err)
	}
	datas := [][]interface{}{data}
	for _, val := range datas {
		res, err := stmt.Exec(val...)
		if err != nil {
			fmt.Print(err)
		}
		num, _ = res.RowsAffected()
	}
	return num
}


