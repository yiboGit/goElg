### 基于es的docker集中日志收集系统
#### 使用Fluentd收集

#### 安装Fluentd
*方式*：Fluentd经由td-agent包，包含Fluentd运行环境，所以你并不需要建立一个Ruby环境中运行Fluentd     
##### *命令：*     
>下载脚本： 
```
\curl -L http://toolbelt.treasuredata.com/sh/install-ubuntu-xenial-td-agent2.sh -o install-td-agent.sh 
```    
>运行脚本安装：
``` 
sh install-td-agent.sh
```       
>启动进程： 
```
systemctl start td-agent
```       
>可以查看是否成功： 
```
tail /var/log/td-agent/td-agent.
```      
>安装Fluentd-ElasticSearch插件：
```
td-agent-gem install fluent-plugin-elasticsearch
```

#### 修改fluent配置文件      

>文件所在地址 
```
vim /etc/td-agent/td-agent.conf 
```
>删除原本不要的文件内容，并增加
```
<source>
  @type forward
  port  24224
</source>

<match docker.**>
  @type elasticsearch
  logstash_format true
  host 172.19.0.34
  port 9200
  logstash_prefix docker
  logstash_dateformat %Y-%m-%d
  type_name docker_container_log
  include_tag_key true
  flush_interval 5s
</match>
```
    <source> 为日志信息来源 包括端口与格式
    <match> 传输目的的类型，地址端口，以及输出格式信息          
>重新启动服务进程 
``` 
sudo systemctl restart td-agent
```

#### 修改docker daemon.yml 配置文件
>
```
"log-driver": "fluentd",
  "log-opts": {
    "fluentd-address": "127.0.0.1:24224",
    "tag": "{{.ImageName}}/{{.ID}}"
   }
```
>log-driver,进行全局配置 指明docker日志驱动       
>log-opts日志的配置,fluentd-address 对日志输出地址指定；tag：日志的标签         

重新加载daemon配置文件：
```
systemctl daemon-reload
```
重启docker,使得配置生效:
```
systemctl restart docker
```
完成之后，重新启动的容器会采取这样的日志收集策略 但是旧的容器并不会

####这是我本次搭建过程，如有问题请联系我