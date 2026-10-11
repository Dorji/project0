# ===== Настройки =====
CONTAINER_NAME = my-postgres
IMAGE          = postgres
DB_USER        = postgres
DB_PASSWORD    = postgres
DB_NAME        = postgres
DB_PORT        = 5432
VOLUME_NAME    = pgdata

GREEN  = \033[0;32m
YELLOW = \033[0;33m
RED    = \033[0;31m
NC     = \033[0m

.PHONY: help init up wait down restart logs ps shell psql \
        seed-all seed-northwind seed-faker seed-sql clean clean-all

# ===== Главная команда: всё в одном =====
init: up wait seed-all ## 🚀 Поднять БД и наполнить всеми данными одной командой
	@echo ""
	@echo "$(GREEN)✅ ГОТОВО! База поднята и наполнена.$(NC)"
	@echo "$(YELLOW)   Подключение: localhost:$(DB_PORT) | user: $(DB_USER) | pass: $(DB_PASSWORD)$(NC)"
	@echo "$(YELLOW)   Зайти в psql: make psql$(NC)"

help: ## Показать список команд
	@echo "$(GREEN)Доступные команды:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-18s$(NC) %s\n", $$1, $$2}'

# ===== Управление контейнером =====
up: ## Поднять контейнер PostgreSQL
	@echo "$(GREEN)▶ [1/5] Запускаем контейнер $(CONTAINER_NAME)...$(NC)"
	@docker rm -f $(CONTAINER_NAME) >/dev/null 2>&1 || true
	@docker run --name $(CONTAINER_NAME) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_PORT):5432 \
		-v $(VOLUME_NAME):/var/lib/postgresql \
		-d $(IMAGE) >/dev/null
	@echo "$(GREEN)  ✔ Контейнер запущен$(NC)"

wait: ## Дождаться готовности PostgreSQL
	@echo "$(GREEN)▶ [2/5] Ждём готовности PostgreSQL...$(NC)"
	@until docker exec $(CONTAINER_NAME) pg_isready -U $(DB_USER) >/dev/null 2>&1; do \
		sleep 1; \
	done
	@echo "$(GREEN)  ✔ PostgreSQL готов$(NC)"

down: ## Остановить и удалить контейнер (данные сохранятся)
	@docker rm -f $(CONTAINER_NAME) >/dev/null 2>&1 || true
	@echo "$(GREEN)✔ Контейнер удалён (volume сохранён)$(NC)"

restart: down init ## Перезапустить с полным наполнением

logs: ## Показать логи контейнера
	@docker logs -f $(CONTAINER_NAME)

ps: ## Показать статус контейнера
	@docker ps -a --filter "name=$(CONTAINER_NAME)"

shell: ## Зайти в bash внутри контейнера
	@docker exec -it $(CONTAINER_NAME) bash

psql: ## Зайти в psql внутри контейнера
	@docker exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)

# ===== Наполнение данными =====
seed-all: seed-northwind seed-faker seed-sql ## Наполнить всеми датасетами

seed-northwind: ## Загрузить датасет Northwind
	@echo "$(GREEN)▶ [3/5] Загружаем Northwind...$(NC)"
	@curl -sO https://raw.githubusercontent.com/YugaByte/yugabyte-db/master/sample/northwind_ddl.sql
	@curl -sO https://raw.githubusercontent.com/YugaByte/yugabyte-db/master/sample/northwind_data.sql
	@docker cp northwind_ddl.sql $(CONTAINER_NAME):/tmp/ >/dev/null
	@docker cp northwind_data.sql $(CONTAINER_NAME):/tmp/ >/dev/null
	@docker exec -i $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -q -f /tmp/northwind_ddl.sql >/dev/null
	@docker exec -i $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -q -f /tmp/northwind_data.sql >/dev/null
	@rm -f northwind_ddl.sql northwind_data.sql
	@echo "$(GREEN)  ✔ Northwind загружен (14 таблиц)$(NC)"

seed-faker: ## Сгенерировать 1000 пользователей через Faker
	@echo "$(GREEN)▶ [4/5] Генерируем данные через Faker...$(NC)"
	@pip3 install -q --user faker psycopg2-binary
	@python3 seed_faker.py
	@echo "$(GREEN)  ✔ 1000 пользователей добавлено$(NC)"

seed-sql: ## Сгенерировать 10000 строк чистым SQL
	@echo "$(GREEN)▶ [5/5] Генерируем practice_data через SQL...$(NC)"
	@docker exec -i $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -q -c "\
		CREATE TABLE IF NOT EXISTS practice_data (id SERIAL PRIMARY KEY, random_text TEXT, random_num INT); \
		INSERT INTO practice_data (random_text, random_num) SELECT md5(random()::text), (random() * 10000)::int FROM generate_series(1, 10000);"
	@echo "$(GREEN)  ✔ 10000 строк в practice_data$(NC)"
# ===== Очистка =====
clean: ## Удалить контейнер (данные в volume остаются)
	@docker rm -f $(CONTAINER_NAME) >/dev/null 2>&1 || true
	@echo "$(GREEN)✔ Контейнер удалён$(NC)"

clean-all: clean ## Полный сброс: удалить контейнер + volume
	@docker volume rm $(VOLUME_NAME) >/dev/null 2>&1 || true
	@echo "$(RED)✔ Volume $(VOLUME_NAME) удалён — все данные стёрты$(NC)"