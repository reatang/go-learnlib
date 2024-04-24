package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	data := Data{1, "reatang"}

	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))

	data2 := Data{}
	_ = json.Unmarshal(jsonStr, &data2)
	fmt.Println(data2)
}
