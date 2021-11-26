package routes

import (
	"github.com/gin-gonic/gin"
)

func Bind_routes(r *gin.Engine) {
	r.GET("organizations", Get_organizations)
	r.POST("organization", Add_organization)
	r.POST("user_config", Update_user_config)
	r.POST("user", Update_user)
	r.POST("workday", Update_workday)
}
