package database

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func InitMySQL() error {
	sqlConn, _ := beego.AppConfig.String("sqlConn")

	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", sqlConn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL")
	return nil
}
