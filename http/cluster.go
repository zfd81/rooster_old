package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zfd81/rooster/cluster"
	"net/http"
)

func ListClusterNode(c *gin.Context) {
	members := cluster.GetMembers()
	nodes := make([]*cluster.Node, 0, len(members))
	for _, v := range members {
		nodes = append(nodes, v)
	}
	c.JSON(http.StatusOK, nodes)
}
