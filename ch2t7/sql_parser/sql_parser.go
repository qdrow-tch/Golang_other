package sql_parser

import (
	"reflect"
	"strings"
)

func Parse_query(q string, args ...interface{}) (string, []interface{}) {

	outsl := []interface{}{}

	slq := strings.Split(q, "?")

	var fslq []string

	for i := range args {
		rv := reflect.ValueOf(args[i])
		if rv.Type().ConvertibleTo(reflect.TypeOf([]int{})) {
			sl := rv.Interface().([]int)
			cashstr := ""
			for j := range sl {
				outsl = append(outsl, sl[j])
				cashstr += "?"
			}
			fslq = append(fslq, slq[i], cashstr)

		} else {
			outsl = append(outsl, args[i])
			fslq = append(fslq, slq[i], "?")
		}
	}
	fslq = append(fslq, slq[len(slq)-1])

	return strings.Join(fslq, ""), outsl
}

//Вызов: func ( "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555 )

//Ответ: "SELECT * FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ?", []interface{}{ false, 1, 6, 234, 555 }
