.PHONY: default all deps fmt vet staticcheck test cover build programs clean success

RED     := \033[31m
GREEN   := \033[32m
ORANGE  := \033[33m
BLUE    := \033[34m
MAGENTA := \033[35m
CYAN    := \033[36m
RESET   := \033[0m

define announce
	@printf "\n$(MAGENTA)==>$(RESET) $(1) $(BLUE)$(2)$(RESET)\n"
	@FIRST_JOB=0; export FIRST_JOB
endef

define ok_rainbow
	@printf "\n"
	@printf "$(RED)*$(ORANGE)*$(GREEN)*$(CYAN)*$(MAGENTA)* $(GREEN)$(1)$(RESET) $(RED)*$(ORANGE)*$(GREEN)*$(CYAN)*$(MAGENTA)*\n"
	@printf "\n"
endef

default: deps fmt vet staticcheck test build success

all: deps fmt vet staticcheck test build programs success

deps:
	$(call announce,📦,Downloading dependencies)
	go mod download

fmt:
	$(call announce,🎨,Formatting code)
	go fmt ./...

vet:
	$(call announce,🔍,Running go vet)
	go vet ./...

staticcheck:
	$(call announce,🧹,Running staticcheck)
	staticcheck ./...

test:
	$(call announce,🧪,Running tests)
	go test -coverprofile cover.out ./...

cover:
	$(call announce,📊,Generating coverage report)
	go tool cover -html=cover.out

build:
	$(call announce,🚀,Building emulator)
	go build -o 6502_emulator main.go

programs:
	$(call announce,🧩,Building example programs)
	make -C programs

clean:
	$(call announce,🧼,Cleaning build artifacts)
	rm -f 6502_emulator cover.out
	make -C programs clean

success:
	$(call ok_rainbow,BUILD SUCCESSFUL)
