Description:

The project has three components to it
1) The Products service  Built on NestJS(Typescript)- for managing product information like name, price, availabe quantity and descriptions.
2) The Order service Built on NestJS(Typescript) - for placing and listing orders
3) the API gateway provides a unified API for accessing both services.


Comunication between services and the API gate are made through grpc with data encoded using the protobuf encoding format.




Usage Instruction:

The project provides 3 Docker files, one for each of the services and a single `docker-compose.yml` file as a single entrypoint to build, boot and  orchestrate the services.
To get started `docker compose up --build`

the API Gateway is exposed on port `8080` while the product and order services are hidden on ports `50051` and `50052`

Once booted and running

Create an account 
`POST http://localhost:8080/register`

Login
`POST http://localhost:8080/login`

List Orders
`GET http://localhost:8080/orders`

CREATE an Order
`POST http://localhost:8080/orders`

View an Order
`POST http://localhost:8080/orders/:id`


**ADMIN ZONE**
These actions require you to use the command line of the `ecommerce-api-gateway` service to create admin users
`docker  exec -it <service-id>`
`cmd user create`
Create a product 
The create an admin user
`POST http://localhost:8080/product`

**Further Documentation comming soon**