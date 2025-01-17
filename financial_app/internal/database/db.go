package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getConnectionString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%v/%s",
		viper.Get("database.node_1.admin_user"),
		viper.Get("database.node_1.admin_pass"),
		viper.Get("database.node_1.host"),
		viper.GetInt("database.node_1.port"),
		viper.Get("database.node_1.database_name"))
}

func pgxPoolConnection(maxConns int32) (*pgxpool.Pool, error) {
	var connStr string = getConnectionString()
	ctx := context.Background()
	// Configurar el pool con opciones adicionales
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logrus.Error("Error al parsear la configuración de la conexión:", err)
		return nil, err
	}
	// Ajustar parámetros opcionales del pool (ajusta según tu carga)
	config.MaxConns = maxConns // Máximo de conexiones simultáneas
	config.MinConns = 1        // Mínimo de conexiones activas
	config.MaxConnLifetime = 0 // No cerrar conexiones automáticamente

	// Crear el pool de conexiones
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logrus.Error("Error al conectar con el pool:", err)
		return nil, err
	}

	return pool, nil
}

func InitDatabasePool(maxConns int, numGoroutines int) *pgxpool.Pool {
	if maxConns <= 0 {
		logrus.Fatal("Please, enter a valid number for max-conns")
	}
	if maxConns < numGoroutines {
		logrus.Fatal("max-conns must be >= goroutines")
	}
	maxConnsInt32 := int32(maxConns)
	pool, err := pgxPoolConnection(maxConnsInt32)
	if err != nil {
		logrus.Fatal("No se pudo conectar a la base de datos:", err)
	}
	return pool
}
