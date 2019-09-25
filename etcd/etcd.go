package etcd

import (
	"context"
	"github.com/zfd81/rooster/conf"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

var (
	config = conf.NewConfig()
	ctx    = context.TODO()
	client *clientv3.Client
	lease  clientv3.Lease
)

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Etcd.Endpoints,
		DialTimeout: time.Duration(config.Etcd.DialTimeout) * time.Second,
	})
	if err != nil {

	}
	client = cli
	lease = clientv3.NewLease(cli)
}

func GetClient() *clientv3.Client {
	return client
}

func Put(key, value string) (revision int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Etcd.RequestTimeout)*time.Second)
	resp, err := client.Put(ctx, key, value)
	cancel()
	if err != nil {
		return -1, err
	}
	return resp.Header.Revision, nil
}

func Del(key, value string) (revision int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Etcd.RequestTimeout)*time.Second)
	resp, err := client.Delete(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return -1, err
	}
	return resp.Header.Revision, nil
}

func Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Etcd.RequestTimeout)*time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) < 1 {
		return nil, nil
	}
	return resp.Kvs[0].Value, nil
}

func GetWithPrefix(prefix string) ([]*mvccpb.KeyValue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Etcd.RequestTimeout)*time.Second)
	resp, err := client.Get(ctx, prefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}
	return resp.Kvs, nil
}

//func Grant(ttl int64) error {
//	//设置租约时间
//	resp, err := lease.Grant(context.TODO(), ttl)
//	if err != nil {
//		return err
//	}
//}
