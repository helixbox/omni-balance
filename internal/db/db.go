package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"omni-balance/utils/configs"
)

var (
	globalDb *gorm.DB
)

func InitDb(conf configs.Config) (err error) {
	var (
		dialector gorm.Dialector
	)
	switch conf.Db.Type {
	case configs.MySQL:
		// do something
		c := conf.Db.MySQL
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database))
	case configs.PostgreSQL:
		c := conf.Db.PostgreSQL
		dialector = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.Host, c.User, c.Password, c.Database, c.Port))
	case configs.SQLite:
		dialector = sqlite.Open(conf.Db.SQLite.Path)
	default:
		return errors.New("unsupported db type")
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	db.Logger = db.Logger.LogMode(logger.Silent)
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	globalDb = db
	return sqlDb.Ping()
}

func DB() *gorm.DB {
	if globalDb == nil {
		panic("db is nil")
	}
	return globalDb
}
