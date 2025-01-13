package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getConnectionString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%v/%s",
		viper.Get("database.admin_user"),
		viper.Get("database.admin_pass"),
		viper.Get("database.host"),
		viper.GetInt("database.port"),
		viper.Get("database.database_name"))
}

func PgxConnection() (*pgx.Conn, context.Context) {
	var connStr string = getConnectionString()
	contxt := context.Background()
	conn, err := pgx.Connect(contxt, connStr)
	if err != nil {
		logrus.Errorf("%v Unable to connect to database: %v", os.Stderr, err)
		os.Exit(1)
	}
	return conn, contxt
}

func PgxPoolConnection() (*pgxpool.Pool, error) {
	var connStr string = getConnectionString()
	ctx := context.Background()
	// Configurar el pool con opciones adicionales
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logrus.Error("Error al parsear la configuración de la conexión:", err)
		return nil, err
	}
	// Ajustar parámetros opcionales del pool (ajusta según tu carga)
	config.MaxConns = 5        // Máximo de conexiones simultáneas
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
