package main
import (
	"fmt"
"net/http"
"github.com/gin-gonic/gin")
func main(){
	fmt.Println("Hello world")
	r := gin.New()
	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"Hello world",
		})
	})
	r.Run()
}