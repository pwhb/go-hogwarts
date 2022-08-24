package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pwhb/go-gin-hogwarts/pkg/controllers/house"
	"github.com/pwhb/go-gin-hogwarts/pkg/controllers/user"
)

func SetupRoutes(r *gin.Engine) {
	userPath := "/api/v1/users"

	r.GET(userPath+"/refresh", user.Refresh)

	r.POST(userPath, user.Create)

	r.GET(userPath, user.ReadAll)

	r.GET(userPath+"/:id", user.ReadOne)

	r.PUT(userPath+"/:id", user.Update)

	r.DELETE(userPath+"/:id", user.Delete)

	housePath := "/api/v1/houses"

	// r.GET(housePath+"/refresh", house.Refresh)

	r.POST(housePath, house.Create)

	r.GET(housePath, house.ReadAll)

	r.GET(housePath+"/:id", house.ReadOne)

	r.PUT(housePath+"/:id", house.Update)

	r.DELETE(housePath+"/:id", house.Delete)

}
