//R3b00t8oC4L
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSmallestString(strings []string) string {
	if len(strings) == 0 {
		return "" // Handle case where no strings are provided
	}
	var small string
	smallestCount := len(strings[0])
	for _, str := range strings[1:] {
		strCount := len(str)
		if strCount < smallestCount {
			smallestCount = strCount
			small = str
		}
	}

	return small
}

func GetData(args []string) (string, string, string, string, string, string) {
	str := ""
	color := ""
	font := "standard.txt"
	outputFile := ""
	align := "left"
	toBeColored := findSmallestString(args)
	for _, arg := range args {
		if len(args) == 1 {
			str = arg
			break
		}
		if arg == "shadow" || arg == "thinkertoy" || arg == "standard" {
			font = arg + ".txt"
			continue
		}
		if strings.Contains(arg, ".txt") {
			outputFile = arg
			continue
		}
		if arg == toBeColored {
			continue
		}
		if arg[0:2] == "--" && arg[0:8] == "--color=" {
			color = arg[8:]
			continue
		}
		if arg[0:2] == "--" && arg[0:8] == "--align=" {
			align = arg[8:]
			continue
		}
		str += arg + " "
	}
	return str, toBeColored, color, font, outputFile, align
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No Enough arguments !!")
		fmt.Println("Usage:\nEx: go run main.go \"string\"\nEx: go run main.go --color=<color> \"string\"")
		return
	}
	str, toBeColored, color, font, opFile, align := GetData(args)
	fmt.Println("String is :	", str)
	fmt.Println("Smallest:	", toBeColored)
	fmt.Println("Color is:	", color)
	fmt.Println("Font is :	", font)
	fmt.Println("OutPut File :	", opFile)
	fmt.Println("Alignment is :	", align)
	// 2. Checking weather str contain "\n" or not ---> executing the ascii-art
	prev := 'a'
	severallines := false
	nline := 0
	for _, v := range str {
		if v == 'n' && prev == '\\' {
			nline++
			severallines = true
		}
		prev = v
	}
	// 3. Writing text line by line into res
	res := ""
	if severallines {
		args := strings.Split(str, "\\n")
		for _, word := range args {
			if word == "" {
				fmt.Println()
				continue

			}

			for i := 0; i < 8; i++ {

				for _, letter := range word {
					res += GetLine(2+int(letter-' ')*9+i, font)
				}

				fmt.Println(res)
				res = ""
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				res += GetLine(2+(int(letter)-32)*9+i, font)
			}
			fmt.Println(res)
			res = ""
		}
	}
}

// Reading the file
func GetLine(num int, font string) string {
	str := ""
	f, e := os.Open("../fonts/" + font)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		str, _ = content.ReadString('\n')
	}
	str = strings.TrimSuffix(str, "\n")
	return str
}
