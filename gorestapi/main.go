// @title Customer crud api
//@version 2.0
// @description Customer stand alone server
//@contact.email sidhanta.jena@quest-global.com
// @host localhost:8081
// @BasePath /api
// @Success 200 {array} entity.Customer
// @Success 404 {array} entity.Customer
// @Success 500 {array} entity.Customer

package main

import (
	"gorestapi/api"
	"gorestapi/controller"
	"gorestapi/repository"
	"gorestapi/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/example/basic/docs"
)

var (
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository()
	customerService    service.CustomerService       = service.New(customerRepository)
	customerController controller.CustomerController = controller.New(customerService)
	customerApi        api.CustomerApi               = *api.NewCustomerAPI(customerController)
)

func main() {

	defer customerRepository.CloseDB()
	server := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	server.GET("/customers", customerApi.GetCustomers)
	server.POST("/customers", customerApi.SaveCustomers)

	server.PUT("/customers/:id", customerApi.UpdateCustomers)
	server.DELETE("/customers/:id", customerApi.DeleteCustomers)

	server.Run(":2021")

}
