package tools

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB
var testStore db.Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = db.NewStore(testDB)

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}