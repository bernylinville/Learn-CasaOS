package v1

import (
	"Learn-CasaOS/model"
	oasis_err2 "Learn-CasaOS/pkg/utils/oasis_err"
	"Learn-CasaOS/service"
	"net/http"
	"strings"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/disk"
)

// @Summary 获取cpu信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getcpuinfo [get]
func CpuInfo(c *gin.Context) {
	// 检查参数是否正确
	cpu := service.MyService.ZiMa().GetCpuPercent()
	num := service.MyService.ZiMa().GetCpuCoreNum()
	data := make(map[string]interface{})
	data["percent"] = cpu
	data["num"] = num
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    data,
		})
}

// @Summary 获取内存信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getmeminfo [get]
func MemInfo(c *gin.Context) {
	// 检查参数是否正确
	mem := service.MyService.ZiMa().GetMemInfo()
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    mem,
		})
}

// @Summary 获取磁盘信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getdiskinfo [get]
func DiskInfo(c *gin.Context) {
	// 检查参数是否正确
	disk := service.MyService.ZiMa().GetDiskInfo()
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    disk,
		})
}

// @Summary 获取网络信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getnetinfo [get]
func NetInfo(c *gin.Context) {
	netList := service.MyService.ZiMa().GetNetInfo()

	newNet := []model.IOCountersStat{}
	for _, n := range netList {
		for _, netCardName := range service.MyService.ZiMa().GetNet(true) {
			if n.Name == netCardName {
				item := *(*model.IOCountersStat)(unsafe.Pointer(&n))
				item.State = strings.TrimSpace(service.MyService.ZiMa().GetNetState(n.Name))
				item.DateTime = time.Now()
				newNet = append(newNet, item)
				break
			}
		}
	}

	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    newNet,
		})
}

// @Summary 获取信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/getinfo [get]
func Info(c *gin.Context) {
	var data = make(map[string]interface{}, 4)

	var diskArr []*disk.UsageStat
	diskArr = append(diskArr, service.MyService.ZiMa().GetDiskInfo())
	data["disk"] = diskArr

	cpu := service.MyService.ZiMa().GetCpuPercent()
	num := service.MyService.ZiMa().GetCpuCoreNum()
	cpuData := make(map[string]interface{})
	cpuData["percent"] = cpu
	cpuData["num"] = num
	data["cpu"] = cpuData

	data["mem"] = service.MyService.ZiMa().GetMemInfo()

	netList := service.MyService.ZiMa().GetNetInfo()
	newNet := []model.IOCountersStat{}
	for _, n := range netList {
		for _, netCardName := range service.MyService.ZiMa().GetNet(true) {
			if n.Name == netCardName {
				item := *(*model.IOCountersStat)(unsafe.Pointer(&n))
				item.State = strings.TrimSpace(service.MyService.ZiMa().GetNetState(n.Name))
				item.DateTime = time.Now()
				newNet = append(newNet, item)
				break
			}
		}
	}
	data["net"] = newNet

	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    data,
		})
}

// @Summary 获取系统信息
// @Produce  application/json
// @Accept application/json
// @Tags zima
// @Security ApiKeyAuth
// @Success 200 {string} string "ok"
// @Router /zima/sysinfo [get]
func SysInfo(c *gin.Context) {
	info := service.MyService.ZiMa().GetSysInfo()
	c.JSON(http.StatusOK,
		model.Result{
			Success: oasis_err2.SUCCESS,
			Message: oasis_err2.GetMsg(oasis_err2.SUCCESS),
			Data:    info,
		})
}
