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
		ip:= config.Host
		//config ends here
		fmt.Print("Opening connection to mySQL server...")
		db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/"+dbn)
		if err != nil {
			panic(err)
		}
		fmt.Println("Done.")
		defer db.Close()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What file do you want to import?")
		filePath, _ := reader.ReadString('\n')
		filePath = strings.TrimSuffix(filePath, "\n")
		filePath, _ = filepath.Abs(filePath)
		mysql.RegisterLocalFile(filePath)
		fmt.Print("Importing Data...")
		_, err = db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE "+tn+ " FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '\"' LINES TERMINATED BY '\n'  IGNORE 1 LINES (@dummy,BIN,CNSTRCT_YR,@dummy,@dummy,LSTSTATTYPE,@dummy,HEIGHTROOF,FEAT_CODE,GROUNDELEV,SHAPE_AREA,SHAPE_LEN,BASE_BBL,@dummy,@dummy)")
		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Done.")
}
