package main

import (
	"log"
)

type Server struct {
	connections map[string]*Client
	tasks       chan *Task
	register    chan *Client
	unregister  chan *Client
}

func (s *Server) spin() {
	for {
		select {
		case client := <-s.register:
			s.connections[client.user] = client
		case client := <-s.unregister:
			if _, ok := s.connections[client.user]; ok {
				delete(s.connections, client.user)
				close(client.message)
			}
		case task := <-s.tasks: // Error handle ?
			if client, ok := s.connections[task.To]; ok {
				client.message <- task
			} else {
				log.Printf("SERVER: No such user: %s", task.To)
			}
		}
	}
}
