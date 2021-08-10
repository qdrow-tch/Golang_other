package main

import (
	"fmt"

	"github.com/qdrow-tch/Golang/tree/ch2t7/sql_parser"
)

func main() {

	query := "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?"
	goodquery, values := sql_parser.Parse_query(query, false, []int{1, 6, 234}, 555)
	fmt.Println(goodquery, values)

}
