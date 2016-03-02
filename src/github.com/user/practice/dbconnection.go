package main

import (
    "database/sql"
    "fmt"
    _"github.com/lib/pq"
)

const (
    DB_USER     = "malik"
    DB_PASSWORD = ""
    DB_NAME     = "malik"
)
func main() {
    var choice string
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    fmt.Println("\n1. insert\n2. select\n3. update\n4. alter\n5.delete ")
    fmt.Print("enter your choice : ")
    fmt.Scan(&choice)
    switch choice {
      case "1":
      fmt.Println("# Inserting values")
      stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3)");
      res,err := stmt.Exec(choice, "it", "2016-02-25")
      if res == nil {
        fmt.Println(stmt)
      }
      checkErr(err)
      fmt.Println("# Row inserted")
      break
    }
}
func checkErr(err error) {
  if err != nil {
      panic(err)
  }
}
