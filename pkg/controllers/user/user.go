package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pwhb/go-gin-hogwarts/pkg/database"
	"github.com/pwhb/go-gin-hogwarts/pkg/models/user"
)

func Refresh(ctx *gin.Context) {

	var rows []user.User

	db := database.DB
	db.Find(&rows)
	for _, v := range rows {
		v.Slug = strings.Join([]string{strings.ToLower(v.FirstName), strings.ToLower(v.LastName)}, "-")
		db.Updates(&v)
	}
	ctx.JSON(http.StatusOK, rows)
}

func Create(ctx *gin.Context) {
	var row user.User
	if err := ctx.ShouldBindJSON(&row); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row.Slug = strings.Join([]string{strings.ToLower(row.FirstName), strings.ToLower(row.LastName)}, "-")
	db := database.DB
	db.Create(&row)
	ctx.JSON(http.StatusCreated, row)
}

func ReadAll(ctx *gin.Context) {
	var rows []user.PublicUser
	db := database.DB
	q := db.Model(&user.User{})
	if name := ctx.Query("name"); name != "" {
		q = q.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
	}
	if first_name := ctx.Query("first_name"); first_name != "" {
		q = q.Where("first_name LIKE ?", first_name)
	}
	if last_name := ctx.Query("last_name"); last_name != "" {
		q = q.Where("last_name LIKE ?", last_name)
	}
	if blood_status := ctx.Query("blood_status"); blood_status != "" {
		q = q.Where("blood_status = ?", blood_status)
	}
	if house := ctx.Query("house"); house != "" {
		q = q.Where("house = ?", house)
	}
	if species := ctx.Query("species"); species != "" {
		q = q.Where("species = ?", species)
	}
	if occupation := ctx.Query("occupation"); occupation != "" {
		q = q.Where("occupation = ?", occupation)
	}
	q.Find(&rows)
	// db.Model(&user.User{}).Find(&rows)
	ctx.JSON(http.StatusOK, rows)
}

func ReadOne(ctx *gin.Context) {
	var row user.User
	id := ctx.Param("id")
	db := database.DB
	db.Find(&row, id)
	if row.ID == 0 {
		res := fmt.Sprintf("User with the ID %s doesn't exist.", id)
		ctx.String(http.StatusNotFound, res)
	}
	ctx.JSON(http.StatusOK, row)
}

func Update(ctx *gin.Context) {
	var row user.User
	id := ctx.Param("id")
	db := database.DB
	db.Find(&row, id)
	if row.ID == 0 {
		res := fmt.Sprintf("User with the ID %s doesn't exist.", id)
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
	var row user.User
	id := ctx.Param("id")
	db := database.DB
	db.Find(&row, id)
	if row.ID == 0 {
		res := fmt.Sprintf("User with the ID %s doesn't exist.", id)
		ctx.String(http.StatusNotFound, res)
		return
	}
	db.Delete(&row)
	res := fmt.Sprintf("User with the ID %s was deleted successfully.", id)
	ctx.String(http.StatusNoContent, res)
}
