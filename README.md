# experiencia_laboral_crud
API de registro de experiencia laboral

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `experiencia_laboral_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/experiencia_laboral_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `EXPERIENCIA_LABORAL_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `EXPERIENCIA_LABORAL_CRUD__PGUSER`: Usuario de la base de datos
 - `EXPERIENCIA_LABORAL_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `EXPERIENCIA_LABORAL_CRUD__PGURLS`: Host de conexión
 - `EXPERIENCIA_LABORAL_CRUD__PGDB`: Nombre de la base de datos
 - `EXPERIENCIA_LABORAL_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
EXPERIENCIA_LABORAL_HTTP_PORT=8099 EXPERIENCIA_LABORAL_CRUD__PGUSER=user EXPERIENCIA_LABORAL_CRUD__PGPASS=password EXPERIENCIA_LABORAL_CRUD__PGURLS=localhost EXPERIENCIA_LABORAL_CRUD__PGDB=bd EXPERIENCIA_LABORAL_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![modelo_experiencia_laboral_crud](https://user-images.githubusercontent.com/14035745/66067295-e07c7100-e510-11e9-98f5-8234185d2b82.png)
