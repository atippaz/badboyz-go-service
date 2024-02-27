package api
import (
	"fmt"
"net/http"
"github.com/gin-gonic/gin")

var (
	app *gin.Engine
)
func myRoute(r *gin.RouterGroup){
	r.GET("/test",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"Hello Test",
		})
	})
	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"Index",
		})
	})
}
func init(){
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
fmt.Println("init")

}
func Handler(w http.ResponseWriter,r *http.Request){
app.ServeHTTP(w,r)
fmt.Println("handler")
}
