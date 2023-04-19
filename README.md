Итоговая работа по курсу Skillbox "Go-разработчик"

Сборка

Используя go

go build cmd/main.go

Используя make

make build 
Сторонние Пакеты
github.com/spf13/viper - для файлов конфигурации
github.com/go-chi/chi/v5 - маршрутизатор для http
Конфигурация
Файл конфигурации config.yaml в папке config

# Настройки основной программы
server:
  addr: "localhost"
  port: "8484"


