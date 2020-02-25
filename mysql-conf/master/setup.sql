BEGIN;
--
-- Create entities User
--
CREATE TABLE IF NOT EXISTS users
(
    id           bigint PRIMARY KEY AUTO_INCREMENT,
    login        varchar(255) NOT NULL UNIQUE check ( login <> '' ),
    password     varchar(255) NOT NULL CHECK (password <> ''),
    email        varchar(255) NOT NULL CHECK (email <> ''),
    age          int          NOT NULL default 30,
    city         varchar(255),
    gender       varchar(255),
    interests    text,
    first_name   varchar(255),
    last_name    varchar(255),
    date_created datetime     NOT NULL default NOW(),
    date_modify  datetime     NOT NULL default NOW(),
    check (gender in ('male', 'female', 'Unknown')),
    KEY f (first_name),
    KEY l (last_name)
);
CREATE TABLE IF NOT EXISTS `subscriber`
(
    id           bigint PRIMARY KEY AUTO_INCREMENT,
    user_id      bigint  NOT NULL,
    subscribe_id bigint  NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ,
    FOREIGN KEY (subscribe_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE INDEX subscriber_user_id on subscriber (user_id);
CREATE INDEX subscriber_subscriber_id on subscriber (subscribe_id);

COMMIT;

GRANT REPLICATION SLAVE ON *.* TO 'slave_user'@'%' IDENTIFIED BY 'qwerty';
FLUSH PRIVILEGES;
INSTALL PLUGIN rpl_semi_sync_master SONAME 'semisync_master.so';
SET GLOBAL rpl_semi_sync_master_enabled = 1;
SET GLOBAL rpl_semi_sync_master_timeout = 1000;
