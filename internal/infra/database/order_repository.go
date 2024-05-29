package database

import (
	"database/sql"

	"github.com/Gustavo-RF/desafio-3/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT * FROM orders")
	if err != nil {
		return []entity.Order{}, err
	}

	var orders []entity.Order

	for rows.Next() {
		var order entity.Order

		err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return []entity.Order{}, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// rows, err := db.Query("SELECT * FROM products")
// if err != nil {
// 	panic(err)
// }
// defer rows.Close()

// var products []Product

// for rows.Next() {
// 	var p Product

// 	err = rows.Scan(&p.Id, &p.Name, &p.Price)
// 	if err != nil {
// 		panic(err)
// 	}

// 	products = append(products, p)
// }

// return products, nil

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
