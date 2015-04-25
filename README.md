# meta

接收ops-updater汇报上来的agent real state，返回最新的agent desired state

## FAQ

- **如何支持某个agent小流量测试？**

针对某个agent，比如falcon-agent，线上很可能出现两个版本，一个是当前大部分机器使用的版本，一个是新版的小流量版本。做得更通用一些就是可以支持某个agent在线上存在多个版本。为了使系统依赖性小些，先期姑且采用strings.contains方法来判断当前机器应该使用哪个agent的哪个版本