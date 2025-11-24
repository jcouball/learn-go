.PHONY: help fmt vet test tidy list build-all run-all clean ci docs docs-serve

GO ?= go
BIN_DIR ?= bin
MKDOCS ?= $(shell asdf which python) -m mkdocs
MODULE := $(shell $(GO) list -m -f '{{.Path}}')
ALL_PKGS := $(shell $(GO) list ./...)
MAIN_IMPORTS := $(shell $(GO) list -f '{{if eq .Name "main"}}{{.ImportPath}}{{end}}' ./...)
MAIN_PKGS := $(patsubst $(MODULE)/%,./%,$(filter $(MODULE)/%,$(MAIN_IMPORTS)))
SAFE_MAIN_PKGS := $(foreach pkg,$(MAIN_PKGS),$(subst /,-,$(pkg:./%=%)))

help:
	@printf "Targets:\n"
	@printf "  make fmt         # gofmt every package\n"
	@printf "  make vet         # go vet ./...\n"
	@printf "  make test        # go test ./...\n"
	@printf "  make tidy        # go mod tidy\n"
	@printf "  make list        # list runnable packages\n"
	@printf "  make build-all   # build all main packages\n"
	@printf "  make run-all     # go run each main package\n"
	@printf "  make ci          # fmt, vet, test, build-all\n"
	@printf "  make docs        # build MkDocs site\n"
	@printf "  make docs-serve  # serve MkDocs with live reload\n"
	@printf "  make clean       # remove $(BIN_DIR)/ and site/\n"

fmt:
	@$(GO) fmt $(ALL_PKGS)

vet:
	@$(GO) vet ./...

test:
	@$(GO) test ./...

tidy:
	@$(GO) mod tidy

list:
	@printf "Runnable targets (target -> package):\n"
	@$(foreach pkg,$(MAIN_PKGS),printf "  %s -> %s\n" $(subst /,-,$(pkg:./%=%)) $(pkg);)

build-all: $(foreach safe,$(SAFE_MAIN_PKGS),build-$(safe))

run-all: $(foreach safe,$(SAFE_MAIN_PKGS),run-$(safe))

ci: fmt vet test build-all

docs:
	@$(MKDOCS) build

docs-serve:
	@PYTHON_GIL=1 $(MKDOCS) serve

# Template targets generated for each runnable package.
define MAIN_template
build-$(3):
	@mkdir -p $(BIN_DIR)
	@printf "Building $(1)...\n"
	@$(GO) build -o $(BIN_DIR)/$(3) $(1)

run-$(3):
	@printf "Running $(1)...\n"
	@$(GO) run $(1)
endef

$(foreach pkg,$(MAIN_PKGS),$(eval $(call MAIN_template,$(pkg),$(pkg:./%=%),$(subst /,-,$(pkg:./%=%)))))

clean:
	@rm -rf $(BIN_DIR) site
