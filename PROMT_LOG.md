Prompt Log
Лабораторная работа №10
Веб-разработка: FastAPI (Python) vs Gin (Go)

Студент: Бондаренко Полина Кирилловна
Группа: 221331
Вариант: 10
Дата выполнения: 03.04.2026


Промпт 1
Инструмент: DeepSeek-V3.2

Промпт:
Составь детальный пошаговый план для лабораторной работы №10. Задания:
2.Добавить middleware для логирования в Go.
4.Создать FastAPI-сервис, который вызывает Go-сервис через HTTP.
6.Сравнить скорость ответа FastAPI и Gin под нагрузкой (wrk/ab).

4 (повышенная). Использовать WebSocket: реализовать чат на Go и подключиться к нему из Python.
6 (повышенная). Написать тесты производительности и сравнить потребление памяти.

Выполнять строго в порядке: 2 → 4 → 6 → 4(повыш) → 6(повыш).

Из методички используем:
- Go: Gin/Echo фреймворки, middleware, контексты
- Python: FastAPI, асинхронность, Pydantic
- Взаимодействие: REST, WebSocket

План должен быть таким, чтобы можно было двигаться шаг за шагом, проверяя каждый этап.

Результат: Получила план с разбивкой на микрошаги и проверками.


Задание 2: Middleware для логирования в Go

Промпт 2
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. Инициализируй Go-модуль с Gin в папке go-gin-service. Создай базовый сервер с эндпоинтом /ping, который возвращает "pong". Проверяем.

Результат: Создала go-gin-service/, main.go, go.mod, go.sum. Установила Gin. Проверка: curl http://localhost:8080/ping → {"message":"pong"}

Что пришлось исправлять вручную: пришлось удалить .idea из репозитория через git rm -r --cached .idea, так как папка попала в первый коммит. Добавила .gitignore до инициализации Git.

Коммит: chore: add gitignore, init go-gin-service with ping endpoint (c98e6d3)

Промпт 3
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Добавь middleware для логирования в папке middleware/logger.go. Middleware должен логировать метод, путь, статус, время выполнения и IP клиента. Подключи глобально.

Результат: Создала middleware/logger.go с логированием. Подключила через r.Use(). Проверка: в консоли появился лог [GET] /ping - 200 (0s)

Что пришлось исправлять вручную: ничего

Коммит: feat: add logging middleware (f3c9b5f)


Задание 4 (среднее): FastAPI-сервис вызывает Go через HTTP

Промпт 4
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 1. В Go-сервисе (Gin) создай эндпоинт POST /process, который принимает JSON {"data": "text"} и возвращает {"status": "ok", "received": "text"}.

Результат: Добавила эндпоинт /process, структуры ProcessRequest и ProcessResponse. Проверка через curl.

Коммит: feat: add FastAPI service calling Go /process endpoint (e4b0f46, 1b5c53d)

Промпт 5
Инструмент: DeepSeek-V3.2

Промпт:
Шаг 2. Создай FastAPI-сервис в папке fastapi-service с эндпоинтом /call-go, который внутри вызывает Go /process и возвращает ответ.

Результат: Создала fastapi-service/main.py с httpx.AsyncClient, requirements.txt. Проверка через python -c "import requests; print(requests.post(...).json())" → {"status":"ok","received":"hello"}

Что пришлось исправлять вручную: пришлось установить requests, так как модуля не было. PowerShell не экранировал JSON — использовала python -c для проверки.

Коммит: feat: add FastAPI service calling Go /process endpoint (e4b0f46, 1b5c53d)


Задание 6 (среднее): Нагрузочное тестирование ab

Промпт 6
Инструмент: DeepSeek-V3.2

Промпт:
Установи ab и запусти нагрузочное тестирование для Gin и FastAPI на эндпоинте /ping. Сохрани результаты.

Результат: Скачала Apache, использовала C:\Apache24\bin\ab.exe. Go Gin: 14285.71 RPS, FastAPI: 7692.31 RPS. Gin быстрее в 1.86 раза.

Что пришлось исправлять вручную: ab не был установлен — скачала Apache, распаковала.

Коммит: bench: add ab load test results for Gin and FastAPI (69eab3b)


Задание 4 (повышенное): WebSocket чат

Промпт 7
Инструмент: DeepSeek-V3.2

Промпт:
Реализуй WebSocket чат на Go с использованием gorilla/websocket. Сервер должен принимать сообщения и рассылать их всем клиентам.

Результат: Добавила handleWebSocket, handleMessages, канал broadcast. Клиенты хранятся в map.

Что пришлось исправлять вручную: установила gorilla/websocket через go get.

Коммит: feat: add WebSocket chat on Go (292d46a)

Промпт 8
Инструмент: DeepSeek-V3.2

Промпт:
Напиши Python-клиент для WebSocket чата. Клиент должен подключаться, отправлять сообщения и выводить ответы.

Результат: Создала python-websocket-client/ws_client.py с websockets. Два клиента обмениваются сообщениями.

Что пришлось исправлять вручную: установила websockets для Python.

Коммит: feat: add Python client for WebSocket chat (292d46a)

Промпт 9
Инструмент: DeepSeek-V3.2

Промпт:
В WebSocket коде добавила мьютекс для защиты clients map — иначе race condition при параллельных подключениях.

Результат: Добавила mutex.Lock()/Unlock() при записи и удалении из map.

Что пришлось исправлять вручную: найдена проблема по ревью, исправлена добавлением sync.Mutex.

Коммит: fix: add mutex for WebSocket clients map

Промпт 10
Инструмент: DeepSeek-V3.2

Промпт:
Добавь graceful shutdown для WebSocket сервера — чтобы при остановке соединения закрывались корректно.

Результат: Заменила r.Run() на http.Server с обработкой сигналов SIGINT/SIGTERM.

Коммит: feat: add graceful shutdown for WebSocket server


Задание 6 (повышенное): Тесты памяти

Промпт 11
Инструмент: DeepSeek-V3.2

Промпт:
Создай скрипт для измерения потребления памяти Go Gin и FastAPI под нагрузкой. Используй psutil.

Результат: Создала benchmarks/memory_test.py. Go Gin: 32 MB, FastAPI: 49 MB.

Что пришлось исправлять вручную: установила psutil.

Коммиты: docs: add memory benchmark results in txt format (5d8ddc0), bench: add memory test results (3188bd8)


Тесты

Промпт 12
Инструмент: DeepSeek-V3.2

Промпт:
Добавь юнит-тесты для Go сервера: ping, process, invalid json, empty data, large data, middleware.

Результат: Создала handler_test.go с 6 тестами.

Проблема: тест с large data падал из-за нулевых байтов — исправила генерацию строки.

Коммит: test: add 20 unit tests (6 Go + 7 FastAPI + 7 WebSocket), all passing (6220c84)

Промпт 13
Инструмент: DeepSeek-V3.2

Промпт:
Добавь юнит-тесты для FastAPI сервера: ping, call-go success, empty data, none data, invalid json, unicode, large data.

Результат: Создала test_main.py с 7 тестами.

Коммит: test: add 20 unit tests (6 Go + 7 FastAPI + 7 WebSocket), all passing (6220c84)

Промпт 14
Инструмент: DeepSeek-V3.2

Промпт:
Добавь юнит-тесты для WebSocket клиента: connection, send/receive, multiple clients, close, invalid message, empty message, long message.

Результат: Создала test_ws.py с 7 тестами (pytest-asyncio).

Коммит: test: add 20 unit tests (6 Go + 7 FastAPI + 7 WebSocket), all passing (6220c84)


Документация и CI/CD

Промпт 15
Инструмент: DeepSeek-V3.2

Промпт:
Добавь README.md с ФИО, группой, вариантом и результатами тестов (RPS, память).

Результат: Создала README.md.

Коммит: docs: add README with FIO, group, variant (dc83b8a)

Промпт 16
Инструмент: DeepSeek-V3.2

Промпт:
Добавь CI/CD через GitHub Actions для автоматического запуска тестов Go и Python.

Результат: Создала .github/workflows/test.yml.

Коммит: ci: add GitHub Actions workflow for tests

Промпт 17
Инструмент: DeepSeek-V3.2

Промпт:
Создай PROMPT_LOG.md с полной историей всех промптов, результатов, проблем и коммитов.

Результат: Создала PROMPT_LOG.md в корне репозитория.

Коммит: docs: add PROMPT_LOG.md (8b9f156)


Итоговая статистика

Всего промптов: 17
Всего коммитов: 14

Что пришлось исправлять вручную:
- удаление .idea из репозитория
- установка requests, pytest, httpx, psutil, websockets
- установка ab (Apache Bench)
- установка gorilla/websocket
- экранирование JSON в PowerShell (замена на python -c)
- добавление мьютекса в WebSocket (race condition)
- добавление graceful shutdown
- исправление теста TestProcessEndpointLargeData


Полный список коммитов

c98e6d3 chore: add gitignore, init go-gin-service with ping endpoint
34e5bce chore: remove .idea from repository
f3c9b5f feat: add logging middleware
e4b0f46 feat: add FastAPI service calling Go /process endpoint
1b5c53d feat: add FastAPI service calling Go /process endpoint
69eab3b bench: add ab load test results for Gin and FastAPI
292d46a feat: add WebSocket chat on Go with Python client
5d8ddc0 docs: add memory benchmark results in txt format
3188bd8 bench: add memory test results
dc83b8a docs: add README with FIO, group, variant
6220c84 test: add 20 unit tests (6 Go + 7 FastAPI + 7 WebSocket), all passing
8b9f156 docs: add PROMPT_LOG.md


Выводы по результатам тестирования

Нагрузочное тестирование (ab, 10000 запросов, concurrency 100)

Go Gin: 14285.71 RPS, 7.000 ms
FastAPI: 7692.31 RPS, 13.000 ms

Вывод: Go Gin быстрее FastAPI в 1.86 раза.

Потребление памяти

Go Gin: 32 MB
FastAPI: 49 MB

Вывод: FastAPI потребляет больше памяти (49 MB против 32 MB у Go Gin).