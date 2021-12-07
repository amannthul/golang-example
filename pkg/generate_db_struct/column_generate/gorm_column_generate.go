package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

var (
	PrimaryColumn = "   `gorm:\"primary_key;column:%s\"`"
	CommonColumn  = "   `gorm:\"column:%s\"`"
)

// TODO: Go AST改造
func main() {
	p := os.Args[1:]
	for _, v := range p {
		fi, err := os.Open(v)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer fi.Close()
		br := bufio.NewReader(fi)
		// table := ""
		flag, first, generate := false, true, true
		// var wordName, wordType []string
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			str := string(a)
			if flag && str == "}" {
				fmt.Println(str)
				break
			}
			if strings.HasPrefix(str, "type") {
				flag = true
				fmt.Println(str)
			} else if flag {
				res := strings.Fields(str)
				if res[0] == "CreatedAt" {
					generate = false
				}
				// 对已经写过的行进行兼容
				if len(res) >= 3 {
					if first {
						first = false
					}
					fmt.Println(str)
					continue
				}
				if generate {
					// 默认第1行是主键
					if first {
						first = false
						fmt.Printf(str+PrimaryColumn+"\n", ToSnakeCase(res[0]))
						continue
					}
					fmt.Printf(str+CommonColumn+"\n", ToSnakeCase(res[0]))
				} else {
					fmt.Println(str)
				}
			}
		}
	}
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
