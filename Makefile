run :
	@tailwindcss -i ./view/css/input.css -o ./static/styles.css --minify
	@templ generate
	@go run cmd/main.go