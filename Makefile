include .env

init: remove
	go get -u github.com/nfnt/resize
	go get -u github.com/fatih/color
	go get -u github.com/wcharczuk/go-chart
	go get -u gonum.org/v1/gonum/mat
	go build -o $(PROJECTNAME)
	./$(PROJECTNAME)

clean: 
	rm $(PROJECTNAME)

remove:
	rm -rf src/gonum.org
	rm -rf src/golang.org
	rm -rf src/github.com/fatih
	rm -rf src/github.com/wcharczuk
	rm -rf src/github.com/nfnt

run:
	./$(PROJECTNAME)

all: clean init