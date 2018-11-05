-- data for testing


-- CREATE DATABASE taxi_traker;

DROP TABLE IF EXISTS taxi_service;

DROP TABLE IF EXISTS vehicle_driver;

DROP TABLE IF EXISTS vehicle_position_history;

DROP TABLE IF EXISTS vehicle;

DROP TABLE IF EXISTS driver;

DROP TABLE IF EXISTS customer;

CREATE TABLE vehicle(
	id 			VARCHAR(8)  NOT NULL,
	make		VARCHAR(50) NOT NULL,
	model		VARCHAR(50) NOT NULL,
	"year"		INT2		NOT NULL,
	latitude	FLOAT		NOT NULL,
	longitude	FLOAT		NOT NULL,
	CONSTRAINT pk_vehicle PRIMARY KEY (id)
);

CREATE TABLE driver(
	id 			VARCHAR(36) NOT NULL,
	first_name 	VARCHAR(50) NOT NULL,
	last_name 	VARCHAR(50)	NOT NULL,
	email 		VARCHAR(320) NOT NULL,
	password 	TEXT		NOT NULL,
	CONSTRAINT pk_driver PRIMARY KEY(id)
);

CREATE TABLE vehicle_driver(
	vehicle_id VARCHAR(8) NOT NULL,
	driver_id VARCHAR(36) NOT NULL,
	CONSTRAINT pk_vehicle_driver PRIMARY KEY(vehicle_id, driver_id),
	CONSTRAINT fk_vehicle_vehicledriver_id FOREIGN KEY(vehicle_id)
		REFERENCES vehicle(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	CONSTRAINT fk_driver_vehicledriver_id FOREIGN KEY(driver_id)
		REFERENCES driver(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

CREATE TABLE customer (
	id 			VARCHAR(36) NOT NULL,
	first_name 	VARCHAR(50) NOT NULL,
	last_name 	VARCHAR(50) NOT NULL,
	email 		VARCHAR(320) NOT NULL,
	password 	TEXT		NOT NULL,
	CONSTRAINT pk_costumer PRIMARY KEY(id)
);

CREATE TABLE vehicle_position_history(
	id			SERIAL		NOT NULL,
	vehicle_id	VARCHAR(8)	NOT NULL,
	latitude	FLOAT		NOT NULL,
	longitude	FLOAT		NOT NULL,
	"date" 		TIMESTAMPTZ	NOT NULL DEFAULT NOW(),
	CONSTRAINT pk_historical PRIMARY KEY(id),
	CONSTRAINT fk_vehicle_vehiclepositionhistory_id FOREIGN KEY(vehicle_id)
		REFERENCES vehicle(id)
		ON UPDATE CASCADE
		ON DELETE CASCADE
);

CREATE TABLE taxi_service(
	customer_id VARCHAR(36) NOT NULL,
	vehicle_id 	VARCHAR(8)  NOT NULL,
	CONSTRAINT pk_taxi_service PRIMARY KEY(customer_id, vehicle_id),
	CONSTRAINT fk_customer_taxiservice_id FOREIGN KEY(customer_id)
		REFERENCES customer(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	CONSTRAINT fk_vehicle_taxiservice_id FOREIGN KEY(vehicle_id)
		REFERENCES vehicle(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

-- =========================================================================================
-- =========================================================================================
--
-- Functions:

/* -----------------------------------------------------------------------------------------
 * Result Codes:
 * Succes codes - httpStatusOk: 200
 * 0 = OK
 
 * User Errors - httpStatusOk: 200
 * -1 = EUS001 - Email already exists 
*/
DROP FUNCTION IF EXISTS customer_insert(VARCHAR, VARCHAR, VARCHAR, VARCHAR, TEXT);

CREATE OR REPLACE FUNCTION customer_insert(
        p_id 			VARCHAR(36),
        p_first_name 	VARCHAR(50),
        p_last_name 	VARCHAR(50),
        p_email 		VARCHAR(320),
        p_password		TEXT
)
RETURNS INT AS
$$
DECLARE
    counter INT2;
BEGIN    
    -- ***** *************************************************************
    -- ***** Validations: User Errors - httpStatusOk: 200            *****
    -- ***** *************************************************************
    -- If the email is not null, validate that not exist
    IF p_email IS NOT NULL THEN
        SELECT COUNT(*) INTO counter FROM customer WHERE email = p_email;
        IF counter <> 0 THEN
            Return -1;
        END IF;
    END IF;

    INSERT INTO customer
        ( id,  first_name,   last_name,   email,   password)
    VALUES
        (p_id, p_first_name, p_last_name, p_email, p_password); 
    RETURN 0;   
END    
$$ LANGUAGE plpgsql;






