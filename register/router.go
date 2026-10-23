package register
import(
	"github.com/gin-gonic/gin"
)
func RegisterRouter(router *gin.Engine){
	router.POST("/register",RegisterHandler)
}