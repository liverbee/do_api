package ctrl

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	//"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/liverbee/do_api/model"
	"net/http"
)

const (
	host     = "192.168.0.163"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "backup"
)

func init(){
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
}

func RptGetDoDetail(c *gin.Context){
	log.Println("call GET DoDetail()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	user_code := c.Request.URL.Query().Get("user_code")

	u := new(model.User)
	u.UserCode = user_code
	err := u.RptGetDoDetail(dbc,access_token,u.UserCode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UserCode = ",user_code,u.UserCode,u.UserName)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = u
		c.JSON(http.StatusOK, rs)
	}

}