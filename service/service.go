package service

import (
	"github.com/gin-gonic/gin"

	"fasttest/repo"
)

type service struct {
	Products map[string]repo.Products
	Orders   []repo.Orders
}

func NewService() Service {
	return service{
		Products: repo.NewProducts(),
		Orders:   []repo.Orders{},
	}
}

type Service interface {
	GetAllProduct(context *gin.Context)
	GetAllOrder(context *gin.Context)
	CreateOrder(context *gin.Context)
}

func (s service) GetAllProduct(context *gin.Context) {
	res := []repo.Products{}
	for _, product := range s.Products {
		res = append(res, product)
	}
	context.JSON(200, res)
}

func (s service) GetAllOrder(context *gin.Context) {
	res := repo.GetOrders()
	context.JSON(200, res)
}

func (s service) CreateOrder(context *gin.Context) {
	orders := repo.OrderRequest{}
	if err := context.BindJSON(&orders); err != nil {
		context.JSON(400, map[string]interface{}{
			"error": "Bad Request",
		})
	}

	orderID := repo.NewOrderID(5)
	newOrder := repo.Orders{
		OrderID: orderID,
		Order:   []repo.Order{},
	}

	//
	for _, order := range orders.Data {
		stock := s.Products[order.SKU]

		if stock.Stock-order.Qty < 0 {
			context.JSON(400, map[string]interface{}{
				"error": "Stock is not enough",
			})
			return
		}
		stock.Stock -= order.Qty
		s.Products[stock.SKU] = stock
		newOrder.Order = append(newOrder.Order, repo.Order{
			Products: stock,
			Quantity: order.Qty,
		})
	}

	repo.AddOrder(newOrder)
	context.JSON(200, map[string]interface{}{
		"success": "Order has been created",
	})

}
