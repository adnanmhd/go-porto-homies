package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/adnanmhd/go-porto-homies/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbClient      *sql.DB
	onceConnected sync.Once
)

func dsn(cfg config.MySQL) string {
	parseTime := "?parseTime=True&loc=Asia%2FJakarta"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s%s", cfg.Username, cfg.Password, cfg.Host, cfg.Name, parseTime)
}

func NewConnection(cfg config.MySQL) (*sql.DB, error) {
	var err error
	onceConnected.Do(func() {
		dbClient, err = sql.Open("mysql", dsn(cfg))
		if err != nil {
			log.Printf("Error %s when opening DB\n", err)
			return
		}

		dbClient.SetMaxOpenConns(20)
		dbClient.SetMaxIdleConns(20)
		dbClient.SetConnMaxLifetime(time.Minute * 5)

		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		err = dbClient.PingContext(ctx)
		if err != nil {
			log.Printf("Errors %s pinging DB", err)
			return
		}
		log.Printf("Connected to DB %s successfully\n", cfg.Name)
	})

	return dbClient, nil
}
