package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0)
	unsigned := flag.Bool("u", false, "Print unsigned varint representation")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}
	n, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		log.Fatalf("Cannot parse %s as integer", flag.Arg(0))
	}
	b := make([]byte, binary.MaxVarintLen64)
	if *unsigned {
		if n < 0 {
			log.Fatalln("Must provide nonnegative integer with -u")
		}
		nn := binary.PutUvarint(b, uint64(n))
		b = b[:nn]
	} else {
		nn := binary.PutVarint(b, n)
		b = b[:nn]
	}
	fmt.Printf("0x%x\n", b)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [flags] N\n", os.Args[0])
	flag.PrintDefaults()
}
