package meta

import (
	"encoding/json"
	"fmt"
	"github.com/zfd81/rooster/conf"
	"github.com/zfd81/rooster/etcd"
	"strings"
)

type Meta map[string]*Instance
type SchemaState byte
type EventType int32

const (
	// StateNone means this schema element is absent and can't be used.
	StateNone SchemaState = iota
	// StateDeleteOnly means we can only delete items for this schema element.
	StateDeleteOnly
	// StateWriteOnly means we can use any write operation on this schema element,
	// but outer can't read the changed data.
	StateWriteOnly
	// StateWriteReorganization means we are re-organizing whole data after write only state.
	StateWriteReorganization
	// StateDeleteReorganization means we are re-organizing whole data after delete only state.
	StateDeleteReorganization
	// StatePublic means this schema element is ok for all write and read operations.
	StatePublic

	CREATE EventType = 0
	MODIFY EventType = 1
	REMOVE EventType = 2

	Separator = "/" // 路径分隔符（分隔路径元素）
)

var (
	meta   = make(Meta)
	config = conf.NewConfig()
)

func (s SchemaState) String() string {
	switch s {
	case StateDeleteOnly:
		return "delete only"
	case StateWriteOnly:
		return "write only"
	case StateWriteReorganization:
		return "write reorganization"
	case StateDeleteReorganization:
		return "delete reorganization"
	case StatePublic:
		return "public"
	default:
		return "none"
	}
}

func parsePath(path string) (iname string, dname string, tname string, suffix string) {
	strs := strings.Split(path[len(config.Meta.Root)+1:], "/")
	length := len(strs)
	if length == 3 {
		iname, _ = parseName(strs[0])
		dname, _ = parseName(strs[1])
		tname, suffix = parseName(strs[2])
	} else if length == 2 {
		iname, _ = parseName(strs[0])
		dname, suffix = parseName(strs[1])
	} else if length == 1 {
		iname, suffix = parseName(strs[0])
	}
	return
}

func parseName(mname string) (name string, suffix string) {
	index := strings.LastIndex(mname, ".")
	return mname[:index], mname[index:]
}

func GetMeta() Meta {
	return meta
}

func CreateInstance(name string) *Instance {
	ins := &Instance{
		InstanceInfo: InstanceInfo{
			Name: name,
			Text: name,
		},
		Databases: make(map[string]*Database),
	}
	meta[name] = ins
	return ins
}

func CreateInstanceWithInfo(info InstanceInfo) *Instance {
	ins := &Instance{
		InstanceInfo: info,
		Databases:    make(map[string]*Database),
	}
	meta[info.Name] = ins
	return ins
}

func FindInstance(name string) *Instance {
	return meta[name]
}

func LoadInstance(name string) error {
	ins := &Instance{}
	ins.Name = name
	path := fmt.Sprintf("%s%s%s", config.Meta.Root, Separator, ins.GetMName())
	data, err := etcd.Get(path)
	if data != nil {
		err = json.Unmarshal(data, ins)
		if err == nil {
			meta[name] = ins
		}
	}
	return err
}

func metaWatcher(operType etcd.OperType, key []byte, value []byte, createRevision int64, modRevision int64, version int64) {
	iname, dname, tname, suffix := parsePath(string(key))
	if operType == etcd.CREATE {
		switch suffix {
		case config.Meta.TableSuffix:
			tbl := meta[iname].GetDatabase(dname).CreateTable(tname)
			tbl.Load(value)
			break
		case config.Meta.DatabaseSuffix:
			db := meta[iname].CreateDatabase(dname)
			db.Load(value)
			break
		case config.Meta.InstanceSuffix:
			ins := CreateInstance(iname)
			ins.Load(value)
			break
		}
	} else if operType == etcd.MODIFY {
		switch suffix {
		case config.Meta.TableSuffix:
			tbl := meta[iname].GetDatabase(dname).GetTable(tname)
			tbl.Load(value)
			break
		case config.Meta.DatabaseSuffix:
			db := meta[iname].GetDatabase(dname)
			db.Load(value)
			break
		case config.Meta.InstanceSuffix:
			ins := meta[iname]
			ins.Load(value)
			break
		}
	} else if operType == etcd.DELETE {
		switch suffix {
		case config.Meta.TableSuffix:
			db := meta[iname].GetDatabase(dname)
			db.RemoveTable(tname)
			break
		case config.Meta.DatabaseSuffix:
			break
		case config.Meta.InstanceSuffix:
			break
		}
	}
}

func WatchMeta() {
	etcd.WatchWithPrefix(config.Meta.Root, metaWatcher)
}

func LoadMeta() error {
	kvs, err := etcd.GetWithPrefix(config.Meta.Root)
	if err == nil {
		for _, kv := range kvs {
			iname, dname, tname, suffix := parsePath(string(kv.Key))
			if suffix == config.Meta.TableSuffix {
				tbl := meta[iname].GetDatabase(dname).CreateTable(tname)
				err = tbl.Load(kv.Value)
			} else if suffix == config.Meta.DatabaseSuffix {
				db := meta[iname].CreateDatabase(dname)
				err = db.Load(kv.Value)
			} else if suffix == config.Meta.InstanceSuffix {
				ins := CreateInstance(iname)
				err = ins.Load(kv.Value)
			}
			if err != nil {
				break
			}
		}
	}
	return err
}

func storeInstance(ins *Instance) error {
	data, err := json.Marshal(ins)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s%s%s", config.Meta.Root, Separator, ins.GetMName())
	_, err = etcd.Put(path, string(data))
	return err
}

func storeDatabase(db *Database) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s%s%s", db.Instance.GetPath(), Separator, db.GetMName())
	_, err = etcd.Put(path, string(data))
	return err
}

func storeTable(tbl *Table) error {
	data, err := json.Marshal(tbl)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s%s%s", tbl.Database.GetPath(), Separator, tbl.GetMName())
	_, err = etcd.Put(path, string(data))
	return err
}

func Remove(path string) (err error) {
	_, err = etcd.Del(path)
	return
}
