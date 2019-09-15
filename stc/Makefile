CMD_PATH := $(PWD)/cmd
PKG_PATH := $(PWD)/pkg

BUILD_DIR := $(PWD)/build
TESTS_DIR := $(PWD)/tests
TOOLS_DIR := $(PWD)/tools

STC_BINARY_NAME := stc
STVM_BINARY_NAME := stvm

COVERAGE_PROFILE_FILE := $(TESTS_DIR)/c.out
COVERAGE_REPORT_FILE := $(TESTS_DIR)/c.html

TEST_FLAGS := -mod=vendor -v -race -coverprofile='$(COVERAGE_PROFILE_FILE)' -covermode=atomic
COVER_FLAGS := -html='$(COVERAGE_PROFILE_FILE)' -o '$(COVERAGE_REPORT_FILE)'
INSTALL_FLAGS := -mod=vendor

ANTLR_VERSION := 4.7.1
ANTLR_JAR_NAME := antlr-$(ANTLR_VERSION)-complete.jar
ANTLR_JAR_PATH := $(TOOLS_DIR)/$(ANTLR_JAR_NAME)

GRAMMAR_NAME := SmallTalk

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: dep
dep:
	go mod tidy && go mod vendor && go mod verify

.PHONY: build
build: dep gen
	go build -o $(BUILD_DIR)/$(STC_BINARY_NAME) $(CMD_PATH)/stc/main.go
	go build -o $(BUILD_DIR)/$(STVM_BINARY_NAME) $(CMD_PATH)/stvm/main.go

.PHONY: gen
gen:
	java -jar $(ANTLR_JAR_PATH) -Dlanguage=Go -o $(PKG_PATH)/parser $(PWD)/$(GRAMMAR_NAME).g4
	PATH='$(TOOLS_DIR):$(PATH)' go generate ./...

.PHONY: test
test: mkdirs dep gen
	go test $(TEST_FLAGS) ./...
	GOFLAGS='-mod=vendor' go tool cover $(COVER_FLAGS)

.PHONY: clean
clean:
	rm '$(COVERAGE_PROFILE_FILE)'
	rm '$(COVERAGE_REPORT_FILE)'

.PHONY: mkdirs
mkdirs:
	mkdir -p '$(BUILD_DIR)'
	mkdir -p '$(TESTS_DIR)'
	mkdir -p '$(TOOLS_DIR)'

.PHONY: bootstrap
bootstrap: install_antlr4 install_easyjson

.PHONY: install_easyjson
install_easyjson: mkdirs
	GOBIN='$(TOOLS_DIR)' go install $(INSTALL_FLAGS) $(PWD)/vendor/github.com/mailru/easyjson/easyjson

.PHONY: install_antlr4
install_antlr4: mkdirs
	wget -O $(ANTLR_JAR_PATH) https://www.antlr.org/download/$(ANTLR_JAR_NAME)
