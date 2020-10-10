package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/data", func(c *gin.Context) {
		// file, err := os.Open("./data/short-data.json")
		file, err := os.Open("./data/questionConfig.json")
		if err != nil {
			fmt.Println(err)
		}
		// * 内存指针
		// & 内存地址提取符号
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)
		fmt.Println(string(byteValue))
		fmt.Println(byteValue)
		// c.JSON(200, string(byteValue))
		c.Data(200, "data", byteValue)

		// file, err := ioutil.ReadFile("./data/short-data.json")
		// if err != nil {
		// 	fmt.Println("read fail--", err)
		// }
		// fmt.Println("read--", file)
		// c.JSON(200, string(file))
		// c.Data(200, "data", file)
	})
	r.Run()
	// mux := http.NewServeMux()
	// http.HandleFunc("/data", Cors)
	// http.ListenAndServe(":8080", nil)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Method", "GET, POST, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json;charset=UTF-8")
			// w.Write([]byte("Hello, World!"))
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}

}
