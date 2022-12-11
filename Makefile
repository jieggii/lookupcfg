SOURCES := ./lookupcfg.go ./util.go ./convert.go

.PHONY: all
all: help

.PHONY: help
help:
	@echo "COMMAND      DESCRIPTION       "
	@echo "-------------------------------"
	@echo "make fmt     format source code"
	@echo "make todo    grep TODOs        "

.PHONY: fmt
fmt:
	gofmt -w $(SOURCES)
	golines --max-len 105 -w $(SOURCES)

.PHONY: todo
todo:
	grep -irn todo ./examples/*.go ./*.go
