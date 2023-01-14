package testutil

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kwtryo/go-sample/config"
)

func OpenDbForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	// address := "docker.for.mac.localhost"
	// port := 33306
	// if _, defined := os.LookupEnv("CI"); defined {
	// 	address = "127.0.0.1"
	// 	port = 3306
	// }

	cfg, err := config.CreateForTest()
	if err != nil {
		t.Fatalf("cannot get config: %v", err)
	}

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
