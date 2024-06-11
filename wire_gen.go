// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Gustavo-RF/desafio-3/internal/entity"
	"github.com/Gustavo-RF/desafio-3/internal/event"
	"github.com/Gustavo-RF/desafio-3/internal/infra/database"
	"github.com/Gustavo-RF/desafio-3/internal/infra/web"
	"github.com/Gustavo-RF/desafio-3/internal/usecase"
	"github.com/Gustavo-RF/desafio-3/pkg/events"
	"github.com/google/wire"
	"gorm.io/gorm"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewListOrderUseCase(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	listOrderUseCase := usecase.NewListOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return listOrderUseCase
}

func NewWebOrderHandler(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreated)
	return webOrderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))
