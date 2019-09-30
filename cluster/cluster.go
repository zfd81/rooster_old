package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zfd81/rooster/conf"
	"github.com/zfd81/rooster/etcd"
	"github.com/zfd81/rooster/meta"
	"go.etcd.io/etcd/clientv3"
	"strconv"
)

var (
	leaseID clientv3.LeaseID
	node    Node = Node{}
	cluster      = make(map[string]*Node)
	config       = conf.GetGlobalConfig()
)

func GetLeaseID() *clientv3.LeaseID {
	return &leaseID
}

func GetNode() *Node {
	return &node
}

func GetCluster() map[string]*Node {
	return cluster
}

func Register() error {
	cli := etcd.GetClient()
	resp, err := cli.Grant(context.TODO(), conf.GetGlobalConfig().Cluster.HeartbeatInterval)
	if err != nil {
		return err
	}

	//获得leaseID
	leaseID = resp.ID

	//续租
	ch, err := cli.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		return err
	}

	//监听租约
	go func() {
		for {
			select {
			case leaseKeepResp := <-ch:
				if leaseKeepResp == nil {
					fmt.Printf("已经关闭续租功能\n")
					return
				} else {
					fmt.Printf("续租成功\n")
					goto END
				}
			}
		END:
			//time.Sleep(500 * time.Millisecond)
		}

	}()

	path := config.Cluster.Root

	//加载现有结点
	kvs, err := etcd.GetWithPrefix(path)
	if err == nil {
		for _, kv := range kvs {
			addNode(kv.Key, kv.Value, kv.CreateRevision)
		}
	}

	//集群结点监听
	etcd.WatchWithPrefix(path, clusterWatcher)

	//结点注册
	data, err := json.Marshal(node)
	if err != nil {
		return err
	}
	path = fmt.Sprintf("%s%s%s", path, meta.Separator, strconv.FormatInt(node.StartUpTime, 10))
	_, err = cli.Put(context.TODO(), path, string(data), clientv3.WithLease(leaseID))

	return err
}

func clusterWatcher(operType etcd.OperType, key []byte, value []byte, createRevision int64, modRevision int64, version int64) {
	//index, _ := strconv.Atoi(strconv.FormatInt(createRevision, 10))
	if operType == etcd.CREATE {
		node := &Node{}
		err := json.Unmarshal(value, node)
		if err == nil {
			node.Id = createRevision
			cluster[string(key)] = node
		}
	} else if operType == etcd.MODIFY {

	} else if operType == etcd.DELETE {
		delete(cluster, string(key))
	}
}

func addNode(key []byte, value []byte, id int64) {
	node := &Node{}
	err := json.Unmarshal(value, node)
	if err == nil {
		node.Id = id
		cluster[string(key)] = node
	}
}
