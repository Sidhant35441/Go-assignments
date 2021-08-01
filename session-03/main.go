package main

import (
	"fmt"
)

// Declare package level variable for storing map
var data map[int]Customer
var customer Customer

type datas map[int]Customer

type Customer struct {
	CustomerId int
	Name       string
	Age        int
	Mobile     int
}

// init function will be automatically invoked before main function
// init function is used to initialise package level variables
func init() {
	// Initialise map with make
	data = make(map[int]Customer)
}
func (c Customer) add(k int, v Customer) {
	// Check if key exists or not
	if _, check := data[k]; check { //if else short
		//fmt.Println("key exist :", k)
		fmt.Println(k, "Customer not added..")

	} else {
		data[k] = v
		fmt.Println(k, "Customer added..")
	}
}
func (c Customer) update(k int, v Customer) {
	if _, check := data[k]; check {
		data[k] = v
		fmt.Println(k, "Customer updated.")
	} else {
		fmt.Println("Not Updated..Please enter a valid key!")
	}

}
func (c Customer) delete(k int) {
	if _, check := data[k]; check {
		delete(data, k)
		fmt.Println(k, "Customer deleted..")
	} else {
		fmt.Println("Not Deleted...Enter a valid Key!!")
	}
}
func (c Customer) get(k int) (Customer, string) {
	if _, check := data[k]; check {
		return data[k], "success.."
	}
	return Customer{}, "failed.."
}
func (c Customer) getAll() datas {
	return data
}

func main() {
	//customer instaces
	c1 := Customer{101, "Sidhanta", 25, 7008895276}
	c2 := Customer{102, "Ranjan", 23, 8852533664}
	c3 := Customer{103, "Alexa", 24, 7908755606}
	fmt.Println(data)
	//add operation
	customer.add(101, c1)
	customer.add(102, c2)
	customer.add(103, c3)
	customer.add(104, Customer{104, "Pradeep", 29, 9090554445})
	fmt.Println(data)

	//customer.update(106, Customer{103, "Mayesa", 21, 895080745})
	//fmt.Println(data)
	//customer.delete(104)
	//fmt.Println(data)
	customers := customer.getAll()
	fmt.Println(customers)
	//a1, msg := customer.get(109)
	//fmt.Println(a1, msg)

}
