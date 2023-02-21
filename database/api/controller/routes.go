package controller

func (s *Server) InitRoutes() {
	s.Router.GET("/products", s.GetProducts)
	s.Router.GET("/products/:id", s.GetProduct)
	s.Router.GET("/product/:name", s.GetProductsByName)
	s.Router.POST("/create", s.CreateProduct)
	s.Router.PUT("/updated/:id", s.UpdateProducts)
	s.Router.DELETE("/deleteproduct/:id", s.DeleteProducts)
	s.Router.POST("/add", s.CreateOrder)
	s.Router.GET("/orders", s.GetOrders)
	s.Router.GET("/orders/:id", s.GetOrder)
	s.Router.GET("/order/:id", s.GetOrderById)
	s.Router.PUT("/update/:id", s.UpdateOrders)
	s.Router.DELETE("/deleteorder/:id", s.DeleteOrders)
	s.Router.POST("/transaction", s.CreateTransactions)
	s.Router.GET("/transactions/:id", s.GetTransactions)
	s.Router.GET("/transactions", s.GetAllTransactions)
	s.Router.PUT("/updatetransaction/:id", s.UpdateTransactions)
	s.Router.DELETE("/deletetransaction/:id", s.DeleteTransactions)
	s.Router.POST("/user", s.CreateInventoryTrackers)
	s.Router.GET("/get", s.GetAllInventoryTrackers)
	s.Router.GET("/users/:id", s.GetInventoryTrackers)
	s.Router.PUT("/updateuser/:id", s.UpdateInventoryTrackers)
	s.Router.DELETE("/deleteuser/:id", s.DeleteInventoryTrackers)
	
	

}
