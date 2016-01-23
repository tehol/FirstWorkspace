// Playground project main.go
package main

import (
	"fmt"

	"database/sql"

	"log"

	"github.com/Startup/Playground/FirstSubPkg"
	"github.com/Startup/Playground/hash"
	"github.com/Startup/Playground/rest"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {

	total := 0

	for i := 0; i <= 10; i++ {
		total = total + i

	}
	fmt.Printf("1-10 total:%v\n", total)
	fmt.Println(FirstSubPkg.Test())
	//insertTestData()
	fmt.Printf("Hashed abcd:%v\n", hash.HashString("abcd"))
	fmt.Printf("Base 64 encode abcd:%v\n", hash.EncodeBase64([]byte("abcd")))

	rest.StartRestServer(":8443")

}

func insertTestData() {
	db, err := sql.Open("pgx", "postgres://postgres:aq1sw2de@localhost:5432/postgres")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("Insert into secrets(username,salt,secret) VALUES($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("luke", "now", "test")
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Affected = %d\n", rowCnt)
}
