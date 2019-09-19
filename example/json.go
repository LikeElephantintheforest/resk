package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type User struct {
	//[string]					 于json中的数据格式为string
	Id int `json:"id,string"`
	//【omitempty】				「0值不输出」
	Age int `json:"age,omitempty"`
	//【-】 					「不进行转换」
	Address string `json:"-"`
	Name    string `json:"name"`
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//u := User{
	//	Id:      12,
	//	Name:    "冯菲菲",
	//	Age:     1,
	//	Address: "甜橙金融",
	//}
	//bytes, e := json.Marshal(&u)
	//
	//if e != nil {
	//	fmt.Println(e)
	//} else {
	//	fmt.Println(string(bytes))
	//}

	bytes := []byte(`{"id":"12","name":"冯菲菲","age":1}`)

	u2 := &User{}
	unmarshal := json.Unmarshal(bytes, u2)

	if unmarshal != nil {
		fmt.Println(unmarshal)
	} else {
		fmt.Printf("%+v \n", u2)
		fmt.Print(u2)
	}

}
