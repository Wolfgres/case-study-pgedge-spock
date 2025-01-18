# Despliegue de cluster pgedge/Spock  
 Esta es una guiá para desplegar un cluster con dos nodos en pgedge con la extensión spock base para la prueba de la aplicación financial app.  

# Clonar proyecto pgedge-docker
- En caso de aun no contar con el proyecto  se deberá de clonar el siguiente repositorio
	> git clone https://github.com/pgEdge/pgedge-docker.git

# Ejecutar ejemplo docker swarm

Una vez clonado el proyecto, los pasos para paso para poder efectuar el despliege de un cluster de pgedge son los soguientes:

- Acceder a la carpeta  **pgedge-docker -> examples -> swarm**
	> cd pgedge-docker 
	> cd examples
	> cd swarm
- Dentro de la carpeta swarm deberá existir un archivo con el nombre Makefile el cual contiene las instrucciones para generar el despliegue del cluster mismas que resumen a continuación
  - Almacenar y exportar dentro de las variables de entorno la contraseña correspondiente al usuario admin 
	> export PGPASSWORD=uFR44yr69C4mZa72g3JQ37GX
  - Agregar las siguientes direcciones dentro del archivo hosts del servidor
	> 127.0.0.1 n1  
    > 127.0.0.1 n2
  - Crear las carpetas data, n1 y n2 las cuales servirán para montar los volumenes de cada nodo
	> mkdir -p ./data/n1 ./data/n2
  - Iniciar el despliegue del cluster 
	> docker stack deploy -c ./stack.yaml db
  - Para validar el despliegue exitoso del cluster pgedge ejecutar el siguiente comando
	> docker stack ps db
  - Como lo muestra la imagen los accesos a cada componente del cluster quedarían definidos de la siguiente manera

|       NODO         |PUERTO                      |COMPONENTE                  |
|----------------|--------------------------------|-----------------------------|
|n1							 |:6431                           |Servidor postgresql primario |
|n2          		 |:6441                           |Servidor postgresql primario |
|pgcat1          |:64432                          |Balanceador                  |
|pgcat2          |:6442                           |Balanceador                  |
|traefic         |http://localhost:8080/dashboard                            |Enrouteador                  |

  
  - Una vez validado el despliegue exitoso se podrá acceder a cada nodo y realizar la carga de la base de datos o establecer una conexión mediante los comandos:
	> psql -h node_ip -p 5432 -U admin wolfgres_db < wolfgres_db.sql
	> psql -h node_name -p 5432 -U admin wolfgres_db
  - En caso de mostrar algun error se puede ejecutar los siguientes comandos para hacer una limpieza y volver a intentar el despliegue una vez revisados los pasos anteriores
	> docker stack rm db
	> docker volume rm db_n1
	> docker volume rm db_n2
	> rm -rf data/		
