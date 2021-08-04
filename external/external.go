package external

import "context"

type Product struct {
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Discount int    `json:"discount"`
}

func GetProductByIDFromDB(ctx context.Context, id int) {

}
