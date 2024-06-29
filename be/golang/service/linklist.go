package service

import (
	"k8s-go-gin/dal"
	"k8s-go-gin/models"
)

type LinkListService struct {
	mongoDAL *dal.MongoDAL
}

func NewLinkListService(mongoDAL *dal.MongoDAL) *LinkListService {
	return &LinkListService{mongoDAL: mongoDAL}
}

func (s *LinkListService) GetLinkList() ([]models.LinkData, error) {
	return s.mongoDAL.GetLinkList()
}
