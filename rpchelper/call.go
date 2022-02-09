package rpchelper

import (
	"context"

	etcdC "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CallService(service, method string, args interface{}) (*Reply, error) {
	d, err := etcdC.NewEtcdDiscovery("mall", service, []string{"127.0.0.1:2379"}, true, nil)
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON
	xClient := client.NewXClient(service, client.Failtry, client.RandomSelect, d, opt)
	defer func() {
		d.Close()
		err = xClient.Close()
		if err != nil {
			panic("rpc client close err")
		}
	}()

	reply := &Reply{}
	err = xClient.Call(context.Background(), method, args, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
