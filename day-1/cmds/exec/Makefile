bin/execme:
	go build -o bin/execme

.PHONY: dlv-exec
dlv-exec:
# Let you debug an existing binary
	dlv exec ./bin/execme

.PHONY: build-showing-inlining
build-showing-inlining:
# shows the optimizations made by the Go compiler
	go build -gcflags=-m=2 -o bin/prog

.PHONY: search-dwarf-info
search-dwarf-info: bin/execme
	objdump --dwarf=info bin/execme | grep -C 2 Rectangle

bin/execmenooptimizationsexecme:
	go build -o bin/execmenooptimizations -gcflags="all=-N -l"

.PHONY: dlv-exec-no-optimizations
dlv-exec-no-optimizations:
	dlv exec ./bin/execmenooptimizations

# make bin/execme
# make dlv-exec
# (dlv) transcript exec-1.txt
# We can set breakpoints on file line, functions
# (dlv) break Rectangle.Area
#### Breakpoint 1 set at 0x481c05,0x481c65,0x481ba5 for main.main() <multiple locations>:0 -> Numerous memory addresses are present due to the binary being optimized with inline optimization.
# (dlv) break Rectangle.Width
#### Command failed: location "Rectangle.Width" not found -> Since the compiler detected that this function returns a constant, it was removed during the optimization process
# (dlv) funcs
# Search functions using regular expression
# (dlv) funcs Rectangle
# (dlv) exit
# make bin/execmenooptimizationsexecme
# make dlv-exec-no-optimizations
# (dlv) break Rectangle.Area
#### Breakpoint 1 set at 0x49d5aa for main.(*Rectangle).Area() ./main.go:20
# (dlv) break Rectangle.Width
# -l means location
# (dlv) disassemble -l main.main
# (dlv) exit
