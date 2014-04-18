//The XMLParse pkg contains a handfull of function to 'parse' xml to make the data stored in the tags' attribute usefull.
//The functions always successed, only files which don't start with the '<?xml ' are cought as not legal, the data returned is up
//to you to catch if corruped, it is meant to be used as a parse for self created xml
package XMLParse

import (
	"strings"
	"io/ioutil"
)

// Get the data from ONE field in the tags of one xml file returns a empty list if the file does not start with <?xml
func GetDataFieldFromFile(file, field string) []string {
	var data []string
	str := getFileString(file)
	if !(strings.HasPrefix(str, "<?xml")){
		return data
	}
	str2 := splitToLines(str)
	_, str3 := trimToContentLines(str2)
	for _, m := range str3{
		i :=  LineToMap(m)
		final := i[field]
		if final == ""{
		data = append(data, "empty" )
		}
		data = append(data, final )
	}
	return data

}

//Creates a map for each line xml containing the attribute as keys and the string as value, always successeds as well empty map if the file does not start with "<?xml"
func FileToMapPreLine(file string) []map[string]string{
	str := getFileString(file)
	if !(strings.HasPrefix(str, "<?xml")){
		return make([]map[string]string, 0)
	}
	str2 := splitToLines(str)
	_, str3 := trimToContentLines(str2)
	data := make([]map[string]string, 0)
	for _, m := range str3{
		i := LineToMap(m)
		data = append(data, i)
	}
	return data
}

func getFileString(file string) string{
	bytesStr, er := ioutil.ReadFile(file)
	if er != nil{
		return ""
	}
	str := string(bytesStr)
	return str
}

func splitToLines( content string) []string{
	contentLines := strings.Split(content, "\n")
	return contentLines
}
func trimToContentLines(content []string)( []string, []string){
	var con []string
	var notCon []string
	for _,m := range content{
		if strings.HasPrefix(m, "<?") {
			notCon = append(notCon , m)
		}else{
			con = append(con, m)
		}
	}
	return notCon, con
}

// takes one line of xml and splits the line into a slice of strings each containing the attribute=string
func ToFields (line string)  []string{
	var result []string
	firstWhiteSpace := strings.Index(line, " ")
	lastWhiteSpace := strings.LastIndex(line, " ")
	if firstWhiteSpace==-1 || lastWhiteSpace==-1{
		return result
	}
	tagType  := line[:firstWhiteSpace] + line[lastWhiteSpace:]
	result = append(result, tagType)
	workOnLine := line[firstWhiteSpace:lastWhiteSpace]
	seps := strings.Split(workOnLine, "\"")
	numElements := (len(seps)-2)
	for n:=0; n<=numElements; n++{
		if len(seps)<n+1 {
			return result
		}else{
			cache := seps[n]+seps[n+1]
			result = append(result, cache)
			n++
		}
	}
	return result
}

//takes a slice of keys and one field you are looking for and return the field attribute=string  if present otherwise an empty string
func FindField(fields []string, field string) string{
	var str string
	for _ , m := range fields{
		d := strings.Contains(m, field + "=")
		if d{
			str = m
		}
	}
	return str
}

//Takes a field and return the value as a string
func FieldValue(field string) string{
	i := strings.Index(field, "=")
	if i>0 {
		return field[i+1:]
	}else{
		return ""
	}
}
// Takes one line and return a map with attribute as key and string as value 
func LineToMap(line string) map[string]string{
	fields := ToFields(line)
	result := FieldsToMap(fields)
	return result
}

//Takes a slice of fields and creates a map of it attribute as key and string as value 
func FieldsToMap(fields []string)map[string]string {
	result :=  make(map[string]string)
	for _, m := range fields{
		i:= strings.Index(m, "=")
		if i<0 {
		}else{
			n := strings.Replace(m[:i], " ", "", -1)
			result[n]=m[i+1:]
		}
	}
	return result
}


