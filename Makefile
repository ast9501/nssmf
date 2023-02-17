all: generate-api-doc build-app

SOURCE = cmd
INSTDIR = bin

generate-api-doc:
	swag init -g $(SOURCE)/main.go

build-app:
	go build -o $(INSTDIR)/nssmf $(SOURCE)/main.go

clean:
	rm -rf bin
	rm -rf docs