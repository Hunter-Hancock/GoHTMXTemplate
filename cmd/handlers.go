package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var store DB

type PageData struct {
	Orders []*Order
}

type Server struct {
	Data  PageData
	store DB
}

func NewServer(store *Store) *Server {
	return &Server{
		store: store,
		Data:  PageData{},
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.Renderer = NewTemplate()

	e.GET("/", s.HandleHome)
	e.GET("/orders", s.HandleGetAllOrders)
	e.GET("/createOrder", s.HandleCreateOrder)

	e.POST("/createOrder", s.HandleCreateOrder)
	e.DELETE("/deleteOrder/:id", s.HandleDeleteOrder)

	e.Logger.Fatal(e.Start(":3000"))
}

func (s *Server) HandleHome(c echo.Context) error {

	orders, err := s.store.GetAllOrders()
	if err != nil {
		return fmt.Errorf("Error Getting Orders %s", err)
	}

	s.Data.Orders = orders
	data := s.Data

	return c.Render(http.StatusOK, "index", data)
}

func (s *Server) HandleGetAllOrders(c echo.Context) error {
	orders, err := s.store.GetAllOrders()
	if err != nil {
		fmt.Println("Error Getting Orders ", err)
	}

	s.Data.Orders = orders
	data := s.Data

	return c.Render(http.StatusOK, "orders", data)
}

func (s *Server) HandleCreateOrder(c echo.Context) error {
	r := c.Request()
	if r.Method == "GET" {
		return c.Render(http.StatusOK, "CreateOrder", nil)
	}

	if r.Method == "POST" {
		method := c.FormValue("method")
		order, err := s.store.CreateOrder(method)
		if err != nil {
			return err
		}

		// orders, err := s.store.GetAllOrders()
		// if err != nil {
		// 	return fmt.Errorf("Error Getting Orders %s", err)
		// }

		s.Data.Orders = append(s.Data.Orders, order)
		data := s.Data

		return c.Render(http.StatusOK, "orders", data)
	}

	return nil
}

func (s *Server) HandleDeleteOrder(c echo.Context) error {
	id := c.Param("id")
	orderId, _ := strconv.Atoi(id)

	fmt.Println("Deleting Order With Id: ", orderId)

	err := s.store.DeleteOrder(orderId)
	if err != nil {
		return fmt.Errorf("Error deleting: %s", err)
	}

	orders := s.Data.Orders

	for i, order := range orders {
		if order.Id == orderId {
			orders = append(orders[:i], orders[i+1:]...)
			break
		}
	}

	s.Data.Orders = orders
	data := s.Data

	return c.Render(http.StatusOK, "orders", data)
}
