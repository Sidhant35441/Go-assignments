package main

import (
	"gorestapi/controller"
	"gorestapi/repository"
	"gorestapi/service"
	"net/http"

	"github.com/gin-gonic/gin"
	//swagger embedded files
)

var (
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository()
	customerService    service.CustomerService       = service.New(customerRepository)
	customerController controller.CustomerController = controller.New(customerService)
)

func main() {
	//swagger 2.0 meta information
	// docs.SwaggerInfo.Title = "Customer CRUD-API"
	// docs.SwaggerInfo.Description = "Customer crud operation api"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:5000"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http"}

	server := gin.Default()

	//url := ginSwagger.URL("http://localhost:2021/swagger/doc.json") // The url pointing to API definition
	//server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("/customers", func(c *gin.Context) {
		c.JSON(200, customerController.FindAll())
	})
	server.POST("/customers", func(ctx *gin.Context) {
		err := customerController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Created Successfully!"})
		}

	})

	server.PUT("/customers/:id", func(ctx *gin.Context) {
		err := customerController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Updated Successfully!"})
		}

	})
	server.DELETE("/customers/:id", func(ctx *gin.Context) {
		err := customerController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully!"})
		}

	})

	server.Run(":2021")

}
