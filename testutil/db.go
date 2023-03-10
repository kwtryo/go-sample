package testutil

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

func OpenDbForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	cfg := CreateConfigForTest(t)
	driver := "mysql"
	db, err := sql.Open(driver, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	))
	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, driver)
}
