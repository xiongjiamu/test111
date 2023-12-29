# dkdns

基于golang实现一个dnsServer，若接收的请求域名解析后的IP在白名单列表中，则返回IP，否则将递归至最优的DNS server重新解析



<span style="color:#ff9900;">主要功能就是：
1、当接收到域名解析时，若解析的IP地址为白名单内IP，就返回
2、若不是白名单内IP，则直接递归至指定的DNS server 继续查询
主要应用场景：
1、本地dns解析和dns 加速用</span>


### output.txt whitelist 格式如下 ###

1.0.1.0/24
1.0.2.0/23
1.0.8.0/21
1.0.32.0/19
1.1.0.0/24
1.1.2.0/23
1.1.4.0/22
1.1.8.0/21
1.1.16.0/20
1.1.32.0/19
1.2.0.0/23
1.2.2.0/24
1.2.4.0/22
1.2.8.0/24
1.2.10.0/23
1.2.12.0/22
1.2.16.0/20
1.2.32.0/19
1.2.64.0/18
1.3.0.0/16
1.4.1.0/24
1.4.2.0/23
1.4.4.0/22
1.4.8.0/24
1.4.10.0/23
1.4.12.0/22