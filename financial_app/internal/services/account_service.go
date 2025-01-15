package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	idCounterAccount int
	idMutexAccount   sync.Mutex
)

func ValidateAccountId(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountIDObject(pool)
	idCounterAccount = Id
}

func CreateAccountObject(pool *pgxpool.Pool) {
	accountId := GenerateAccountID()
	account := models.Account{
		AccountID:     accountId,
		CustomerID:    1,
		AccountTypeID: 1,
		Balace:        1000.0,
	}
	repositories.InsertAccountObject(pool, account)
}

// Genera un ID único de manera segura
func GenerateAccountID() int {
	idMutexAccount.Lock()         // Bloquea el Mutex para evitar accesos concurrentes
	defer idMutexAccount.Unlock() // Asegura que el Mutex se libere después de la función
	idCounterAccount++            // Incrementa el contador global
	return idCounterAccount       // Retorna el nuevo ID
}
