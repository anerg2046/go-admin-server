package database

import (
	"errors"
	"go-app/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBTYPE uint16

const (
	_ DBTYPE = iota
	MYSQL
	MSSQL
	POSTGRES
)

func ConnDB(dsn string, dbtype DBTYPE) (*gorm.DB, error) {
	var (
		dialector gorm.Dialector
		db        *gorm.DB
		err       error
	)

	dialector = GenDialector(dsn, dbtype)

	db, err = gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Pool.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Pool.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(config.Pool.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(config.Pool.ConnMaxLifetime)
	return db, nil
}

func GenDialector(dsn string, dbtype DBTYPE) (dialector gorm.Dialector) {
	switch dbtype {
	case MYSQL:
		dialector = mysql.Open(dsn)
	case MSSQL:
		dialector = sqlserver.Open(dsn)
	case POSTGRES:
		dialector = postgres.Open(dsn)
	default:
		panic(errors.New("请配置正确的数据库类型"))
	}
	return
}
