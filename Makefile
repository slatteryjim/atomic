

.PHONY: depend \
        fmt \
        test

depend:
	go get github.com/onsi/gomega
	go get github.com/onsi/ginkgo/ginkgo

fmt:
	@go get golang.org/x/tools/cmd/goimports
	@goimports -w -l $$(go list -f '{{.Dir}}')

test:
	$(GOPATH)/bin/ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending
