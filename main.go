package main
import(
	"fmt"
	
	
	"github.com/joho/godotenv"
	"medical_inventory/db"
)

func main(){
	err:=godotenv.Load()
	if err!=nil{
		panic("error loadin env")

	}
	// connect with pg
	err=db.Connect()
	if err!=nil{
		panic("connection failed with database")
		
	}
	fmt.Println("sucessfully connected with databse ")

}
