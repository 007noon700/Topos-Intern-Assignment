package main

import (
    "fmt"
    "database/sql"
	"github.com/go-sql-driver/mysql"
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// var(
		// id int
		// name string
		// hot int
	// )
    fmt.Println("Go MySQL Working...")
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("What file do you want to enter?")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSuffix(filePath, "\n")
	filePath, _ = filepath.Abs(filePath)
	
	fmt.Println(filePath)

	file, err := os.Open(filePath)
    if err != nil {
        panic(err.Error())
	}
	fmt.Println(file.Name(), "opened successfully")
	db, err := sql.Open("mysql", "root:administrator@tcp(localhost)/testDB")
	//catch errors
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	// filePath := "testcsv.csv"
	mysql.RegisterLocalFile(filePath)
	_, err = db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE people FIELDS TERMINATED BY ',' LINES TERMINATED BY '\r\n' IGNORE 1 LINES (id,name,@dummy,hot)")
    // perform a db.Query 
	// insert, err := db.Query("INSERT INTO people (name, hot) VALUES ('John', 69)")
	// defer insert.Close()

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
	}



}