package utils

import (
	"sync"
	"sync/atomic"
	"time"
)

type snowFlake struct {
	fixedTimeStamp int64 //填充时间戳
	lastTimeStamp  int64 //上次生成时间戳
	machineID      int64 //数据中心+机器号组成 共十位
	sequenceNumber int64 //序列号
	mutex          sync.Mutex
}

var SNOW_FLAKE *snowFlake

func InitSnowFlake(machineID int64) {
	SNOW_FLAKE = NewSnowFlake(machineID)
}

// machine小于等于1023
func NewSnowFlake(machineID int64) *snowFlake {
	return &snowFlake{
		fixedTimeStamp: int64(1677210908105),
		lastTimeStamp:  int64(-1),
		machineID:      machineID,
		sequenceNumber: int64(0),
	}
}

func (s *snowFlake) GenSnowID() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	currentTimeStamp := time.Now().UnixMilli()
	if currentTimeStamp == atomic.LoadInt64(&s.lastTimeStamp) {
		atomic.AddInt64(&s.sequenceNumber, 1)
		if atomic.LoadInt64(&s.sequenceNumber) > 4096 {
			for currentTimeStamp <= atomic.LoadInt64(&s.lastTimeStamp) {
				currentTimeStamp = time.Now().UnixMilli()
			}
			atomic.StoreInt64(&s.sequenceNumber, 0)
		}
	} else {
		atomic.StoreInt64(&s.sequenceNumber, 0)
	}
	atomic.StoreInt64(&s.lastTimeStamp, currentTimeStamp)
	return (currentTimeStamp-s.fixedTimeStamp)<<22 | s.sequenceNumber | s.machineID<<12
}
