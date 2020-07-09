package router

import (
	"base-server/middleware"
	"time"

	"runtime/pprof"
	"runtime/trace"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// use middleware here
	engine.Use(middleware.Recover)
	engine.Use(middleware.RequestStart)
	engine.Use(middleware.RequestOut)

	// router here
	engine.Any("/", HealthCheck)

	v1 := engine.Group("/debug")
	v1.Any("/goroutine", goroutineProfile)
	v1.Any("/thread", threadProfile)
	v1.Any("/heap", heapProfile)
	v1.Any("/allocs", allocsProfile)
	v1.Any("/block", blockProfile)
	v1.Any("/mutex", mutexProfile)
	v1.Any("/trace", traceOut)
	return engine
}

func HealthCheck(c *gin.Context) {
	c.String(200, "hello world")
	return
}

func goroutineProfile(c *gin.Context) {
	p := pprof.Lookup("goroutine")
	p.WriteTo(c.Writer, 1)
	//c.String(200, fmt.Sprintf("profile:%+v", p))
	return
}

func threadProfile(c *gin.Context) {
	p := pprof.Lookup("threadcreate")
	p.WriteTo(c.Writer, 1)
	//c.String(200, fmt.Sprintf("profile:%+v", p))
	return
}

func heapProfile(c *gin.Context) {
	p := pprof.Lookup("heap")
	p.WriteTo(c.Writer, 1)
	//c.String(200, fmt.Sprintf("profile:%+v", p))
	return
}

func allocsProfile(c *gin.Context) {
	p := pprof.Lookup("allocs")
	p.WriteTo(c.Writer, 1)
	//c.String(200, fmt.Sprintf("profile:%+v", p))
	return
}

func blockProfile(c *gin.Context) {
	p := pprof.Lookup("block")
	p.WriteTo(c.Writer, 1)
	//c.String(200, fmt.Sprintf("profile:%+v", p))
	return
}

func mutexProfile(c *gin.Context) {
	p := pprof.Lookup("mutex")
	p.WriteTo(c.Writer, 1)
	//c.String(200, fmt.Sprintf("profile:%+v", p))
	return
}

func traceOut(c *gin.Context) {
	trace.Start(c.Writer)
	time.Sleep(10 * time.Second)
	trace.Stop()
	return
}
