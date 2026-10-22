package register
import (

	"medical_inventory/db"
)

func CreateUser(req RegisterRequest)error{
	// i already made a table in database so i need to write the querry only 

	query:= `INSERT INTO users(
		business_name,
		
		username,
		address,
		city,
		email,
		phonenumber,
		password

	)
		values($1,$2,$3,$4,$5,$6,$7)  

	
	`
	_, err:=db.DB.Exec(
		query,
		req.Business_name,
		req.Username,
		req.Address,
		req.City,
		req.Email,
		req.PhoneNumber,
		req.Password,
	)
	if err!=nil{
		return err
	}


		return nil
	


}