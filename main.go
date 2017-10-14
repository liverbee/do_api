package do_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liverbee/do_api/ctrl"
	"gopkg.in/gin-contrib/cors.v1"
	//"github.com/gin-gonic/contrib/cors"
	"fmt"
	"database/sql"
)

const (
	host     = "192.168.0.163"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "backup"
)

func main(){

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	///////////////////////////////////
	r := gin.New()

	r.Use(cors.Default())
	r.GET("/doRPT",ctrl.RptGetDoDetail)
	r.Run(":9001")

}