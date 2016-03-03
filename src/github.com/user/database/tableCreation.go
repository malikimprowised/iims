package main

import "github.com/DavidHuie/gomigrate"
import "database/sql"
import _"github.com/lib/pq"
import "fmt"

const (
    DB_USER     = "malik"
    DB_PASSWORD = ""
    DB_NAME     = "malik"
)

func main() {
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
      DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err)
  defer db.Close()
  migrator, err := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "../migrations")
  checkErr(err)
  migrator.Migrate()
  // migrator.Rollback()
}


func checkErr(err error) {
  if err != nil {
      panic(err)
  }
}
