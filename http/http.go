package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zfd81/rooster/conf"
	"strconv"
)

func Start() error {
	router := gin.Default()
	ddl := router.Group("/ddl")
	{
		ddl.GET("/ins", ListInstance)         //查看所有实例
		ddl.GET("/ins/:iname", FindInstance)  //查看实例信息
		ddl.POST("/ins", CreateInstance)      //创建实例
		ddl.PUT("/ins/:iname", AlterInstance) //修改实例

		ddl.GET("/ins/:iname/db", ListDatabase)         //查看实例下的所有数据库
		ddl.GET("/ins/:iname/db/:dname", FindDatabase)  //查看数据库信息
		ddl.POST("/ins/:iname/db", CreateDatabase)      //创建数据库
		ddl.PUT("/ins/:iname/db/:dname", AlterDatabase) //修改数据库

		ddl.GET("/ins/:iname/db/:dname/tbl", ListTable)           //查看数据库下的所有数据表
		ddl.GET("/ins/:iname/db/:dname/tbl/:tname", FindTable)    //查看数据表信息
		ddl.POST("/ins/:iname/db/:dname/tbl", CreateTable)        //创建数据表
		ddl.PUT("/ins/:iname/db/:dname/tbl/:tname", AlterTable)   //修改数据表
		ddl.DELETE("/ins/:iname/db/:dname/tbl/:tname", DropTable) //删除数据表
	}
	cluster := router.Group("/cluster")
	{
		cluster.GET("/", ListClusterNode)
	}
	return router.Run(":" + strconv.Itoa(conf.GetGlobalConfig().Http.Port))
}
