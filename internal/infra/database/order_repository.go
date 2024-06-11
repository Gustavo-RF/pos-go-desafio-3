package database

import (
	"github.com/Gustavo-RF/desafio-3/internal/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	err := r.Db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	var orders []entity.Order

	err := r.Db.Find(&orders).Error

	if err != nil {
		return []entity.Order{}, err
	}

	return orders, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.Select("count(*) as total").Find(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
