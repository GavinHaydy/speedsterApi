.PHONY: help user role gateway dev stop

RUN_DIR := .run
LOG_DIR := .logs

help:
	@echo "make dev     Start all services"
	@echo "make stop    Stop all services"

user:
	cd app/user && go run . -f etc/user-api.local.yaml

role:
	cd app/role && go run . -f etc/role-api.local.yaml

gateway:
	cd app/gateway && go run . -f etc/gateway-api.local.yaml

dev:
	@mkdir -p $(LOG_DIR) $(RUN_DIR)

	@nohup sh -c 'cd app/user && go run . -f etc/user-api.local.yaml' \
		> $(LOG_DIR)/user.log 2>&1 & echo $$! > $(RUN_DIR)/user.pid

	@nohup sh -c 'cd app/role && go run . -f etc/role-api.local.yaml' \
		> $(LOG_DIR)/role.log 2>&1 & echo $$! > $(RUN_DIR)/role.pid

	@nohup sh -c 'cd app/gateway && go run . -f etc/gateway-api.local.yaml' \
		> $(LOG_DIR)/gateway.log 2>&1 & echo $$! > $(RUN_DIR)/gateway.pid

	@echo "Services started."

stop:
	@kill $$(cat $(RUN_DIR)/user.pid 2>/dev/null) 2>/dev/null || true
	@kill $$(cat $(RUN_DIR)/role.pid 2>/dev/null) 2>/dev/null || true
	@kill $$(cat $(RUN_DIR)/gateway.pid 2>/dev/null) 2>/dev/null || true

	@rm -f $(RUN_DIR)/*.pid

	@echo "Services stopped."

logs:
	tail -f .logs/*.log
