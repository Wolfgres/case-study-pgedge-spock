# Prueba de stress mediante la  pgbeanch
 El presente documento presenta una serie de pruebas empleadas para medir las diferencias entre un cluster con una configuración Primario -> Secundario mediante Streaming y un cluster con una configuración Primario -> Primario empleando replicación mediante Spock. 



# Pasos a seguir para la ejecución de la prueba
A continuación se describen los pasos a seguir para la carga de múltiples operaciones INSERT sobre el cluster Priario -> Primario con Spock: 

 1. Definir de la configuración del cluster

     > {
      "json_version": "1.0",
      "cluster_name": "wfg_cluster",
      "log_level": "debug",
      "update_date": "2025-01-22 15:20:00GMT",
      "pgedge": {
        "pg_version": 16,
        "auto_start": "off",
        "spock": {
          "spock_version": "",
          "auto_ddl": "on"
        },
        "databases": [
          {
            "db_name": "wolfgres_db",
            "db_user": "wolfgres_user",
            "db_password": 123456
          }
        ]
      },
      "node_groups": [
        {
          "ssh": {
            "os_user": "spock",
            "private_key": ""
          },
          "name": "n1",
          "is_active": "on",
          "public_ip": "192.168.0.97",
          "private_ip": "127.0.0.1",
          "port": "5432",
          "path": "/home/spock/wfg_cluster/n1",
          "backrest": {
            "stanza": "demo_stanza",
            "repo1-path": "/var/lib/pgbackrest",
            "repo1-retention-full": "7",
            "log-level-console": "info",
            "repo1-cipher-type": "aes-256-cbc"
          }
        },
        {
          "ssh": {
            "os_user": "spock",
            "private_key": ""
          },
          "name": "n2",
          "is_active": "on",
          "public_ip": "192.168.0.98",
          "private_ip": "127.0.0.1",
          "port": "5432",
          "path": "/home/spock/wfg_cluster/n2",
          "backrest": {
            "stanza": "demo_stanza",
            "repo1-path": "/var/lib/pgbackrest",
            "repo1-retention-full": "7",
            "log-level-console": "info",
            "repo1-cipher-type": "aes-256-cbc"
          }
        }
      ]
    }

 2. Inicializar el cluster

   > ./pgedge cluster init wfg_cluster

 3. Ejecutar carga de esquema de base de datos

   > psql -h 192.168.0.97 -U wolfgres_user -d wolfgres_db < wolfgres_db_financial.sql 

 4. Configurar el archivo .pgpass para el acceso a los nodos de bd
   > echo '*:5432:wolfgres_db:wolfgres_user:123456' >> ~/.pgpass
   > chmod 0600 ~/.pgpass

 5. Agregar las siguientes lineas en el archivo /etc/hosts del cliente

   > 192.168.0.97  n1
   > 192.168.0.98  n2

 6. Ejecutar  pgbench con la siguiente configuración y script
        personalizado
> pgbench --client=2 \
    --connect \
    --jobs=2 \
    --time=60 \
    --progress=5 \
    --file=load_wolfgres_db.sql \
    "host=n1,n2 port=5432,5432 load_balance_hosts=random dbname=wolfgres_db user=wolfgres_user"
    

 7. Al finalizar se mostrara una salida similar a la siguiente

   > transaction type: load_wolfgres_db.sql
    scaling factor: 1
    query mode: simple
    number of clients: 2
    number of threads: 2
    maximum number of tries: 1
    duration: 60 s
    number of transactions actually processed: 54
    number of failed transactions: 0 (0.000%)
    latency average = 287.619 ms
    latency stddev = 118.373 ms
    average connection time = 1993.651 ms
    tps = 0.873632 (including reconnection times)
