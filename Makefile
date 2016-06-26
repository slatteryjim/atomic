

.PHONY: depend \
        test

depend:
	go get github.com/onsi/gomega
	go get github.com/onsi/ginkgo/ginkgo

test:
	$(GOPATH)/bin/ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending