.PHONY: debug
debug:
# replace the standard output terminal by a file
	dlv debug -r stdin.txt

# make debug
# help
# (dlv) transcript session-1.txt
# Pressing tab after this command, makes Delve executes a fuzziness search
# (dlv) continue mama
# (dlv) breakpoints
# (dlv) list
# (dlv) next
# (dlv) next
# If you are over a function, next will executes that function without step over that function
# (dlv) next
# If you are over a function, step will step over that function
# (dlv) step
# If you are inside a function, step out will continue to the next line after the function was called
# (dlv) stepout
# If hit enter without specifying a command, Delve will executes the last command entered. For this example stepout
# (dlv)
# (dlv) disassemble
# (dlv) step-instruction
# (dlv) continue mama
# (dlv) next
# (dlv) print data
# (dlv) print string(data)
# In another terminal
# echo "and now for something completely different" > stdin2.txt
# Change the stdout and stdin standard redirections
# (dlv) restart <stdin2.txt
# (dlv) continue mama
# Execute next until you reach the fmt.Println("got data") statement
# (dlv) print string(data)
# (dlv) help restart
# You can set a breakpoint name
# (dlv) break mainfunc main.main
# (dlv) breakpoints
# (dlv) clear mainfunc
# (dlv) break main.main
# (dlv) clear 5
# Create a breakpoint on every function on the main package
# (dlv) break /^main.*/
# (dlv) clearall
# (dlv) break m33 main.go:33
# (dlv) help cond
# (dlv) restart
# (dlv) cond -hitcount m33 == 5
# (dlv) continue
#### [m33] main.produceValues() ./main.go:33 (hits goroutine(37):5 total:5) (PC: 0x4c82a1)
# (dlv) cond -clear m33
# (dlv) clear m33
# Executes a command everytime a breakpoint is hit
# (dlv) help on
# (dlv) break m33 main.go:33
# (dlv) on m33 print id
# (dlv) continue
# (dlv) toggle m33
# (dlv) breakpoints
# (dlv) continue
# (dlv) toggle m33
# (dlv) continue
# (dlv) clearall
# (dlv) continue main.go:61
# (dlv) continue main.go:31
# (dlv) continue main.go:33
# (dlv) list
# (dlv) watch -rw id
#### Command failed: can not watch stack allocated variable for reads
# (dlv) watch -w id
# Display the args of the actual function
# (dlv) args
# (dlv) locals
# Display all global variables defined in all packages
# (dlv) vars
# Display global variables defined in the reflect packages
# (dlv) vars reflect
# (dlv) display -a id
# (dlv) next
# (dlv) display -a id == 5
# (dlv) next
# (dlv) locals
# (dlv) print &p
#### (*main.packet)(0xc000188e80)
# Check the raw bytes in mem of the address specified
# (dlv) examinemem -fmt hex -count 20 0xc000188e80
# (dlv) help x
# Display the variable type
# (dlv) whatis p
# (dlv) print p.id
# (dlv) print p.val
# (dlv) print p.val[13]
# (dlv) set p.id = 500
# (dlv) print p.id
# Display CPU registers information
# (dlv) regs
# Display CPU registers extended information
# (dlv) regs -a
# Display all goroutines
# (dlv) goroutines
# Show the goroutine 36 stack
# (dlv) goroutine 36 stack
# Switch the actual goroutine to the goroutine 36
# (dlv) goroutine 36
# (dlv) list
# (dlv) stack
# Display all threads
# (dlv) threads
# Switch the actual thread to the thread 15215
# (dlv) threads 15215

# help stack
# up
# down
# stack
# frame 1 list
# frame 1 locals
# frame 1 args
# exit



