package web

// Server WebServer sturecture
type Server struct {
}

// NewServer Make new struct
func NewServer() *Server {
	return &Server{}
}

// Start Start web server
func (s *Server) Start(ip string, port int) error {

}
