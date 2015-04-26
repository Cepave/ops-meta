# meta

接收ops-updater汇报上来的agent real state，返回最新的agent desired state

## 注意

- 虽然ops-meta提供了http服务，可以直接用来提供tarball下载，但是不推荐这样用，最好单独再搭建一个服务（比如nginx，如果觉得麻烦再搭一个ops-meta专门用于下载都可以）专门用于文件下载，这样ops-meta做的事情少，稳定。
