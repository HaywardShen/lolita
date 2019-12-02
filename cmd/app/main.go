package main

import (
	"sync"
	"t-mk-opentrace/cmd/grpc"
	"t-mk-opentrace/cmd/http"
	ant "t-mk-opentrace/ext/ants-driver/ants"
	"t-mk-opentrace/ext/log-driver/log"
)

// main main
func main() {
	// tracer, conn := middleware.NewOpenTraceClient()
	_, closer := NewOpenTraceClient()
	defer func() {
		if err := closer.Close(); err != nil {
			log.Warn("opentracer collecter 关闭失败", err)
		}
	}()
	defer ant.Release()
	group := sync.WaitGroup{}
	defer group.Wait()
	group.Add(1)
	ant.Go(func() {
		if err := http.GinInitServer(); err != nil {
			panic(err)
		}
		group.Done()
	})
	group.Add(1)
	ant.Go(func() {
		if err := grpc.RPCServer(); err != nil {
			panic(err)
		}
		group.Done()
	})
	ant.Go(func() {
		CleanUp()
	})
}