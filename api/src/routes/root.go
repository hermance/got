package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Bind_routes(r *gin.Engine, db *gorm.DB) {
	r.GET("organizations", Get_organizations(db))
	r.GET("users", Get_users(db))
	r.POST("organization", Add_organization(db))
	r.POST("user_config", Update_user_config(db))
	r.POST("user", Create_user(db))
	r.POST("workday", Update_workday(db))
}
