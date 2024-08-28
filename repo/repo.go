package repo

import "math/rand"

type Products struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type Order struct {
	Products
	Quantity int `json:"quantity"`
}

type Orders struct {
	OrderID string  `json:"order_id"`
	Order   []Order `json:"order"`
}

func NewProducts() map[string]Products {
	return map[string]Products{
		"PASTA": {
			Name:  "Pasta gigi",
			SKU:   "PASTA",
			Price: 10000,
			Stock: 10,
		},
		"SIKAT": {
			Name:  "Sikat gigi",
			SKU:   "SIKAT",
			Price: 10000,
			Stock: 10,
		},
		"SABUN": {
			Name:  "Sabun",
			SKU:   "SABUN",
			Price: 10000,
			Stock: 10,
		},
	}
}

var (
	AllOrders = []Orders{}
)

func GetOrders() []Orders {
	return AllOrders
}

func AddOrder(order Orders) {
	AllOrders = append(AllOrders, order)
}

type OrderRequest struct {
	Data []DataOrder `json:"data"`
}

type DataOrder struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

const letter = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func NewOrderID(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
