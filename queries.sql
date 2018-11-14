-- name: get-all-recipes
SELECT
  Recipe.id,
  Recipe.title,
  Recipe.created_at,
  Step_Ingredient.id,
  Step_Ingredient.quantity,
  Step_Ingredient.unit,
  Step_Ingredient.created_at,
  Ingredient.id,
  Ingredient.name,
  Ingredient.created_at,
  Step.id,
  Step.data,
  Step.step_number,
  Step.created_at,
  Trigger.id,
  Trigger.action_params,
  Trigger.action,
  Trigger.service,
  Trigger.trigger_params,
  Trigger.created_at,
  Trigger_Type.id,
  Trigger_Type.created_at,
  Trigger_Type.key,
  Trigger_Type.sensor_type,
  U.id,
  U.name,
  U.created_at
FROM Recipe
LEFT JOIN Step ON Step.recipe_fk = Recipe.id
LEFT JOIN Step_Ingredient ON Step_Ingredient.step_fk = Step.id
LEFT JOIN Ingredient ON Ingredient.id = Step_Ingredient.ingredient_fk
LEFT JOIN Trigger on Step.id = Trigger.step_fk
LEFT JOIN Trigger_Type on Trigger.tigger_type_fk = Trigger_Type.id
LEFT JOIN Step_Utensil on Step.id = Step_Utensil.step_fk
LEFT JOIN Utensil U on Step_Utensil.utensil_fk = U.id;



-- name: something else
-- select Recipe.*
-- from Recipe
-- inner join Recipe_Ingredient on Recipe_Ingredient.recipe_fk = Recipe.id
-- inner join Ingredient on Ingredient.id = recipe_fk;
SELECT id, title, created_at FROM Recipe;


-- name: find-one-recipe-by-id
SELECT id, title, created_at FROM Recipe WHERE id = ? LIMIT 1;

--name: get-all-ingredients
SELECT id, name, created_at FROM Ingredient;

--name: get-recipe-ingredients-by-recipe-id
SELECT
  DISTINCT Ingredient.id,
  Ingredient.name,
  Ingredient.created_at
FROM Ingredient
LEFT JOIN Step_Ingredient SI on Ingredient.id = SI.ingredient_fk
LEFT JOIN Step S on SI.step_fk = S.id
WHERE S.recipe_fk = ?
ORDER BY Ingredient.name ASC;

--name: get-recipe-ingredients-by-recipe-step-number
SELECT
  Ingredient.id,
  Ingredient.name,
  Ingredient.created_at,
  SI.id,
  SI.unit,
  SI.quantity,
  SI.created_at
FROM Ingredient
LEFT JOIN Step_Ingredient SI on Ingredient.id = SI.ingredient_fk
LEFT JOIN Step S on SI.step_fk = S.id
WHERE S.recipe_fk = ? AND S.step_number = ?
ORDER BY Ingredient.name ASC;