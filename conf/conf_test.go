package conf

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
	"testing"
)

func TestGetGlobalConfig(t *testing.T) {
	conf := NewConfig()
	fmt.Println(conf.Name)
	fmt.Println(conf.Age)
	fmt.Printf("%p", &conf.Age)
	conf.Load("config.toml")
	fmt.Println(conf.Name)
	fmt.Println(conf.Age)
	fmt.Printf("%p", &conf.Age)
	list := arraylist.New()
	list.Add("a")                     // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]
	fmt.Println(list)
}
