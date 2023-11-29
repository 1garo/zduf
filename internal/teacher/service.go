package teacher

import (
	"errors"

	"github.com/1garo/zduf/internal"
)

var (
	ErrTeacherShouldBeOldEnough = errors.New("Teacher should be old enough (over 21)")
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(t *internal.Teacher) error {

	if t.Age <= 21 {
		return ErrTeacherShouldBeOldEnough
	}

	return s.Repository.Create(t)
}
