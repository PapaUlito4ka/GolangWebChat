Онлайн веб-чат на Golang

Сайт доступен по [ссылке](http://limitless-meadow-49696.herokuapp.com/)

Основные использованные технологии:
- gorilla/mux - удобная маршрутизация с доп. функциями, которых нет в стандартной библиотеке
- gorilla/sessions - простая реализация веб сессий в Go
- gorilla/websockets - веб сокеты (для получения сообщений в реальном времени)
- все остальные зависимости описаны в файле <code>go.mod</code>

Использованная БД - postgresql

Бизнес логика базы данных описана в папке <code>cmd/db</code>

Реализации WebSocket и REST API находятся в папках <code>cmd/routes</code> и <code>cmd/controllers</code>

Клиентские контроллеры - <code>cmd/views</code>



