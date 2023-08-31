package main

import (
	"fmt"
	_ "gitee.com/chunanyong/dm"
	"github.com/jmoiron/sqlx"
	"time"
)

var db *sqlx.DB
var err error

func main() {

	if db, err := connectSqlx(); err != nil {
		fmt.Println(err)
		db.Ping()
		return
	}

	if err = insertTable(); err != nil {
		fmt.Println(err)
		return
	}
	if err = updateTable(); err != nil {
		fmt.Println(err)
		return
	}
	if err = queryTable(); err != nil {
		fmt.Println(err)
		return
	}
	if err = deleteTable(); err != nil {
		fmt.Println(err)
		return
	}
	if err = disconnect(); err != nil {
		fmt.Println(err)
		return
	}
}

func connectSqlx() (*sqlx.DB, error) {
	driverName := "dm"
	dataSourceName := "dm://SYSDBA:SYSDBA001@172.16.104.165:5236?logLevel=all"
	db, err = sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/* 往产品信息表插入数据 */
func insertTable() error {
	var sql = `INSERT INTO agent_list("id","sidecar_id","type","name","version","is_uninstalled","deploy_date","auto_deployment","last_update_date","auto_updated") VALUES(?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sql, "06c2f181-4195-438c-a661-46d03dc98307", "1e04d624-b831-4ae2-a028-23f06449e0c0", 1, "hdfs_datanode", "v1", 1, time.Now(), 1, time.Now(), 1)
	if err != nil {
		return err
	}
	fmt.Println("insertTable succeed")
	return nil
}

/* 修改产品信息表数据 */
func updateTable() error {
	var sql = `UPDATE agent_list SET "sidecar_id" = ? WHERE "sidecar_id" = ?;`
	if _, err := db.Exec(sql, "1e04d624-b831-4ae2-a028-23f06449e0c1", "1e04d624-b831-4ae2-a028-23f06449e0c0"); err != nil {
		return err
	}
	fmt.Println("updateTable succeed")
	return nil
}

/* 查询产品信息表 */
func queryTable() error {
	var info AgentList
	var sql = `SELECT * FROM agent_list where "sidecar_id" = ?`
	err := db.Get(&info, sql, "1e04d624-b831-4ae2-a028-23f06449e0c1")
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 15)
	fmt.Println(info)
	return nil
}

/* 删除产品信息表数据 */
func deleteTable() error {
	var sql = `DELETE FROM agent_list WHERE "sidecar_id" = ?;`
	if _, err := db.Exec(sql, "1e04d624-b831-4ae2-a028-23f06449e0c1"); err != nil {
		return err
	}
	fmt.Println("deleteTable succeed")
	return nil
}

/* 关闭数据库连接 */
func disconnect() error {
	if err := db.Close(); err != nil {
		fmt.Printf("db close failed: %s.\n", err)
		return err
	}
	fmt.Println("disconnect succeed")
	return nil
}
