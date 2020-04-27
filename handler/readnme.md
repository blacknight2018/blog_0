### 【handler】服务层
#### 所有的handler在处理完数据后都要设置要context的【ContextFiledName】域，如果要返回对象给也是设置这个域
##### context.Set(configure.ContextFiledName, utils.JsonGoParseWithThrowException(&node))   解析成json后设置这个域
##### utils.SetSuccessRetObjectToJSONWithThrowException(context, ret)                       这个函数会解析这个对象再设置这个域
##### handler所有函数也是四大原子操作开头
#### handler层抛出异常给 Except 捕获 Except 层处理后给 RequestMiddle层捕获返回给前端 RequestMiddle 是最后一层
#### out.go 是每一层服务层提供给别的服务层调用的接口
#### 每一层的dao只会跟每一层对应的ser打关系，out是提供给别的层的接口
#### utils 函数一律 get set开头 