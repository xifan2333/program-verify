package handlers

import (
	"net/http"
	"time"

	"license-verify/database"
	"license-verify/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateLicenseRequest struct {
	Count     int       `json:"count" binding:"required,min=1,max=100"`
	ExpiredAt time.Time `json:"expired_at" binding:"required"`
}

type VerifyLicenseRequest struct {
	LicenseKey  string `json:"license_key" binding:"required"`
	MachineCode string `json:"machine_code,omitempty"`
}

func CreateLicenses(c *gin.Context) {
	var req CreateLicenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var licenses []models.License
	for i := 0; i < req.Count; i++ {
		license := models.License{
			LicenseKey: uuid.New().String(),
			Status:     "unused",
			ExpiredAt:  req.ExpiredAt,
		}
		licenses = append(licenses, license)
	}

	if err := database.GetDB().Create(&licenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建卡密失败"})
		return
	}

	c.JSON(http.StatusOK, licenses)
}

func GetLicenses(c *gin.Context) {
	var licenses []models.License
	if err := database.GetDB().Find(&licenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取卡密列表失败"})
		return
	}

	c.JSON(http.StatusOK, licenses)
}

func DeleteLicense(c *gin.Context) {
	id := c.Param("id")
	if err := database.GetDB().Delete(&models.License{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除卡密失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func VerifyLicense(c *gin.Context) {
	var req VerifyLicenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var license models.License
	if err := database.GetDB().Where("license_key = ?", req.LicenseKey).First(&license).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "卡密不存在"})
		return
	}

	// 检查卡密状态
	if license.Status == "used" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "卡密已被使用"})
		return
	}

	if license.Status == "expired" || time.Now().After(license.ExpiredAt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "卡密已过期"})
		return
	}

	// 如果提供了机器码，进行绑定
	if req.MachineCode != "" {
		license.MachineCode = req.MachineCode
		license.Status = "used"
		license.UsedAt = time.Now()
		if err := database.GetDB().Save(&license).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "绑定机器码失败"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":        true,
		"expired_at":   license.ExpiredAt,
		"machine_code": license.MachineCode,
	})
}
