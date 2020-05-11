USE `chillit_store`;

INSERT INTO city (`id`, `title`)
VALUES (1, 'Казань'),
       (2, 'Йошкар-Ола');

INSERT INTO place (`title`, `address`, `description`, `city_id`)
VALUES ('Test title 1', 'Test address 1', 'Test description Test description and more 1', 1),
       ('Test title 2', 'Test address 2', 'Test description Test description and more 2', 2),
       ('Test title 3', 'Test address 3', 'Test description Test description and more 3', 1),
       ('Test title 4', 'Test address 4', 'Test description Test description and more 4', 1)