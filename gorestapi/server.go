package main

import (
	"gorestapi/controller"
	"gorestapi/repository"
	"gorestapi/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository()
	customerService    service.CustomerService       = service.New(customerRepository)
	customerController controller.CustomerController = controller.New(customerService)
)

func main() {
	server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hello go lang",
	// 	})
	// })

	server.GET("/customers", func(c *gin.Context) {
		c.JSON(200, customerController.FindAll())
	})
	server.POST("/customers", func(ctx *gin.Context) {
		err := customerController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})
	server.Run(":2021")

}
