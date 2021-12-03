.PHONY: go rust typescript

all: go rust typescript

go:
	@echo "Go" && cd go && go run ./sonar-sweep.go

rust:
	@echo "\nRust" && cd rust && rustc ./sonar-sweep.rs && ./sonar-sweep && rm ./sonar-sweep

typescript:
	@echo "\nTypeScript" && cd typescript && ts-node ./sonar-sweep.ts
