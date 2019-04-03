package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"os"
)

type config struct{
	Username string
	Password string
	DBName string
	TableName string
	Host string
	Port int
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter()
    router.HandleFunc("/getData", getData).Methods("GET")
	router.HandleFunc("/getData/{BIN}", getSomeData).Methods("GET")
	// router.HandleFunc("/getColumn/{column}", getAColumn).Methods("GET")
	router.HandleFunc("/aggregate/{op}/{column}", aggregateAColumn).Methods("GET")
	router.HandleFunc("/getRandom/{number}", getRandom).Methods("GET")
    log.Fatal(http.ListenAndServe(":8086", router))
}

type buildingData struct{
	BIN int `json:"Building Identification Number"`
	CNSTRUCT_YR int `json:"Year Constructed:"`
	LSTSTATTYPE string `json:"Last Status Type"`
	HEIGHTROOF float64 `json:"Roof Height"`
	FEAT_CODE int `json:"Feature Code"`
	GROUNDELEV int `json:"Ground Elevation"`
	SHAPE_AREA float64 `json:"Shape Area"`
	SHAPE_LEN float64 `json: "Shape Length"`
	BASE_BBL int `json:"Base BBL`	
}


var buildings []buildingData
var(
	internalID int
	BIN int
	CNSTRUCT_YR int
	LSTSTATTYPE string
	HEIGHTROOF float64
	FEAT_CODE int
	GROUNDELEV int
	SHAPE_AREA float64
	SHAPE_LEN float64
	BASE_BBL int
	)

func getData(w http.ResponseWriter, r *http.Request) {
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
	//config ends here, goes to open connection to DB
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/"+dbn)
    if err != nil {
        panic(err.Error())
    }
	defer db.Close()
	rows, err := db.Query("SELECT * FROM "+tn)
	if err != nil {
        panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&internalID, &BIN, &CONSTRUCT_YR, &LSTSTATTYPE, &HEIGHTROOF, &FEAT_CODE, &GROUNDELEV, &SHAPE_AREA, &SHAPE_LEN, &BASE_BBL)
		if err != nil {
			panic(err.Error())
		}
		buildings = append(buildings, buildingData{BIN: BIN, CNSTRUCT_YR: CNSTRUCT_YR, LSTSTATTYPE: LSTSTATTYPE, HEIGHTROOF: HEIGHTROOF, FEAT_CODE: FEAT_CODE, GROUNDELEV: GROUNDELEV, SHAPE_AREA: SHAPE_AREA, SHAPE_LEN: SHAPE_LEN, BASE_BBL: BASE_BBL})
	}
	json.NewEncoder(w).Encode(buildings)
}
func getSomeData(w http.ResponseWriter, r *http.Request) {
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
	//config ends here, goes to open connection to DB
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/"+dbn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	params := mux.Vars(r)
	rows, err := db.Query("SELECT * FROM "+tn+" WHERE BIN="+params["BIN"])
	if err != nil {
        panic(err.Error())
	}
	defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		err := rows.Scan(&internalID, &BIN, &LSTSTATTYPE, &HEIGHTROOF, &FEAT_CODE, &GROUNDELEV, &SHAPE_AREA, &SHAPE_LEN, &BASE_BBL)
		if err != nil {
			panic(err.Error())
		}
		buildings = append(buildings, buildingData{BIN: BIN, LSTSTATTYPE: LSTSTATTYPE, HEIGHTROOF: HEIGHTROOF, FEAT_CODE: FEAT_CODE, GROUNDELEV: GROUNDELEV, SHAPE_AREA: SHAPE_AREA, SHAPE_LEN: SHAPE_LEN, BASE_BBL: BASE_BBL})
	}
	json.NewEncoder(w).Encode(buildings)
}

// func getAColumn(w http.ResponseWriter, r *http.Request){
// 	config := config{}
// 	file, err := os.Open("config.json") 
// 	defer file.Close()
// 	if err != nil {
// 		panic(err)
// 	}  
// 	decoder := json.NewDecoder(file) 
// 	err = decoder.Decode(&config) 
// 	if err != nil {
// 		panic(err) }
// 		user := config.Username
// 		pass := config.Password
// 		dbn := config.DBName
// 		tn := config.TableName
// 		ip := config.Host
// 	//config ends here, goes to open connection to DB
// 	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/"+dbn)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	params := mux.Vars(r)
// 	rows, err := db.Query("SELECT "+ params["column"] +"  FROM "+tn)
// 	if err != nil {
//         panic(err.Error())
// 	}
// 	defer rows.Close()
// 	fmt.Println(params["column"])
// 	if(params["column"] == "LSTSTATTYPE"){
// 		for rows.Next() {
// 			err := rows.Scan(&LSTSTATTYPE)
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			buildings = append(buildings, buildingData{BIN: BIN, LSTSTATTYPE: LSTSTATTYPE, HEIGHTROOF: HEIGHTROOF, FEAT_CODE: FEAT_CODE, GROUNDELEV: GROUNDELEV, SHAPE_AREA: SHAPE_AREA, SHAPE_LEN: SHAPE_LEN, BASE_BBL: BASE_BBL})
// 		}
// 	}
// 	json.NewEncoder(w).Encode(buildings)
// }
type aggregateData struct{
	Column string `json:"Column"`
	Operation string `json: "Operation"`
	Value float64 `json: Value`
}
var aggregate []aggregateData

func aggregateAColumn(w http.ResponseWriter, r *http.Request){
	var Value float64
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
	//config ends here, goes to open connection to DB
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+")/"+dbn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	params := mux.Vars(r)
	rows, err := db.Query("SELECT "+params["op"]+"("+params["column"]+")  FROM "+tn)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Value)
		aggregate = append(aggregate, aggregateData{Column: params["column"], Operation: params["op"], Value: Value})
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(aggregate)
}

func getRandom(w http.ResponseWriter, r *http.Request){

}

func main() {
    handleRequests()
}