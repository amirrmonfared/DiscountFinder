package scrap

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
)

var testQueries *db.Queries
var testDB *sql.DB
var TestStore db.Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	TestStore = db.NewStore(testDB)

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}

