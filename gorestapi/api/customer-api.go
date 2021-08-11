package api

import (
	"gorestapi/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerApi struct {
	customerController controller.CustomerController
}

func NewCustomerAPI(customerController controller.CustomerController) *CustomerApi {
	return &CustomerApi{
		customerController: customerController,
	}
}

//get all customers
func (api *CustomerApi) GetCustomers(ctx *gin.Context) {
	ctx.JSON(200, api.customerController.FindAll())
}

//create a customer
func (api *CustomerApi) SaveCustomers(ctx *gin.Context) {
	err := api.customerController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Created Successfully!"})
	}

}

//update a customer
func (api *CustomerApi) UpdateCustomers(ctx *gin.Context) {
	err := api.customerController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Updated Successfully!"})
	}

}

//delete a customer
func (api *CustomerApi) DeleteCustomers(ctx *gin.Context) {
	err := api.customerController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully!"})
	}

}
