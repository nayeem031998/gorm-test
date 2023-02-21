package controller

import "database/api/middleware"

func (s *Server) InitRoutes() {
	s.Router.GET("/products", s.GetProducts)
	s.Router.GET("/products/:id", s.GetProduct)
	s.Router.GET("/product/:name", s.GetProductsByName)
	s.Router.POST("/createproduct", middleware.SetMiddlewareFormHeader, s.CreateProduct)
	s.Router.PUT("/updateproduct/:id", middleware.SetMiddlewareFormHeader, s.UpdateProducts)
	s.Router.DELETE("/deleteproduct/:id", middleware.SetMiddlewareFormHeader, s.DeleteProducts)
	s.Router.POST("/addorder", middleware.SetMiddlewareFormHeader, s.CreateOrder)
	s.Router.GET("/orders", s.GetOrders)
	s.Router.GET("/orders/:id", s.GetOrder)
	s.Router.GET("/order/:id", s.GetOrderById)
	s.Router.PUT("/updateorder/:id", middleware.SetMiddlewareFormHeader, s.UpdateOrders)
	s.Router.DELETE("/deleteorder/:id", middleware.SetMiddlewareFormHeader, s.DeleteOrders)
	s.Router.POST("/createtransaction", middleware.SetMiddlewareJSON, s.CreateTransactions)
	s.Router.GET("/transactions/:id", s.GetTransactions)
	s.Router.GET("/alltransactions", s.GetAllTransactions)
	s.Router.PUT("/updatetransaction/:id", middleware.SetMiddlewareJSON, s.UpdateTransactions)
	s.Router.DELETE("/deletetransaction/:id", middleware.SetMiddlewareJSON, s.DeleteTransactions)
	s.Router.POST("/createuser", middleware.SetMiddlewareJSON, s.CreateInventoryTrackers)
	s.Router.GET("/get", s.GetAllInventoryTrackers)
	s.Router.GET("/users/:id", s.GetInventoryTrackers)
	s.Router.PUT("/updateuser/:id", middleware.SetMiddlewareJSON, s.UpdateInventoryTrackers)
	s.Router.DELETE("/deleteuser/:id", middleware.SetMiddlewareJSON, s.DeleteInventoryTrackers)

}
