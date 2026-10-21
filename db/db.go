
package db
import(
	"database/sql"
	"fmt"
	"os"
	

	_ "github.com/jackc/pgx/v5/stdlib"

)
var DB *sql.DB
func Connect()error {
	dsn:=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
)

	database,err:=sql.Open("pgx",dsn)
	DB=database

if err!=nil{
	fmt.Print("connection with databse fail")
	return err
}
if err:=DB.Ping(); 
err!=nil{
DB.Close()
return err
}

fmt.Println("you are connected with database",DB)
return nil
	

}
