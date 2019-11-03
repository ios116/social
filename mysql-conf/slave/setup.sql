-- BEGIN;
-- --
-- -- Create entities User
-- --
-- CREATE TABLE IF NOT EXISTS `users`
-- (
--     `id`           serial PRIMARY KEY,
--     `login`        varchar(255) NOT NULL UNIQUE check ( login <> '' ),
--     `password`     varchar(255) NOT NULL CHECK (password <> ''),
--     `email`        varchar(255) NOT NULL CHECK (email <> ''),
--     `age`          int NOT NULL default 30,
--     `city`         varchar(255),
--     `gender`       varchar(255),
--     `interests`    text,
--     `first_name`   varchar(255),
--     `last_name`    varchar(255),
--     `date_created` datetime     NOT NULL default NOW(),
--     `date_modify`  datetime     NOT NULL default NOW(),
--     check (gender in ('male', 'female', 'Unknown')),
--     KEY `f` (`first_name`),
--     KEY `l` (`last_name`)
-- );
-- COMMIT;
CHANGE MASTER TO MASTER_HOST = 'master', MASTER_PORT = 3306,  MASTER_USER = 'slave_user', MASTER_PASSWORD = 'qwerty', MASTER_AUTO_POSITION = 1;
START SLAVE;