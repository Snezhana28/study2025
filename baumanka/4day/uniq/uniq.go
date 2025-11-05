package uniq

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	c      bool
	d      bool
	u      bool
	i      bool
	f      int
	s      string
	input  bool
	output bool
}

func Run() {
	options, err := GetOptions(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	} else {
		counts, unique := Lines(options)
		Output(options, counts, unique)
	}

}

func GetOptions(args []string) (Options, error) {
	options := Options{}
	fs := flag.NewFlagSet("options", flag.ContinueOnError)
	fs.BoolVar(&options.c, "c", false, "")
	fs.BoolVar(&options.d, "d", false, "")
	fs.BoolVar(&options.u, "u", false, "")
	fs.BoolVar(&options.i, "i", false, "")
	fs.IntVar(&options.f, "f", 0, "")
	fs.StringVar(&options.s, "s", "", "")

	err := fs.Parse(args)
	if err != nil {
		return options, err
	}

	count := 0
	if options.c {
		count++
	}
	if options.d {
		count++
	}
	if options.u {
		count++
	}

	if count > 1 {
		return options, fmt.Errorf("usage c d u")
	}
	return options, nil
}

func Lines(options Options) (map[string]uint, map[string]string) {
	//uniqueSlice := make([]int, 10)
	counts := make(map[string]uint)
	//unique := make(map[string]uint)
	unique := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		key := buildKey(line, options)
		if options.i {
			key = strings.ToLower(line)
		}
		counts[key]++
		if _, exists := unique[key]; !exists {
			unique[key] = line
		}
	}
	return counts, unique
}

func buildKey(line string, options Options) string {
	var key string
	if options.f > 0 {
		fields := strings.Fields(line)
		if len(fields) > options.f {
			fields = fields[options.f:]
			key = strings.Join(fields, "_")
		} else {
			key = line
		}

	} else {
		key = line
	}
	if options.i {
		key = strings.ToLower(key)
	}
	if options.s != "" {
		numChars, err := strconv.Atoi(options.s)
		if err == nil {
			if len(key) > numChars {
				key = key[numChars:]
			} else {
				key = ""
			}
		}
	}
	return key
}

func Output(options Options, counts map[string]uint, unique map[string]string) {
	for key, num := range counts {
		line := unique[key]
		if options.d {
			if num > 1 {
				fmt.Println(line)
			}
		} else if options.u {
			if num == 1 {
				fmt.Println(line)
			}
		} else {
			if options.c {
				fmt.Print(num, " ")
			}
			fmt.Println(line)
		}
	}
}
