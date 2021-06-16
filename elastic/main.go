package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/olivere/elastic/v7"
	"reflect"
	"time"
)

type User struct {
	Name     string                `json:"name"`
	Age      int                   `json:"age"`
	Married  bool                  `json:"married"`
	Sex      string                `json:"sex"`
	Created  time.Time             `json:"created, "`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

var ctx = context.Background()
var esUrl string = "http://localhost:9200"

func main() {

	//连接客户端
	client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	// Ping the Elasticsearch server to get e.g. the version number
	// ping通服务端，并获得服务端的es版本,本实例的es版本为version 7.6.1
	info, code, err := client.Ping(esUrl).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code>: %d and version %s\n", code, info.Version.Number)
	// 获取版本号的直接API
	esVersion, err := client.ElasticsearchVersion(esUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("es的版本为%s\n", esVersion)

	// 创建index前，先查看es引擎中是否存在自己想要创建的索引index
	exists, err := client.IndexExists("user").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		// 如果不存在，就创建
		createIndex, err := client.CreateIndex("user").BodyString(mapping1).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged ,创建失败
		}
	}

	//为已有的索引添加字段
	_, err = client.PutMapping().Index("user").BodyString(mapping1).Do(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 添加文档方法1
	//user1 := User{Name:"bob",Sex:"male",Married:false,Age:23}
	//put1,err :=client.Index().Index("user").BodyJson(user1).Id("1").Do(ctx)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type) //Indexed user 1 to index user, type _doc

	//添加文档方法2
	//user2 := `{"name":"mike","sex":"male","married":true,"age":"22"}`
	//put2, err := client.Index().Index("user").BodyString(user2).Do(ctx)// 不指定id
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Printf("Indexed user %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)//Indexed user 4-K2wXIB33YuyEzPYoAi to index user, type _doc

	// 查询
	get1, err := client.Get().Index("user").Id("1").Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		// Got document 1 in version 824633838776 from index user, type _doc
	}

	// Flush to make sure the documents got written.将文档涮入磁盘
	//_, err = client.Flush().Index("user").Do(ctx)
	//if err != nil {
	//	panic(err)
	//}

	// 按"term"搜索Search with a term query
	termQuery := elastic.NewTermQuery("name", "mike")
	searchResult, err := client.Search().
		Index("user").     // 搜索的索引"user"
		Query(termQuery).  // specify the query
		Sort("age", true). //按字段"age"排序，升序排列
		From(0).Size(10).  // 分页，单页显示10条
		Pretty(true).      // pretty print request and response JSON以json的形式返回信息
		Do(ctx)            // 执行
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis) // Query took 3 milliseconds
	var user User
	//Each是一个简便函数，此函数忽略了错误输出
	for _, item1 := range searchResult.Each(reflect.TypeOf(user)) {
		if u, ok := item1.(User); ok {
			fmt.Printf("Person by %s,age:%d,married:%t,Sex:%s\n", u.Name, u.Age, u.Married, u.Sex) //Person by bob,age:23,married:false,Sex:male
		}
	}
	// 搜索文档方法2
	// 使用hits，获得更详细的输出结果
	if searchResult.Hits.TotalHits.Value > 0 {
		fmt.Printf("找到的数据总数是 %d \n", searchResult.Hits.TotalHits.Value)
		for _, hits := range searchResult.Hits.Hits {
			u := User{}
			err := json.Unmarshal([]byte(hits.Source), &u)
			if err != nil {
				fmt.Println("反序列化失败", err)
			}
			fmt.Printf("User by %s,age:%d,married:%t,Sex:%s\n", u.Name, u.Age, u.Married, u.Sex)
		}
	} else {
		fmt.Println("没有搜到用户")
	}

	// 更新文档 update
	//update, err := client.Update().Index("user").Id("1").
	//	Script(elastic.NewScriptInline("ctx._source.age += params.num").Lang("painless").Param("num", 1)).
	//	//Upsert(map[string]interface{}{"created": "2020-06-17"}). // 插入未初始化的字段value
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("New version of user %q is now %d\n", update.Id, update.Version)
	// 更新方法2
	//update,err := client.Update().Index("user").Id("1").
	//	Script(elastic.NewScriptInline("ctx._source.created=params.date").Lang("painless").Param("date","2020-06-17")).
	//	Do(ctx)
	termQuery := elastic.NewTermQuery("name", "bob")
	update, err := client.UpdateByQuery("user").Query(termQuery).
		Script(elastic.NewScriptInline("ctx._source.age += params.num").Lang("painless").Param("num", 1)).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("New version of user %q is now %d\n", update.Id, update.Version)
	fmt.Println(update)
	// 删除文档

	//termQuery := elastic.NewTermQuery("name", "mike")
	//_, err = client.DeleteByQuery().Index("user"). // search in index "user"
	//	Query(termQuery). // specify the query
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
}
