package service

import (
	"context"

	"github.com/Gustavo-RF/desafio-3/internal/infra/grpc/pb"
	"github.com/Gustavo-RF/desafio-3/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	dto := usecase.OrderInputDTO{
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(context.Context, *pb.Empty) (*pb.ListOrderResponse, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.Order

	for _, order := range output {
		orderPb := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}

		ordersResponse = append(ordersResponse, orderPb)
	}

	return &pb.ListOrderResponse{
		Order: ordersResponse,
	}, nil
}
