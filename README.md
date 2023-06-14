# urlShortener
Проект представляет собой сервис по укорачиванию ссылок, используя grpc-gateway и docker compose. Хранилище предоставляется на выбор при запуске программы -  Redis или PostgreSQL.

# Инструкция по запуску
```shell
#Склонировать репозиторий и перейти в рабочую директорию
https://github.com/exist03/urlShortener
cd urlShortener
#Кодогенерация
make generate-gateway 
#Выбор PostgreSQL в качестве хранилища
make psql
#Выбор Redis в качестве хранилища
make redis
#Выбор im-memory в качестве хранилища
make in-memory
```
# Запросы обрабатываемые сервисом
`POST /create body{"url": "job.ozon.ru"}`<br/>
`GET /get/{hash}`
# Примеры запросов
```shell
#POST
curl -X POST localhost:8080/create -H "Content-Type: application/json" -d '{"url": "job.ozon.ru"}'
#GET
curl -X POST localhost:8080/get/someHash'
```