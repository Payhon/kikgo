package kik

import "github.com/gin-gonic/gin"

// H 是 map[string]interface{} 的类型别名，类似于 gin.H
type H map[string]interface{}

func Get(engine *gin.Engine,path string, handler func(c *Context)) *gin.Engine {
	engine.GET(path, func(c *gin.Context) {
		handler(&Context{Gc: c})
	})
	return engine
}

func Post(engine *gin.Engine,path string, handler func(c *Context)) *gin.Engine {
	engine.POST(path, func(c *gin.Context) {
		handler(&Context{Gc: c})
	})
	return engine
}

type Context struct {
	Gc *gin.Context
}

func (c *Context) GetQuery(key string) string {
	return c.Gc.Query(key)
}

func (c *Context) JsonResult(code int, obj interface{}) {
	c.Gc.JSON(code, obj)
}

// Json 方法返回与 gin.H 相同类型的 map[string]interface{}
func (c *Context) Json(data map[string]interface{}) map[string]interface{} {
	return data
}

func RegisterGetRoute(engine *gin.Engine, handler interface{}) {
	// 获取函数类型
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() != reflect.Func {
		return
	}

	handlerValue := reflect.ValueOf(handler)
	// 获取函数名作为路径
	funcName := runtime.FuncForPC(handlerValue.Pointer()).Name()
	// 提取最后一个/后的函数名
	if idx := strings.LastIndex(funcName, "/"); idx >= 0 {
		funcName = funcName[idx+1:]
	}

	// 检查返回值数量
	if handlerType.NumOut() == 0 {
		// 第一种情况:无返回值,直接使用函数名作为路径
		Get(engine, "/"+funcName, func(c *Context) {
			handlerValue.Call([]reflect.Value{reflect.ValueOf(c)})
		})
	} else if handlerType.NumOut() == 2 {
		// 第二种情况:有返回值,执行函数获取路径和处理函数
		results := handlerValue.Call(nil)
		if len(results) == 2 {
			path, ok1 := results[0].Interface().(string)
			handler, ok2 := results[1].Interface().(func(*Context))
			if ok1 && ok2 {
				Get(engine, path, handler)
			}
		}
	}
}


