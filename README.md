# Touch Go Blueprint 😸

# A starter kit for Golang API project development  

#### Feature
- [x]  Company creation
- [x]  Staff creation
- [x]  Staff update
- [x]  Get staff by company ID


#### System requirements Development
- [x]  Docker
- [x]  Jaeger
- [x]  MongoDB

#### Prototype
<p>
    <a href="https://github.com/touchtechnologies-product/go-blueprint-clean-architecture">Touch Go Blueprint</a>
</p>

### Api Specification

URL: <a href="http://example.swagger-api-touch.com">example.swagger-api-touch.com</a>

### Pre-Require

Mockery
```
GO111MODULE=off go get github.com/vektra/mockery/.../
```
Swagger
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### Installation

```
git clone https://github.com/touchtechnologies-product/go-blueprint-clean-architecture
cd go-blueprint-clean-architecture
go mod download
```



### Testing 
unit testing command

```
  go test ./... -cover
```

integrating testing command

```
  go test ./... -tags integration
```


### Generate Mocks

generate mocks from interfaces for unit testing

```
  go generate ./...
```


### Local development
development in local start mongodb jaeger

```
cd development
source ./local.env
docker-compose up -d
```

### Tracing with Jaeger
please see in the example code implement jaeger wrap service ```service/company/withtracer```


### gRPC service
- Generate protobuf from staff.proto and company.proto files
```
protoc --go_out=plugins=grpc:./service/grpc/company ./service/grpc/company/protobuf/company.proto
protoc --go_out=plugins=grpc:./service/grpc/staff ./service/grpc/company/protobuf/staff.proto
```


### Others

- Uber golang style guide [link](https://github.com/uber-go/guide)
- Practical Go: Real world advice for writing maintainable Go programs [link](https://dave.cheney.net/practical-go/presentations/qcon-china.html?fbclid=IwAR2_D2Y2HXVYUNiG3LctB0kF64YKzGUatcIHm_sLYwm9SEqEKWAd76G7NAU)