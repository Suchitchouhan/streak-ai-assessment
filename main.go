package main

import (
	"assessment/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Routing Path
	r := gin.Default()
	r.POST("/find-pairs", controller.PairedNumberAPI)
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//curl -X POST http://localhost:8080/find-pairs -H "Content-Type: application/json" -d '{"numbers": [1, 2, 3, 4, 5], "target": 6}'
