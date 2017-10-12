package do_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liverbee/do_api/ctrl"
	"gopkg.in/gin-contrib/cors.v1"
	//"github.com/gin-gonic/contrib/cors"
)

func main(){
	r := gin.New()

	r.Use(cors.Default())
	r.Run(":9000")

}