package XMLParse

import (
	"strings"
	"io/ioutil"
)


func GetDataFieldFromFile(file, field string) []string {
	str := getFileString(file)
	str2 := splitToLines(str)
	_, str3 := trimToContentLines(str2)
	var data []string
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

func FileToMapPreLine(file string) []map[string]string{
	str := getFileString(file)
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
	bytesStr, _ := ioutil.ReadFile(file)
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
func FieldValue(field string) string{
	i := strings.Index(field, "=")
	if i>0 {
		return field[i+1:]
	}else{
		return ""
	}
}

func LineToMap(line string) map[string]string{
	fields := ToFields(line)
	result := FieldsToMap(fields)
	return result
}

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


