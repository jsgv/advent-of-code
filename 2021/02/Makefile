.PHONY: go rust typescript

all: go rust typescript

go:
	@echo "Go" && cd go && go run ./dive.go

rust:
	@echo "\nRust" && cd rust && rustc ./dive.rs && ./dive && rm ./dive

typescript:
	@echo "\nTypeScript" && cd typescript && ts-node ./dive.ts
