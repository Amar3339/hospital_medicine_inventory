package business
type Company struct{
	CompanyName string `json:"company_name" binding:"required"`;
	CompanyAddress string `json:"company_address"  binding:"required"`;
	PhoneNumber string `json:"phone"`;
	CompanyEmail string `json:"company_email"`;
Registrationid string `json:"registraion_id"`;
BusinessUnit string `json:"business_unit"`;
Description string `json:"description"`;


}