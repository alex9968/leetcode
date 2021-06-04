package main


import (
"context"
"fmt"
"github.com/olivere/elastic"
)

// 创建索引：
func main(){
	Client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	fmt.Println(Client, err)
	name := "people2"
	Client.CreateIndex(name).Do(context.Background())
	Insert()
	Select()
}

// 插入数据
func Insert(){
	Client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	fmt.Println(Client, err)
	name := "people2"
	data := `{
    "name": "wali",
    "country": "Chian",
    "age": 30,
    "date": "1987-03-07"
    }`
	_, err = Client.Index().Index(name).Type("man1").Id("1").BodyJson(data).Do(context.Background())

}

// 查找数据： //通过id查找
func Select(){
	Client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	fmt.Println(Client, err)
	name := "people2"
	get, err := Client.Get().Index(name).Type("man1").Id("1").Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n",get)

}



//修改
func Update() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	res, err := client.Update().
		Index("megacorp").
		Type("employee").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}
// 删除数据
func Delete(){
	type Employee interface{}
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.7.6:9200"))

	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	//创建
	put1, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	//删除
	get, err := client.Get().Index("megacorp").Type("employee").Id("1").Do(context.Background())
	fmt.Println(get, err)
}