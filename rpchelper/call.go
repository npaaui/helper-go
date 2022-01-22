package rpchelper

import (
	"context"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CallService(service, method string, args interface{}) (*Reply, error) {
	d, err := client.NewPeer2PeerDiscovery(GetServerAddr(service), "")
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

func GetServerAddr(service string) string {
	switch service {
	case "goods_service":
		return "tcp@127.0.0.1:8071"
	case "order_service":
		return "tcp@127.0.0.1:8074"
	}
	return ""
}
