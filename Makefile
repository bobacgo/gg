.PHONY: build
build: 
	go build

# cobra-cli add <name>
.PHONY: add
add:
	cobra-cli add $(name)
