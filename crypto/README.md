go test -c
mv crypto.test crypto.old
go test -c
mv crypto.test crypto.new
./crypto.old -test.benchmem -test.bench=BenchmarkCrcHash -test.count=10 > old.txt
./crypto.new -test.benchmem -test.bench=BenchmarkCrcHash -test.count=10 > new.txt
benchstat old.txt new.txt