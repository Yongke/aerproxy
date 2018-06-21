# aerproxy

使用免费的Google App Engine作为反向代理, 使用电信家庭宽带提供WEB服务。

Credit to `pokstad` for:

https://gist.github.com/pokstad/936ace2c6fc563105c17

## 问题说明
自动中国电信去年(2017年8月?)封杀了家庭网络的443端口以后，使用家里不稳定的小水管自建WEB服务，基本已经不可能了。

如何使用家里正在吃灰中Raspberry Pi自建一个HTTPS（或者HTTP）的小型WEB服务呢? 并且速度要快，时延要低。

## 解决方法

你需要准备：
* 自有域名， appengine分配的域名大概率会被DNS劫持
* 家里电信需要用外网IP， 拨打10000号找人工客服可解决（随便找个要装监控之类的理由）
* 需要一个DDNS服务(www.pubyun.com下面免费的f3322.net的就很好用)
* 你需要申请一个免费的Google App Engine， 选在日本节点（asia-northeast1）
* 你需要安装gcloud SDK（翻墙看[这里](https://cloud.google.com/sdk/gcloud/)）

## 使用步骤

* 下载本项目
    ```
    git clone https://github.com/Yongke/aerproxy    
    ```
* 修改app.yaml, 将PROXY_PASS_HOST_PORT设置为你家里暴露出来的任意可访问的地址和端口，DDNS的域名
    ```
    PROXY_PASS_HOST_PORT: 'your_ddns_domain_name:port'
    ```
* 部署该项目
    ```
    cd aerproxy && gcloud app deploy
    ```
* 登录app engine控制台， "设置 - 自定义网域 - 添加自定义网域", 绑定你的域名。 Google会先验证你对DNS域名的拥有权。
* 按照控制台的提示，去你的域名服务商配置A记录
  ```
   your_domain A 216.239.32.21 # 不通， 不要这个
   your_domain A 216.239.34.21 # 速度超快
   your_domain A 216.239.36.21 # 不通， 不要这个
   your_domain A 216.239.38.21 # 速度超快
  ```
  
现在你可以用 https://your_domain 访问你家里的WEB服务了。
Enjoy it~

## 其他说明
* Google使用let's encrypt的证书为你提供https服务
* http://your_domain会被重定向至https， 如果你不想强制使用https，可以修改app.yaml，删除"secure: always"这一行
* Google每天提供1G的流量和657084的免费urlfetch调用次数，足够绝大多数人的需求了
