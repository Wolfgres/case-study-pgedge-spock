# Wolfgres Finalcial App (Demo)

Esta aplicación nos permitirá realizar una prueba de estrés a una base de datos postgresql, que nos permitirá especificar el numero de transacciones por tabla.

## Configuración

1. Antes de comenzar se deberá crear el archivo de configuración `config.yaml` en el directorio `financial_app/config/`.
2. Copiar el contenido del archivo `config_example.yaml` al nuevo archivo `config.yaml`.
3. Realizado los pasos anteriores, para la configuración de los parametros, en la sección de `test`se deberá especificar el numero de transacciones y el maximo de transacciones concurrentes. El parametro al que debemos prestar más atención es a `maxCurrent`, ya que este debera estar en el rango especificado por el parametro de configuración `max_connections` de postgresql.

```
database:
  host: 127.0.0.1
  port: 5432
  admin_user: postgres
  admin_pass: postgres
  database_name: mydatabase
test:
  type: stress
  transactions: 1000
  maxCurrent: 50
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

