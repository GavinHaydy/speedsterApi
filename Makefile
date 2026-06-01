.PHONY: help user role gateway local-dev local-stop local-logs dev \
	check-runtime compose-up compose-build compose-stop compose-logs compose-clean compose-ps \
	docker-dev docker-build docker-stop docker-logs docker-clean

RUN_DIR := .run
LOG_DIR := .logs
COMPOSE_FILE := docker-compose.dev.yml
CONTAINER_RUNTIME := $(shell if command -v podman >/dev/null 2>&1; then echo podman; elif command -v docker >/dev/null 2>&1; then echo docker; fi)
COMPOSE_DEV := $(CONTAINER_RUNTIME) compose -f $(COMPOSE_FILE)
DEV_COMMAND := $(word 2,$(MAKECMDGOALS))

help:
	@echo "make dev up       Start container development stack"
	@echo "make dev build    Build container development images"
	@echo "make dev stop     Stop container development stack"
	@echo "make dev logs     Tail container development logs"
	@echo "make dev clean    Stop stack and remove volumes"
	@echo "make dev ps       Show container status"
	@echo "make local-dev    Start local Go services"
	@echo "make local-stop   Stop local Go services"
	@echo "make local-logs   Tail local Go service logs"

user:
	cd app/user && go run . -f etc/user-api.local.yaml

role:
	cd app/role && go run . -f etc/role-api.local.yaml

gateway:
	cd app/gateway && go run . -f etc/gateway-api.local.yaml

local-dev:
	@mkdir -p $(LOG_DIR) $(RUN_DIR)
	@setsid sh -c 'cd app/user && exec go run . -f etc/user-api.local.yaml' \
		> $(LOG_DIR)/user.log 2>&1 & echo $$! > $(RUN_DIR)/user.pid
	@setsid sh -c 'cd app/role && exec go run . -f etc/role-api.local.yaml' \
		> $(LOG_DIR)/role.log 2>&1 & echo $$! > $(RUN_DIR)/role.pid
	@setsid sh -c 'cd app/gateway && exec go run . -f etc/gateway-api.local.yaml' \
		> $(LOG_DIR)/gateway.log 2>&1 & echo $$! > $(RUN_DIR)/gateway.pid
	@echo "Local services started."

local-stop:
	@for service in user role gateway; do \
		pid_file="$(RUN_DIR)/$$service.pid"; \
		if [ -s "$$pid_file" ]; then \
			pid=$$(cat "$$pid_file"); \
			kill -TERM -$$pid 2>/dev/null || kill -TERM $$pid 2>/dev/null || true; \
		fi; \
	done
	@pkill -TERM -f '[u]ser-api.local.yaml' 2>/dev/null || true
	@pkill -TERM -f '[r]ole-api.local.yaml' 2>/dev/null || true
	@pkill -TERM -f '[g]ateway-api.local.yaml' 2>/dev/null || true
	@rm -f $(RUN_DIR)/*.pid
	@echo "Local services stopped."

local-logs:
	tail -f .logs/*.log

dev:
	@case "$(DEV_COMMAND)" in \
		up|start) $(MAKE) --no-print-directory compose-up ;; \
		build) $(MAKE) --no-print-directory compose-build ;; \
		stop|down) $(MAKE) --no-print-directory compose-stop ;; \
		logs) $(MAKE) --no-print-directory compose-logs ;; \
		clean) $(MAKE) --no-print-directory compose-clean ;; \
		ps|status) $(MAKE) --no-print-directory compose-ps ;; \
		"") $(MAKE) --no-print-directory help ;; \
		help) : ;; \
		*) echo "Unknown dev command: $(DEV_COMMAND)"; $(MAKE) --no-print-directory help; exit 2 ;; \
	esac

check-runtime:
	@if [ -z "$(CONTAINER_RUNTIME)" ]; then \
		echo "Neither podman nor docker was found. Please install one of them and retry."; \
		exit 1; \
	fi
	@echo "Using $(CONTAINER_RUNTIME) compose"

compose-up: check-runtime
	$(COMPOSE_DEV) up -d --build

compose-build: check-runtime
	$(COMPOSE_DEV) build

compose-stop: check-runtime
	$(COMPOSE_DEV) down

compose-logs: check-runtime
	$(COMPOSE_DEV) logs -f

compose-clean: check-runtime
	$(COMPOSE_DEV) down -v --remove-orphans

compose-ps: check-runtime
	$(COMPOSE_DEV) ps

docker-dev: compose-up
docker-build: compose-build
docker-stop: compose-stop
docker-logs: compose-logs
docker-clean: compose-clean

ifeq ($(firstword $(MAKECMDGOALS)),dev)
EXTRA_GOALS := $(filter-out help,$(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)))
$(eval $(EXTRA_GOALS):;@:)
endif
