.PHONY: run
run:
	go run last-character.go

.PHONY: install-delve
install-delve:
	go install github.com/go-delve/delve/cmd/dlv@latest

# works like the go build command, it compiles, run the program and attach the debugger
.PHONY: debug
debug:
	dlv debug

# Save the debug session to a file
# (dlv) transcript debugpt1-1.txt

# Continue running the program until it either finishes, reaches a breakpoint, or crashes.
# If a panic occurs, Delve preserves the entire program state before exiting.
# (dlv) continue

# Print the stack trace at that point, and capture key information such as the file and line number where the panic occurred.
#(dlv) stack

# Go back into the beginning.
#(dlv) restart

# Set a breakpoint.
#(dlv) break ./last-character.go:11
# Breakpoint 1 set at 0x49b32e for main.main() ./last-character.go:11
# 1 is the ID of the break point
# 0x49b32e is the address of the break point

#(dlv) continue

# Inspect the variables in your program (os.Args for this example)
#(dlv) print os.Args

# Once you detect the error, add some code to fix it. Then you need to clear all breakpoints
#(dlv) clearall

# Then you need to rebuild the program
#(dlv) rebuild

# Run the program
#(dlv) continue

# Restart the program passing arguments
#(dlv) restart "debugging is so fun"

# Run the program
#(dlv) continue

# Check the stack
#(dlv) stack

# Set a breakpoint.
#(dlv) break ./last-character.go:18

# Restart.
#(dlv) restart

# Continue.
#(dlv) continue

# Print variables.
#(dlv) print arg

# Add fix code.
# Clear breakpoints.
# Rebuild.
# Continue.
# Exit.
#(dlv) exit

.PHONY: debug-with-args
debug-with-args:
	dlv debug -- "debugging is so fun"