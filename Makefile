.PHONY : test_bench

test_bench:
	go test . -run=Bench -bench=.