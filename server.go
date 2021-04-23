package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type employee struct {
	Name string `json:"name"`
	Year string `json:"year"`
}
type Customer struct {
	Name string `json:"name"`
	Year string `json:"year"`
}
type Product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func inicio(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func getCustomer(c echo.Context) error {
	customerName := c.QueryParam("name")
	customerYears := c.QueryParam("year")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("nombre del cliente %s\ny su edad es %s", customerName, customerYears))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": customerName,
			"year": customerYears,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "elija json o string",
	})
}

func getEmployees(c echo.Context) error {
	employeeName := c.QueryParam("name")
	employeeYears := c.QueryParam("year")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("nombre del empleado %s\ny su edad es %s", employeeName, employeeYears))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": employeeName,
			"year": employeeYears,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "elija json o string",
	})
}

func getProducts(c echo.Context) error {
	productName := c.QueryParam("name")
	productPrice := c.QueryParam("price")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("nombre del producto %s\ny su precio es %s", productName, productPrice))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name":  productName,
			"price": productPrice,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "elija json o string",
	})
}

func addEmployee(c echo.Context) error {
	employee := employee{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("no se pudo leer el request body por addEmployee: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &employee)
	if err != nil {
		log.Printf("fallo al desarmar en addEmployee: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("Este es tu empleado %#v", employee)
	return c.String(http.StatusOK, "tenemos tu empleado")
}

func addCustomer(c echo.Context) error {
	customer := Customer{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Printf("fallo al procesar addCustomer: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("Este es un cliente: %#v", customer)
	return c.String(http.StatusOK, "tenemos tu cliente")
}

func addProducts(c echo.Context) error {
	product := Product{}

	err := c.Bind(&product)
	if err != nil {
		log.Printf("fallo al procesar addProducts: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("Este es un producto: %#v", product)
	return c.String(http.StatusOK, "tenemos tu producto")
}

func main() {
	e := echo.New()
	e.GET("/", inicio)
	e.GET("/employees/:id", getEmployees)
	e.GET("/customers/:id", getCustomer)
	e.GET("/products/:id", getProducts)
	e.POST("/employees", addEmployee)
	e.POST("/customers", addCustomer)
	e.POST("/products", addProducts)
	e.Logger.Fatal(e.Start(":1323"))
}
