/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"financial_app/cmd"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

// Configuración de las conexiones a las bases de datos
var dbConfigs = []string{
	"postgres://user:password@192.168.0.140:5432/wolfgres_db?sslmode=disable",
}

// Estructura de conexión
type DBConnection struct {
	DB *sql.DB
}

// Variable que mantiene el último ID generado
var idCounter int64 = 0
var idMutex sync.Mutex

// Función para generar IDs incrementales de manera concurrente
func generateIDAccount() int64 {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
	return idCounter
}

// Obtener el último ID de la tabla account
func getLastID(dbConn *DBConnection) int64 {
	var lastID int64
	err := dbConn.DB.QueryRow("SELECT COALESCE(MAX(account_id), 0) FROM wfg.account").Scan(&lastID)
	if err != nil {
		log.Fatalf("Error al obtener el último ID: %v", err)
	}
	return lastID
}

// Conexión a la base de datos
func connectToDB(dsn string) (*DBConnection, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbConn := &DBConnection{DB: db}

	// Obtener el último ID de la tabla Account
	lastID := getLastID(dbConn)
	idCounter = lastID

	return dbConn, nil
}

// Realizar una transacción con escritura
func performTransaction(dbConn *DBConnection) {
	account_id := generateIDAccount()

	// Iniciar la transacción
	tx, err := dbConn.DB.Begin()
	if err != nil {
		log.Printf("Error al comenzar transacción: %v", err)
		return
	}

	// Ejecutar operación de escritura dentro de la transacción
	_, err = tx.Exec("INSERT INTO wfg.account (account_id, customer_id, account_type_id, balace) VALUES ($1, $2, $3, $4)", account_id,123, 1, 1000.0)
	if err != nil {
		log.Printf("Error al ejecutar operación en transacción: %v", err)
		tx.Rollback()
		return
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		log.Printf("Error al hacer commit de la transacción: %v", err)
	}
}

// Función principal que genera estrés a través de transacciones concurrentes
func stressTest(dbConn *DBConnection) {
	var wg sync.WaitGroup
	numTransactions := 1000 // Número de transacciones
	maxConcurrent := 50    // Número máximo de transacciones concurrentes

	// Control de concurrencia con un canal
	semaphore := make(chan struct{}, maxConcurrent)

	for i := 0; i < numTransactions; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // Bloquear si hay demasiadas transacciones en progreso

		go func() {
			defer wg.Done()
			performTransaction(dbConn)
			<-semaphore // Liberar espacio en el canal
		}()
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
}

func main() {
	cmd.Execute()

	// Conectar a una base de datos
	dbConn, err := connectToDB(dbConfigs[0])
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer dbConn.DB.Close()

	// Iniciar el estrés
	fmt.Println("Generando estrés con transacciones concurrentes...")
	stressTest(dbConn)

	fmt.Println("Prueba de estrés completada.")
}

