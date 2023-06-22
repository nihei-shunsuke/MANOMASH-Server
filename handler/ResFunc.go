package handler

import (
	"MANOMASH-Server/model"
)

func ResFlgCreate(ResStatus int, ResResult string, ResUID uint, ResUserName string, ResIntroduce string) (model.ResFlg) {
	var ResData model.ResFlg
	ResData.Status = ResStatus
	ResData.Result = ResResult
	ResData.User_ID = ResUID
	ResData.UserName = ResUserName
	ResData.Introduce = ResIntroduce
	return ResData
}