-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS Users (
    id      text primary key,
    name    text not null default '',
    surname text not null default ''
);

CREATE TABLE IF NOT EXISTS Houses (
    id          text primary key,
    name        text not null,
    owner_id    text not null,
    allowance   text[],
    CONSTRAINT houses_Constraint
        FOREIGN KEY (owner_id) REFERENCES Users (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS House_parts (
    id          text primary key,
    name        text not null,
    house_id    text not null,
    CONSTRAINT house_part_Constraint
        FOREIGN KEY (house_id) REFERENCES Houses (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Controller_types (
    id                     text primary key,
    name                   text not null,
    photo                  text not null,
    digital_pin_cnt        int not null,
    analog_pin_cnt         int not null
);

CREATE TABLE IF NOT EXISTS Controllers (
    id                     text primary key,
    controller_type_id     text not null,
    ip                     text not null,
    CONSTRAINT controllers_Constraint
        FOREIGN KEY (controller_type_id) REFERENCES Controller_types (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS Device_types (
    id          text primary key,
    name        text not null
);


CREATE TABLE IF NOT EXISTS Devices (
    id                 text primary key,
    device_type_id     text not null,
    house_part_id       text not null,
    CONSTRAINT devices_Constraint
        FOREIGN KEY (device_type_id) REFERENCES Device_types (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT house_part_Constraint
        FOREIGN KEY (house_part_id) REFERENCES House_parts (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Pins (
    id              text primary key,
    controller_id   text not null,
    device_id       text not null,
    value           int not null,
    type            int not null,
    CONSTRAINT pins_constraint
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
