.PHONY: build
build:
	go build -o bin/myclient

.PHONY: run
run:
	./bin/myclient localhost:9999 | (jq . || echo "Non-JSON output detected, skipping jq")

.PHONY: run-testprog-with-dlv
run-testprog-with-dlv:
	cd ./testprog && dlv debug --continue --api-version=2 --headless --accept-multiclient --listen :9999

.PHONY: kill-dlv-server
kill-dlv-server:
	pkill dlv

# make build
# In another terminal
# make run-testprog-with-dlv
# In the previous terminal
# make run