package house

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pwhb/go-gin-hogwarts/pkg/database"
	"github.com/pwhb/go-gin-hogwarts/pkg/models/house"
)

// func Refresh(ctx *gin.Context) {

// 	var rows []house.House

// 	db := database.DB
// 	db.Find(&rows)
// 	for _, v := range rows {
// 		v.Slug = strings.Join([]string{strings.ToLower(v.FirstName), strings.ToLower(v.LastName)}, "-")
// 		db.Updates(&v)
// 	}
// 	ctx.JSON(http.StatusOK, rows)
// }

func Create(ctx *gin.Context) {
	var row house.House
	if err := ctx.ShouldBindJSON(&row); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.DB
	db.Create(&row)
	ctx.JSON(http.StatusCreated, row)
}

func ReadAll(ctx *gin.Context) {
	var rows []house.House
	db := database.DB
	q := db.Model(&house.House{})
	q.Find(&rows)
	// db.Model(&house.House{}).Find(&rows)
	ctx.JSON(http.StatusOK, rows)
}

func ReadOne(ctx *gin.Context) {
	var row house.House
	id := ctx.Param("id")
	db := database.DB
	db.Find(&row, id)
	if row.ID == 0 {
		res := fmt.Sprintf("House with the ID %s doesn't exist.", id)
		ctx.String(http.StatusNotFound, res)
	}
	ctx.JSON(http.StatusOK, row)
}

func Update(ctx *gin.Context) {
	var row house.House
	id := ctx.Param("id")
	db := database.DB
	db.Find(&row, id)
	if row.ID == 0 {
		res := fmt.Sprintf("House with the ID %s doesn't exist.", id)
		ctx.String(http.StatusNotFound, res)
		return
	}
	if err := ctx.ShouldBindJSON(&row); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Updates(row)
	ctx.JSON(http.StatusCreated, row)
}

func Delete(ctx *gin.Context) {
	var row house.House
	id := ctx.Param("id")
	db := database.DB
	db.Find(&row, id)
	if row.ID == 0 {
		res := fmt.Sprintf("House with the ID %s doesn't exist.", id)
		ctx.String(http.StatusNotFound, res)
		return
	}
	db.Delete(&row)
	res := fmt.Sprintf("House with the ID %s was deleted successfully.", id)
	ctx.String(http.StatusNoContent, res)
}
