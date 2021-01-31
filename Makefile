GO_CMD?=go
GO_VERSION_MIN=1.15.3
CGO_ENABLED?=0
GOFMT_FILES?=$$(find . -name '*.go' | grep -v pb.go | grep -v vendor)

EXTERNAL_TOOLS_CI=\
	github.com/mitchellh/gox

TEST?=$$($(GO_CMD) list ./... | grep -v /vendor/ | grep -v /integ)
TEST_TIMEOUT?=45m

default: prep

# test runs the unit tests and vets the code
test: ci-bootstrap
	@CGO_ENABLED=$(CGO_ENABLED) \
	$(GO_CMD) test -tags='$(BUILD_TAGS)' $(TEST) $(TESTARGS) -timeout=$(TEST_TIMEOUT) -parallel=20

# bootstrap the build by downloading additional tools needed to build
ci-bootstrap:
	@for tool in  $(EXTERNAL_TOOLS_CI) ; do \
		echo "Installing/Updating $$tool" ; \
		GO111MODULE=off $(GO_CMD) get -u $$tool; \
	done

bootstrap: ci-bootstrap
	@for tool in  $(EXTERNAL_TOOLS_CI) ; do \
		echo "Installing/Updating $$tool" ; \
		GO111MODULE=off $(GO_CMD) get -u $$tool; \
	done
	@sh -c "'$(CURDIR)/scripts/goversioncheck.sh' '$(GO_VERSION_MIN)'"
	@$(GO_CMD) generate $($(GO_CMD) list ./... | grep -v /vendor/)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

# dev creates binaries for testing Vault locally. These are put
# into ./bin/ as well as $GOPATH/bin
prep: bootstrap fmtcheck
	@CGO_ENABLED=$(CGO_ENABLED) BUILD_TAGS='$(BUILD_TAGS)' sh -c "'$(CURDIR)/scripts/build.sh'"

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@$(GO_CMD) list -f '{{.Dir}}' ./... | grep -v /vendor/ \
		| xargs $(GO_CMD) vet ; if [ $$? -eq 1 ]; then \
			echo ""; \
			echo "Vet found suspicious constructs. Please check the reported constructs"; \
			echo "and fix them if necessary before submitting the code for reviewal."; \
		fi

# lint runs vet plus a number of other checkers, it is more comprehensive, but louder
lint:
	@$(GO_CMD) list -f '{{.Dir}}' ./... | grep -v /vendor/ \
		| xargs golangci-lint run; if [ $$? -eq 1 ]; then \
			echo ""; \
			echo "Lint found suspicious constructs. Please check the reported constructs"; \
			echo "and fix them if necessary before submitting the code for reviewal."; \
		fi

.PHONY: bootstrap prep fmt vet
