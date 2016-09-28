# sampler
Implementation of reservoir sampling in golang.

Reading godoc and souce code of cmd/sampler help to understand usage.

## cmd/sampler
Thi program reservior-sample lines given from stdin.
A number of samples can specified command-line argument.

### example
```
$ seq 100 | sampler -k 5
28
15
52
7
42
```

### installation
```
go get -u github.com/zaltoprofen/sampler/cmd/sampler
```
