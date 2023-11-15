package main

import (
	"errors"
	"fmt"
	_"unittest/service"

	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/assert"
)

func main() {
fmt.Println("")

	// ***** การ Mock ******
	c := CustomerRepository_Mock{}
	c.On("GetCustomer", 1).Return("Russy", 39, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	//****** การใช้งาน ******
	name,age,err:=c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name,age)
}

type CustomerRepository_Mock struct {
	mock.Mock
}

func (m *CustomerRepository_Mock) GetCustomer(id int) (name string, age int, err error) {
	// Called จะ  call  ไปที่ Func GetCustomer
	args := m.Called(id)
	fmt.Println("********")
	fmt.Println(args...)
	fmt.Println("**********")

	// c.On("GetCustomer", 1).Return("Russy", 39, nil)
	// args.String(0) = russy
	//args.Int(1) = 39
	// args.Error(2) = nil
	return args.String(0), args.Int(1), args.Error(2)
}
