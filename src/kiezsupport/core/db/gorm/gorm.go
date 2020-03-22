package gorm

import (
	"fmt"
	"kiezsupport/core/bootstrap"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

// initialize database
func Init() {
	var adapter string
	adapter = bootstrap.App.DBConfig.String("adapter")
	switch adapter {
	case "mysql":
		mysqlConn()
		break
	case "postgre":
		postgresConn()
		break
	default:
		panic("Undefined connection on config.yaml")
	}
}

// setupPostgresConn: setup postgres database connection using the configuration from config.yaml
func postgresConn() {
	var (
		connectionString string
		err              error
	)
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"), bootstrap.App.DBConfig.String("sslmode"))
	if db, err = gorm.Open("postgres", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.Exec("CREATE EXTENSION postgis")

	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

// mysqlConn: setup mysql database connection using the configuration from config.yaml
func mysqlConn() {
	var (
		connectionString string
		err              error
	)

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"))

	if db, err = gorm.Open("mysql", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

/*
 * Gorm: return GORM's database connection instance.
 */
func DBManager() *gorm.DB {
	var adapter string
	adapter = bootstrap.App.DBConfig.String("adapter")
	switch adapter {
	case "mysql":
		mysqlConn()
		break
	case "postgre":
		postgresConn()
		break
	default:
		panic("Undefined connection on config.yaml")
	}

	return db
}
