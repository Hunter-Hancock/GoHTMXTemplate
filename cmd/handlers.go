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

func HandleHome(c echo.Context) error {
	orders, err := store.GetAllOrders()
	if err != nil {
		return fmt.Errorf("Error Getting Orders %s", err)
	}

	data := PageData{Orders: orders}

	return c.Render(http.StatusOK, "index", data)
}

func HandleGetAllOrders(c echo.Context) error {
	orders, err := store.GetAllOrders()
	if err != nil {
		fmt.Println("Error Getting Orders ", err)
	}

	data := PageData{Orders: orders}

	return c.Render(http.StatusOK, "orders", data)
}

func HandleCreateOrder(c echo.Context) error {
	r := c.Request()
	if r.Method == "GET" {
		return c.Render(http.StatusOK, "CreateOrder", nil)
	}

	if r.Method == "POST" {
		method := c.FormValue("method")
		err := store.CreateOrder(method)
		if err != nil {
			return err
		}

		orders, err := store.GetAllOrders()
		if err != nil {
			return fmt.Errorf("Error Getting Orders %s", err)
		}
		data := PageData{Orders: orders}

		return c.Render(http.StatusOK, "orders", data)
	}

	return nil
}

func HandleDeleteOrder(c echo.Context) error {
	id := c.Param("id")
	orderId, _ := strconv.Atoi(id)

	fmt.Println("Deleting Order With Id: ", orderId)

	err := store.DeleteOrder(orderId)
	if err != nil {
		return fmt.Errorf("Error deleting: %s", err)
	}

	orders, err := store.GetAllOrders()
	if err != nil {
		return fmt.Errorf("Error Getting Orders %s", err)
	}

	for i, order := range orders {
		if order.Id == orderId {
			orders = append(orders[:i], orders[i+1:]...)
			break
		}
	}

	data := PageData{Orders: orders}

	return c.Render(http.StatusOK, "orders", data)
}
