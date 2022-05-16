package config

import (
	"clickcash_backend/logs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type SqlLogger struct {
	logger.Interface
}

var openConnectionDB *gorm.DB
var err error

func PostgresConnection() (*gorm.DB, error) {
	myDSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Vientiane",
		GetEnv("postgres.host", "localhost"),
		GetEnv("postgres.user", "postgres"),
		GetEnv("postgres.password", "123456"),
		GetEnv("postgres.database", "naga"),
		GetEnv("postgres.port", "5432"),
	)
	//dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	fmt.Println("CONNECTING_TO_POSTGRES_DB")
	openConnectionDB, err = gorm.Open(postgres.Open(myDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Vientiane")
			return time.Now().In(ti)
		},
	})
	//DryRun: false,
	if err != nil {
		logs.Error(err)
		log.Fatal("ERROR_PING_POSTGRES", err)
		return nil, err
	}
	fmt.Println("POSTGRES_CONNECTED")
	return openConnectionDB, nil
}
