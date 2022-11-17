# 项目开发过程中的查漏补缺
1. http.Handle()与http.HandleFunc()的区别?

  前者需要自行定义struct来实现Handler接口；后者则直接传写好的方法即可。
2. log.Fatalln()与panic()的区别?

前者：打印输出内容 退出应用程序 **defer函数不会执行**
后者：打印输出内容 函数立刻停止执行 (注意是函数本身，不是应用程序停止) defer函数被执行 返回给调用者(caller)

3. router,controller,handler的区别?

router：将浏览器发送的请求根据URL分发至对应的controller

controller：将请求转至对应的handler进行处理

handler：处理对应请求

4. 解析body中的json数据所使用的json包中数据解码的两种方法NewDecoder与Unmarshal有什么区别？

前者：是从一个流里面直接进行解码，代码精干，用于http连接与socket连接的读取与写入，或者文件读取；
后者：是从已存在与内存中的json进行解码，用于直接是byte的输入；
！！！相对于解码，json.NewEncoder进行大JSON的编码比json.marshal性能高，因为内部使用pool

5. 数据库的相关表之间需要设置外键吗？

设置外键的缺点如下： 
**性能影响**: 大型互联网项目或者分布式项目,进行更新操作时外键会影响数据库的性能

**强耦合**: 如果数据库中存在外键,会导致表与表之间的强耦合;其实可以再物理上删除外键;逻辑上还是存在的.

(举个例子:用户表与信用卡表;信用卡表是有user_id这个字段的;我们再查询的时候还是通过user_id字段进行查询,但是表中并不会存在外键进行关联)

**热更新**: 如果数据库存在外键,会导致新更新的代码无法运行,会有产生冲突,sql报错的隐患

**分库分表难度增加**: 外键会导致分库分表难度增加(举个例子:用户表与信用卡表之间有外键关联,如果需要做分库,由于有外键的存在,是没有办法实现分布操作的)
6. 为什么需要jwt中间件？

    在之前的web项目中，我使用的是Cookie-Session模式实现用户认证。相关流程大致如下：

    a.用户在浏览器端填写用户名和密码，并发送给服务端

    b.服务端对用户名和密码校验通过后会生成一份保存当前用户相关信息的session数据和一个与之对应的标识（通常称为session_id）

    c.服务端返回响应时将上一步的session_id写入用户浏览器的Cookie

    d.后续用户来自该浏览器的每次请求都会自动携带包含session_id的Cookie

    e.服务端通过请求中的session_id就能找到之前保存的该用户那份session数据，从而获取该用户的相关信息。
    这种方案依赖于客户端（浏览器）保存 Cookie，并且需要在服务端存储用户的session数据。

    在移动互联网时代，我们的用户可能使用浏览器也可能使用APP来访问我们的服务，我们的web应用可能是前后端分开部署在不同的端口，有时候我们还需要支持第三方登录，这下Cookie-Session的模式就有些力不从心了。
    JWT就是一种基于Token的轻量级认证模式，服务端认证通过后，会生成一个JSON对象，经过签名后得到一个Token（令牌）再发回给用户，用户后续请求只需要带上这个Token，服务端解密之后就能获取该用户的相关信息了。
7. 解析 JSON 数据时，默认将数值当做哪种类型？

    当成float64.已亲自测试（踩坑）
8. json包里使用的时候，结构体里的变量不加tag能不能正常转成json里的字段？

    a.如果变量首字母小写，则为private。无论如何不能转，因为**取不到反射信息**。
    b.如果变量首字母大写，则为public。
    不加tag，可以正常转为json里的字段，json内字段名跟结构体内字段原名**一致**。
    加了tag，从struct转json的时候，json的字段名就是**tag里的字段名**，原字段名已经没用。
9. Gorm更新字段使用updates，某字段为false，为什么没有被更新？

   `Updates` 方法支持 `struct` 和 `map[string]interface{}` 参数。当使用 `struct` 更新时，默认情况下，GORM 只会更新非零值的字段

```go
// 根据 `struct` 更新属性，只会更新非零值的字段
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 根据 `map` 更新属性
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
```
10. Gorm删除时，迁移的结构体中有DeletedAt一定是软删除吗？软删除后一定不能被查询到吗？

不是。如下情况，指定某个字段删除时，会直接删除表中的这一条记录

```go
utils.MDB.Table("articles").Where("aid=?", aid).Delete(aid).Error
```

也不是。GORM中指出使用普通的查询（Find）无法查询到，使用如下的Scan则可以查询到软删除的记录

```go
utils.MDB.Table("articles").Select("aid,author,title,cover,summary").Scan(&articleRes).Error
```

也可以使用 `Unscoped` 找到被软删除的记录

```go
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

**永久删除**

您也可以使用 `Unscoped` 永久删除匹配的记录

```go
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

11. 
