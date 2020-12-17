package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
  "microservice/cloudbucket"
)

func main() {
  //gin server
	r := gin.Default()

  //simple test endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "TFLite microservice endpoints",
		})
	})

  //query image inference results
  r.POST("/queryImage", queryImage)

  //get confidence interval
	r.POST("/queryConfidenceInterval")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func queryImage(c *gin.Context) {

	var err error
	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		});
	}
	defer f.close()

  key:= cloudbucket.HandleUploadtoCloudBucket(c)

	ctx := appengine.NewContext(c.Request)

	queryCF := CallCloudFunction(f, key)
}

func CallCloudFunction(f File, key string) {
	var err error

	#Put in config
  url = "https://us-west2-vikcraft.cloudfunctions.net/function-tfl-1"
	req, err := http.NewRequest("POST", url, body)

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, fi.Name())

	fileContents, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	writer.Close()
	request, err := http.NewRequest("POST", url, body)

	if err != nil {
			log.Fatal(err)
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
			log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
			log.Fatal(err)
	}

	return content
}
