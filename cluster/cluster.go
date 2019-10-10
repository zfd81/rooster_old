package cluster

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zfd81/rooster/conf"
	"github.com/zfd81/rooster/etcd"
	"github.com/zfd81/rooster/meta"
	"go.etcd.io/etcd/clientv3"
	"net"
	"strconv"
)

var (
	leaseID clientv3.LeaseID
	node    Node = Node{}
	members      = make(map[string]*Node)
	config       = conf.GetGlobalConfig()
)

func GetLeaseID() *clientv3.LeaseID {
	return &leaseID
}

func GetNode() *Node {
	return &node
}

func GetMembers() map[string]*Node {
	return members
}

func Register(startUpTime int64) error {
	ip, err := externalIP()
	if err != nil {
		return err
	}
	node.Address = ip.String()
	node.Port = config.Http.Port
	node.StartUpTime = startUpTime
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
	path = fmt.Sprintf("%s%s%s:%s", path, meta.Separator, node.Address, strconv.Itoa(node.Port))

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
			members[string(key)] = node
		}
	} else if operType == etcd.MODIFY {
	} else if operType == etcd.DELETE {
		delete(members, string(key))
	}
}

func addNode(key []byte, value []byte, id int64) {
	node := &Node{}
	err := json.Unmarshal(value, node)
	if err == nil {
		node.Id = id
		members[string(key)] = node
	}
}

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}
	return ip
}
