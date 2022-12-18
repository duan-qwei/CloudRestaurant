package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PointService struct{}

func (pointService *PointService) add(c *gin.Context, userId int64, points int32) error {
	var point *model.Point

	db := common.DB.Where("user_id=?", userId).Find(&point)
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected > 0 {
		point.Points = point.Points + points
		if err := db.Model(&point).Where("user_id=?", userId).Update("points", point.Points).Error; err != nil {
			return err
		}
	} else {
		db.Create(model.Point{
			Model:  gorm.Model{},
			Points: points,
			UserId: userId,
		})

		if err := db.Create(&points).Error; err != nil {
			return err
		}
	}
	return nil
}

func (pointService *PointService) getPointsByUserId(c *gin.Context, userId int64) (point *model.Point, error error) {
	db := common.DB.Where("user_id=?", userId).Find(&point)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RowsAffected == 0 {
		return nil, nil
	}

	return point, nil
}
