hot:
	air

run: generate build_css
	go build -o ./tmp/main ./main.go
	export APP_ENV=dev && ./tmp/main -db-path=./sms-trap.db

build: generate build_css
	go build -o ./tmp/main ./main.go

build_css:
	npx tailwindcss -i ./assets/app.css -o ./public/css/app.css

generate:
	templ generate; go mod tidy


PHONY: hot run build_css generate