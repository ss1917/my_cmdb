package models

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	DbInit()
}

func DbInit() {
	read_db := "read_db"
	rdbHost := beego.AppConfig.String(read_db + "::db_host")
	rdbName := beego.AppConfig.String(read_db + "::db_name")
	rdbPort := beego.AppConfig.String(read_db + "::db_port")
	rdbUser := beego.AppConfig.String(read_db + "::db_user")
	rdbPawd := beego.AppConfig.String(read_db + "::db_pawd")
	if rdbPort == "" {
		rdbPort = "3306"
	}

	write_db := "write_db"
	wdbHost := beego.AppConfig.String(write_db + "::db_host")
	wdbName := beego.AppConfig.String(write_db + "::db_name")
	wdbPort := beego.AppConfig.String(write_db + "::db_port")
	wdbUser := beego.AppConfig.String(write_db + "::db_user")
	wdbPawd := beego.AppConfig.String(write_db + "::db_pawd")
	if wdbPort == "" {
		wdbPort = "3306"
	}
	wdbCharset := beego.AppConfig.String(write_db + "::db_charset")
	if wdbCharset == "" {
		wdbCharset = "utf8"
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	conn := wdbUser + ":" + wdbPawd + "@tcp(" + wdbHost + ":" + wdbPort + ")/" + wdbName + "?charset=" + wdbCharset + "&loc=Local"
	orm.RegisterDataBase("default", "mysql", conn)

	conn_read := rdbUser + ":" + rdbPawd + "@tcp(" + rdbHost + ":" + rdbPort + ")/" + rdbName + "?charset=" + wdbCharset + "&loc=Local"
	orm.RegisterDataBase("read", "mysql", conn_read)

	orm.RegisterModelWithPrefix("cmdb_", new(ProjectConf), new(ProjectEnvi))
	orm.RunSyncdb("default", false, false)
}
