
build: main/main.go 
	@go build -o main main/main.go 

run:
	@go run main/main.go
.PHONY: run

clean:
	@rm -f main/main 
.PHONY: clean

foo = 
bar = $(foo)
if:
ifdef bar
	@echo yes
else
	@echo no
endif
.PHONY: if
