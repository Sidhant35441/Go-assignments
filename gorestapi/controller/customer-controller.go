package controller

import (
	"gorestapi/entity"
	"gorestapi/service"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	FindAll() []entity.Customer
	Save(ctx *gin.Context) entity.Customer
}
type controller struct {
	service service.CustomerService
}

func New(service service.CustomerService) CustomerController {
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.Customer {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Customer {
	var customer entity.Customer
	ctx.BindJSON(&customer)
	c.service.Save(customer)
	return customer
}
