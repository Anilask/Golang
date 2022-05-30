package users
import(
	"net/http"
	"github.com/gin-gonic/gin"
)
var (
	counter int
)
func GetUser(c *gin.Context){
c.String(http.StatusNotImplemented,"implement me!")
}
func CreateUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"implement me!")

}
// func SearchUser(c *gin.Context){
// 	c.String(http.StatusNotImplemented,"implement me!")

// }
