package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liverbee/do_api/ctrl"
	"gopkg.in/gin-contrib/cors.v1"
	//"github.com/gin-gonic/contrib/cors"
	_ "github.com/lib/pq"
)



func main(){
	r := gin.New()

	r.Use(cors.Default())
	r.GET("/doRPT",ctrl.RptGetDoDetail)
	r.Run(":9001")

}