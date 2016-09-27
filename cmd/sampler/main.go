package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/zaltoprofen/sampler"
)

var (
	k = flag.Int("k", 1, "number of samples")
)

func main() {
	os.Exit(_main())
}

func printError(msg interface{}) {
	fmt.Printf("error:\t%s\n", msg)
}

func _main() int {
	flag.Parse()
	if *k < 1 {
		printError("k must be larger than 0")
		return 1
	}

	r := bufio.NewReader(os.Stdin)
	src := func() (interface{}, error) {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			return nil, err
		}
		return line[:len(line)-1], nil
	}

	sampled, err := sampler.Sample(*k, sampler.IteratorFunc(src))
	if err != nil {
		printError(err)
		return 1
	}
	for _, s := range sampled {
		if ss, ok := s.(string); ok {
			fmt.Println(ss)
		} else {
			panic("failed type assertion")
		}
	}
	return 0
}
