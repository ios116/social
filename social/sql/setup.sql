BEGIN;
--
-- Create entities User
--
CREATE TABLE IF NOT EXISTS `users`
(
    `id`           serial PRIMARY KEY,
    `login`        varchar(255) NOT NULL UNIQUE check ( login <> '' ),
    `password`     varchar(255) NOT NULL CHECK (password <> ''),
    `email`        varchar(255) NOT NULL CHECK (email <> ''),
    `age`          int NOT NULL default 30,
    `city`         varchar(255),
    `gender`       varchar(255),
    `interests`    text,
    `first_name`   varchar(255),
    `last_name`    varchar(255),
    `date_created` datetime     NOT NULL default NOW(),
    `date_modify`  datetime     NOT NULL default NOW(),
    check (gender in ('male', 'female', 'Unknown'))
);
COMMIT;
