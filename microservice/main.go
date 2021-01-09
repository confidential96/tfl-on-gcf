package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
  "microservice/cloudbucket"
	"microservice/cloudsql"
)
type Config struct {
    Server struct {
        Port string `yaml:"port"`
        Host string `yaml:"host"`
				Url  string `yaml:"url"`
    } `yaml:"server"`
}

func main() {
  //gin server
	r := gin.Default()
  config := &Config{}
	configFile, err := os.Open("config.yml")
	// Init new YAML decode
  d := yaml.NewDecoder(configFile)

  // Start YAML decoding from file
  if err := d.Decode(&config); err != nil {
		return nil, err
  }

  if err != nil {
		log.Fatal(err)
	}
	defer config.Close()
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

  c.Request.ParseMultipartForm(10 << 20)
	var err error
	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		});
	}
	defer f.close()

	queryCF := CallCloudFunction(f, key)

  key:= cloudbucket.HandleUploadtoCloudBucket(c)

	ctx := appengine.NewContext(c.Request)

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
