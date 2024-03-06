# .env => .env.template 

envTemplate:
	@echo "Updating .env.template"
	@awk -F '=' '/=/ {print $$1 "="}' .env > .env.template