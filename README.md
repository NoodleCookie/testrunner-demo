### Testrunner 测试执行器

##### 概念

测试执行器被用于执行定义好的DSL文件（测试声明文件）。根据测试声明文件内容，执行相应的动作（如发送请求，暂停等待等）；

测试测试器被用于记录发送的请求，持久化完整的返回值；

测试执行器被用于执行断言，判断API返回值是否与断言内容匹配。

#### 测试声明（迭代中）


description.yaml
```yaml
import:
  - hello-world
```

hello-world.yaml
```yaml
stages:
  - type: api
    request:
      url: http://127.0.0.1:8884/hello
      method: GET
    response:
      status: 201 Created
    assert:
      status: /^20.* [a-zA-Z]+$/
```

#### 代码层级

```yaml
- testrunner
  - common          全局功能方法，变量
  - demo            存放测试样例以及demo代码
  - mock            测试使用的mock服务
  - test_assertion  负责执行断言的代码
  - test_report     负责报告收集、整理的代码
  - test_suite      测试用例代码 (基本结构 suite -> case -> stage)
  - test_runner     测试执行器代码
  - util            工具包
```

#### 基本结构

使用 `suite, case, stage` 来描述一套测试， 1个suite通过import字段引入n个case，1个case通过stages字段编写n个stage，testrunner通过遍历和识别root路径下所有的suite执行测试。

因此 `suite, case, stage` 都是可执行的， 都应该实现`excutable`接口，在`Execute()`中进行读取环境变量（如系统环境变量`testrunner_phase`以及测试用例变量`testcase_varibales`）完成用例的初始化以及执行测试和生成报告的逻辑。


#### Dev

构建
```shell
 go build
```

运行
```shell
 ./testrunner exec demo
```

mock服务

```shell
docker-compose -f mock/docker-compose.yml up -d
```


单元测试

测试用例维护在每个模块下的unit_test包下

```shell
go test -v ./.../unit_test
```

