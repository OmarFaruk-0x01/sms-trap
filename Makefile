hot:
	air

run: generate build_css
	go run .

build_css:
	npx tailwindcss -i ./assets/app.css -o ./public/app.css

generate:
	templ generate; go mod tidy