SOURCES := ./*.go ./internal/*.go
MAX_LINE_LENGTH := 120

.PHONY: all
all: help

.PHONY: help
help:
	@echo "COMMAND      DESCRIPTION                                 "
	@echo "---------------------------------------------------------"
	@echo "make fmt     format source code (using gofmt and golines)"
	@echo "make test    run go test for lookupcfg package           "
	@echo "make doc     run go doc for lookupcfg package            "
	@echo "make todo    grep TODOs                                  "

.PHONY: gofmt
gofmt:
	gofmt -w $(SOURCES)

.PHONY: golines
golines:
	golines --max-len $(MAX_LINE_LENGTH) -w $(SOURCES)

.PHONY: fmt
fmt: gofmt golines

.PHONY: todo
todo:
	grep -irn todo $(SOURCES)

.PHONY: test
test:
	go test ./internal/

.PHONY: doc
doc:
	go doc -all
