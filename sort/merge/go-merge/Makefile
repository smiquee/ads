all: mergesort

mergesort:
	GOPATH=${PWD} go install ms/mergesort

run_mergesort: mergesort
	@bin/mergesort ../../data.in
