package controller

func (s *Server) initRouter() {
	s.Router.GET("/home", s.IndexController)
	s.Router.GET("/register", s.RegisterController)
	s.Router.GET("/student", s.StudentController)
	s.Router.GET("/class", s.ClassController)
	s.Router.GET("/subject", s.SubjectController)
	s.Router.GET("/signup", s.SignUpController)
}
