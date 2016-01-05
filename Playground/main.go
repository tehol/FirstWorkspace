// Playground project main.go
package main

import (
	"fmt"

	"crypto/sha256"
	"database/sql"
	"encoding/base64"

	"log"
	"strings"

	"github.com/Startup/Playground/FirstSubPkg"
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
	fmt.Printf("Hashed abcd:%v\n", hashString("abcd"))
	fmt.Printf("Base 64 encode abcd:%v\n", encodeBase64([]byte("abcd")))

}

func hashString(s string) (hashString string) {

	reader := strings.NewReader(s)
	size := reader.Size()
	data := make([]byte, size)
	reader.Read(data)
	hash := sha256.Sum256(data)
	fmt.Println(hash)
	hashString = encodeBase64(hash[:])
	return
}

func encodeBase64(data []byte) string {
	hashString := base64.StdEncoding.EncodeToString(data)
	return hashString
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
