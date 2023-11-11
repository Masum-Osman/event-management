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

/*
func InitMySQL() error {
	dsn, _ := beego.AppConfig.String("sqlConn")

	var db *sql.DB
	var err error

	maxAttempts := 12 // Number of attempts (adjust as needed)
	attempt := 0

	for attempt < maxAttempts {
		// db, err = sql.Open("mysql", dsn)
		err = orm.RegisterDriver("mysql", orm.DRMySQL)
		err := orm.RegisterDataBase("default", "mysql", dsn)
		if err != nil {
			return fmt.Errorf("failed to connect to database: %v", err)
		}
		if err == nil {
			break
		}
		fmt.Printf("Attempt %d: Database connection failed. Retrying in 5 seconds...\n", attempt+1)
		time.Sleep(5 * time.Second)
		attempt++
	}

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return err
	}

	defer db.Close()
	fmt.Println("Database connection successful")

	return nil
}
*/
