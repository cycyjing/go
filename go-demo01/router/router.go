package router

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	// r := gin.Default()
	r.GET("/short-data", func(c *gin.Context) {
		jsonFile, err := os.Open("/data/short-data.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		fmt.Println(string(byteValue))

		c.JSON(200, string(byteValue))
		// jsonFile, err := ioutil.ReadFile("/data/short-data.json")
		// if err != nil {
		// 	fmt.Println("read fail--", err)
		// }
	})
	r.Run()
}
