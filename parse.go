package XMLParse

import (
	"strings"
	"io/ioutil"
)


func GetDataFieldFromFile(file, field string) []string {
	bytesStr, _ := ioutil.ReadFile(file)
	str := string(bytesStr)
	mes := strings.SplitAfter(str, field + "=\"" )
	var data []string
	for _, m := range mes{
		b := strings.SplitAfter(m, "\"")
		data = append(data, strings.Replace(b[0], "\"", " ", 1))

	}
	return data

}



