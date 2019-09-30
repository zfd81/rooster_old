package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zfd81/rooster/cluster"
	"net/http"
)

func ListClusterNode(c *gin.Context) {
	clus := cluster.GetCluster()
	nodes := make([]*cluster.Node, 0, len(clus))
	for _, v := range clus {
		nodes = append(nodes, v)
	}
	c.JSON(http.StatusOK, nodes)
}
