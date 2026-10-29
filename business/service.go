package business

import (
	"errors"
)

func CompanyService(x Company) error {
	exists, err := CompanyExist(x)

	
		if err != nil {
		return err
	}
	if exists {
		return errors.New("business unit ,reg id already exist ")
	}
err=CompanyData(x)
	if err!=nil{
		return err
	}
	return err
}
