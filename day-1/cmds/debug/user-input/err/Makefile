PHONY: run
run:
	go run main.go

.PHONY: debug
debug:
	dlv debug

# dlv debug
# (dlv) continue
# (dlv) stack

# Display the source code around line 23 of the file.
# (dlv) list ./main.go:23

# Show the exact location within the source code where we are currently.
# (dlv) list

# Print all the local variables from the stack 2 (see dlv stack)
# (dlv) frame 2 locals

# Print the variable err
# (dlv) frame 2 print err

# Exit
# (dlv) exit

.PHONY: debug-with-tty
debug-with-tty:
	dlv debug --tty=/dev/pts/2

.PHONY: install-ptyme
install-ptyme:
	git clone https://github.com/derekparker/ptyme.git && cd ptyme && cargo build --release

.PHONY: run-ptyme
run-ptyme: install-ptyme
	cd ptyme && cargo run

.PHONY: debug-headless
debug-headless:
	# Delve (dlv) consists of a client (CLI) and a server. When the --headless option is specified, it instructs Delve to start only the server.
	dlv debug --headless --listen :9090 --api-version 2

.PHONY: connect-to-headless-delve
connect-to-headless-delve:
	dlv connect :9090


# In one terminal run: make debug-headless
# In another terminal run: make connect-to-headless-delve
# On the client side:
# (dlv) continue
# (dlv) transcript debugpt2.txt
# (dlv) stack
# (dlv) break main.go:38
# (dlv) restart
# (dlv) continue
# (dlv) list -10
# (dlv) print i
# (dlv) list -15
# (dlv) break main.go:26
# (dlv) restart
# (dlv) continue
# (dlv) print i
# Watch is breakpoint set in memory location
# (dlv) help watch
# (dlv) watch -w i
##### Watchpoint i set at 0xc000110020 -> 0xc000110020 variable i memory location
# (dlv) continue
# (dlv) breakpoints
# (dlv) list
##### fix the code
# (dlv) rebuild
# exit
# Restart the server: make debug-headless
# Connect to the server: make connect-to-headless-delve
# (dlv) continue