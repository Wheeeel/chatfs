package model

import (
	irc "github.com/Wheeeel/chatfs/server"
)

type Server struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
	Client   *irc.IRC
}

func GetServers() (servers []Server, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	sis := DB["server"]
	for _, si := range sis {
		s, ok := si.(Server)
		if !ok {
			continue
		}

		servers = append(servers, s)
	}
	return
}

// Update the config, if it has a nickname, connect to it
func UpdateConfig(server Server) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	sis := DB["server"]
	var idx int
	for k, si := range sis {
		s, ok := si.(Server)
		if !ok {
			continue
		}
		if s.Name == server.Name {
			idx = k
		}
	}
	sis = append(sis[:idx], sis[idx+1:])

	// Server user not empty, we can connect to
	if server.User != "" {
		server.Client = irc.NewIRC(server.User, server.Name)
	}
	sis = append(sis, server)
	DB["server"] = sis
	return
}

func GetServer(name string) (server *Server, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	sis := DB["server"]
	for _, si := range sis {
		s, ok := si.(Server)
		if !ok {
			continue
		}
		if s.Name == name {
			server = &s
			return
		}
	}
	server = nil
	return
}

func AddServer(srv Server) (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	sis := DB["server"]
	sis = append(sis, srv)
	DB["server"] = sis
	return
}
