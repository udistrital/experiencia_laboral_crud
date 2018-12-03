# experiencia_laboral_crud

--Api de personas con CI--
CI deploy with lambda - S3
Drone 0.8 
experiencia_laboral_crud master/develop

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/experiencia_laboral_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_EXPERIENCIA_LABORAL_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `EXPERIENCIA_LABORAL_CRUD__PGUSER`: Usuario de la base de datos
 - `EXPERIENCIA_LABORAL_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `EXPERIENCIA_LABORAL_CRUD__PGURLS`: Host de conexión
 - `EXPERIENCIA_LABORAL_CRUD__PGDB`: Nombre de la base de datos
 - `EXPERIENCIA_LABORAL_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_EXPERIENCIA_LABORAL_HTTP_PORT=8083 EXPERIENCIA_LABORAL_CRUD__PGUSER=user EXPERIENCIA_LABORAL_CRUD__PGPASS=password EXPERIENCIA_LABORAL_CRUD__PGURLS=localhost EXPERIENCIA_LABORAL_CRUD__PGDB=udistrital_core_db EXPERIENCIA_LABORAL_CRUD__SCHEMA=core_new bee run

## MODELO
![image]().
