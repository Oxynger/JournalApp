# JournalApp

Это сервер предоставляющий API для сервиса электронных журналов

### Запуск

---

Для запуска приложения введите следущие команды в корне прокта

```sh
$ go get
$ go get -u github.com/swaggo/swag/cmd/swag
$ swag init
$ go run ./main.go
```

Для проверки работы можно перейти по url `http://localhost:8080/swagger/index.html`

#### Команды

- `go get`: Получение зависимостей
- `go get -u github.com/swaggo/swag/cmd/swag`: Установка swag для генерации документации
- `PATH=$PATH:<Path/to/go/bin>`(optional): Добавить бинарные файлы go где храниться собранная swag утилита.
- `swag init`: Генерация swagger файлов для оображения документации
- `go run ./main.go`: Запуск сервера
- `go build ./main.go`: Компиляция бинарного файла

### Настройка

---

#### Пременные среды

- `PORT`: Для изменения порта на котором будет хоститься сервер надо поменять переменую окружения `PORT = <port nmber>`

- `MongoURI`: Если сервер базы данных расположен не по стандартному локальному пути `mongodb://localhost:27017` то его надо указать `MongoURI = <mongo path>`
