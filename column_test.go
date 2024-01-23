package slice

import (
	"fmt"
	"testing"
)

type columnStruct struct {
	Id   int
	Name string
	Ok   bool
}

var sampleColunmData = []*columnStruct{
	{Id: 10, Name: "aa", Ok: false},
	{Id: 20, Name: "bb", Ok: true},
	{Id: 30, Name: "cc", Ok: false},
}
var notSlice = columnStruct{Id: 10, Name: "aa", Ok: false}

func TestColunm(t *testing.T) {
	rs := Colunm(sampleColunmData, "Name")
	fmt.Println(rs)
	fmt.Println(rs.String())
	ok := Colunm(sampleColunmData, "Ok")
	fmt.Println(ok.Unique())

	fmt.Println(Colunm(notSlice, "Name"))
}
