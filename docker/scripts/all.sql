CREATE SCHEMA profile;

CREATE TABLE profile.userdata (
    tg_id BIGINT PRIMARY KEY,
    profile_name VARCHAR(100) NOT NULL
);

create schema income;

create table income.income_type (
    name_type varchar(100) primary key
);

--должно быть
create table income.income_standart (
    id_income_standart serial primary key,
    name_income varchar(100) not NULL,
    tg_id bigint not NULL,
    price_standart float not NULL,
    date_income_standart_date timestamp not NULL, --здесь должно быть число
    date_income_standart_month timestamp --здесь должно быть месяц
    date_actual_from timestamp not NULL,
    date_actual_to timestamp not NULL, -- '9999-12-31' если актуальная запись
);

--как всегда...
create table income.income_real (
    id_income_real serial primary key,
    name_income varchar(100) not NULL,
    id_income_standart bigint not NULL,
    tg_id bigint not NULL,
    price_real float not NULL,
    date_income_real_date timestamp not NULL, --здесь должно быть число
    date_income_real_month timestamp --здесь должно быть месяц
    date_actual_from timestamp not NULL,
    date_actual_to timestamp not NULL, -- '9999-12-31' если актуальная запись
);

create schema bank;

create table bank.remnant ( -- остаток
    id_remnant serial primary key
    tg_id bigint not null,
    price float not NULL,
    date_actual_from timestamp not NULL,
    date_actual_to timestamp not NULL, -- '9999-12-31' если актуальная запись
);

create schema expences;

create table expences.expences_categories ( --категории трат
    name_category varchar(150)
);

--todo переделать 
create table expences.expences ( --траты
    id_expences serial primary key,
    expences_category varchar(150) not null,
    expences_name varchar(500) not null,
    expences_description text,
    expences_money float not null,
    expences_date timestamp not null,
    flag_constancy tinyint not null,
    date_actual_from timestamp not NULL,
    date_actual_to timestamp not NULL
)