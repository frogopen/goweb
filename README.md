# goweb
项目目录结构的组织https://github.com/golang-standards/project-layout。<br>
单元测试的写法https://learnku.com/docs/learn-go-with-tests。<br>
错误处理https://www.cnblogs.com/qcrao-2018/p/11538387.html。<br>
## 模块拆分
* 很多其它面向对象的编程语言都非常推崇MVC的架构模式，但在go语言中最好是按照职责对模块进行拆分。例如一个博客系统会有user、comment、article等模块，每一个模块都应对外提供相应的功能，如user模块就应包含相关的模型以及处理API请求的服务。
* go的每个文件目录都代表一个独立的命名空间，如果采用MVC架构的话，通常会有model、view、和controller。如在controller中使用model中的方法，会有大量重复代码，如model.insert、model.delete、model.xxx这样的话代码就会很冗余。所以应该遵循顶层的设计用这种方式进行拆分。而且这样划分的话在遇到瓶颈时非常有利于转微服务架构。根据已经拆分好的模块进行分布式部署。
## 待续...
