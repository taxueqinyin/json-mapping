English|[简体中文](https://github.com/taxueqinyin/json-mapping/blob/main/README_cn.md)



JsonMapping is a Go package that provides a method to quickly map json data to a structure. You only need to fill in the json path corresponding to JsonMaping in the tag of the structure, and you can directly map different levels of json to the structure. You can also distinguish different jsons by setting labels to map different jsons to the same structure type, so that the json data returned by the same type of functional interfaces provided by different manufacturers can be accepted by the same structure, realizing the generalization of the interface structure.



# Getting Started

## Installing

To start using json-mapping , install Go and run `go get`:

```go
$ go get -u github.com/taxueqinyin/json-mapping
```

## map to struct

```go
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

func main() {
	testBaidu := TestStruct{}
	testAli := TestStruct{}
	var err error
	err = MappingStruct("baidu", []byte(jsonBaidu), &testBaidu)
	if err != nil {
		fmt.Println("baidu err", err)
		return
	}
	err = MappingStruct("ali", []byte(jsonAli), &testAli)
	if err != nil {
		fmt.Println("ali err", err)
		return
	}
	fmt.Println("baidu = ", testBaidu)
	fmt.Println("ali = ", testAli)
}

```

This will print:

```go
baidu =  {name-baidu cg bug 24 answer 44444444444}
ali =  {name-ali sd aliDeBug 18 dsanswer 3344868}
```





## Path Syntax

The package uses GJSON, so it also uses GJSON syntax. For details, please refer to [GJSON ](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)



## Label Syntax

Each label is separated by ";", and the label uses ":" to correspond to the label and path, such as label:path
