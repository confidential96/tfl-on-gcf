package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
  "microservice/cloudbucket"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "TFLite microservice endpoints",
		})
	})

  r.POST("/queryImage", queryImage)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func queryImage(c *gin.Context) {

	var err error

  key:= cloudbucket.HandleUploadtoCloudBucket(c)
	ctx := appengine.NewContext(c.Request)

}
