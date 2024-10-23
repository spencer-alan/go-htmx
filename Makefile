build:
	templ generate
	tailwindcss -i ./static/main.css -o ./static/output.css
	go build  -o ./tmp/bin/main ./main.go

clean: 
	go clean
	rm -rf ./tmp/bin
