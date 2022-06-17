package json_mapping

import (
	"fmt"
	"testing"
)

var jsonStr = `
{
    "name":"音",
	"data":["sd","hp","cg","sb"],
	"bug":"bug",
    "baseInfo":{
        "address":"福建省",
        "age":12,
        "first":{
            "second":{
                "third":[
                    "none",
                    "answer"
                ]
            }
        },
        "f":{
            "s":{
                "t":{
                    "four":44444444444
                }    
            }
        }
    }
}
`

type TestStruct struct {
	Name1  string `JsonMapping:"name"`
	Data   string `json:"data" JsonMapping:"data.2"`
	Bug    string `json:"bug"`
	Age    int    `JsonMapping:"baseInfo.age"`
	Answer string `JsonMapping:"baseInfo.first.second.thi*.1"`
	Four   int    `JsonMapping:"baseInfo.f.s.t.fou*"`
}

func TestMappingStruct(t *testing.T) {
	test := TestStruct{}
	var err error
	for i := 0; i < 100000; i++ {
		err = MappingStruct([]byte(jsonStr), &test)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test)
}
