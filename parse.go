package XMLParse

import (
	"strings"
	"io/ioutil"
)


func GetDataFieldFromFile(file, field string) []string {
	bytesStr, _ := ioutil.ReadFile(file)
	str := string(bytesStr)
	mes := strings.SplitAfter(str, field + "\"" )
	var data []string
	for n, m := range mes{
		b := strings.SplitAfter(m, "\"")
		data[n] = strings.Replace(b[0], "\"", " ", 1)

	}
	return data

}



