# Wolfgres Finalcial App (Demo)

Esta aplicación nos permitirá realizar una prueba de estrés a una base de datos postgresql, que nos permitirá especificar el numero de transacciones por tabla.

## Configuración

1. Antes de comenzar se deberá crear el archivo de configuración `config.yaml` en el directorio `financial_app/config/`.
2. Copiar el contenido del archivo `config_example.yaml` al nuevo archivo `config.yaml` y modificar los parametros de conexión a la base de datos.

```
database:
  host: 127.0.0.1
  port: 5432
  admin_user: postgres
  admin_pass: postgres
  database_name: mydatabase
```

## Ejecución

Para ejecutar la aplicación debemos seguir los siguientes pasos:

1. Tener previamente configurado el archivo de configuración `config.yaml` especificado anteriormente.
2. Compilar la aplicación con el comando siguiente:

```
go build -o ./bin/wfg_financial_app

```
3. Ejecutar la aplicación con el siguiente comando:
```
./bin/wfg_financial_app start
```

