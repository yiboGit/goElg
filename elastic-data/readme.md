## ElasticSearch 的使用

#### 新建有ik中文分词的索引
*格式：*    
http://elastic:yourPassword@< yourHost >:< yourPort >/yourIndex    
请求类型POST，请求体：索引结构映射 
eg:    
http://elastic:changeme@101.132.73.244:9200/msku
Body:
```    
{    
 "mappings": {
    "glass": {
      "properties": {
        "brand": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_max_word"
        },
        "name": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_max_word"
        },
        "refraction":{
        	"type":"text",
        	"index": "analyzed"
        }
      }
    }
  }
}  
```   
此处映射包含了，index(索引)下的指定的type（类别）。Document（文档）的字段是否要要进行分词分析，指定的分词器，若指定ik分词器，设定分词策略等操作。    

#### 删除指定索引
*格式：*    
http://elastic:yourPassword@< yourHost >:< yourPort >/yourIndex     
请求类型：DELETE
eg:     
http://elastic:changeme@101.132.73.244:9200/member    
删除索引与索引下的数据     

#### 新增记录数据 
*格式：*     
http://elastic:yourPassword@< yourHost >:< yourPort >/yourIndex/yourType/id    
请求类型：PUT，请求体：要新增的记录     
eg：      
http://elastic:changeme@101.132.73.244:9200/msku/glass/1      
Body:
```     
{
    "brand": "蔡司",
    "id": 3042,
    "name": "1.6新三维博锐（钻立方晶彩膜）",
    "refraction": 1.6
}   
```    
**新增记录时候，也可以不指定id,但此时请求类型要改为POST**    
         
#### 删除记录数据      
*格式：*
http://elastic:yourPassword@< yourHost >:< yourPort >/yourIndex/yourType/id     
请求类型：DELETE       

#### 查询记录数据   
*格式：*
http://elastic:yourName@ < yourHost > : < yourPort> /yourIndex/yourType/_search
请求格式：POST，Body：查询字句
eg:         
http://elastic:changeme@101.132.73.244:9200/member/optometryUser/_search
Body:
```      
{
  "size": 100,
  "query": {
    "bool" : {
     "filter": {
      	"match":{"third_id":"201"}
      },
      "should" : [
        { "match" : { "name" : "156" } },
        { "match" : { "remark" : "156" } },
        {"wildcard": {"phone_wx":"* 156 *"} }
      ],
       "minimum_should_match" : 1
    }
  }
}
```

query，bool字句这里filter为必然要遵守的过滤条件，should内条件可以理解为or关系，字句间关系为and，minimum_should_match表示should中条件至少有一项必须满足，wildcard为模糊查询，match为精确过滤，size:每次查询的容量。       
#### 修改mapping,添加字段(不可以修改字段)
*格式：*       
http://elastic:yourPassword@< yourHost >:< yourPort >/yourIndex/yourType/_mapping?pretty       
请求类型：POST， 请求体：添加的字段
eg:      
http://elastic:changeme@101.132.73.244:9200/member/optometryUser/_mapping?pretty         
Body:       
```
{
	"testmap":{
		"properties":{
			"address":{
				"type":"text",
				"analyzer":"keyword"
			}
		}
	}
}
```


### 目前使用到的查询请求：

#### msku中：

请求类型：POST
http://elastic:changeme@101.132.73.244:9200/msku/glass/_search   
Body : 
```       
{
  "size": 100,
  "query": {
    "bool" :{
      "should" : [
        { "match" : { "name" : "蔡司新三维博锐1.5" } },
        { "match" : { "brand" : "蔡司新三维博锐1.5" } },
        {"wildcard": {"refraction":"*蔡司新三维博锐1.5*"} }
      ],
       "minimum_should_match" : 1
    }
  }
}
```

#### member中：     
请求类型：POST      
http://elastic:changeme@101.132.73.244:9200/member/optometryUser/_search     
Body:   
```     
{
  "size": 100,
  "query": {
    "bool" : {
     "filter": {
      	"match":{"third_id":"201"}
      },
      "should" : [
        { "match" : { "name" : "156" } },
        { "match" : { "remark" : "156" } },
        {"wildcard": {"phone_wx":"*156*"} }
      ],
       "minimum_should_match" : 1
    }
  }
}
```


