package ctrl

import (
	_ "github.com/lib/pq"
	"fmt"
	//"database/sql"
	//"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "192.168.0.163"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "backup"
)

func ConnectDB(dbName string)(db *sqlx.DB,err error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//db, err := sql.Open("postgres", psqlInfo)
	db, err = sqlx.Connect("postgres",psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db,nil
}

var headerKeys = make(map[string]interface{})

func setHeader(){

	headerKeys = map[string]interface{}{
		"Server":"Do_API",
		"Host":"DO:9000",
		"Content_Type":"application/json",
		"Access-Control-Allow-Origin":"*",
		"Access-Control-Allow-Methods":"GET, POST, PUT, DELETE",
		"Access-Control-Allow-Headers":"Origin, Content-Type, X-Auth-Token",
	}
}