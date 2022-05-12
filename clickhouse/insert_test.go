package clickhouse

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCreateTable(t *testing.T) {
	conn := initClickHouse()
	var (
		ctx = context.Background()
	)
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		logrus.Error(err)
	}
	err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS example (
			  Col1 UInt8
			, Col2 String
			, Col3 FixedString(3)
			, Col4 UUID
			, Col5 Map(String, UInt8)
			, Col6 Array(String)
			, Col7 Tuple(String, UInt8, Array(Map(String, String)))
			, Col8 DateTime
		) Engine = Null
	`)
	if err != nil {
		logrus.Error(err)
	}
}

