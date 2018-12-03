INSERT INTO Ingredient (name, created_at) VALUES ('Water', '2018-11-14 20:10:36');
INSERT INTO Ingredient (name, created_at) VALUES ('Pasta', '2018-11-14 20:10:36');
INSERT INTO Ingredient (name, created_at) VALUES ('Butter', '2018-11-19 04:18:13');
INSERT INTO Ingredient (name, created_at) VALUES ('Heavy Whipping Cream', '2018-11-26 18:46:37');
INSERT INTO Ingredient (name, created_at) VALUES ('Salt', '2018-11-26 18:46:37');
INSERT INTO Ingredient (name, created_at) VALUES ('Nutmeg', '2018-11-26 18:46:37');
INSERT INTO Ingredient (name, created_at) VALUES ('Freshly Ground Black Pepper', '2018-11-26 18:46:37');
INSERT INTO Ingredient (name, created_at) VALUES ('Block of Parmesan', '2018-11-26 18:46:37');
INSERT INTO Ingredient (name, created_at) VALUES ('Chicken Marinade', '2018-11-26 18:46:37');
INSERT INTO Ingredient (name, created_at) VALUES ('Chicken Breasts', '2018-11-26 18:46:37');
INSERT INTO Recipe (title, created_at) VALUES ('Grilled Chicken Alfredo', '2018-11-14 20:10:44');
INSERT INTO Recipe (title, created_at) VALUES ('Maruchan Ramen Noodle Soup', '2018-11-19 04:17:46');
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'To make pasta, first measure out 3 quarts water into a large pot', '2018-11-14 20:11:07', 1);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Bring water to a boil over high heat', '2018-11-14 20:11:14', 2);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Add 12 ounces of pasta to the boiling water', '2018-11-16 13:56:50', 3);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Boil the pasta for 10 minutes', '2018-11-19 04:21:11', 4);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Turn off heat and strain the water out of the pot, leaving the pasta in', '2018-11-19 04:21:12', 5);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Now remove pasta from the pot and set it aside. Try putting it in a bowl and keeping it in a microwave to keep the pasta warm. Do not turn the microwave on.', '2018-11-19 04:21:12', 6);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'To make the Alfredo sauce, start by grating one and a half cups of parmesan cheese onto a spare plate.', '2018-11-19 04:21:12', 7);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'In the same pot as the pasta, melt 6 tablespoons of butter over medium heat.', '2018-11-19 04:21:12', 8);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Add heavy whipping cream, one half teaspoon salt, one fourth teaspoon nutmeg, 1 teaspoon black ground pepper to the melted butter.', '2018-11-19 04:21:12', 10);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Stir the sauce constantly until it thickens. One way to tell when the sauce is thickened is when the bubbles in the sauce travel noticeable slower and the sauce is a litter harder to stir. Be careful not to boil the cream as it will spoil.', '2018-11-19 04:21:12', 11);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Reduce temp to low, slowly stir in parmesan cheese a little at a time with a whisk.', '2018-11-26 19:01:07', 12);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Turn off heat and set the Alfredo sauce aside', '2018-11-26 19:01:07', 13);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Add chicken marinade to skillet over medium heat', '2018-11-26 19:01:07', 15);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Add chicken breasts to the skillet', '2018-11-26 19:01:07', 16);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Cook until internal temperature of the chicken is 165 degrees turning often, (about 10 minutes). Be sure to keep the thermometer probe in the chicken.', '2018-11-26 19:01:07', 17);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Turn off heat, move chicken breasts to a cutting board.', '2018-11-26 19:01:07', 18);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Slice chicken into one fourth inch strips', '2018-11-26 19:01:07', 19);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Place 1 serving of pasta onto a plate', '2018-11-26 19:01:07', 20);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Drizzle pasta on plate with Alfredo sauce', '2018-11-26 19:01:07', 21);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Add several strips of chicken', '2018-11-26 19:01:07', 22);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Enjoy!', '2018-11-26 19:01:07', 23);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'Measure 2 cups heavy whipping cream', '2018-11-26 19:01:07', 9);
INSERT INTO Step (recipe_fk, data, created_at, step_number) VALUES (1, 'To grill chicken, start by measuring one half cup marinade', '2018-11-26 19:01:07', 14);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (3, 'Quart', '2018-11-14 20:11:46', 1, 1);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (12, 'Ounce', '2018-11-14 20:11:46', 3, 2);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (1, 'Block', '2018-11-19 04:22:56', 7, 8);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (6, 'Tablespoon', '2018-11-19 04:22:56', 8, 3);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (2, 'Cup', '2018-11-19 05:53:56', 9, 4);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (0.5, 'Teaspoon', '2018-11-19 05:53:56', 9, 5);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (0.25, 'Teaspoon', '2018-11-19 05:53:56', 9, 6);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (1, 'Teaspoon', '2018-11-19 05:53:56', 9, 7);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (0.5, 'Cups', '2018-11-19 05:53:56', 13, 9);
INSERT INTO Step_Ingredient (quantity, unit, created_at, step_fk, ingredient_fk) VALUES (4, 'Breasts', '2018-11-19 05:53:56', 14, 10);
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 1, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 2, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 3, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 4, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 5, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (7, 5, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (9, 6, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (8, 7, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (10, 7, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 8, '2018-11-19 05:51:06');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 9, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (2, 10, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (2, 11, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 11, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (3, 12, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (4, 13, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (4, 14, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (4, 15, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (5, 16, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (5, 17, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (6, 17, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (8, 18, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (8, 19, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (11, 19, '2018-11-26 18:36:51');
INSERT INTO Step_Utensil (utensil_fk, step_fk, created_at) VALUES (8, 19, '2018-11-26 18:36:51');
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (9, 'walk-through', '2018-11-14 20:15:08', '1', 'incrementStepByN', '3785', 1);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (4, 'walk-through', '2018-11-19 04:26:09', '1', 'incrementStepByN', '97', 2);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-19 04:27:41', '1', 'incrementStepByN', '1', 3);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (2, 'walk-through', '2018-11-19 04:32:11', '1', 'incrementStepByN', '600', 4);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-19 04:34:35', '1', 'incrementStepByN', '1', 5);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-19 04:35:26', '1', 'incrementStepByN', '1', 6);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (9, 'walk-through', '2018-11-19 04:37:37', '1', 'incrementStepByN', '150', 7);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 8);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (9, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '462', 22);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 9);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 10);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 11);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 12);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (9, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '115', 23);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 13);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 14);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (2, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '165', 15);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 16);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 17);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 18);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 19);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 20);
INSERT INTO Trigger (tigger_type_fk, service, created_at, action_params, action, trigger_params, trigger_group_fk) VALUES (1, 'walk-through', '2018-11-26 18:43:24', '1', 'incrementStepByN', '1', 21);
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (1, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (2, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (3, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (4, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (5, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (6, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (7, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (8, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (9, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (10, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (11, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (12, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (13, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (14, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (15, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (16, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (17, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (18, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (19, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (20, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (21, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (22, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Group (step_fk, action_key, service, action_params) VALUES (23, 'changeStep', 'walk-through', '{"increment_steps": 1, "send_to_nlp": true}');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-14 20:15:23', 'pass', 'pass');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-16 14:43:09', 'timer', 'timer');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-19 04:30:44', 'temp_>', 'thermometer');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-19 04:33:36', 'time_>=', 'thermometer');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'temp_<', 'thermometer');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'temp_<=', 'thermometer');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'temp_==', 'thermometer');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'weight_>', 'scale');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'weight_>=', 'scale');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'weight_<', 'scale');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'weight_<=', 'scale');
INSERT INTO Trigger_Type (created_at, key, sensor_type) VALUES ('2018-11-26 18:30:35', 'weight_==', 'scale');
INSERT INTO Utensil (name, created_at) VALUES ('Whisk', '2018-11-14 20:12:12');
INSERT INTO Utensil (name, created_at) VALUES ('Wooden Spoon', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Large Pot', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Skillet', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Cutting Board', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Knife', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Strainer', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Plate', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Bowl', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Cheese Grater', '2018-11-26 18:32:48');
INSERT INTO Utensil (name, created_at) VALUES ('Spoon', '2018-11-26 18:32:48');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Ingredient', 'Ingredient', 2, 'CREATE TABLE Ingredient
-- (
-- id integer
-- primary key
-- autoincrement,
-- name varchar,
-- created_at datetime default CURRENT_TIMESTAMP
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'sqlite_sequence', 'sqlite_sequence', 3, 'CREATE TABLE sqlite_sequence(name,seq)');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('index', 'Ingredient_name_uindex', 'Ingredient', 4, 'CREATE UNIQUE INDEX Ingredient_name_uindex
-- on Ingredient (name)
-- ');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Recipe', 'Recipe', 5, 'CREATE TABLE Recipe
-- (
-- id integer
-- primary key
-- autoincrement,
-- title varchar,
-- created_at datetime default CURRENT_TIMESTAMP
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Step', 'Step', 6, 'CREATE TABLE Step
-- (
-- id integer
-- primary key
-- autoincrement,
-- recipe_fk integer
-- constraint recipe_fk
-- references Recipe,
-- data varchar,
-- created_at datetime default CURRENT_TIMESTAMP,
-- step_number integer
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Step_Ingredient', 'Step_Ingredient', 7, 'CREATE TABLE Step_Ingredient
-- (
-- id integer not null
-- primary key
-- autoincrement,
-- quantity float,
-- unit varchar,
-- created_at datetime default CURRENT_TIMESTAMP,
-- step_fk int
-- constraint step__fk
-- references Step,
-- ingredient_fk int
-- constraint ingredient__fk
-- references Ingredient
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Trigger_Type', 'Trigger_Type', 8, 'CREATE TABLE Trigger_Type
-- (
-- id integer
-- primary key
-- autoincrement,
-- created_at datetime default CURRENT_TIMESTAMP,
-- key varchar,
-- sensor_type varchar
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Utensil', 'Utensil', 10, 'CREATE TABLE Utensil
-- (
-- id integer
-- primary key
-- autoincrement,
-- name varchar,
-- created_at datetime default CURRENT_TIMESTAMP
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Step_Utensil', 'Step_Utensil', 11, 'CREATE TABLE Step_Utensil
-- (
-- id integer not null
-- primary key
-- autoincrement,
-- utensil_fk int
-- constraint utensil__fk
-- references Utensil,
-- step_fk int
-- constraint step__fk
-- references Step,
-- created_at datetime default CURRENT_TIMESTAMP
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Trigger', 'Trigger', 12, 'CREATE TABLE "Trigger"
-- (
--     id integer PRIMARY KEY AUTOINCREMENT,
--     tigger_type_fk integer,
--     service varchar,
--     created_at datetime DEFAULT CURRENT_TIMESTAMP,
--     action_params varchar,
--     action varchar,
--     trigger_params varchar,
--     trigger_group_fk integer,
--     CONSTRAINT trigger_type__fk FOREIGN KEY (tigger_type_fk) REFERENCES Trigger_Type (id),
--     CONSTRAINT trigger_group__fk FOREIGN KEY (trigger_group_fk) REFERENCES Trigger_Group (id)
-- )');
-- INSERT INTO sqlite_master (type, name, tbl_name, rootpage, sql) VALUES ('table', 'Trigger_Group', 'Trigger_Group', 9, 'CREATE TABLE "Trigger_Group"
-- (
--     id integer PRIMARY KEY AUTOINCREMENT,
--     step_fk integer,
--     action_key varchar,
--     service varchar,
--     action_params varchar,
--     CONSTRAINT step__fk FOREIGN KEY (step_fk) REFERENCES Step (id)
-- )');
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Ingredient', 10);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Recipe', 2);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Step', 23);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Step_Ingredient', 12);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Utensil', 11);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Step_Utensil', 25);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Trigger_Type', 12);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Trigger', 23);
-- INSERT INTO sqlite_sequence (name, seq) VALUES ('Trigger_Group', 23);