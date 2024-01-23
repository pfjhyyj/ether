package domain

import "sync"

var (
	userRepository *UserRepository

	once sync.Once
)

func Init() {
	once.Do(func() {
		userRepository = &UserRepository{}
	})
}
