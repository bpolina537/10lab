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
2. Добавить middleware для логирования в Go.
4. Создать FastAPI-сервис, который вызывает Go-сервис через HTTP.
6. Сравнить скорость ответа FastAPI и Gin под нагрузкой (wrk/ab).
4 (повышенная). Использовать WebSocket: реализовать чат на Go и подключиться к нему из Python.
6 (повышенная). Написать тесты производительности и сравнить потребление памяти.

Выполнять строго в порядке: 2 → 4 → 6 → 4(повыш) → 6(повыш).

Из методички используем:
- Go: Gin/Echo фреймворки, middleware, контексты
- Python: FastAPI, асинхронность, Pydantic
- Взаимодействие: REST, WebSocket

План должен быть таким, чтобы можно было двигаться шаг за шагом, проверяя каждый этап.

Результат: Получила план с разбивкой на микрошаги и проверками.


Промпт 2
Инструмент: DeepSeek-V3.2

Промпт:
Начни с задания 2. Инициализируй Go-модуль с Gin в папке go-gin-service. Создай базовый сервер с эндпоинтом /ping, который возвращает "pong". Проверяем.

Результат: Создала go-gin-service/, main.go, go.mod, go.sum. Установила Gin. Проверка: go run main.go, curl http://localhost:8080/ping → {"message":"pong"}

Что пришлось исправлять вручную: пришлось удалить .idea из репозитория через git rm -r --cached .idea, так как папка попала в первый коммит. Добавила .gitignore до инициализации Git.

Коммит: chore: add gitignore, init go-gin-service with ping endpoint (c98e6d3)

Итого промптов: 2


Промпт 3
Инструмент: DeepSeek-V3.2

Промпт:
Добавь middleware для логирования в папке middleware/logger.go. Middleware должен логировать метод, путь, статус и время выполнения. Подключи его глобально. Проверяем.

Результат: Создала middleware/logger.go. Подключила глобально через r.Use(middleware.Logger()). Проверка: curl http://localhost:8080/ping → в консоли сервера появился лог [GET] /ping - 200 (0s)

Что пришлось исправлять вручную: ничего

Коммит: feat: add logging middleware (f3c9b5f)

Итого промптов: 3


Промпт 4
Инструмент: DeepSeek-V3.2

Промпт:
Переходим к заданию 4 (среднее). Создай FastAPI-сервис, который вызывает Go-сервис через HTTP. В Go-сервисе (Gin) создай эндпоинт POST /process, который принимает JSON {"data": "text"} и возвращает {"status": "ok", "received": "text"}. В FastAPI создай эндпоинт /call-go, который внутри вызывает Go /process и возвращает ответ. Проверяем.

Результат: Добавила эндпоинт POST /process в Go-сервис. Создала fastapi-service/ с main.py и requirements.txt. FastAPI сервер на порту 8000, эндпоинт /call-go. Проверка через Python requests → вернулось {"status":"ok","received":"hello"}

Что пришлось исправлять вручную: пришлось установить requests в виртуальное окружение, так как модуля не было. Первый запрос через PowerShell давал 422 из-за неправильного экранирования JSON, использовала python -c для проверки.

Коммит: feat: add FastAPI service calling Go /process endpoint (e4b0f46, 1b5c53d)

Итого промптов: 4


Промпт 5
Инструмент: DeepSeek-V3.2

Промпт:
Переходим к заданию 6 (среднее). Сравни скорость ответа FastAPI и Gin под нагрузкой с помощью ab. Запусти тесты для обоих сервисов на эндпоинте /ping. Сохрани результаты.

Результат: Запустила ab для Go Gin (2135 RPS, 4.68 ms) и FastAPI (194 RPS, 51.45 ms). Gin быстрее FastAPI в 11 раз.

Что пришлось исправлять вручную: ab не был установлен — скачала Apache, распаковала, использовала C:\Apache24\bin\ab.exe

Коммит: bench: add ab load test results for Gin and FastAPI (69eab3b)

Итого промптов: 5


Промпт 6
Инструмент: DeepSeek-V3.2

Промпт:
Переходим к заданию 4 (повышенное). Реализуй WebSocket чат на Go с использованием gorilla/websocket. Сервер должен принимать сообщения от клиентов и рассылать их всем подключённым участникам. Напиши Python-клиент для подключения к чату.

Результат: Добавила WebSocket хендлер в Go-сервис. Создала python-websocket-client/ws_client.py. Проверка: два Python-клиента подключились к чату, обмениваются сообщениями.

Что пришлось исправлять вручную: установила gorilla/websocket через go get, установила websockets для Python.

Коммит: feat: add WebSocket chat on Go with Python client (292d46a)

Итого промптов: 6


Промпт 7
Инструмент: DeepSeek-V3.2

Промпт:
Переходим к заданию 6 (повышенное). Напиши тесты производительности и сравни потребление памяти между FastAPI и Gin. Измерь память (RSS) до и после нагрузки. Создай скрипт memory_test.py, сохрани результаты.

Результат: Создала benchmarks/memory_test.py. Замерила память до и после 200 запросов через ab. Go Gin: 32.16 MB, FastAPI: 49.27 MB. Утечек памяти не обнаружено.

Что пришлось исправлять вручную: установила psutil, настроила пути к серверам.

Коммиты: docs: add memory benchmark results in txt format (5d8ddc0), bench: add memory test results (3188bd8)

Итого промптов: 7


Промпт 8
Инструмент: DeepSeek-V3.2

Промпт:
Создай README.md с ФИО, группой, вариантом.

Результат: Создала README.md в корне проекта.

Коммит: docs: add README with FIO, group, variant (dc83b8a)

Итого промптов: 8


Промпт 9
Инструмент: DeepSeek-V3.2

Промпт:
Создай PROMPT_LOG.md в корне репозитория с полной историей всех промптов, результатов, проблем и коммитов.

Результат: Создала PROMPT_LOG.md в корне проекта 10lab. Включила все 8 промптов с описанием шагов, ручных правок и коммитов. Добавила итоговую статистику и полный список коммитов.

Что пришлось исправлять вручную: ничего

Коммит: docs: add PROMPT_LOG.md

Итого промптов: 9

Итоговая статистика

Всего промптов: 9
Всего коммитов: 10

Что пришлось исправлять вручную: удаление .idea, установка requests, установка ab, установка psutil, установка gorilla/websocket, установка websockets


Полный список коммитов

dc83b8a docs: add README with FIO, group, variant
3188bd8 bench: add memory test results
5d8ddc0 docs: add memory benchmark results in txt format
292d46a feat: add WebSocket chat on Go with Python client
69eab3b bench: add ab load test results for Gin and FastAPI
e4b0f46 feat: add FastAPI service calling Go /process endpoint
1b5c53d feat: add FastAPI service calling Go /process endpoint
f3c9b5f feat: add logging middleware
34e5bce chore: remove .idea from repository
c98e6d3 chore: add gitignore, init go-gin-service with ping endpoint