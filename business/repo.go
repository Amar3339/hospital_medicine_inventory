package business

import (
"medical_inventory/db"

)



func CompanyData(x Company) error{

	query:=`insert into company(
	company_name,
	company_address,
		phone_number,
	company_email,
	registration_id,
	business_unit,
	description,

	)
	values($1,$2,$3,$4,$5,$6,$7) 
	`



_, err:=db.DB.Exec(
	query,
	x.CompanyName,
	x.CompanyAddress,
	x.PhoneNumber,
	x.CompanyEmail,
	x.Registrationid,
	x.BusinessUnit,
	x.Description,
)
if err!=nil{
		return err
	}
	return  nil

}
func CompanyExist(x Company) error{
	query:=`select exits(select 1 from company where business_id=$1 or registration_id=$2)
	`
	var exits bool
	err:=db.DB.QueryRow(query,x.BusinessUnit,x.Registrationid ).Scan(&exits)
	if err!=nil{
		return  err

	}
	return nil
}