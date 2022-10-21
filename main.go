package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.engdb.com.br/apigin/infrastructure/handler"
)

func main() {
	r := gin.Default()
	handler.Routes(r)
	r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
