package clickhouse

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"time"
)

func initClickHouse() driver.Conn {
	conn, _ := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"172.16.82.176:9000"},
		Auth: clickhouse.Auth{
			Database: "test",
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	return conn
}
