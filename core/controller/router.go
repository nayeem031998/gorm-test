package controller

func (s *Server) initRouter() {
	s.Router.GET("/student/get", s.UserController)
	s.Router.GET("/employee/get", s.employeeController)
	s.Router.GET("/company/get", s.companyController)

}
