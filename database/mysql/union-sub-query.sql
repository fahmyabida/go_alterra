SELECT id, name FROM users
WHERE id = 1
UNION
SELECT id, name FROM users
WHERE id = 2
UNION
SELECT 100, 'fahmy';

SELECT id, name FROM
    (
        SELECT id, name, status FROM users
    ) result;

