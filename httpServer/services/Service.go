package services

import "github.com/ZhenlyChen/BugServer/httpServer/models"

type Service struct {
	Model *models.Model
	User userService
	Room roomService
	Game gameService
}

func NewService(m *models.Model) *Service {
	service := new(Service)
	service.Model = m
	service.User = userService{
		Model: &m.User,
		Service: service,
		UserInfo: make(map[string]UserBaseInfo),
	}
	service.Room = roomService{
		Service: service,
	}
	service.Game = gameService{
		Service: service,
		Model: &m.Game,
	}
	return service
}

func (s *Service) GetUserService() UserService {
	return &s.User
}

func (s *Service) GetRoomService() RoomService {
	return &s.Room
}

func (s *Service) GetGameService() GameService {
	return &s.Game
}