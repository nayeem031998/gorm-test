package controller

func (s *Server) initRouter() {
	s.Router.GET("/", s.IndexController)
	s.Router.GET("/student", s.StudentController)
	s.Router.GET("/class", s.ClassController)
	s.Router.GET("/subject", s.SubjectController)


}
