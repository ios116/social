BEGIN;
--
-- Create entities User
--
CREATE TABLE IF NOT EXISTS `users`
(
    `id`           serial PRIMARY KEY,
    `login`        varchar(255) NOT NULL UNIQUE check ( login <> '' ),
    `password`     varchar(128) NOT NULL CHECK (password <> ''),
    `email`        varchar(255) NOT NULL CHECK (email <> ''),
    `first_name`   varchar(255),
    `last_name`    varchar(255),
    `is_staff`     boolean      NOT NULL,
    `is_active`    boolean      NOT NULL,
    `date_created` datetime    NOT NULL,
    `date_modify`  datetime    NOT NULL default NOW()
    );
COMMIT;
