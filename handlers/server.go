package handlers

import (
	"github.com/geoairng/Finalyear/model"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) Teardown() {
	s.DB.Migrator().DropTable(&model.User{})
	s.DB.AutoMigrate(&model.User{})
}
