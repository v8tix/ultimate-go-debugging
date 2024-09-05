.PHONY: build
build:
	go build -o bin/test-attach-program

.PHONY: run
run:
	./bin/test-attach-program

.PHONY: curl-app
curl-app:
	curl localhost:8081/foo/bar

.PHONY: attach-debugger
attach-debugger:
	# You will got and error if you run this command
	# FIX: sudo echo "0" | sudo tee /proc/sys/kernel/yama/ptrace_scope
	dlv attach $$(pgrep -f test-attach-program)

# make build
# make attach-debugger
# List all the functions on the main package
# (dlv) funcs ^main.*
# (dlv) break main.setupHandlers.func1

