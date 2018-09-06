package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main(){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://140.143.56.14:32964"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	resp, err := cli.Get(context.TODO(), "testkey")
	fmt.Println(resp)
}
