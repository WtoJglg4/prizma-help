# Запуск

## Установка Go

Напрямую можно с сайта и просто потом запускаете exe-шник
```https://go.dev/doc/install```

## Установка make

Make - утилита для написания и запуска 'make-скриптов'.
Устанавливаем через `pacman -S make` или если есть choco через `choco install make`.

## Запуск системы

В трех разных терминалах ввести:

- ```make addr```

- ```make signal```

- ```make stat```

В каждом терминале запустится свой сервер.

## API

### Address Server

Сервер предоставляет информацию о сервисах.

#### GET /get

Получение списка сервисов.

**Ответ:**

```json
{
    "services":[
        {
            "id":"1915361",
            "key":"signal",
            "value":"MTI3LjAuMC4xOjUwMDE="
        },
        {
            "id":"4352204",
            "key":"statistics",
            "value":"MTI3LjAuMC4xOjUwMDI="
        }
    ]
}
```

**Пример запроса:**

```bash
curl -X GET http://localhost:5000/get
```

### Signal Server

Сервер предоставляет данные о сигналах.

#### GET /get

Получение данных сигнала.

**Ответ:**

```json
{
  "id": 1,
  "name": "sinus",
  "x": [1.0, 2.0, 3.0], // Их здесь будет не 3, а 20-30
  "y": [4.0, 5.0, 6.0]  // Их здесь будет не 3, а 20-30
}
```

**Пример запроса:**

```bash
curl -X GET http://localhost:5001/get
```

### Statistics Server

Сервер для расчета статистических показателей.

#### POST /statistics

Расчет статистических показателей для массива значений.

**Запрос:**

```json
{
  "values": [1.0, 2.0, 3.0, 4.0, 5.0]
}
```

**Ответ:**

```json
{
  "max": 5.0,
  "min": 1.0,
  "average": 3.0
}
```

**Пример запроса:**

```bash
curl -X POST http://localhost:5002/statistics \
  -H "Content-Type: application/json" \
  -d '{"values": [1.0, 2.0, 3.0, 4.0, 5.0]}'
```
