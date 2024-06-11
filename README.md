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
Antes de tudo, vamos subir o banco de dados

### Banco de dados
Para subir o banco de dados, é necessário ter o docker instalado.
Execute o comando
```
docker compose up -d
```
E o banco estará online.

### Executando o projeto
```
go mod tidy
```
```
go run main.go wire_gen.go
```

Com isso, os três servidores estarão online.

## Teste
Para teste, teremos três servidores, cada um com a sua forma de teste:

### Servidor web:
Para o servidor web, basta acessar via postman ou arquivo .http dentro do projeto.
- Listar orders: GET http://localhost:8080/
- Criar order: POST http://localhost:8080/orders, com o corpo da requisição:
```
{
    "price": 50,
    "tax": 3
}
```

### Servidor gRCP
Para executar as chamadas via RCP, é necessário ter o [evans](https://github.com/ktr0731/evans) instalado. 
Execute o comando:
```
evans --proto internal/infra/grpc/protofiles/order.proto repl
```
Com isso, estará dentro do servidor gRCP. 
- Listar orders: call ListOrder
- Criar order: call CreateOrder. Será solicitado o preço e taxa da order.

### Servidor GraphQL
Para executar as chamadas via GraphQL, abra o navegador no endereço http://localhost:8080. 
Com isso, estará dentro do playground do servidor GraphQL.
- Listar orders: Crie a query
```
query listOrders {
  orders{
    id
    Price
    Tax
    FinalPrice
  }
}
```
> As informações de retornos podem ser customizadas.

- Criar order: Crie a mutation
```
mutation createOrder {
  createOrder(input: {
    Price: 123
    Tax: 22
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```
> O preço e taxa são editáveis. As informações de retorno podem ser customizadas.

## Exemplos
### Servidor web
- Listar order:
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/9e009b3b-be9e-4c0a-ad73-d53e9263403d)
- Criar order e listar em seguida:
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/5daaa8dd-68ee-4778-b4e3-8f63defac6ee)
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/e2fb307b-0bcf-42f9-be5e-b3f0c8046452)

### Servidor gRPC
- Listar order:
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/7fb0fbf4-a2d6-4f68-b60f-45b65ba5595c)
- Criar order e listar em seguida:

![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/10eed78e-94df-4b21-9cae-8d857379c083)
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/85bad01d-5c38-4c9f-bfc4-8e6709214ac8)

### Servidor GraphQL
- Listar order:
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/d5ed9b60-c949-4db1-aebd-07786d3639de)
- Criar order e listar em seguida:
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/7952efcf-5ce0-4519-8496-2f6570232706)
![image](https://github.com/Gustavo-RF/pos-go-desafio-3/assets/15891351/7e570293-092a-49c1-9ee4-f975a33b2e32)




