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
	Update(customer entity.Customer)
	Delete(customer entity.Customer)
	CloseDB()
}
type database struct {
	connection *gorm.DB
}

//database configs
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

//close the database connection
func (db *database) CloseDB() {
	mysqlDb, err := db.connection.DB()
	mysqlDb.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

//creating new customer in db
func (db *database) Save(cust entity.Customer) {
	db.connection.Create(&cust)
}

//find all customers from db
func (db *database) FindAll() []entity.Customer {
	var customers []entity.Customer
	db.connection.Set("gorm:auto_preload", true).Find(&customers)
	return customers
}

//updating the customer
func (db *database) Update(cust entity.Customer) {
	db.connection.Save(&cust)
}

//deleting the customer
func (db *database) Delete(cust entity.Customer) {
	db.connection.Delete(&cust)
}
