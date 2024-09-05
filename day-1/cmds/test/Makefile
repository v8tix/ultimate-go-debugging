.PHONY: clone-trie
clone-trie:
	git clone -b empty-remove-bug git@github.com:derekparker/trie.git

.PHONY: test
test:
	GO111MODULE=off go test -v ./trie

.PHONY: debug
debug:
	cd ./trie && dlv debug

.PHONY: dlv-test
dlv-test:
	cd ./trie && dlv test

# make clone-trie
# make test
# make debug -> Throw: Not an executable file (Cause is a library)
# make dlv-test
# (dlv) transcript test.txt
# (dlv) continue
# (dlv) stack
# (dlv) break ./trie.go:113
# (dlv) restart
# (dlv) continue
# (dlv) print node
# (dlv) continue
# (dlv) print node
# (dlv) restart
# (dlv) breakpoints
# Set a condition on the breakpoint with ID 1
# (dlv) cond 1 node == nil
# (dlv) breakpoints
# (dlv) continue
# (dlv) print node
# (dlv) exit