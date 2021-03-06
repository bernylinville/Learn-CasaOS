package route

import (
	"Learn-CasaOS/middleware"
	"Learn-CasaOS/pkg/config"
	jwt2 "Learn-CasaOS/pkg/utils/jwt"
	v1 "Learn-CasaOS/route/v1"
	"Learn-CasaOS/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

var swagHandler gin.HandlerFunc

func InitRouter(swagHandler gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	gin.SetMode(config.ServerInfo.RunMode)

	r.StaticFS("/ui", http.FS(web.Static))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "ui/")
	})

	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}

	// 登录
	r.POST("/v1/user/login", v1.Login)

	r.GET("/debug", v1.GetSystemConfigDebug)

	v1Group := r.Group("/v1")

	v1Group.Use(jwt2.JWT(swagHandler))
	{
		v1UserGroup := v1Group.Group("/user")
		v1UserGroup.Use()
		{
			// 设置用户
			v1UserGroup.POST("/setusernamepwd", v1.Set_Name_Pwd)
			// 修改头像
			v1UserGroup.POST("/changhead", v1.Up_Load_Head)
			// 修改用户名
			v1UserGroup.POST("/changusername", v1.Chang_User_Name)
			// 修改密码
			v1UserGroup.POST("/changuserpwd", v1.Chang_User_Pwd)
			// 修改用户信息
			v1UserGroup.POST("/changuserinfo", v1.Chang_User_Info)
			// 获取用户详情
			v1UserGroup.GET("/info", v1.UserInfo)
		}

		v1ZiMaGroup := v1Group.Group("/zima")
		v1ZiMaGroup.Use()
		{
			// 获取 cpu 信息
			v1ZiMaGroup.GET("/getcpuinfo", v1.CpuInfo)
			// 获取内存信息
			v1ZiMaGroup.GET("/getmeminfo", v1.MemInfo)
			// 获取磁盘信息
			v1ZiMaGroup.GET("/getdiskinfo", v1.DiskInfo)
			// 获取网络信息
			v1ZiMaGroup.GET("/getnetinfo", v1.NetInfo)
			// 获取信息
			v1ZiMaGroup.GET("/getinfo", v1.Info)
			// 获取系统信息
			v1ZiMaGroup.GET("/sysinfo", v1.SysInfo)
		}

		v1ZeroTierGroup := v1Group.Group("/zerotier")
		v1ZeroTierGroup.Use()
		{
			// 登录 zerotier 获取 token
			v1ZeroTierGroup.POST("/login", v1.ZeroTierGetToken)
		}
	}

	return r
}
