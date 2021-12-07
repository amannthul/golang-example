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
		table := ""
		flag := false
		var wordName, wordType []string
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			str := string(a)
			if flag && str == "}" {
				break
			}
			if strings.HasPrefix(str, "type") {
				res := strings.Fields(str)
				table = res[1]
				flag = true
				fmt.Println(buildTableName(table))
			} else if flag {
				res := strings.Fields(str)
				if res[0] == "CreatedAt" {
					break
				}
				fmt.Println(buildGetFunc(table, res[0], res[1]))
				wordName = append(wordName, res[0])
				wordType = append(wordType, res[1])
			}
		}

		fmt.Println(buildStruct(table, wordName, wordType))
		fmt.Println(buildMapper(table))
		fmt.Println(buildToMapper(table))
		fmt.Println(buildFromMapper(table))
		fmt.Println(buildInterface(table, wordName, wordType))
		fmt.Printf("var _ %vIntf = (*%v)(nil)\n\n", table, table)

		fmt.Println("--------------DB Interface---------------------")
		fmt.Println(buildFromoDomain(table, wordName))
		fmt.Println(buildToDomain(table, wordName))

		fmt.Printf("\n\n--------------END---------------------\n\n")
	}
}

func buildGetFunc(tableName, wordName, wordType string) string {
	tablePrefix := strings.ToLower(string(tableName[0]))
	firstLine := fmt.Sprintf("func (%v *%v) Get%v() %v {", tablePrefix, tableName, wordName, wordType)
	secondLine := fmt.Sprintf("    return %v.%v", tablePrefix, wordName)
	mid := buildJudgement(tablePrefix, wordType)
	return firstLine + "\n" + mid + secondLine + "\n" + "}\n\n"
}

func buildJudgement(prefix, wordType string) string {
	var ret string
	switch wordType {
	case "bool":
		ret = "false"
	case "string":
		ret = "\"\""
	case "byte":
		ret = "0"
	case "float32":
		ret = "0.0"
	case "float64":
		ret = "0.0"
	case "xid.ID":
		ret = "xid.NilID()"
	default:
		ret = wordType + "{}"
	}

	if len(wordType) < 3 {
		ret = "nil"
	}
	if wordType[:2] == "[]" {
		ret = "nil"
	}
	if wordType[:3] == "int" {
		ret = "0"
	}
	res := fmt.Sprintf("%6v %v == nil {\n        return %v\n    }\n", "if", prefix, ret)
	return res
}

func buildTableName(tableName string) string {
	tablePrefix := strings.ToLower(string(tableName[0]))
	firstLine := fmt.Sprintf("func (%v %v) TableName() string {", tablePrefix, tableName)
	secondLine := fmt.Sprintf("     return \"%v\"", ToSnakeCase(tableName))
	return firstLine + "\n" + secondLine + "\n" + "}\n\n"
}

func buildInterface(tableName string, wordName, wordType []string) string {
	res := fmt.Sprintf("type %vIntf interface {\n", tableName)
	for index, content := range wordName {
		f := fmt.Sprintf("	Get%v()  %v\n", content, wordType[index])
		res += f
	}
	res += "}\n"
	return res
}

func buildStruct(tableName string, wordName, wordType []string) string {
	res := fmt.Sprintf("type %v struct {\n", tableName)
	for index, content := range wordName {
		f := fmt.Sprintf("	%v  %v\n", content, wordType[index])
		res += f
	}
	res += "}\n"
	return res
}

func buildMapper(tableName string) string {
	res := fmt.Sprintf("type %vMapper interface {\n", tableName)
	res += fmt.Sprintf("    ToDomain%vMapper\n", tableName)
	res += fmt.Sprintf("    FromDomain%vMapper\n", tableName)
	res += "}\n"
	return res
}

func buildToMapper(tableName string) string {
	res := fmt.Sprintf("type ToDomain%vMapper interface {\n", tableName)
	res += fmt.Sprintf("    ToDomain%v() %vIntf\n", tableName, tableName)
	res += "}\n"
	return res
}

func buildFromMapper(tableName string) string {
	res := fmt.Sprintf("type FromDomain%vMapper interface {\n", tableName)
	res += fmt.Sprintf("    FromDomain%v(%vIntf)\n", tableName, tableName)
	res += "}\n"
	return res
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// db-->domain
func buildToDomain(tableName string, wordName []string) string {
	prefix := strings.ToLower(string(tableName[0]))
	header := "// db->domain\n"
	firstLine := fmt.Sprintf("func (%v *%v) ToDomain%v() domain.%vIntf {\n", prefix, tableName, tableName, tableName)
	secondPart := fmt.Sprintf("    if %v == nil{\n      return nil\n    }\n\n", prefix)
	thirdPart := fmt.Sprintf("    p := domain.%v{\n", tableName)
	ret := make([]string, len(wordName))
	for k, v := range wordName {
		ret[k] = fmt.Sprintf("      %v:      %v.%v,\n", v, prefix, v)
	}
	res := header + firstLine + secondPart + thirdPart
	for _, v := range ret {
		res += v
	}
	res += "    }\n    return &p\n}\n\n"
	return res
}

// domain-->db
func buildFromoDomain(tableName string, wordName []string) string {
	prefix := strings.ToLower(string(tableName[0]))
	header := "// domain->db\n"
	firstLine := fmt.Sprintf("func (%v *%v) FromDomain%v(d domain.%vIntf) {\n", prefix, tableName, tableName, tableName)
	secondPart := fmt.Sprintf("    if %v == nil || d == nil{\n      return\n    }\n\n", prefix)
	ret := make([]string, len(wordName))
	for k, v := range wordName {
		ret[k] = fmt.Sprintf(" 	%v.%v = d.Get%v()\n", prefix, v, v)
	}
	res := header + firstLine + secondPart
	for _, v := range ret {
		res += v
	}
	res += "}\n\n"
	return res
}
