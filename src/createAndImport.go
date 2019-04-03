package main

import (
    "fmt"
    "database/sql"
	"github.com/go-sql-driver/mysql"
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
)

type config struct{
	Username string
	Password string
	DBName string
	TableName string
	Host string
	Port int
}

func main() {
	//I wanted the following in a helper function but irritatingly go doesn't work the same way I'm used to in like python or java so unfortunately it's just gonna
	//have to be before every function.
	config := config{}
	file, err := os.Open("config.json") 
	defer file.Close()
	if err != nil {
		panic(err)
	}  
	decoder := json.NewDecoder(file) 
	err = decoder.Decode(&config) 
	if err != nil {
		panic(err) }
		user := config.Username
		pass := config.Password
		dbn := config.DBName
		tn := config.TableName
		ip := config.Host
		//config ends here
		fmt.Print("Opening connection to mySQL server...")
		db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/")
	//I'm aware this likely isn't the best way to go about it, but it works for what I need for now. And boy do I hope it's fine because it's used a lot.
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
	defer db.Close()
	fmt.Print("Creating database...")
	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS "+dbn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
	fmt.Println("Using new database...")
	_,err = db.Exec("USE "+dbn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
	fmt.Print("Creating new table...")
	_,err = db.Exec("create table "+tn+ " (internalID mediumint unsigned not null auto_increment, BIN mediumint unsigned not null, CNSTRCT_YR smallint unsigned, LSTSTATTYPE varchar(30), HEIGHTROOF decimal(8,4), FEAT_CODE smallint unsigned, GROUNDELEV smallint unsigned, SHAPE_AREA decimal(15,9), SHAPE_LEN decimal(12,8), BASE_BBL bigint unsigned, primary key (internalID)) engine = InnoDB default character set = utf8 collate = utf8_general_ci;")
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
	defer db.Close()
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("What file do you want to import?")
	filePath, _ := reader.ReadString('\n')
	fmt.Print("Importing Data...")
	filePath = strings.TrimSuffix(filePath, "\n")
	filePath, _ = filepath.Abs(filePath)
	mysql.RegisterLocalFile(filePath)
	_, err = db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE "+tn+" FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '\"' LINES TERMINATED BY '\n'  IGNORE 1 LINES (@dummy,BIN,CNSTRCT_YR,@dummy,@dummy,LSTSTATTYPE,@dummy,HEIGHTROOF,FEAT_CODE,GROUNDELEV,SHAPE_AREA,SHAPE_LEN,BASE_BBL,@dummy,@dummy)")
    if err != nil {
        panic(err.Error())
	}
	fmt.Println("Done.")
}
