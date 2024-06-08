## Go Simple Microservice
Learning konsep microservice dengan go, dan komunikasi antar service menggunakan message broker, grpc, rest, dan docker.



### Services
| Service name     | Description                |
|------------------|----------------------------|
| Auth Service     | security for services      |
| Broker Service   | service gateaway           |
| Listener Service | service for message broker |
| Logger Service   | write log service          |
| Mail Service     | service for send email     |
| Front end        | front end to test service  |

### Getting Started
1. clone repository
2. `go mod tidy` for all service
3. run `make up_build`
4. run `make build_front`
5. run `make start`


