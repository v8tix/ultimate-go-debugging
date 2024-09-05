bin/traceme:
	go build -gcflags="all=-N -l" -o bin/traceme

.PHONY: strace
strace:
	# Display information about all the system calls made by the program.
	strace ./bin/traceme

.PHONY: trace
trace:
	dlv trace ^main.callMe

# make bin/traceme
# make bin/strace
# make trace
