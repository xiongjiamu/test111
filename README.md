# dkdns

基于golang实现一个dnsServer，若接收的请求域名解析后的IP在白名单列表中，则返回IP，否则将递归至最优的DNS server重新解析



<span style="color:#ff9900;">主要功能就是：
1、当接收到域名解析时，若解析的IP地址为白名单内IP，就返回
2、若不是白名单内IP，则直接递归至指定的DNS server 继续查询
主要应用场景：
1、本地dns解析和dns 加速用</span>