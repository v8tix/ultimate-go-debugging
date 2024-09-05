.PHONY: build-crash
build-crash:
	go build -gcflags="all=-N -l" -o bin/crash crash.go

.PHONY: run-crash
run-crash:
	./bin/crash

.PHONY: run-crash-gotraceback
run-crash-gotraceback:
	GOTRACEBACK=crash ./bin/crash

.PHONY: dlv-core
dlv-core:
	dlv core ./bin/crash core-crash.54818

# core dump is a snapshot of your program's memory at a specific moment
# make build-crash
# make run-crash
# ulimit -c unlimited
# SEE README.MD to save core dumbs in this directory
# make run-crash-gotraceback
# (dlv) stack
# (dlv) frame 11
# (dlv) list
# (dlv) frame 12 locals
# We can't use commands to interact with a life program, cause this is a snapshot

# how to create a core dump from a life process
# dlv debug crash.go
# (dlv) continue main.main
# (dlv) n
# (dlv) help dump
# (dlv) dump crash.debug
# (dlv) exit
# dlv core ./bin/crash crash.debug
# (dlv) stack
# (dlv) threads
# (dlv) thread 55131
# (dlv) stack
# (dlv) list
# (dlv) exit