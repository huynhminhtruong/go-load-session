APP_NAME=instagram-downloader
BUILD_DIR=./build
ENTRY=./cmd/main.go
CACHE_DIR=$(BUILD_DIR)/.go-cache
GOCACHE=$(abspath $(CACHE_DIR))

.PHONY: build-linux build-windows run clean

# build binary on linux and put it in the build/
build-linux:
	@echo "Building for Linux..."
	@if not exist "$(GOCACHE)" mkdir -p "$(GOCACHE)"
	GOOS=linux GOARCH=amd64 GOCACHE=$(GOCACHE) go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY)

# build binary on windows and put it in the build/
build-windows:
	echo "Building for Windows at $(GOCACHE)..."
	if not exist "$(GOCACHE)" mkdir "$(GOCACHE)"
	set GOCACHE=$(GOCACHE) && set GOOS=windows && set GOARCH=amd64
	go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY)

# run from the binary file in build/
run:
	@echo "Starting the server $(APP_NAME)..."
	$(BUILD_DIR)/$(APP_NAME)

# clean up build/ and cache
clean:
	echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

# install dependencies
get-gin:
	@echo "Installing Gin..."
	@go get -u github.com/gin-gonic/gin

# remove Gin dependency in go.sum
soft-remove-gin:
	@echo "Soft removing Gin..."
	@go mod tidy

# remove Gin dependency completely in local cache
hard-remove-gin:
	@echo "Removing Gin..."
	@go clean -modcache
	@rm -rf $(shell go env GOPATH)/pkg/mod/github.com/gin-gonic/gin@*

# install goquery
get-goquery:
	@echo "Installing GoQuery..."
	@go get -u github.com/PuerkitoBio/goquery

# install chromedp
get-chromedp:
	@echo "Installing Chromedp..."
	@go get -u github.com/chromedp/chromedp