# Запуск
Перед запуском необходимо указать порт сервера в .env и выполнить 
```
go mod download
```  
Для запуска необходимо выполнить 
```
go run ./cmd/main.go
```
# Swagger
Документацию для проекта можно посмотреть после запуска сервера по [ссылке](http://localhost:8080/swagger/index.html):  
http://localhost:8080/swagger/index.html

# Тесты
Для запуска теста taskManager.go ввести команду:
```
go test ./internal/services/taskManager/
```
