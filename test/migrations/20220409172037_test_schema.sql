-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS TEST;

CREATE TABLE IF NOT EXISTS TEST.Users (
    id      text primary key,
    name    text not null default '',
    surname text not null default ''
);

CREATE TABLE IF NOT EXISTS TEST.Houses (
    id          text primary key,
    name        text not null,
    owner_id    text not null,
    allowance   text[],
    CONSTRAINT houses_Constraint
        FOREIGN KEY (owner_id) REFERENCES TEST.Users (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS TEST.House_parts (
    id          text primary key,
    name        text not null,
    house_id    text not null,
    CONSTRAINT house_part_Constraint
        FOREIGN KEY (house_id) REFERENCES TEST.Houses (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS TEST.Controller_types (
    id                     text primary key,
    name                   text not null,
    photo                  text not null,
    digital_pin_cnt        int not null,
    analog_pin_cnt         int not null
);

CREATE TABLE IF NOT EXISTS TEST.Controllers (
    id                     text primary key,
    controller_type_id     text not null,
    ip                     text not null,
    CONSTRAINT controllers_Constraint
        FOREIGN KEY (controller_type_id) REFERENCES TEST.Controller_types (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);



CREATE TABLE IF NOT EXISTS TEST.Device_types (
    id          text primary key,
    name        text not null,
    photo       text not null
);


CREATE TABLE IF NOT EXISTS TEST.Devices (
    id                 text primary key,
    device_type_id     text not null,
    house_part_id       text not null,
    CONSTRAINT devices_Constraint
        FOREIGN KEY (device_type_id) REFERENCES TEST.Device_types (id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT house_part_Constraint
        FOREIGN KEY (house_part_id) REFERENCES TEST.House_parts (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS TEST.Pins (
    id              text primary key,
    controller_id   text not null,
    device_id       text not null,
    value           int not null,
    type            int not null,
    CONSTRAINT pins_constraint
        FOREIGN KEY (controller_id) REFERENCES TEST.Controllers (id)
        ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO TEST.Users(id, name, surname) VALUES('9120f91d-17b9-405a-ae74-797c4c9e0119', 'name', 'surname');
INSERT INTO TEST.Device_types (id, name, photo) VALUES ('4fba07cb-7c5e-4a18-a62f-2e9044a50c1b', 'Выключатель', 'https://avselectro.ru/uploads/gallery/44/max/1a2a3ef554d7cd9f32fc6895a6f13d86.jpg');
INSERT INTO TEST.Houses (id, name, owner_id) VALUES ('2e345e6d-b3b9-42e8-a3b4-cf147b037d3c', 'house1', '9120f91d-17b9-405a-ae74-797c4c9e0119');
INSERT INTO TEST.House_parts (id, name, house_id) VALUES ('8120f91d-17b9-405a-ae74-797c4c9e0117', 'house1', '2e345e6d-b3b9-42e8-a3b4-cf147b037d3c');
INSERT INTO TEST.Controller_types (id, name, photo, digital_pin_cnt, analog_pin_cnt) VALUES ('39248a56-18d7-46c1-bbd9-a8139b6bf1fa', 'Orange Pi One', 'https://static.chipdip.ru/lib/736/DOC002736925.jpg', 11, 5);
INSERT INTO TEST.Controllers (id, controller_type_id, ip) VALUES ('4523e684-ad41-4fe3-8835-8d200d164a2f', '39248a56-18d7-46c1-bbd9-a8139b6bf1fa', '10.21.0.21');
INSERT INTO TEST.Devices (id, device_type_id, house_part_id) VALUES ('70d3d531-4041-4d74-8306-bf8e7319b74b', '4fba07cb-7c5e-4a18-a62f-2e9044a50c1b', '8120f91d-17b9-405a-ae74-797c4c9e0117');
INSERT INTO TEST.Pins (id, controller_id, device_id, value, type) VALUES ('730b8d73-8426-4a82-85d9-2c79ce4f3d9c', '4523e684-ad41-4fe3-8835-8d200d164a2f', '70d3d531-4041-4d74-8306-bf8e7319b74b', 0, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS TEST.Pins;
DROP TABLE IF EXISTS TEST.Controllers;
DROP TABLE IF EXISTS TEST.Controller_types;
DROP TABLE IF EXISTS TEST.Devices;
DROP TABLE IF EXISTS TEST.Device_types;
DROP TABLE IF EXISTS TEST.House_parts;
DROP TABLE IF EXISTS TEST.Houses;
DROP TABLE IF EXISTS TEST.Users;
-- +goose StatementEnd
