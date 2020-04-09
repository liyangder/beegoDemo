package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/rand"
	"net/url"
	"os"
	"time"
)

var (
	dbm         *xorm.Engine
	dbs         []*xorm.Engine
	check_count int
)

func init() {
	var err error

	//主库添加
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	m_dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	if timezone != "" {
		m_dsn = m_dsn + "&loc=" + url.QueryEscape(timezone)
	}

	dbm, err = xorm.NewEngine("mysql", m_dsn)
	dbm.SetMaxIdleConns(10)
	dbm.SetMaxOpenConns(200)
	dbm.ShowSQL(true)
	dbm.ShowExecTime(true)

	if err != nil {
		fmt.Printf("Fail to connect to master: %v", err)
		os.Exit(1)
	}

}

func GetMaster() *xorm.Engine {
	return dbm
}

func GetSlave() *xorm.Engine {
	rand.Seed(time.Now().Unix())
	rn := rand.Intn(len(dbs) - 1)
	return dbs[rn]
}

func DbCheck() {
	check_count++
	fmt.Printf("Begin->数据库检查:第%d次\n", check_count)

	if dbm != nil {

		// Raw SQL
		dbm_err := dbm.Ping()

		if dbm_err != nil {
			fmt.Println("=!!!=主库报警处理")
			fmt.Println(dbm_err)

		} else {
			fmt.Println("--主数据库查询正常！")
		}
	} else {
		fmt.Println("=!!!=主数据库连接异常！")
	}

	fmt.Println("==\n")

}

//xorm reverse mysql root:@/blog?charset=utf8 templates/goxorm entity/
