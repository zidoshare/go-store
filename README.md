# go-store

[拙劣的功能概览](https://github.com/zidoshare/go-store/blob/master/doc/%E5%8A%9F%E8%83%BD%E6%A6%82%E8%A7%88.md)
golang实现的微商店，目前正处于开发中...

为什么开发？

最重要的是为了学习golang，其他原因包括：对大面积的外包性类似项目不满意/技术老旧等，类似应用缺乏golang相关，练手，学习docker，试验react-ssr等。。。

## 主要技术


服务端使用golang，restful api风格，使用mux/gorm框架。前端使用自己搭建的[react-ssr-starter](https://github.com/zidoshare/react-ssr-starter)脚手架，服务端渲染，react框架。

## 开发

项目大概分为三大块，分别为服务端/商店客户端/管理客户端，其中两个客户端通过submodule集合在了本项目里。

为了更好的运行开发，我写了[shell脚本(go-store.sh)](https://github.com/zidoshare/go-store/blob/master/go-store.sh)，包含了一些常用操作。

一般来说只需要手动安装golang/nodejs/yarn，其他就交给shell好了。

|命令|功能|备注|
|---|---|----|
|sub|更新submodule|简单的检查go-store-client和go-store-admin-client是否有文件，没有文件就会初始化子模块|
|run:server|启动golang服务端||
|run:client|启动商店客户端||
|run:client:admin|启动管理客户端||
