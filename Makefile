# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main.lambda
VECTY_OUT=vectyui/build/web
VECTY_BUILT=build/web/main.wasm
FLUTTER_OUT=flutterui/build
    
all: clean test build package
build: 
#	$(GOBUILD) ./...
	cd vectyui && GOOS=js GOARCH=wasm $(GOBUILD) -o $(VECTY_BUILT)
	cp `go env GOROOT`"/misc/wasm/wasm_exec.js" $(VECTY_OUT)/
	cd server && go generate && GOOS=linux $(GOBUILD) -o $(BINARY_NAME)
	cd flutterui && flutter build web
package:
	zip -r deploy/awslambda/$(BINARY_NAME).zip server/$(BINARY_NAME) 
	#$(VECTY_OUT) $(FLUTTER_OUT) 
test: 
	$(GOTEST) -v ./...
clean: 
	rm -f server/$(BINARY_NAME)
	rm -f deploy/awslambda/$(BINARY_NAME).zip
	rm -f $(VECTY_OUT)/$(BINARY_NAME).wasm $(VECTY_OUT)/wasm_exec.js
	rm -rf $(FLUTTER_OUT)
