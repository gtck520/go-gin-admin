package repository

import (
	//"github.com/jinzhu/gorm"

	"github.com/konger/ckgo/common/logger"
	//models "github.com/konger/ckgo/models/common"
)

//GroupRepository 注入IDb
type GroupRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}
