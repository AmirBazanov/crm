
# 🧩 CRM System — Микросервисная архитектура

## 📌 Описание

Этот проект представляет собой **CRM-систему**, построенную на **микросервисной архитектуре**. Он предназначен для демонстрации современных подходов в разработке backend-систем с применением:

- 🔧 NestJS (TypeScript)
- 🛠 Go (Golang)
- ⚡ Kafka и RabbitMQ
- 🔗 gRPC
- 🗃 PostgreSQL
- 🔁 Redis
- 🌐 GraphQL
- 🐳 Docker

Проект использует микросервисы для управления пользователями, аутентификацией и клиентами с использованием различных технологий, таких как Go для работы с базой данных и NestJS для API Gateway с GraphQL.

---

## 🏗 Архитектура

Проект состоит из следующих микросервисов:

| Сервис        | Назначение                       | Технология   |
|---------------|----------------------------------|--------------|
| `gateway`     | Точка входа, прокси + GraphQL    | NestJS       |
| `auth`        | Аутентификация и JWT             | NestJS       |
| `users`       | Профили пользователей            | Go           |
| `clients`     | Управление клиентами             | Go           |
| `db`          | PostgreSQL база данных           | PostgreSQL   |
| `redis`       | Кэш, сессии                      | Redis        |
| `kafka`       | Event-streaming и коммуникация   | Kafka + Zookeeper |

---

## 🚀 Запуск проекта

Убедись, что у тебя установлен [Docker](https://www.docker.com/) и [Docker Compose](https://docs.docker.com/compose/).

```bash
git clone https://github.com/your-username/crm-system.git
cd crm-system
docker-compose up --build
```

---

## 📦 Структура проекта

```
crm-system/
│
├── gateway/             # NestJS API Gateway + GraphQL
├── services/
│   ├── auth/            # Auth-сервис
│   ├── users/           # Users-сервис (Go)
│   └── clients/         # Клиенты (Go)
│
├── docker-compose.yml   # Контейнеризация
└── obsidian-vault/      # Документация для Obsidian
```

---

## 🧪 Тестовые пользователи

```
Email: test@example.com
Пароль: password123
```

---

## 📖 Документация

Документация хранится в формате Obsidian Vault — в папке `obsidian-vault/`.

Там содержится информация о:

- Архитектуре
- Базе данных
- Эндпоинтах
- Форматах сообщений
- Примерах использования gRPC и Kafka
- Валидации и DTO

---

## 📚 Полезные материалы

- [NestJS Docs](https://docs.nestjs.com/)
- [Go by Example](https://gobyexample.com/)
- [Kafka Fundamentals](https://www.confluent.io/learn/)
- [GraphQL Learn](https://graphql.org/learn/)
- [Docker Docs](https://docs.docker.com/)
- [gRPC Basics](https://grpc.io/docs/)

---

## 📌 Планы

- [ ] Реализация микросервиса заказов
- [ ] Админ-панель
- [ ] CI/CD pipeline
- [ ] Мониторинг (Grafana + Prometheus)
- [ ] Аутентификация через OAuth

## 🛡 License

Проект открыт для личного использования. Для коммерческого использования — свяжитесь с автором.
