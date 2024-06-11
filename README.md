# Desafio 3 pos Go
Esse desafio consiste em fazer três servidores em torno de duas regras de negócio.

## Servidor web
O servidor web é criado em cima do go-chi para montar as rotas. Ele será exposto na porta 8000 e poderá ser acessado do browser ou do arquivo api.http, dentro da pasta api. Para chamar a rota, é necssário instalar o plugin no vscode chamado Rest client

## Servidor grpc
O servidor gRPC 
    
## Servidor GraphQL

As regras de negócio foram desenvolvidas utilizando clean architecture, separando as em dois casos de uso:
- Listar orders: Esse caso de uso é responsável por listar as orders cadastradas.
- Criar orders: Esse caso de uso é responsável por criar uma nova order.

## Rodando o projeto
```
go mod tidy
```
```
go run main.go wire_gen.go
```

Com isso, os três servidores estarão online.

## Teste

### Exemplos

