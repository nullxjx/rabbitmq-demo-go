# rabbitmq-demo-go

## Usage
1. run rabbitmq using docker
```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 \
-e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=rootroot \
rabbitmq:management
```
open http://localhost:15672/ in browser

2. start producer & consumer
```bash
go run producer.go
go run consumer.go
```