package storage

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type StoredHash struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type Storage struct {
	mu   sync.Mutex
	data []StoredHash
	file string
}

func NewStorage(filename string) *Storage {
	s := &Storage{file: filename}
	s.load()
	return s
}

func (s *Storage) load() {
	file, err := os.ReadFile(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			s.data = []StoredHash{}
			return
		}
		log.Fatal(err)
	}
	json.Unmarshal(file, &s.data)
}

func (s *Storage) save() {
	bytes, _ := json.MarshalIndent(s.data, "", "  ")
	os.WriteFile(s.file, bytes, 0644)
}

func (s *Storage) Add(email, hash string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, StoredHash{Email: email, Hash: hash})
	s.save()
}

func (s *Storage) VerifyAndDelete(hash string) (bool, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, v := range s.data {
		if v.Hash == hash {
			email := v.Email
			// удалить запись
			s.data = append(s.data[:i], s.data[i+1:]...)
			s.save()
			return true, email
		}
	}
	return false, ""
}
