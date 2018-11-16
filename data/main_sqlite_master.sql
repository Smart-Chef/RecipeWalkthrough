INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Ingredient', 'Ingredient', 2, 'CREATE TABLE Ingredient
(
id integer
primary key
autoincrement,
name varchar,
created_at datetime default CURRENT_TIMESTAMP
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'sqlite_sequence', 'sqlite_sequence', 3, 'CREATE TABLE sqlite_sequence(name,seq)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('index', 'Ingredient_name_uindex', 'Ingredient', 4, 'CREATE UNIQUE INDEX Ingredient_name_uindex
on Ingredient (name)
');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Recipe', 'Recipe', 5, 'CREATE TABLE Recipe
(
id integer
primary key
autoincrement,
title varchar,
created_at datetime default CURRENT_TIMESTAMP
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Step', 'Step', 6, 'CREATE TABLE Step
(
id integer
primary key
autoincrement,
recipe_fk integer
constraint recipe_fk
references Recipe,
data varchar,
created_at datetime default CURRENT_TIMESTAMP,
step_number integer
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Step_Ingredient', 'Step_Ingredient', 7, 'CREATE TABLE Step_Ingredient
(
id integer not null
primary key
autoincrement,
quantity float,
unit varchar,
created_at datetime default CURRENT_TIMESTAMP,
step_fk int
constraint step__fk
references Step,
ingredient_fk int
constraint ingredient__fk
references Ingredient
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Trigger_Type', 'Trigger_Type', 8, 'CREATE TABLE Trigger_Type
(
id integer
primary key
autoincrement,
created_at datetime default CURRENT_TIMESTAMP,
key varchar,
sensor_type varchar
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Utensil', 'Utensil', 10, 'CREATE TABLE Utensil
(
id integer
primary key
autoincrement,
name varchar,
created_at datetime default CURRENT_TIMESTAMP
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Step_Utensil', 'Step_Utensil', 11, 'CREATE TABLE Step_Utensil
(
id integer not null
primary key
autoincrement,
utensil_fk int
constraint utensil__fk
references Utensil,
step_fk int
constraint step__fk
references Step,
created_at datetime default CURRENT_TIMESTAMP
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Trigger', 'Trigger', 12, 'CREATE TABLE "Trigger"
(
    id integer PRIMARY KEY AUTOINCREMENT,
    tigger_type_fk integer,
    service varchar,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    action_params varchar,
    action varchar,
    trigger_params varchar,
    trigger_group_fk integer,
    CONSTRAINT trigger_type__fk FOREIGN KEY (tigger_type_fk) REFERENCES Trigger_Type (id),
    CONSTRAINT trigger_group__fk FOREIGN KEY (trigger_group_fk) REFERENCES Trigger_Group (id)
)');
INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Trigger_Group', 'Trigger_Group', 9, 'CREATE TABLE "Trigger_Group"
(
    id integer PRIMARY KEY AUTOINCREMENT,
    step_fk integer,
    action_key varchar,
    service varchar,
    action_params varchar,
    CONSTRAINT step__fk FOREIGN KEY (step_fk) REFERENCES Step (id)
)');