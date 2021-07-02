package main

import (
		"fmt"
		"log"
		"github.com/labstack/echo"
		"errors"
	)

type model struct {
	Matriz [][]int `json:"Matriz"`
}

type service interface {
	voltearMatriz(data *model) ([][]int, error)
}

type svr struct {}

type handler struct {
	svr service
}

func NewHandler(svr service) *handler {
	return &handler{svr}
}


func Route(e *echo.Echo) {
	r := e.Group("/matriz")
	s := svr{}
	h := NewHandler(s)
	r.POST("", h.procesarMatriz)
}

func(s svr) voltearMatriz(data *model) ([][]int, error) {
	idx := len(data.Matriz)
	arr := make([][]int, idx)

	// la matriz tiene que estar bien definida
	for i := 0; i < idx; i++ {
		if len(data.Matriz[i]) != idx { // ixj
			return nil, errors.New("Matriz no valida !")
		}
		arr[i] = make([]int, idx)
	}

	//girar matriz
	for i := 0; i < idx; i++ {
		for j := 0; j < idx; j++ { // jxi
			arr[j][i] = data.Matriz[i][idx-j-1]
		}
	}

	return arr, nil
}

func (h *handler) procesarMatriz(c echo.Context) error {
	data := model{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(400, "erro !")
	}

	result, err := h.svr.voltearMatriz(&data)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	data.Matriz = result


	return c.JSON(200, data)
}


// funcion principal
func main() {

	e := echo.New()
	Route(e)

	// Levantamos el servidor 
	log.Println("Bienvenido al servidor apiMatriz  H: 8090")
	fmt.Println("Esperando matriz ....")
	err := e.Start(":8090")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
		
}
