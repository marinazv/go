package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	producto "github.com/marinazv/go/clase10/internal"
)

const (
	puerto = ":8080"
)

var (
	valueContext any = "user"
)

func main() {

	storage := producto.Storage{
		Productos: loadData(),
	}

	storage.PrintInfo()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	})

	router.GET("/productos/search", func(ctx *gin.Context) {

		precioQuery := ctx.Query("priceGt")
		user := ctx.Query("user")

		if precioQuery != "" {
			precio, err := strconv.ParseFloat(precioQuery, 64)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"mensaje": "precio invalido",
				})
				return
			}

			data := storage.GetProductosMayorPrecio(precio)
			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
			return
		}

		nuevoContexto := addToContext(ctx, user)

		ctx.JSON(http.StatusOK, gin.H{
			"data": storage.GetAll(nuevoContexto),
		})
	})

	router.Run(puerto)

}

func loadData() []producto.Producto {
	productos := []producto.Producto{
		{
			Id:          1,
			Name:        "Banana",
			CodeValue:   "AABBCCC",
			Quantity:    10,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			Id:          2,
			Name:        "Manzana",
			CodeValue:   "AABBDDD",
			Quantity:    5,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.0,
		},
		{
			Id:          3,
			Name:        "Pera",
			CodeValue:   "AAZZZCCC",
			Quantity:    8,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.0,
		},
	}

	return productos
}

func addToContext(ctx context.Context, user string) context.Context {
	nuevoContexto := context.WithValue(ctx, valueContext, user)
	return nuevoContexto
}
