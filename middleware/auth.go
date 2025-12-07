package middleware

import (
	"net/http"
	"productcrud/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token nao informado!!"})
			ctx.Abort()
			return
		}

		// Espera "Bearer token_aqui"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Formato token invalido!!"})
			ctx.Abort()
			return
		}

		tokenStr := parts[1]
		token, claims, err := auth.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalido!!"})
			ctx.Abort()
			return
		}

		// pega o "sub" que colocamos no GenerateToken
		// o claims é o objeto que criamos com a estrutura/config do token
		sub, ok := claims["sub"].(float64) // vem como float64
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido!!"})
			ctx.Abort()
			return
		}

		userID := int(sub)

		// joga no contexto pra usar nos controllers
		ctx.Set("userId", userID)

		ctx.Next()

	}
}
