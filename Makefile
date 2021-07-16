f:
	find . -name '*.go' | grep -Ev 'vendor' | xargs gofmt -s -w && find . -name '*.go' | grep -Ev 'vendor' | xargs ./.tmp/goimports -w -local go_practice
