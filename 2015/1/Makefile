q1:
	go build -o q1 q1.go

q2:
	go build -o q2 q2.go

.PHONY: run1 run2 clean

run1: q1
	./q1 <q1.sample.txt

run2: q2
	./q2 <q2.sample.txt

clean:
	rm -f q1 q2