q1:
	go build -o main1 main.go

q2:
	go build -o q2 q2.go

.PHONY: run1 run2 clean

run1: main1
	./main1 <input.txt

run2: q2
	./q2 <input.txt

clean:
	rm -f q1 q3