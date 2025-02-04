q1:
	go build -o main main.go

.PHONY: run1 run2 clean

run: main
	./main input.txt

runTest: main 
	./main sample.txt

clean:
	rm -f main