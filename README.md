**Go Activity Service API** — це REST API для відстеження активності користувачів та агрегування статистики..  
Проєкт розроблено на **Golang (Gin + PostgreSQL)** 
___
##  Технологічний стек 

- **Go 1.25.1**
- **Gin** — HTTP Framework
- **PostgreSQL** — база даних
- **database/sql + pq driver** 
- **Docker / Docker Compose** 
- **Cron (robfig/cron)** 


## 🗂️ Структура проєкту

📦 ***Go_Activity_Service_API***
- 📁 `cmd/`
    - 📁 `api/`
        +  ***`main.go`*** ←  **Точка входу**
- 📁 `configs/` 
  - + ***`config.go`*** ←  **Конфігурація з .env**
- 📁 `internal/`
    - 📁 `handler/`
       + ***`event_handler.go`*** ← **HTTP Handlers**
       + ***`router`*** 
    - 📁 `model/` ← **Моделі подій та статистики**
        +  ***`event.go`*** 
        +  ***`stat.go`*** 
    - 📁 `reository/` ← **DB репозиторій**
        + ***`event_repository.go`*** 
        + ***`event_repository_interface.go`*** 
    - 📁 `scheduler/`
        + ***`sheduler.go`*** ← **Cron job**
    - 📁 `usecase/` ← **Бізнес-логіка**
        + ***`event_usecase.go`*** 
- 📁 `migrations/` ← **SQL міграції**
   + ***`001_init.sql`*** 
   + ***`002_demo_users.sql`***
- 📁 `pkg/`
  - 📁 `database/`
       + ***`postgre.go`*** ←  **Підключення до PostgreSQL**

## ⚙️ Конфігурація 
Конфігураціф відбувається за допомогою `.env` файла + `config.go`

## 🚀 Локальний запуск
 ***Клонуємо репозиторій:***
 git clone https://github.com/AlexRijikov/go-activity-service.git

 `cd go-activity-service`

***Створюємо базу PostgreSQL та застосовуємо міграції:***

**`psql -U postgres -f migrations/001_create_tables.sql`**

**`psql -U postgres -f migrations/002_seed_users.sql`**

***Запускаємо сервіс:***

`go run cmd/main.go`

## 🐳 Docker 
```yaml
version: '3.8'
services:
  db:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: activity
    ports:
      - "5432:5432"
    volumes:
      - db-data:/var/lib/postgresql/data
volumes:
  db-data:
```
## Запуск
```
docker-compose up -d
```


## 🔹 API Endpoints
| Метод   |  Route  | Опис                                                    |
|:--------|:-------:|:--------------------------------------------------------|
| POST    | /events | Створити нову подію                                     |
| GET     | /events | Отримати події з фільтром: user_id, from, to            |
| GET     | /stats  | Отримати статистику по користувачах за останні 4 години |


## Приклад події:

```json
{
  "user_id": 42,
  "action": "page_view",
  "metadata": {"page": "/home"}
}
```

## 🕒 Cron job 
`Виконується кожні 4 години`

`Агрегує кількість подій по кожному користувачу`

`Результат зберігається в таблиці aggregated_stats`


🧑‍💻 Автор: ***Олександр Рижиков***

📍 Україна

🔗 GitHub: AlexRijikov 