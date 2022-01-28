# Azure Functions Golang
본 레파지스토리는 [**Azure Functions에 Go언어 사용 하기**](https://velog.io/@hazle/Azure-Functions%EC%97%90-Go%EC%96%B8%EC%96%B4-%EC%82%AC%EC%9A%A9-%ED%95%98%EA%B8%B0)를 참고하여서 개발하였습니다.  
Go언어의 경우에는 메인으로 지원하지 않고 있어서 사용 지정 처리기로 개발하였습니다.  

## Stack
- **golang (go version go1.17.6 darwin/amd64)**
    - [**fiber v2**](https://github.com/gofiber/fiber)

## Getted Starting
### Azure Function Project Initialization
```shell
$ func init
Select a number for worker runtime:
1. dotnet
2. dotnet (isolated process)
3. node
4. python
5. powershell
6. custom
Choose option: 6
custom
Writing .gitignore
Writing host.json
Writing local.settings.json
Writing /Users/hyun.sang/dev/GitHub/Azure-Functions-Golang/.vscode/extensions.json
```

```shell
$ func new
Select a number for template:
1. Azure Blob Storage trigger
2. Azure Cosmos DB trigger
3. Azure Event Grid trigger
4. Azure Event Hub trigger
5. HTTP trigger
6. IoT Hub (Event Hub)
7. Kafka trigger
8. Azure Queue Storage trigger
9. RabbitMQ trigger
10. SendGrid
11. Azure Service Bus Queue trigger
12. Azure Service Bus Topic trigger
13. SignalR negotiate HTTP trigger
14. Timer trigger
Choose option: 5
HTTP trigger
Function name: [HttpTrigger] hello-world
Writing /Users/hyun.sang/dev/GitHub/Azure-Functions-Golang/hello-world/function.json
The function "hello-world" was created successfully from the "HTTP trigger" template.
```

### Golang(Fiber) + Azure Function
```shell
$ go mod init github.com/Dev-HyunSang/Azure-Functions-Golang
$ go mod tidy
$ go get -u github.com/gofiber/fiber/v2
```