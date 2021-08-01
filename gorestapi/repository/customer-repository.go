package repository

import (
	"fmt"
	"gorestapi/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var urlDB = "root:Mysql@35441@tcp(localhost:3306)/godb"

type CustomerRepository interface {
	Save(customer entity.Customer)
	FindAll() []entity.Customer
	CloseDB()
}
type database struct {
	connection *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	db, err := gorm.Open(mysql.Open(urlDB), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connection failed..")
	}
	db.AutoMigrate(&entity.Customer{})
	return &database{
		connection: db,
	}

}

// func (db *database) CloseDB() {
// 	db.connection.close()
// }
func (db *database) Save(cust entity.Customer) {
	db.connection.Create(&cust)
}
func (db *database) FindAll() []entity.Customer {
	var customers []entity.Customer
	db.connection.Find(&customers)
	return customers
}
