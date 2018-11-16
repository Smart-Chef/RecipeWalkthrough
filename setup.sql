-- name: create-smart-chef
create table Ingredient
(
	id integer
		primary key
		 autoincrement,
	name varchar,
	created_at datetime default CURRENT_TIMESTAMP
)
;

create unique index Ingredient_name_uindex
	on Ingredient (name)
;

create table Recipe
(
	id integer
		primary key
		 autoincrement,
	title varchar,
	created_at datetime default CURRENT_TIMESTAMP
)
;

create table Step
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
)
;

create table Step_Ingredient
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
)
;

create table Trigger_Group
(
	id integer
		primary key
		 autoincrement,
	step_fk integer
		constraint step__fk
			references Step,
	action_key varchar,
	service varchar,
	action_params varchar
)
;

create table Trigger_Type
(
	id integer
		primary key
		 autoincrement,
	created_at datetime default CURRENT_TIMESTAMP,
	key varchar,
	sensor_type varchar
)
;

create table Trigger
(
	id integer
		primary key
		 autoincrement,
	tigger_type_fk integer
		constraint trigger_type__fk
			references Trigger_Type,
	service varchar,
	created_at datetime default CURRENT_TIMESTAMP,
	action_params varchar,
	action varchar,
	trigger_params varchar,
	trigger_group_fk integer
		constraint trigger_group__fk
			references Trigger_Group
)
;

create table Utensil
(
	id integer
		primary key
		 autoincrement,
	name varchar,
	created_at datetime default CURRENT_TIMESTAMP
)
;

create table Step_Utensil
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
)
;
