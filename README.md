# grpc-printer

`grpc-printer` — Простой gRPC-сервис на Go, который принимает текстовые сообщения и выводит их в консоль сервера.

---

## Структура проекта

```
grpc-printer/
├── client/               # gRPC-клиент, отправляющий текст
│   └── main.go
├── server/               # gRPC-сервер, принимающий текст
│   └── main.go
├── printer.proto         # gRPC контракт (описание сервиса и сообщений)
├── printer.pb.go         # Сгенерированный Go-код из .proto (автоматически)
├── printer_grpc.pb.go    # Сгенерированный gRPC-интерфейс Go (автоматически)
├── go.mod                # Файл зависимостей Go
└── go.sum                # Контрольные суммы зависимостей
```

---

## Требования

- Go ≥ 1.20
- protoc ≥ 3.20
- Плагины: `protoc-gen-go`, `protoc-gen-go-grpc`

Установка плагинов:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

##  Генерация gRPC-кода

Если ты изменишь файл `printer.proto`, сгенерируй `.pb.go` файлы снова:

```bash
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  printer.proto
```

---

##  Как запустить

### 1. Запусти сервер

```bash
go run server/main.go
```

Ты увидишь:

```
gRPC сервер запущен на порту 50051...
```

### 2. В другом терминале — запусти клиента

```bash
go run client/main.go
```

Клиент выведет:

```
Ответ сервера: OK
```

А сервер получит и выведет сообщение:

```
Получено сообщение: Привет, сервер!
```

---

## Что происходит

- `client/main.go` подключается к `localhost:50051` и вызывает `PrintMessage` с текстом.
- `server/main.go` принимает этот запрос, выводит текст в консоль и отправляет ответ `OK`.

---
