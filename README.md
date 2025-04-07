#  ToDoGo


##  Быстрый старт

###  Установка и запуск
```bash
# Клонируем репозиторий
git clone https://github.com/DropsWeb/todogo.git ./todogo

# Переходим в директорию проекта
cd todogo

# Копируем файл окружения (настройки БД и др.)
cp .env.example .env

# Переходим в директорию с Docker-конфигурацией
cd .ops

# Запускаем сервисы в фоновом режиме
docker-compose -p todogo up -d --build

# ссылка на документацию
https://dropsweb.github.io/todogo/