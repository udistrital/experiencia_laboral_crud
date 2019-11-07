-- object: experiencia_laboral | type: SCHEMA --
-- DROP SCHEMA IF EXISTS experiencia_laboral CASCADE;
CREATE SCHEMA experiencia_laboral;
-- ddl-end --
COMMENT ON SCHEMA experiencia_laboral IS 'Esquema para el módulo de experiencia laboral del campus';
-- ddl-end --

SET search_path TO pg_catalog,public,experiencia_laboral;
-- ddl-end --

-- object: experiencia_laboral.cargo | type: TABLE --
-- DROP TABLE IF EXISTS experiencia_laboral.cargo CASCADE;
CREATE TABLE experiencia_laboral.cargo (
	id serial NOT NULL,
	nombre character varying(120) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(6,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_cargo PRIMARY KEY (id),
	CONSTRAINT uq_nombre_cargo UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE experiencia_laboral.cargo IS 'Tabla paramétrica para guardar los cargos de la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.id IS 'Identificador de la tabla paramétrica cargo';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.nombre IS 'Nombre del cargo desempeñado';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.descripcion IS 'Descripción opcional del parámetro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.activo IS 'Campo que indica si el parámetro está activo';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.cargo.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_cargo ON experiencia_laboral.cargo  IS 'Llave primaria de la tabla cargo';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_cargo ON experiencia_laboral.cargo  IS 'Restricción para que no se repita el nombre de los cargos';
-- ddl-end --

-- object: experiencia_laboral.dato_adicional_experiencia_laboral | type: TABLE --
-- DROP TABLE IF EXISTS experiencia_laboral.dato_adicional_experiencia_laboral CASCADE;
CREATE TABLE experiencia_laboral.dato_adicional_experiencia_laboral (
	id serial NOT NULL,
	tipo_dato_adicional integer NOT NULL,
	experiencia_laboral integer NOT NULL,
	valor character varying(50) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_dato_adicional_experiencia_laboral PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE experiencia_laboral.dato_adicional_experiencia_laboral IS 'Tabla que almacena el valor de los datos adicionales de la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.dato_adicional_experiencia_laboral.id IS 'Identificador de la tabla dato_adicional_experiencia_laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.dato_adicional_experiencia_laboral.tipo_dato_adicional IS 'Tipo de dato adicional para la experiencia laboral, referencia a la tabla tipo_dato_adicional del core';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.dato_adicional_experiencia_laboral.valor IS 'Valor para el datos adicional de la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.dato_adicional_experiencia_laboral.activo IS 'Campo para indicar el estado del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.dato_adicional_experiencia_laboral.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.dato_adicional_experiencia_laboral.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_dato_adicional_experiencia_laboral ON experiencia_laboral.dato_adicional_experiencia_laboral  IS 'Llave primaria de la tabla dato_adicional_experiencia_laboral';
-- ddl-end --

-- object: experiencia_laboral.experiencia_laboral | type: TABLE --
-- DROP TABLE IF EXISTS experiencia_laboral.experiencia_laboral CASCADE;
CREATE TABLE experiencia_laboral.experiencia_laboral (
	id serial NOT NULL,
	persona integer NOT NULL,
	actividades text,
	organizacion integer NOT NULL,
	fecha_inicio date NOT NULL,
	fecha_finalizacion date,
	tipo_vinculacion integer NOT NULL,
	tipo_dedicacion integer NOT NULL,
	cargo integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_experiencia_docente PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE experiencia_laboral.experiencia_laboral IS 'Experiencia laboral de la persona';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.id IS 'Identificador de la tabla experiencia_laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.persona IS 'Referencia a la tabla ente';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.actividades IS 'Descripción de las actividades realizadas en la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.organizacion IS 'Referencia a la tabla organizacion del core';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.fecha_inicio IS 'Fecha en la que inicio la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.fecha_finalizacion IS 'Fecha en la que finalizó la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.experiencia_laboral.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_experiencia_docente ON experiencia_laboral.experiencia_laboral  IS 'Llave primaria de la tabla experiencia_laboral';
-- ddl-end --

-- object: experiencia_laboral.soporte_experiencia_laboral | type: TABLE --
-- DROP TABLE IF EXISTS experiencia_laboral.soporte_experiencia_laboral CASCADE;
CREATE TABLE experiencia_laboral.soporte_experiencia_laboral (
	id serial NOT NULL,
	documento integer NOT NULL,
	experiencia_laboral integer NOT NULL,
	descripcion character varying(250),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_soporte_experiencia_laboral PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE experiencia_laboral.soporte_experiencia_laboral IS 'Tabla que almacena los soportes correspondientes a la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.soporte_experiencia_laboral.id IS 'Identificador de la tabla soporte_experiencia_laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.soporte_experiencia_laboral.documento IS 'Referencia a la tabla documento del microservicio de documentos';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.soporte_experiencia_laboral.descripcion IS 'Descripción del soporte asociado a la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.soporte_experiencia_laboral.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.soporte_experiencia_laboral.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.soporte_experiencia_laboral.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_soporte_experiencia_laboral ON experiencia_laboral.soporte_experiencia_laboral  IS 'Llave primaria de la tabla soporte_experiencia_laboral';
-- ddl-end --

-- object: experiencia_laboral.tipo_dedicacion | type: TABLE --
-- DROP TABLE IF EXISTS experiencia_laboral.tipo_dedicacion CASCADE;
CREATE TABLE experiencia_laboral.tipo_dedicacion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_dedicacion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_dedicacion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE experiencia_laboral.tipo_dedicacion IS 'Tabla paramétrica para la definición de los tipos de dedicación de la experiencia laboral, ej: tiempo completo, medio tiempo, tiempo parcial, otra';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.id IS 'Identificador de la tabla tipo_dedicacion';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.nombre IS 'Campo en el que se indica el nombre del tipo de dedicación de la experiencia laboral';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.descripcion IS 'Descripción opcional del parámetro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.activo IS 'Campo que indica si el parámetro está activo';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_dedicacion.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_dedicacion ON experiencia_laboral.tipo_dedicacion  IS 'Llave primaria de la tabla tipo_dedicacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_dedicacion ON experiencia_laboral.tipo_dedicacion  IS 'Restricción para que no se repita el nombre de un tipo de dedicación';
-- ddl-end --

-- object: experiencia_laboral.tipo_vinculacion | type: TABLE --
-- DROP TABLE IF EXISTS experiencia_laboral.tipo_vinculacion CASCADE;
CREATE TABLE experiencia_laboral.tipo_vinculacion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_vinculacion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_vinculacion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE experiencia_laboral.tipo_vinculacion IS 'Tabla paramétrica con los tipos de vinculación (prestacion de servicios,a termino indefinido, de aprendizaje, vinculación especial)';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.id IS 'Identificador de la tabla tipo_vinculacion';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.descripcion IS 'Descripción opcional del parámetro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.activo IS 'Campo que indica si el parámetro esta activo';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN experiencia_laboral.tipo_vinculacion.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_vinculacion ON experiencia_laboral.tipo_vinculacion  IS 'Llave primaria de la tabla tipo_vinculacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_vinculacion ON experiencia_laboral.tipo_vinculacion  IS 'Restricción para que no se repita el nombre de un tipo de vinculación';
-- ddl-end --

-- object: fk_experiencia_laboral_tipo_vinculacion | type: CONSTRAINT --
-- ALTER TABLE experiencia_laboral.experiencia_laboral DROP CONSTRAINT IF EXISTS fk_experiencia_laboral_tipo_vinculacion CASCADE;
ALTER TABLE experiencia_laboral.experiencia_laboral ADD CONSTRAINT fk_experiencia_laboral_tipo_vinculacion FOREIGN KEY (tipo_vinculacion)
REFERENCES experiencia_laboral.tipo_vinculacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_experiencia_laboral_cargo | type: CONSTRAINT --
-- ALTER TABLE experiencia_laboral.experiencia_laboral DROP CONSTRAINT IF EXISTS fk_experiencia_laboral_cargo CASCADE;
ALTER TABLE experiencia_laboral.experiencia_laboral ADD CONSTRAINT fk_experiencia_laboral_cargo FOREIGN KEY (cargo)
REFERENCES experiencia_laboral.cargo (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_dato_adicional_experiencia_laboral_experiencia_laboral | type: CONSTRAINT --
-- ALTER TABLE experiencia_laboral.dato_adicional_experiencia_laboral DROP CONSTRAINT IF EXISTS fk_dato_adicional_experiencia_laboral_experiencia_laboral CASCADE;
ALTER TABLE experiencia_laboral.dato_adicional_experiencia_laboral ADD CONSTRAINT fk_dato_adicional_experiencia_laboral_experiencia_laboral FOREIGN KEY (experiencia_laboral)
REFERENCES experiencia_laboral.experiencia_laboral (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_soporte_experiencia_laboral_experiencia_laboral | type: CONSTRAINT --
-- ALTER TABLE experiencia_laboral.soporte_experiencia_laboral DROP CONSTRAINT IF EXISTS fk_soporte_experiencia_laboral_experiencia_laboral CASCADE;
ALTER TABLE experiencia_laboral.soporte_experiencia_laboral ADD CONSTRAINT fk_soporte_experiencia_laboral_experiencia_laboral FOREIGN KEY (experiencia_laboral)
REFERENCES experiencia_laboral.experiencia_laboral (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_experiencia_laboral_tipo_dedicacion | type: CONSTRAINT --
-- ALTER TABLE experiencia_laboral.experiencia_laboral DROP CONSTRAINT IF EXISTS fk_experiencia_laboral_tipo_dedicacion CASCADE;
ALTER TABLE experiencia_laboral.experiencia_laboral ADD CONSTRAINT fk_experiencia_laboral_tipo_dedicacion FOREIGN KEY (tipo_dedicacion)
REFERENCES experiencia_laboral.tipo_dedicacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

--GRANT USAGE ON SCHEMA experiencia_laboral TO desarrollooas;
--GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA experiencia_laboral TO desarrollooas;
--GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA experiencia_laboral TO desarrollooas;