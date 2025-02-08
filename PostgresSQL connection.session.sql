-- Select all tasks from the tasks table
SELECT *
FROM tasks;
-- Select tasks that are completed
SELECT *
FROM tasks
WHERE is_completed = TRUE;
-- Select tasks of given ID
SELECT *
FROM tasks
WHERE id = 31;