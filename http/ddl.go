package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zfd81/rooster/meta"
	"net/http"
)

var (
	ErrInsNotFound = errors.New("Instance not found")
	ErrInsExists   = errors.New("Instance already exists")
	ErrDBNotFound  = errors.New("Database not found")
	ErrDBExists    = errors.New("Database already exists")
	ErrTblNotFound = errors.New("Table not found")
	ErrTblExists   = errors.New("Table already exists")
)

func ListInstance(c *gin.Context) {
	m := meta.GetMeta()
	inses := make([]*meta.Instance, 0, len(m))
	for _, v := range m {
		inses = append(inses, v)
	}
	c.JSON(http.StatusOK, inses)
}

func FindInstance(c *gin.Context) {
	m := meta.GetMeta()
	iname := c.Param("iname")
	c.JSON(http.StatusOK, m[iname])
}

func CreateInstance(c *gin.Context) {
	var insInfo meta.InstanceInfo
	if err := c.BindJSON(&insInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ins := meta.GetMeta()[insInfo.Name]
	if ins != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsExists.Error()})
		return
	}
	ins = meta.CreateInstanceWithInfo(insInfo)
	ins.Store()
	c.JSON(http.StatusOK, ins)
}

func AlterInstance(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	var insInfo meta.InstanceInfo
	if err := c.BindJSON(&insInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	insInfo.Name = iname
	ins.InstanceInfo = insInfo
	ins.Store()
	c.JSON(http.StatusOK, ins)
}

func ListDatabase(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	dbs := make([]*meta.Database, 0, len(ins.Databases))
	for _, v := range ins.Databases {
		dbs = append(dbs, v)
	}
	c.JSON(http.StatusOK, dbs)
}

func FindDatabase(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	dname := c.Param("dname")
	c.JSON(http.StatusOK, ins.GetDatabase(dname))
}

func CreateDatabase(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	var dbInfo meta.DatabaseInfo
	if err := c.BindJSON(&dbInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ins.GetDatabase(dbInfo.Name)
	if db != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrDBExists.Error()})
		return
	}
	db = ins.CreateDatabaseWithInfo(dbInfo)
	db.Store()
	c.JSON(http.StatusOK, db)
}

func AlterDatabase(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	var dbInfo meta.DatabaseInfo
	if err := c.BindJSON(&dbInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dname := c.Param("dname")
	db := ins.GetDatabase(dname)
	if db == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrDBNotFound.Error()})
		return
	}
	dbInfo.Name = dname
	db.DatabaseInfo = dbInfo
	db.Store()
	c.JSON(http.StatusOK, db)
}

func ListTable(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	dname := c.Param("dname")
	db := ins.GetDatabase(dname)
	if db == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrDBNotFound.Error()})
		return
	}
	tbls := make([]*meta.Table, 0, len(db.Tables))
	for _, v := range db.Tables {
		tbls = append(tbls, v)
	}
	c.JSON(http.StatusOK, tbls)
}

func FindTable(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	dname := c.Param("dname")
	db := ins.GetDatabase(dname)
	if db == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrDBNotFound.Error()})
		return
	}
	tname := c.Param("tname")
	c.JSON(http.StatusOK, db.GetTable(tname))
}

func CreateTable(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	dname := c.Param("dname")
	db := ins.GetDatabase(dname)
	if db == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrDBNotFound.Error()})
		return
	}
	var tblInfo meta.TableInfo
	if err := c.BindJSON(&tblInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tbl := db.GetTable(tblInfo.Name)
	if tbl != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrTblExists.Error()})
		return
	}
	tbl = db.CreateTableWithInfo(tblInfo)
	tbl.Store()
	c.JSON(http.StatusOK, db)
}

func AlterTable(c *gin.Context) {
	iname := c.Param("iname")
	ins := meta.GetMeta()[iname]
	if ins == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInsNotFound.Error()})
		return
	}
	dname := c.Param("dname")
	db := ins.GetDatabase(dname)
	if db == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrDBNotFound.Error()})
		return
	}
	var tblInfo meta.TableInfo
	if err := c.BindJSON(&tblInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tname := c.Param("tname")
	tbl := db.GetTable(tname)
	if tbl == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrTblNotFound.Error()})
		return
	}
	tblInfo.Name = tname
	tbl.TableInfo = tblInfo
	tbl.Store()
	c.JSON(http.StatusOK, tbl)
}
