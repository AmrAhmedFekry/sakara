package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(g *gin.Context) {
	limit := g.DefaultQuery("limit", "10")
	offset := g.DefaultQuery("offset", "0")
	var {moduleInPlural} []{modelName}
	db.Limit(limit).Offset(offset).Find(&{moduleInPlural})
	g.JSON(http.StatusOK, gin.H{"data": {moduleInPlural}})
}

func Store(g *gin.Context) {
	var {moduleInSingular} {modelName}
	if err := g.ShouldBindJSON(&{moduleInSingular}); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	db.Create(&{moduleInSingular})
	g.JSON(http.StatusCreated, gin.H{"message": "Post has been created successfully", "data": {moduleInSingular}})
}

func Update(g *gin.Context) {
}

func Delete(g *gin.Context) {
	{modelNameInSmallCase} := getById(g)
	if {modelNameInSmallCase}.ID == 0 {
		return
	}
	db.Unscoped().Delete(&{modelNameInSmallCase})
	g.JSON(http.StatusOK, gin.H{"data": ""})
}

func Show(g *gin.Context) {
	{modelNameInSmallCase} := getById(g)
	if {modelNameInSmallCase}.ID == 0 {
		return
	}
	g.JSON(http.StatusCreated, gin.H{"data": {modelNameInSmallCase}})
}

func getById(g *gin.Context) {modelName} {
	id := g.Param("id")
	var {modelNameInSmallCase} {modelName}
	db.First(&{modelNameInSmallCase}, id)
	if {modelNameInSmallCase}.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{"message": "Not found this resource"})
	}
	return {modelNameInSmallCase}
}
