-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS Users (
    id      bigserial primary key,
    name    text not null default '',
    surname text not null default ''
);

CREATE TABLE IF NOT EXISTS Houses (
    id          bigserial primary key,
    name        text not null default '',
    owner       int not null,
    allowance   int[]
);

CREATE TABLE IF NOT EXISTS House_parts (
    id          bigserial primary key,
    name        text not null default '',
    house_id    int not null default 0,
    CONSTRAINT house_part_Constraint
        FOREIGN KEY (house_id) REFERENCES Houses (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Controller_types (
    id                     bigserial primary key,
    name                   text not null default '',
    photo                  text not null default '',
    digital_pin_cnt        int not null default 0,
    analog_pin_cnt         int not null default 0
);

CREATE TABLE IF NOT EXISTS Controllers (
    id                     bigserial primary key,
    controller_type_id     int not null default 0,
    ip                     text not null default '',
    CONSTRAINT controllers_Constraint
        FOREIGN KEY (controller_type_id) REFERENCES Controller_types (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS Device_types (
    id          bigserial primary key,
    name        text not null default ''
);


CREATE TABLE IF NOT EXISTS Devices (
    id                 bigserial primary key,
    device_type_id     int not null default 0,
    CONSTRAINT devices_Constraint
        FOREIGN KEY (device_type_id) REFERENCES Device_types (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Pins (
    id              bigserial primary key,
    controller_id   int not null,
    device_id       int not null,
    value           int not null,
    CONSTRAINT analog_pin_constraint
        FOREIGN KEY (controller_id) REFERENCES Controllers (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Pins;
DROP TABLE IF EXISTS Controllers;
DROP TABLE IF EXISTS Controller_types;
DROP TABLE IF EXISTS Devices;
DROP TABLE IF EXISTS Device_types;
DROP TABLE IF EXISTS House_parts;
DROP TABLE IF EXISTS Houses;
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd
