### ElasticSearch Docker的单点以及集群部署

#### 单节点部署
```
docker run --name myelastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:5.6.1 
```

##### 使用本地的ElasticSearch配置文件Docker部署：
```
docker run --name myelastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -v /root/elasticsearch.yaml:/usr/share/elasticsearch/config/elasticsearch.yml docker.elastic.co/elasticsearch/elasticsearch:5.6.1
```

*查看当前ElasticSearch的健康度*
curl -u elastic http://127.0.0.1:9200/_cat/health

#### 集群部署      

此次集群部署，我决定部署两个节点，一个节点为master节点，一个节点为data节点，部署一个最小集群。

*master节点local配置文件：*  
```    
cluster.name: "eglass-es-cluster"
node.name: myes1
node.master: true
node.data: true
network.host: 0.0.0.0
network.publish_host: 172.19.0.34
discovery.zen.ping.unicast.hosts: ["172.19.157.62","172.19.0.34"]
discovery.zen.minimum_master_nodes: 1  
```      
          
*data节点local配置文件：*
```
cluster.name: "eglass-es-cluster"
node.name: myes2
node.master: false
node.data: true
network.host: 0.0.0.0
network.publish_host: 172.19.157.62
discovery.zen.ping.unicast.hosts: ["172.19.0.34","172.19.0.34"]      
```
Docker 部署：  

*master节点：*  
```
docker run -d --name=myes1 -p 9200:9200 -p 9300:9300  -v /usr/share/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml elasticsearch:5.6.1
```

*data节点*
```
docker run -d --name=myes2 -p 9200:9200 -p 9300:9300  -v /usr/share/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml  elasticsearch:5.6.1  
```      

*查看集群节点情况：*      
curl -u elastic http://127.0.0.1:9200/_cat/nodes     

*查看集群健康度：*    
 curl -u elastic http://127.0.0.1:9200/_cat/health     

 *查看集群内数据分片情况：*      
 curl -u elastic http://127.0.0.1:9200/_cat/shards     






