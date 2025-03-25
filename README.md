# orders-api


go build -o api main.go : to build the project and give name of build:api

go mod tidy : to install any dependency used in project


FOLDER STRUCTURE : 
Application : anything related to the lifecycle of the app . 

CURL COMMANDS :
[GET] curl localhost:3000/hello
[POST] curl -X POST localhost:3000/hello -v


[GET] curl localhost:3000/orders
[POST]curl -X POST localhost:3000/orders -v
[GetByID]curl localhost:3000/orders/myorder   // pretending that myorder is the ID
[UpdateByID]curl -X PUT  localhost:3000/orders/myorder
[DeleteByID]curl -X DELETE  localhost:3000/orders/myorder