package main

import (
	"base-server/common"
	"base-server/config"
	"base-server/logger"
	"base-server/router"
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/judwhite/go-svc/svc"
	_ "net/http/pprof"
)

type BaseServer struct {
	server *http.Server
}

func (s *BaseServer) Init(env svc.Environment) error {
	runtime.GOMAXPROCS(1)

	config.InitConfig()
	println("RunMode:", config.CURMODE)
	for key, val := range config.GetSection("system") {
		println(key, val)
	}

	// init log
	logger.ConfigLogger(
		config.CURMODE,
		config.GetConfig("system", "app_name"),
		config.GetConfig("logs", "dir"),
		config.GetConfig("logs", "file_name"),
		config.GetConfigInt("logs", "keep_days"),
		config.GetConfigInt("logs", "rate_hours"),
	)
	println("logger init success")

	// init mysql
	dbInfo := config.GetSection("dbInfo")
	for name, info := range dbInfo {
		if err := common.AddDB(
			name,
			info,
			config.GetConfigInt("mysql", "maxConn"),
			config.GetConfigInt("mysql", "idleConn"),
			time.Hour*time.Duration(config.GetConfigInt("mysql", "maxLeftTime"))); err != nil {
			return err
		}
	}
	println("mysql init success")

	// init redis
	if err := common.AddRedisInstance(
		"",
		config.GetConfig("redis", "addr"),
		config.GetConfig("redis", "port"),
		config.GetConfig("redis", "password"),
		config.GetConfigInt("redis", "db_num")); err != nil {
		return err
	}
	println("redis init success")

	return nil
}

func (s *BaseServer) Start() error {
	s.server = &http.Server{
		Addr:    ":8087",
		Handler: router.NewEngine(),
	}
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()
	println("http service start, listen on 8087")

	go http.ListenAndServe(":9999", nil)
	println("pprof service start, listen on 9999")


	return nil
}

func (s *BaseServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		println(err.Error())
	}

	// release source
	common.ReleaseMysqlDBPool()
	common.ReleaseRedisPool()

	return nil
}

func main() {
	if err := svc.Run(&BaseServer{}); err != nil {
		println(err.Error())
	}

	http.Get("www.baidu.com")
}
