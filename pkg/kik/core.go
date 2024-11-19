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



