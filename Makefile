SOURCES := ./lookupcfg.go ./util.go


.PHONY: fmt
fmt:
	gofmt -w $(SOURCES)
	golines --max-len 105 -w $(SOURCES)

