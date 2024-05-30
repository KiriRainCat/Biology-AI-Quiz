package internal

import (
	"stp/internal/pkg/config"
	res "stp/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var Controller = &controller{s: s}

type controller struct {
	s *service
}

func (controller) GenerateQuiz(ctx *gin.Context) {
	list, err := s.GenerateQuiz()
	if err != nil {
		res.InternalErr(ctx)
	}

	res.Success(ctx, list)
}

func (*controller) GetRecords(ctx *gin.Context) {
	var list []*record
	if DB.Select("name", "accuracy", "time_taken", "updated_at").Find(&list).Error != nil {
		res.InternalErr(ctx)
	}

	res.Success(ctx, list)
}

func (*controller) PutRecord(ctx *gin.Context) {
	// 请求数据模型
	type json struct {
		Name       string  `json:"name" binding:"required"`
		Passphrase string  `json:"passphrase" binding:"required"`
		Accuracy   float64 `json:"accuracy" binding:"required"`
		TimeTaken  float64 `json:"time_taken" binding:"required"`
	}

	// 解析并校验请求数据
	var data *json
	if err := ctx.ShouldBindJSON(&data); err != nil {
		res.ParamErr(ctx)
		return
	}

	// 校验 passphrase
	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Passphrase+config.Config.Server.EncryptSalt), bcrypt.MinCost)
	if err != nil {
		res.InternalErr(ctx)
		return
	}
	var storedRecord *record
	if DB.First(&storedRecord, "name = ?", data.Name).Error == nil {
		if string(bytes) != storedRecord.Passphrase {
			res.ParamErrM(ctx, "Passphrase is incorrect")
			return
		}
	}

	// 保存记录
	record := &record{
		Name:       data.Name,
		Accuracy:   data.Accuracy,
		Passphrase: string(bytes),
		TimeTaken:  data.TimeTaken,
	}
	if DB.Save(&record).Error != nil {
		res.InternalErr(ctx)
		return
	}

	res.Success(ctx, record)
}
