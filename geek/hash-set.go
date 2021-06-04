package main

import "sync"

type HashSet interface {
	Set(key string)
	Size() int
	Exist(key string) bool
}

type hashset struct {
	m map[string]interface{}

}

func (h *hashset) Set(key string) {
	h.m[key] = "alex"
}

func (h *hashset) Size() int {
	return len(h.m)
}

func (h *hashset) Exist(key string) bool {
	_, ok := h.m[key]
	return  ok
}

type safeset struct {
	HashSet
	mutex sync.RWMutex
}

func (s *safeset) Size() int {
	s.mutex.RLock()
	defer s.mutex.Unlock()
	return s.HashSet.Size()
}

func (s *safeset) Set(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.HashSet.Set(key)
}

func(s *safeset)Exist(key string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.HashSet.Exist(key)
}


