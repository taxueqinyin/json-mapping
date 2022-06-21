package json_mapping

import (
	"fmt"
	"testing"
)

var (
	jsonBaidu = `
{
    "name":"name-baidu",
	"data":["sd","hp","cg","sb"],
	"bug":"bug",
    "baseInfo":{
        "address":"China",
        "age":12,
		"baidu_age":24,
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

	jsonAli = `
{
    "a": {
        "b": "sd",
        "age": 18
    },
	"bug":"aliDeBug",
    "random": [
        "dsd",
        "name-ali"
    ],
    "data": {
		"ffff":3344868,
        "test": [
            {
                "ds": {
                    "ds": "dsanswer"
                }
            }
        ]
    }
}
`
)

type TestStruct struct {
	Name1  string `JsonMapping:"baidu:name;ali:random.1"`
	Data   string `json:"data" JsonMapping:"baidu:data.2;ali:a.b"`
	Bug    string `json:"bug"`
	Age    int    `JsonMapping:"baidu:baseInfo.baidu_age;ali:a.age"`
	Answer string `JsonMapping:"baidu:baseInfo.first.second.thi*.1;ali:data.test.0.ds.ds"`
	Four   int    `JsonMapping:"baidu:baseInfo.f.s.t.fou*;ali:data.ffff"`
}

func TestMappingStruct(t *testing.T) {
	testBaidu := TestStruct{}
	testAli := TestStruct{}
	var err error
	for i := 0; i < 500000; i++ {
		err = MappingStruct("baidu", []byte(jsonBaidu), &testBaidu)
		if err != nil {
			fmt.Println("baidu err", err)
			return
		}
	}
	err = MappingStruct("ali", []byte(jsonAli), &testAli)
	if err != nil {
		fmt.Println("ali err", err)
		return
	}
	fmt.Println("baidu = ", testBaidu)
	fmt.Println("ali = ", testAli)
}
