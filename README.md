# casbin

### 模型的定义（model.conf）

#### 请求定义
```
[request_definition]
r = sub, obj, act
分表是
访问实体 (Subject)，访问资源 (Object) 和访问方法 (Action)

举个栗子：
GET /users 获取用户列表
    譬如 用户sjt 要访问/users  ,GET 请求
则:

sjt    sub
/users obj
GET    act
```

#### 策略定义和角色定义
```
[policy_definition]
p = sub, obj, act

效果一样 ,是对策略的定义

[role_definition]  角色定义
g = _, _

_, _表示角色继承关系的前项和后项，即前项继承后项角色的权限(前项拥有后项的所有权限)

sjt    >= admin
admin  >= member
lisi   >= member
```

#### 生效范围的定义
```
[policy_effect]
e = some(where (p.eft == allow))
 
对policy生效范围的定义
上面表示：如果存在任意一个决策结果为allow的匹配规则，则最终决策结果为allow

p.eft就是决策结果

示例
!some(where (p.eft == deny))  表示 任何一个决策结果都不能是deny
```

#### 请求和策略的匹配规则
```
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
```

---

### 策略定义(p.csv)
``` 
g, sjt, admin
g, lisi, member
//上面这两个 代表定义了两个人：sjt和lisi ，分表角色是admin和member

p, memeber, /depts, GET
p, memeber, /depts/:id, GET
//这代表 member可以访问的 path和请求方式
p, admin, /depts, POST
p, admin, /depts/:id, PUT
p, admin, /depts/:id, DELETE
//同上 
 g, admin, member
 g,member,guest
这代表 admin同时也拥有member的关系 

```