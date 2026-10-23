package register

import (
	"errors"
	"fmt"
	"regexp" //used for phone number validation 

	"golang.org/x/crypto/bcrypt"
)
func RegisterUser(req RegisterRequest)error{
	if req.Password!=req.ConfirmPassword{
		return errors.New("password and confirm pass doesnot match ")

	}
	matched, err:= regexp.MatchString(`^\+?[0-9]{8,12}$`,req.PhoneNumber)
	if err!=nil{
		return err
	}
	if !matched{
		fmt.Printf("phone number should contains natural number:  %q\n",req.PhoneNumber)
		return errors.New("fone number should contain natural number ")
	}
	
hashedPassword, err:=bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,

	)
	if err!=nil{
		return err
	
	}
	
	
	req.Password=string(hashedPassword)

	
	
	err=CreateUser(req)
	if err!=nil{
		return err 
	
	}
	return nil 
}