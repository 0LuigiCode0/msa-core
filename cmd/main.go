package main

import (
	"fmt"

	"github.com/0LuigiCode0/msa-core/helper"
	"github.com/0LuigiCode0/msa-core/observer/client"
	"github.com/0LuigiCode0/msa-core/observer/server"
)

func main() {
	observer, err := server.NewObserverServer("tets", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	helper.Wg.Add(1)
	go func() {
		defer helper.Wg.Done()
		if err = observer.Start(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	helper.Wg.Add(1)
	go func() {
		defer helper.Wg.Done()
		clientObserver, err := client.NewObserverClient("127.0.0.1:8090")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("connect")
		clientObserver.Close()
	}()

	<-helper.C
	observer.Close()
	helper.CloseCtx()
	helper.Wg.Wait()
}
