# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=testapp

all: clean compile run
print:
	@echo "***** $@ $(BINARY_NAME) *****"
compile: 
	@echo "***** $@ $(BINARY_NAME) *****"
	cd $(GOPATH) && \
	$(GOCMD) install ./src/$(BINARY_NAME)/testlib/ && \
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) ./src/$(BINARY_NAME)
test: 
	@echo "***** $@ $(BINARY_NAME) *****"
clean: 	
	@echo "***** $@ $(BINARY_NAME) *****"
	$(GOCLEAN)
	[ -f $(GOBIN)/$(BINARY_NAME) ] && rm $(GOBIN)/$(BINARY_NAME) || true
run:
	@echo "***** $@ $(BINARY_NAME) *****"
	$(GOBIN)/$(BINARY_NAME)