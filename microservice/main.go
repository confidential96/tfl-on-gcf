package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
  "microservice/cloudbucket"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "TFLite microservice endpoints",
		})
	})

  r.POST("/queryImage", queryImage)
	r.GET("/queryConfidenceInterval")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func queryImage(c *gin.Context) {

	var err error
	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return ""
	}
	defer f.close()

  key:= cloudbucket.HandleUploadtoCloudBucket(c)

	ctx := appengine.NewContext(c.Request)


}
