package database

import (
	"context"
	"financial_app/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetConnectionBodies() []models.Connection {
	nodes := viper.GetStringMap("database")
	var connections []models.Connection
	for index, _ := range nodes {
		var connection models.Connection
		idx := "database." + parseAnyToString(index)
		connection.User = viper.GetString(idx + ".admin_user")
		connection.Password = viper.GetString(idx + ".admin_pass")
		connection.Host = viper.GetString(idx + ".host")
		connection.Port = viper.GetInt(idx + ".port")
		connection.DBName = viper.GetString(idx + ".database_name")
		connections = append(connections, connection)
	}
	return connections
}

func GetConnectionString(connection models.Connection) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%v/%s",
		connection.User,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.DBName)
}

func pgxPoolConnection(maxConns int32, connStr string) (*pgxpool.Pool, error) {
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

func InitDatabasePool(maxConns int, numGoroutines int, connStr string) *pgxpool.Pool {
	if maxConns <= 0 {
		logrus.Fatal("Please, enter a valid number for max-conns")
	}
	if maxConns < numGoroutines {
		logrus.Fatal("max-conns must be >= goroutines")
	}
	maxConnsInt32 := int32(maxConns)
	pool, err := pgxPoolConnection(maxConnsInt32, connStr)
	if err != nil {
		logrus.Fatal("No se pudo conectar a la base de datos:", err)
	}
	return pool
}

func parseAnyToString(value any) string {
	// Convertir usando type assertion
	str, ok := value.(string)
	if !ok {
		logrus.Fatal("Value isn't a string")
	}
	return str
}
