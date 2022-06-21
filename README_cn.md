[English](https://github.com/taxueqinyin/json-mapping/blob/main/README.md)|简体中文



JsonMapping是一个Go包，提供了一个将json数据快速映射到结构体的方法，只需要在结构体的tag里填写JsonMaping对应的json路径，就可以把不同层级的json直接对应到结构体中。还可以通过设置label来区分不同json，以将不同json数据映射到同一个结构体类型，使得不同厂商提供的同类型功能接口返回的json数据可以用同一个结构体承接，实现接口结构体的通用化



# 使用

## 安装

要开始使用json-mapping ，请安装Go并运行go get:

```go
$ go get -u github.com/taxueqinyin/json-mapping
```

## 映射到结构体

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

这将打印

```go
baidu =  {name-baidu cg bug 24 answer 44444444444}
ali =  {name-ali sd aliDeBug 18 dsanswer 3344868}
```





## 路径语法

包使用了GJSON，所以也采用GJSON语法，详情可参阅[GJSON 语法](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)

## 标签语法 

每个label用";"隔开，label使用":"来对应label和路径,例如label:path
