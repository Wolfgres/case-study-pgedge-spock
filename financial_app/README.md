# Wolfgres Finalcial App (Demo)

Esta aplicación nos permitirá realizar una prueba de estrés a una base de datos postgresql, que nos permitirá especificar el numero de transacciones por tabla.

## Configuración

1. Antes de comenzar se deberá crear el archivo de configuración `config.yaml` en el directorio `financial_app/config/`.
2. Copiar el contenido del archivo `config_example.yaml` al nuevo archivo `config.yaml` y modificar los parametros de conexión a la base de datos.

```
database:
  node_1:
    host: localhost
    port: 5432
    admin_user: postgres
    admin_pass: postgres
    database_name: wolfgres_db
```

## Ejecución

Para ejecutar la aplicación debemos seguir los siguientes pasos:

1. Tener previamente configurado el archivo de configuración `config.yaml` especificado anteriormente.
2. (Opcional) Compilar la aplicación con el comando siguiente:

```
go build -o ./bin/wfg_financial_app
```

3. Ejecutar la aplicación con los siguientes comando:

3.1 Con archivo compilado:

```
./bin/wfg_financial_app [flags]
```

3.2 Con el archivo `main.go`

```
go run main.go [flags]
```

3.1 Flags

```
  -d, --duration int       Test duration in seconds
  -g, --goroutines int     Number of concurrent goroutines
  -h, --help               help for financial_app
  -m, --max-conns int      Maximum number of connections in the pool
  -o, --operation int      Choose a stress test transactions: INSERT=1, SELECT=2, UPDATE=3
  -t, --transactions int   Number of transactions. Must be a number that are multiples of goroutine
```

4. Ejemplos de ejecución.

4.1 Ejecutar por número de transacciones por tabla:

```
go run main.go -m 50 -g 20 -t 100 -o 1
```

4.2 Ejecutar por duración en segundos:

```
go run main.go -m 50 -g 20 -d 60 -o 1
```
