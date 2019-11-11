package main

import (
	 "fmt"
	"encoding/json"
	"os"
)

type PersonInfo struct {
    Key    string
    Value     string
}


func main() {
	f,_ := os.Create("/root/go/6.824/src/test/test")
	
	enc :=json.NewEncoder(f)

	personInfo := PersonInfo{"1",""}
	err := enc.Encode(personInfo)
	if err != nil {
        fmt.Println(err)
    }
	personInfo2 := PersonInfo{"2",""}
	err = enc.Encode(personInfo2)

	f,err = os.Open("/root/go/6.824/src/test/1")
	if err != nil {
        fmt.Println(err)
	}
	var person PersonInfo
	decoder := json.NewDecoder(f)
    err = decoder.Decode(&person)
	if err != nil {
        fmt.Println(err)
	}
	fmt.Println(person)
	err = decoder.Decode(&person)
	if err != nil {
        fmt.Println(err)
	}
	fmt.Println(person)
	
	f.Close()
	


}
