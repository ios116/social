BEGIN;
CHANGE MASTER TO MASTER_HOST = 'master', MASTER_PORT = 3306,  MASTER_USER = 'slave_user', MASTER_PASSWORD = 'qwerty', MASTER_AUTO_POSITION = 1;
START SLAVE;
COMMIT;

