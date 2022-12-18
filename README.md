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
    response:
      status: 201 Created
    assertion:
      equals:
        status: /^20.* [a-zA-Z]+$/
```

#### 代码层级

```yaml
- testrunner
  - common          全局功能方法，变量
  - mock            测试使用的mock服务
  - test_assertion  测试断言代码
  - test_case       测试用例代码
  - test_runner     测试执行器代码
  - util            工具包
```

#### Dev

构建
```shell
 go build
```

运行
```shell
 ./testrunner
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

