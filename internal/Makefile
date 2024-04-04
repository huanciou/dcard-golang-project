# .env.prod => .env.template 

envTemplate:
	@echo "Updating .env.template"
	@awk -F '=' '/=/ {print $$1 "="}' .env.prod > .env.template