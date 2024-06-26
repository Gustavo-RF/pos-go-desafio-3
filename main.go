package main

import (
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Gustavo-RF/desafio-3/configs"
	"github.com/Gustavo-RF/desafio-3/internal/entity"
	"github.com/Gustavo-RF/desafio-3/internal/event/handler"
	"github.com/Gustavo-RF/desafio-3/internal/infra/graph"
	"github.com/Gustavo-RF/desafio-3/internal/infra/grpc/pb"
	"github.com/Gustavo-RF/desafio-3/internal/infra/grpc/service"
	"github.com/Gustavo-RF/desafio-3/internal/infra/web/webserver"
	"github.com/Gustavo-RF/desafio-3/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// carrega as configs
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// conexão com o BD
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Order{})

	// conexão com o RabbitMQ
	rabbitMQChannel := getRabbitMQChannel(configs)

	// Registra os eventos
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// cria o Use case, usando o wire
	// wire
	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db, eventDispatcher)

	// inicia o servidor web
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/", webOrderHandler.List)
	webserver.AddHandler("/order", webOrderHandler.Create)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// inicia o servidor grpc
	// https://grpc.io/docs/languages/go/quickstart/
	// protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

	// https://github.com/ktr0731/evans
	//  evans --proto internal/infra/grpc/protofiles/order.proto repl
	grpcServer := grpc.NewServer()
	orderGrpcService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderGrpcService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}

	go grpcServer.Serve(lis)

	// inicia o servidor graphQL
	//go run github.com/99designs/gqlgen generate
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GraphQLServerPort), nil)
}

func getRabbitMQChannel(configs *configs.Conf) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", configs.RabbitMqUser, configs.RabbitMqPassword, configs.RabbitMqHost, configs.RabbitMqPort))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
