hot:
	air

run: generate build_css
	go run main.go serve

build: generate build_css
	go build -o ./tmp/main ./main.go

build_css:
	npx tailwindcss -i ./assets/app.css -o ./public/app.css

generate:
	templ generate; go mod tidy

migrate-reset:
	go run main.go migrate:reset

migrate-up:
	go run main.go migrate:up

migrate-down:
	go run main.go migrate:down

PHONY: hot run build_css generate migrate-reset migrate-up migrate-down