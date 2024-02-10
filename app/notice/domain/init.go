package domain

import "sync"

var (
	noticeRepository *NoticeRepository

	once sync.Once
)

func Init() {
	once.Do(func() {
		noticeRepository = &NoticeRepository{}
	})
}
