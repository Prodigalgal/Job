package connect

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	datasource := os.Getenv("DATASOURCE")

	if datasource == "" {
		datasource = "root:fuckharkadmin@(192.168.100.130:31306)/Job?charset=utf8"
	}

	engineInit(datasource)
}

func engineInit(path string) {
	var err error
	var e *xorm.Engine

	e, err = xorm.NewEngine("mysql", path)
	if err != nil {
		log.Panicln("engine open error:", err)
		return
	}

	Engine = e
	log.Println("engine init over")
}
