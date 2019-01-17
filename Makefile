HOOKS_PATH=hooks
DIFF_FILES=$(shell git diff-index --cached --name-only origin/develop | xargs printf -- '--include=%s\n')
MODIFIED_FILES=$(shell git diff-index --cached --name-only HEAD | xargs printf -- '--include=%s\n')

.PHONY: dep setup test coverage mocks

dep:
	@go get -v github.com/onsi/ginkgo/ginkgo
	@go get -v github.com/onsi/gomega/...
	@go get -v github.com/axw/gocov/gocov
	@go get -v gopkg.in/matm/v1/gocov-html
	@go get -v github.com/vektra/mockery/.../
	@go get -v github.com/alecthomas/gometalinter
	@gometalinter --install > NUL
	@rm NUL

fmt:
	@go fmt ./...

setup: dep
	@cp -f $(HOOKS_PATH)/** .git/hooks

run: dep
	@go run main.go

check: setup
	@gometalinter ./... --aggregate --fast $(MODIFIED_FILES)

deep-check: setup
	@gometalinter ./... --aggregate $(DIFF_FILES)

full-check: setup
	@gometalinter ./... --aggregate

test: setup
	@ginkgo -gcflags=-l ./...

test-integ:
	@ginkgo -gcflags=-l --tags=integration ./...

cov: setup
	@gocov test -gcflags=-l ./... | gocov report

cov-html: setup
	@gocov test -gcflags=-l ./... | gocov-html > cov.html
