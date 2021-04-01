
run-tests:
        # build runner
	cd ../runner && go build .
	###############################
	# test demo
	gcc demo.c -o demo
	../runner/runner json --file demo.json

	@echo

	###############################
	# test demo with flush
	gcc demo_with_flush.c -o demo_with_flush
	../runner/runner json --file demo_with_flush.json

	@echo

	###############################
	# test demo with line
	gcc demo_with_line.c -o demo_with_line
	../runner/runner json --file demo_with_line.json
