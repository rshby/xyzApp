CREATE DATABASE xyz_app;

USE xyz_app;

-- create table konsumer
CREATE TABLE konsumer
(
    nik VARCHAR(255) NOT NULL ,
    full_name VARCHAR(255) NOT NULL ,
    legal_name VARCHAR(255) NOT NULL ,
    tempat_lahir VARCHAR(255) NOT NULL ,
    tanggal_lahir TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    gaji DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    foto_ktp VARCHAR(5000) NOT NULL ,
    foto_selfie VARCHAR(5000) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (nik)
) engine = InnoDB;

-- create table accounts
CREATE TABLE accounts
(
    nik VARCHAR(255) NOT NULL ,
    email VARCHAR(255) NOT NULL ,
    password VARCHAR(1000) NOT NULL ,
    PRIMARY KEY (nik),
    FOREIGN KEY (nik) REFERENCES konsumer(nik),
    INDEX idx_email (email)
) engine InnoDB;

-- create table tenor
CREATE TABLE tenor
(
    id INT NOT NULL AUTO_INCREMENT,
    nik VARCHAR(255) NOT NULL NULL ,
    bulan VARCHAR(255) NOT NULL ,
    start_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    tenor DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (nik) REFERENCES konsumer(nik),
    INDEX idx_nik (nik)
) engine = InnoDB;

-- create table transaction
CREATE TABLE transactions
(
    id INT NOT NULL AUTO_INCREMENT,
    reff_number VARCHAR(255) NOT NULL UNIQUE ,
    nik VARCHAR(255) NOT NULL ,
    date_transaction TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    otr DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    admin_fee DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    jumlah_cicilan INT NOT NULL ,
    jumlah_bunga DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    aset VARCHAR(255) NOT NULL ,
    total_debet DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    PRIMARY KEY (id),
    FOREIGN KEY (nik) REFERENCES konsumer(nik),
    INDEX idx_reff_number (reff_number),
    INDEX idx_nik (nik)
) engine = InnoDB;