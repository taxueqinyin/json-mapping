JsonMapping是一个Go包，提供了一个从json数据快速映射到结构体的方法，只需要在结构体的tag里填写JsonMaping对应的json路径，就可以把不同层级的json直接对应到结构体中



# 使用

## 安装

要开始使用json-mapping ，请安装Go并运行go get:

```go
$ go get -u github.com/taxueqinyin/json-mapping
```

## 映射到结构体

```go

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

func main(){
    test := TestStruct{}
    _ = MappingStruct([]byte(jsonStr), &test)
    fmt.Println(test)
}
```

这将打印

```go
{音 cg bug 12 answer 44444444444}
```





## 路径语法

包使用了GJSON，所以也采用GJSON语法，详情可参阅[GJSON 语法](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)

