package ctrl

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"net/http"
	"github.com/liverbee/do_api/model"
	"github.com/jmoiron/sqlx"
	"github.com/liverbee/do_api/api"
)

var  dbc *sqlx.DB

func init(){
	db, err := ConnectDB("backup")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbc = db
}

func RptGetDoDetail(c *gin.Context) {
	log.Println("call GET DoDetail()")
	c.Keys = headerKeys

	date_start := c.Request.URL.Query().Get("date_start")
	date_stop := c.Request.URL.Query().Get("date_stop")
	branch := c.Request.URL.Query().Get("branch")

	r := new(model.Rpt)
	r.DocDate = date_start
	r.DocDate = date_stop
	r.Branch = branch
	rss,err := r.RptGetDoDetail(dbc, r.DocDate, r.DocDate, r.Branch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Docdate = ", date_start, date_stop)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rss
		c.JSON(http.StatusOK, rs)
	}
}

