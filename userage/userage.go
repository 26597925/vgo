package main

import (
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

/*func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}*/

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age,ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
/*
考点：map线程安全
解答：
可能会出现fatal error: concurrent map read and map write. 修改一下看看效果*/