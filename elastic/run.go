package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	Host     = "https://10.20.156.102:9200,https://10.20.156.102:9201"
	Username = "elastic"
	Password = "muniulab"
)

var ES2 *elastic.Client

func init() {
	//连接客户端
	var err error
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: 1 * time.Minute,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	ES2, err = elastic.NewClient(
		elastic.SetHttpClient(httpClient),
		elastic.SetSniff(false),
		elastic.SetURL(strings.Split(Host, ",")...),
		elastic.SetBasicAuth(Username, Password),
	)
	if err != nil {
		log.Panicf("[InitEs] ES init fail, err: " + err.Error())
		return
	}
	// 获取版本号的直接API
	var esVersion string
	esVersion, err = ES2.ElasticsearchVersion(strings.Split(Host, ",")[0])
	if err != nil {
		log.Panicf("[InitEs] ES init fail, err: " + err.Error())
		return
	}
	log.Printf("[InitEs] init ES2 success,ES2 info %s", esVersion)
}

func main() {

}

// FieldDetailBucket 根据ip、uri等获取访问详情
func FieldDetailBucket2(req EsReq) (respBucket *RespBucket, err error) {
	bq := elastic.NewBoolQuery()
	for k, v := range req.Fields {
		bq.Must(elastic.NewMatchQuery(k, v))
	}
	//bq.Filter(elastic.NewRangeQuery())
	bq.Must(elastic.NewRangeQuery(Time).Gte(req.StartTime).Lte(req.EndTime))

	agg := elastic.NewTermsAggregation()
	for _, v := range req.SubFieldName {
		agg.Field(v)
	}
	res, err := ES2.Search().
		Index(req.Indexs...).
		Query(bq).
		Aggregation("group", agg).
		From(int(req.From)).Size(int(req.Size)).
		Pretty(true).            // pretty print request and response JSON以json的形式返回信息
		Do(context.Background()) // 执行
	if err != nil {
		return nil, err
	}
	group := Group{}
	if err = json.Unmarshal(res.Aggregations["group"], &group); err != nil {
		return nil, err
	}

	return &RespBucket{
		TotalDocCount: res.Hits.TotalHits.Value,
		Took:          int(res.TookInMillis),
		Group:         group,
	}, nil
}

// GetBucketCount2  获取某个field bucket的数量
func GetBucketCount2(req EsReq) (count int64, hitsCount int64, err error) {
	var res *elastic.SearchResult
	bq := elastic.NewBoolQuery()
	boolFilter(req, bq)
	if len(req.Fields) != 0 {
		for k, v := range req.Fields {
			bq.Must(elastic.NewMatchQuery(k, v))
		}
	}
	agg := elastic.NewCardinalityAggregation().Field(req.SubFieldName[0])
	res, err = ES2.Search(req.Indexs...).
		Query(bq).
		Size(0).
		Aggregation("group_count", agg).
		Do(context.Background())
	if err != nil {
		return
	}
	v, _ := res.Aggregations.ValueCount("group_count")
	count = int64(*v.Value)
	hitsCount = res.Hits.TotalHits.Value
	return
}

// ListFieldBucketRank 获取某个字段 bucket排名
func ListFieldBucketRank2(req EsReq) (respBucket *RespBucket, err error) {
	bq := elastic.NewBoolQuery()
	boolFilter(req, bq)
	for k, v := range req.Fields {
		bq.Must(elastic.NewMatchQuery(k, v))
	}

	if len(req.SubFieldName) == 0 {
		return
	}
	agg := elastic.NewTermsAggregation().Field(req.SubFieldName[0]).Size(int(req.Count)).
		SubAggregation("page", elastic.NewBucketSortAggregation().From(0).Size(int(req.Size)))

	var res *elastic.SearchResult
	res, err = ES2.Search().
		Index(req.Indexs...).
		Query(bq).
		Aggregation("group", agg).
		From(0).Size(20).
		Pretty(true).
		Do(context.Background()) // 执行
	if err != nil {
		return
	}
	group := Group{}
	jsonGroup, _ := res.Aggregations["group"].MarshalJSON()
	err = json.Unmarshal(jsonGroup, &group)
	if err != nil {
		return
	}
	respBucket = &RespBucket{
		//数据总量， 只和索引有关，和过滤条件无关
		TotalDocCount: res.Hits.TotalHits.Value,
		Took:          int(res.TookInMillis),
		Group:         group,
	}
	return
}

// GetOneInfo 通过关键字获取一条信息, 然后获取地理位置信息
func GetOneInfo2(req EsReq) (map[string]interface{}, error) {
	var (
		mq *elastic.MatchQuery
	)
	for k, v := range req.Fields {
		mq = elastic.NewMatchQuery(k, v)
	}
	res, err := ES2.Search().
		Index(req.Indexs...).
		Query(mq).
		Size(1).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		logger.ErrorF("[GetValueCount][%s] %s", res.Status, err.Error())
		return nil, err
	}
	hits := res.Hits.Hits
	if len(hits) == 0 {
		return nil, err
	}
	source := hits[0].Source
	body := map[string]interface{}{}
	err = json.Unmarshal(source, &body)
	if err != nil {
		return nil, err
	}
	return body["location"].(map[string]interface{}), nil
}

func boolFilter(req EsReq, bq *elastic.BoolQuery) {
	if req.StartTime != 0 {
		bq.Filter(elastic.NewRangeQuery(Time).Gte(req.StartTime))
	}
	if req.EndTime != 0 {
		bq.Filter(elastic.NewRangeQuery(Time).Lte(req.EndTime))
	}
	if req.IsAttack {
		bq.Filter(elastic.NewExistsQuery(Attack))
	}
}

//func GetWebBucket() (interface{}, error) {
//	// 按"term"搜索Search with a term query
//	//termQuery := elastic.NewTermQuery("name", "mike")
//	bq := elastic.NewBoolQuery()
//	for k,v := range fields{
//		bq.Must(elastic.NewMatchQuery(k, v))
//	}
//	bq.Must(elastic.NewRangeQuery(Time).Gte(startTime).Lte(endTime))
//	res, err := ES2.Search().
//		Index(indexs...).
//		Query(bq).
//		Sort("age", true). //按字段"age"排序，升序排列
//		From(0).Size(10).
//		Pretty(true).            // pretty print request and response JSON以json的形式返回信息
//		Do(context.Background()) // 执行
//	if err != nil {
//
//	}
//	var weblog WebLog
//	//Each是一个简便函数，此函数忽略了错误输出
//	for _, item1 := range res.Each(reflect.TypeOf(weblog)) {
//		if w, ok := item1.(WebLog); ok {
//			fmt.Printf("WebLog: %#v", w)
//		}
//	}
//	return nil, nil
//}
//
//// 统计数量
//func Count(field string, indexs []string) (count int64, err error) {
//	cq := elastic.NewConstantScoreQuery(elastic.NewExistsQuery(field))
//	count, err = ES2.Count(indexs...).Query(cq).Do(context.Background())
//	if err != nil {
//	}
//	return
//}

//func GetHitsHits(req EsReq) {
//if hitSize != 0 {
//	hits := make([]HitWithId, 0)
//	for _, item := range res.Each(reflect.TypeOf(WebLog{})) {
//		if val, ok := item.(WebLog); ok {
//			hits  = append(hits, HitWithId{val.ID, nil})
//		}
//	}
//	respHit = &RespHitWithId{
//		Took: int(res.TookInMillis),
//		Hits: hits,
//	}
//}
