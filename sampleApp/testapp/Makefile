GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
PROJECT_NAME=sampleApp
BINARY_NAME=testapp

all: clean compile exec
print:
	@echo "***** $@ $(BINARY_NAME) *****"
compile: 
	@echo "***** $@ $(BINARY_NAME) *****"
	cd $(GOPATH) && \
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) ./src/$(PROJECT_NAME)/$(BINARY_NAME)
test: 
	@echo "***** $@ $(BINARY_NAME) *****"
clean: 	
	@echo "***** $@ $(BINARY_NAME) *****"
	$(GOCLEAN)
	[ -f $(GOBIN)/$(BINARY_NAME) ] && rm $(GOBIN)/$(BINARY_NAME) || true
exec:
	@echo "***** $@ $(BINARY_NAME) *****"
	$(GOBIN)/$(BINARY_NAME)
run:
	@echo "***** $@ $(BINARY_NAME) *****"
	cd $(GOPATH) && $(GORUN) ./src/$(PROJECT_NAME)/$(BINARY_NAME)/*.go