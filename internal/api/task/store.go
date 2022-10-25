package task

import (
	"artemzinchenko.com/auth/pb"
	"errors"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("already exists")

type StorageService interface {
	SaveOrUpdate(task *pb.Task) (*pb.Task, error)
}

type inMemoryStorageService struct {
	mutex sync.RWMutex
	data  map[string]*pb.Task
}

func NewInMemoryStorageService() StorageService {
	return &inMemoryStorageService{}
}

func (s inMemoryStorageService) SaveOrUpdate(task *pb.Task) (*pb.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, found := s.data[task.Id]; !found {
		s.data[task.Id] = task
	} else {
		s.data[task.Id].Message = task.Message
	}
	return s.data[task.Id], nil
}
