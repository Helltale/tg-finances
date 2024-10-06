CREATE SCHEMA profile;

CREATE TABLE profile.userdata (
    tg_id BIGINT PRIMARY KEY,
    profile_name VARCHAR(100) NOT NULL
);

CREATE SCHEMA income;

CREATE TABLE income.income_type (
    name_type VARCHAR(100) PRIMARY KEY
);

CREATE TABLE income.income_standart (
    id_income_standart SERIAL PRIMARY KEY,
    name_income VARCHAR(100) NOT NULL REFERENCES income.income_type(name_type),
    tg_id BIGINT NOT NULL,
    price_standart DECIMAL(10, 2) NOT NULL,
    date_income_standart_date DATE NOT NULL,
    date_income_standart_month DATE,
    date_actual_from TIMESTAMP NOT NULL,
    date_actual_to TIMESTAMP NOT NULL
);

CREATE TABLE income.income_real (
    id_income_real SERIAL PRIMARY KEY,
    name_income VARCHAR(100) NOT NULL REFERENCES income.income_type(name_type),
    id_income_standart BIGINT NOT NULL,
    tg_id BIGINT NOT NULL,
    price_real DECIMAL(10, 2) NOT NULL,
    date_income_real_date DATE NOT NULL,
    date_income_real_month DATE,
    date_actual_from TIMESTAMP NOT NULL,
    date_actual_to TIMESTAMP NOT NULL
);

CREATE SCHEMA bank;

CREATE TABLE bank.remnant (
    id_remnant SERIAL PRIMARY KEY,
    tg_id BIGINT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    date_actual_from TIMESTAMP NOT NULL,
    date_actual_to TIMESTAMP NOT NULL
);

CREATE SCHEMA expences;

CREATE TABLE expences.expences_categories (
    name_category VARCHAR(150) PRIMARY KEY
);

CREATE TABLE expences.expences (
    id_expences SERIAL PRIMARY KEY,
    expences_category VARCHAR(150) NOT NULL REFERENCES expences.expences_categories(name_category),
    expences_name VARCHAR(500) NOT NULL,
    expences_description TEXT,
    expences_money DECIMAL(10, 2) NOT NULL,
    expences_date TIMESTAMP NOT NULL,
    flag_constancy TINYINT NOT NULL,
    date_actual_from TIMESTAMP NOT NULL,
    date_actual_to TIMESTAMP NOT NULL
);

CREATE SCHEMA cashback;

CREATE TABLE cashback.cashback_info (
    id_cashback SERIAL PRIMARY KEY,
    bank_name VARCHAR(100) NOT NULL,
    category VARCHAR(150) NOT NULL,
    cashback_percentage DECIMAL(5, 2) NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL
);
