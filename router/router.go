package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	initilizeRoutes(r)
	r.Run(":8083")
}
