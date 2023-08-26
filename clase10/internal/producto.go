package producto

import (
	"context"
	"fmt"
	"time"
)

type Producto struct {
	Id          int       `json:"id"`
	Name        string    `json:"nombre"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_plublished"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

type Storage struct {
	Productos []Producto
}

func (s *Storage) PrintInfo() {
	fmt.Println(s.Productos)
}

func (s *Storage) GetAll(ctx context.Context) []Producto {

	user, ok := ctx.Value("user").(string)
	if ok && user != "" {
		fmt.Println("Valor de contexto en package producto:", user)
	}
	return s.Productos
}

func (s *Storage) GetProductosMayorPrecio(precio float64) []Producto {
	var resultado []Producto

	for _, producto := range s.Productos {
		if producto.Price >= precio {
			resultado = append(resultado, producto)
		}
	}

	return resultado
}