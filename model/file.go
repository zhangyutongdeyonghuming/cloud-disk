package model

import "gorm.io/gorm"

// File 文件信息
type File struct {
	gorm.Model

	Filename  string `gorm:"varchar(64) index"` // 文件名
	Extension string `gorm:"varchar(32)"`       // 扩展名
	Path      string `gorm:"varchar(255)"`      // 上传全路径
}
