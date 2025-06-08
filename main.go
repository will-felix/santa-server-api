package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura para representar um binário
type Binary struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	User   string `json:"user"`
	Policy string `json:"policy"` // allow, block, monitor
}

var binaries = []Binary{} // nossa "base de dados" temporária

func main() {
	r := gin.Default()

	// Rota de boas-vindas
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Santa está rodando!"})
	})

	// Listar todos os binários
	r.GET("/binaries", func(c *gin.Context) {
		c.JSON(http.StatusOK, binaries)
	})

	// Criar novo binário
	r.POST("/binaries", func(c *gin.Context) {
		var newBinary Binary
		if err := c.ShouldBindJSON(&newBinary); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		binaries = append(binaries, newBinary)
		c.JSON(http.StatusCreated, newBinary)
	})

	// Deletar binário por ID
	r.DELETE("/binaries/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, b := range binaries {
			if b.ID == id {
				binaries = append(binaries[:i], binaries[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Binário removido"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Binário não encontrado"})
	})

	// Iniciar o servidor
	r.Run(":8080") // http://localhost:8080
}
