package ctrl

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/liverbee/do_api/model"
	"net/http"
)


func RptGetDoDetail(c *gin.Context){
	log.Println("call GET DoDetail()")
	c.Keys = headerKeys

	access_token := c.Request.URL.Query().Get("access_token")
	date_start := c.Request.URL.Query().Get("date_start")
	date_stop := c.Request.URL.Query().Get("date_stop")
	branch := c.Request.URL.Query().Get("branch")

	r := new(model.Rpt)
	r.DocDate = date_start
	r.DocDate = date_stop
	r.Branch = branch
	err := r.RptGetDoDetail(db,access_token,r.DocDate,r.DocDate,r.Branch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Docdate = ",date_start,date_stop)

	rs := api.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content: " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = r
		c.JSON(http.StatusOK, rs)
	}

}