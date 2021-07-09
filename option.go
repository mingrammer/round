package main

import (
	"flag"
	"log"
	"strings"
)

type corner struct {
	topL    bool
	topR    bool
	bottomL bool
	bottomR bool
}

type option struct {
	rate   float64
	owrite bool
	base64 bool
	output string
	prefix string
	suffix string
	corner *corner
}

func defaultOptions() *option {
	return &option{
		rate:   0.25,
		output: "",
		prefix: "",
		suffix: "_rounded",
	}
}

func parseOptions() *option {
	var err error

	opts := defaultOptions()

	rate := flag.Float64("r", opts.rate, "rounding rate. 1 means circular. (0~1)")
	corner := flag.String("c", "tl,tr,bl,br", "comma separated corners to round.")
	owrite := flag.Bool("w", false, "if true, will overwrite the original files.")
	base64 := flag.Bool("b", false, "if true, will read base64-encoded bytes from stdin and print output to stdout.")
	output := flag.String("o", opts.output, "output file name for a single file.")
	prefix := flag.String("p", opts.prefix, "prefix for the output file names.")
	suffix := flag.String("s", opts.suffix, "suffix for the output file names.")
	flag.Parse()

	opts.rate, err = parseRate(*rate)
	if err != nil {
		log.Fatalln(err)
	}
	opts.corner, err = parseCorner(*corner)
	if err != nil {
		log.Fatalln(err)
	}
	opts.owrite = *owrite
	opts.base64 = *base64
	opts.output = *output
	opts.prefix = *prefix
	opts.suffix = *suffix
	return opts
}

func parseRate(rate float64) (float64, error) {
	if rate < 0 || rate > 1 {
		return 0, errInvalidRate
	}
	return rate, nil
}

func parseCorner(corners string) (*corner, error) {
	cn := new(corner)
	cs := strings.Split(corners, ",")
	if len(cs) == 0 {
		cn.topL = true
		cn.topR = true
		cn.bottomL = true
		cn.bottomR = true
	}
	for _, c := range cs {
		switch c {
		case "tl":
			cn.topL = true
		case "tr":
			cn.topR = true
		case "bl":
			cn.bottomL = true
		case "br":
			cn.bottomR = true
		default:
			return nil, errInvalidCorner
		}
	}
	return cn, nil
}
