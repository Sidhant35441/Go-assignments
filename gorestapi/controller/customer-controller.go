package controller

import (
	"gorestapi/entity"
	"gorestapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	FindAll() []entity.Customer
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
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

func (c *controller) Save(ctx *gin.Context) error {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer) //for binding and validations
	if err != nil {
		return err
	}
	c.service.Save(customer)
	return nil
}
func (c *controller) Update(ctx *gin.Context) error {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	customer.ID = id
	c.service.Update(customer)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var customer entity.Customer
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	customer.ID = id
	c.service.Delete(customer)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	customers := c.service.FindAll()
	data := gin.H{
		"title":     "Customer Page",
		"customers": customers,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
