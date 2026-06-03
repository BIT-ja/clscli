BINARY := clscli
VERSION ?= dev
DIST_DIR := dist
SKILL_DIR := SKILL
SKILL_BIN_DIR := $(SKILL_DIR)/bin

GO := go
MKDIR_P := mkdir -p
RM_RF := rm -rf
CP := cp

LOCAL_BINARY := $(BINARY)$(if $(filter Windows_NT,$(OS)),.exe,)
LDFLAGS := -s -w -X main.version=$(VERSION)

DIST_TARGETS := \
	$(DIST_DIR)/$(BINARY)-linux-amd64 \
	$(DIST_DIR)/$(BINARY)-linux-arm64 \
	$(DIST_DIR)/$(BINARY)-darwin-amd64 \
	$(DIST_DIR)/$(BINARY)-darwin-arm64 \
	$(DIST_DIR)/$(BINARY)-windows-amd64.exe \
	$(DIST_DIR)/$(BINARY)-windows-arm64.exe

.PHONY: build run dist skill tidy clean

build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(LOCAL_BINARY) .

run: build
	./$(LOCAL_BINARY) $(ARGS)

dist: $(DIST_TARGETS)

$(DIST_DIR):
	$(MKDIR_P) $(DIST_DIR)

$(SKILL_BIN_DIR):
	$(MKDIR_P) $(SKILL_BIN_DIR)

$(DIST_DIR)/$(BINARY)-linux-amd64: | $(DIST_DIR)
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $@ .

$(DIST_DIR)/$(BINARY)-linux-arm64: | $(DIST_DIR)
	GOOS=linux GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $@ .

$(DIST_DIR)/$(BINARY)-darwin-amd64: | $(DIST_DIR)
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $@ .

$(DIST_DIR)/$(BINARY)-darwin-arm64: | $(DIST_DIR)
	GOOS=darwin GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $@ .

$(DIST_DIR)/$(BINARY)-windows-amd64.exe: | $(DIST_DIR)
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $@ .

$(DIST_DIR)/$(BINARY)-windows-arm64.exe: | $(DIST_DIR)
	GOOS=windows GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $@ .

skill: dist | $(SKILL_BIN_DIR)
	$(CP) $(DIST_DIR)/$(BINARY)-linux-amd64 $(SKILL_BIN_DIR)/$(BINARY)-linux-amd64
	$(CP) $(DIST_DIR)/$(BINARY)-linux-arm64 $(SKILL_BIN_DIR)/$(BINARY)-linux-arm64
	$(CP) $(DIST_DIR)/$(BINARY)-darwin-amd64 $(SKILL_BIN_DIR)/$(BINARY)-darwin-amd64
	$(CP) $(DIST_DIR)/$(BINARY)-darwin-arm64 $(SKILL_BIN_DIR)/$(BINARY)-darwin-arm64
	$(CP) $(DIST_DIR)/$(BINARY)-windows-amd64.exe $(SKILL_BIN_DIR)/$(BINARY)-windows-amd64.exe
	$(CP) $(DIST_DIR)/$(BINARY)-windows-arm64.exe $(SKILL_BIN_DIR)/$(BINARY)-windows-arm64.exe

tidy:
	$(GO) mod tidy

clean:
	$(RM_RF) $(DIST_DIR) $(SKILL_BIN_DIR) $(BINARY) $(BINARY).exe
