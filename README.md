AnswerService
Сервис для управления вопросами и ответами с REST API.

Описание
AnswerService — бэкенд-приложение на Go, реализующее хранение вопросов и ответов (модели Question и Answer), с использованием REST API и поддержкой миграций базы данных.

Основные возможности:

CRUD-операции для вопросов и ответов

Проверка существования вопроса перед добавлением ответа

Удаление вопроса с каскадным удалением ответов

Использование миграций (через goose) и ORM (GORM)

Поддержка конфигурации, логирования и разных профилей (PostgreSQL или in-memory)

Докеризация приложения

Технологический стек
Язык: Go 1.21+

База данных: PostgreSQL

ORM: GORM

Миграции: Goose

Контейнеризация: Docker, Docker Compose

Логирование: Custom logger

Тестирование: Testify

Структура проекта
text
├── cmd/
│   └── server/                 # точка входа приложения
├── internal/
│   ├── config/                # загрузка конфигурации
│   ├── dto/                   # структуры для API (запросы/ответы)
│   ├── handlers/              # HTTP-обработчики
│   ├── models/                # модели данных (Question, Answer)
│   ├── service/               # бизнес-логика
│   ├── storage/               # слой доступа к данным (PostgreSQL, in-memory)
│   └── logger/                # логирование
├── migrations/                # файлы миграций (goose)
├── tests/                     # интеграционные тесты
├── docker-compose.yml         # настройка контейнеров (БД, сервис)
├── Dockerfile                 # сборка образа приложения
└── README.md                  # документация
Быстрый запуск
Локальная разработка
Предварительные условия
Go версии 1.21+

PostgreSQL 13+

Установленный goose для миграций

bash
# Установка goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# Клонирование и настройка проекта
git clone <repository-url>
cd answerservice

# Настройка конфигурации
cp config/local.yaml.example config/local.yaml
# отредактируйте config/local.yaml при необходимости

# Применение миграций
goose -dir ./migrations postgres "user=dbuser password=dbpass dbname=answerservice sslmode=disable" up

# Запуск приложения
go run cmd/server/main.go --config=./config/local.yaml
Запуск через Docker
bash
# Запуск всего стека (приложение + БД)
docker-compose up --build

# Только для разработки
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build
