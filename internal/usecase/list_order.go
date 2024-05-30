package usecase

import (
	"github.com/Gustavo-RF/desafio-3/internal/entity"
	"github.com/Gustavo-RF/desafio-3/pkg/events"
)

type OrderListOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderListOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	var ordersResponse []OrderListOutputDTO

	if err != nil {
		return ordersResponse, err
	}

	for _, order := range orders {
		orderListOutput := OrderListOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		order.CalculateFinalPrice()

		ordersResponse = append(ordersResponse, orderListOutput)
	}

	return ordersResponse, nil
}
