package model

// Definindo a estrutura do objeto
// Tambem definindo como sera lido para cpnverter em json
type User struct {
	ID       int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
