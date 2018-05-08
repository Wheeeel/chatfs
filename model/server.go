package model

type Server struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
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

func AddServer(srv Server) (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	sis := DB["server"]
	sis = append(sis, srv)
	DB["server"] = sis
	return
}
