package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/payhon/kikgo/pkg/kik"
)

var (
	router *gin.Engine
)
func init() {
	router = gin.Default()
}

func pingHandler(c *kik.Context) {
	fmt.Println(c.GetQuery("name"))
	c.JsonResult(200, kik.H{"message": "pong","name":c.GetQuery("name")})
}



func bootstrap(){
	kik.Get(router,"/ping",ping$Handler)
	router.Run()
}

func main() {
	bootstrap()
}
