build:
	go build -o t-selling-service main.go
run:
	./t-selling-service
clear:
	rm t-selling-service
execute:
	go build -o t-selling-service main.go
	./t-selling-service