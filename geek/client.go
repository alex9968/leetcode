package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var CfgMap map[string]*Config

type Config struct {
	Endpoint string
}

func init() {
	CfgMap = make(map[string]*Config, 4)

	CfgMap["hello"] = &Config{
		Endpoint: "http://localhost:8080/",
	}
}


type HelloService interface{
	SayHello(string) (string, error)
}

var _ HelloService = &hello{}

type hello struct {
	endpoint string
	FuncFiled func( string) (string, error)
}


type Service interface {
	ServiceName() string
}

func (h *hello) ServiceName() string {
	return "hello"
}


func (h *hello) SayHello(s string) (string, error) {

	client := http.Client{}
	res, err := client.Get(h.endpoint + "/" + s)

	if err  != nil {
		println("err")
		return "", nil
	}

	data,err := ioutil.ReadAll(res.Body)
	if err !=nil {
		fmt.Println("%+v", err)
		return "", nil
	}

	return  string(data), nil

}

//SetFuncFiled
func SetFuncFiled(val interface{}) {
	// val 不是指针, 可以直接获取
	// t := reflect.TypeOf(val)

	//val  是指针，需要间接获取
	v := reflect.ValueOf(val)
	ele := v.Elem()
	t := ele.Type()

	num := t.NumField()
	for i:=0;i<num;i++ {
		f := ele.Field(i)
		if f.CanSet() {
			// fmt.Printf("aaa")
			fn := func(args []reflect.Value)(results []reflect.Value) {
				name := args[0].Interface().(string)
				client := http.Client{}

				serviceName := val.(Service).ServiceName()
				endpoint := CfgMap[serviceName].Endpoint

				res, err := client.Get(endpoint + name)

				if err != nil {
					println("err")
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println("%+v", err)
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				fmt.Println(string(data))
				return []reflect.Value{reflect.ValueOf(string(data)),
					reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			f.Set(reflect.MakeFunc(f.Type(), fn))
		}
	}
}

func PrintFiledName(val interface{}) {
	v := reflect.ValueOf(val)
	ele := v.Elem()
	t := ele.Type()

	num := t.NumField()
	fmt.Println("num:", num)
	for i:=0;i<num; i++{
		f := t.Field(i)
		fieldValue := ele.Field(i)
		if fieldValue.CanSet() {
				fmt.Printf("%s 可以被设置", f.Name)
		}
		// fmt.Println(field.Name)
		// fmt.Println(fieldValue.Name)
	}

}
func SetFiledName(val interface{}) {
	v := reflect.ValueOf(val)
	ele := v.Elem()
	t := reflect.TypeOf(val)

	numField := t.NumField()
	for i:=0;i<numField; i++{
		field := t.Field(i)
		fieldValue := ele.Field(i)
		if fieldValue.CanSet() {
			fieldValue.Set(reflect.ValueOf(func() {
				fmt.Print("你在调用方法", field.Name)
			}))
		}
	}
}




func main(){
	h := &hello {
		endpoint: "http://localhost:8080",
	}

	// msg, err := h.SayHello("golang")
	// if err  != nil {
	// 	fmt.Println("%+v", err)
	// 	return
	// }
	// println(msg)

	SetFuncFiled(h)
	data,_ := h.FuncFiled("aa")
	fmt.Print(data)

	// PrintFiledName(h)

}
