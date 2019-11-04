## Social network 

### TASK3

**Добавить m/s репликацию. Сделать балансирование запросов на чтение. Провести нагрузочное тестирование**

- конфиг мастера
```
# master
server-id = 1 # идентификатор мастер сервера
binlog_do_db = soc_db # база для репликации
gtid_mode=ON # включает GTID
binlog_format = ROW # формат ведения журнала row base
log_bin=mysql-bin # Ведение бинарного лога для мастера (с него читает слейв).
enforce-gtid-consistency=ON
```
- конфиг слейва
```
#salve
binlog_do_db = soc_db # база для репликации
server_id = 2  # идентификатор slave сервера
binlog_format = ROW # формат ведения журнала row base
gtid_mode = on # GTID mod
enforce_gtid_consistency
read-only=on # только в режиме чтения
```

На мастере создан пользователь для реплики
```mysql
GRANT REPLICATION SLAVE ON *.* TO 'slave_user'@'%' IDENTIFIED BY 'qwerty';
FLUSH PRIVILEGES;
```
На слейве указываем мастер
```mysql
CHANGE MASTER TO MASTER_HOST = 'master', MASTER_PORT = 3306,  MASTER_USER = 'slave_user', MASTER_PASSWORD = 'qwerty', MASTER_AUTO_POSITION = 1;
START SLAVE;
```

Нагрузка на чтение
```wrk -c 200 -t 16 -d 30s "http://212.109.223.229/search?query=Tomas"```

1) В случае где запросы на чтение идут на master:

![master](social/assets/img/rep_master.png)

2) В случае где запросы на чтение идут на slave

![slave](social/assets/img/rep_slave.png) 

### Task2

**Без индекса**
```mysql
explain SELECT id, first_name, last_name, city FROM users WHERE id>22481 AND (first_name LIKE 'tom%' or last_name LIKE 'tom%') ORDER BY id ASC LIMIT 201;
```
```shell script
+----+-------------+-------+------------+-------+---------------+---------+---------+------+--------+----------+-------------+
| id | select_type | table | partitions | type  | possible_keys | key     | key_len | ref  | rows   | filtered | Extra       |
+----+-------------+-------+------------+-------+---------------+---------+---------+------+--------+----------+-------------+
|  1 | SIMPLE      | users | NULL       | range | PRIMARY,id    | PRIMARY | 8       | NULL | 472495 |    20.99 | Using where |
+----+-------------+-------+------------+-------+---------------+---------+---------+------+--------+----------+-------------+
```
```json
 {
  "query_block": {
    "select_id": 1,
    "cost_info": {
      "query_cost": "189809.78"
    },
    "ordering_operation": {
      "using_filesort": false,
      "table": {
        "table_name": "users",
        "access_type": "range",
        "possible_keys": [
          "PRIMARY",
          "id"
        ],
        "key": "PRIMARY",
        "used_key_parts": [
          "id"
        ],
        "key_length": "8",
        "rows_examined_per_scan": 472495,
        "rows_produced_per_join": 99156,
        "filtered": "20.99",
        "cost_info": {
          "read_cost": "169978.52",
          "eval_cost": "19831.26",
          "prefix_cost": "189809.78",
          "data_read_per_join": "511M"
        },
        "used_columns": [
          "id",
          "city",
          "first_name",
          "last_name"
        ],
        "attached_condition": "((`soc_db`.`users`.`id` > 22481) and ((`soc_db`.`users`.`first_name` like 'tom%') or (`soc_db`.`users`.`last_name` like 'tom%')))"
      }
    }
  }
}
```

**С индексом**
```mysql
create index f on users(first_name);
create index l on users(last_name);
explain SELECT id, first_name, last_name, city FROM users WHERE id>22481 AND (first_name LIKE 'tom%' or last_name LIKE 'tom%') ORDER BY id ASC LIMIT 201;
```
```shell script
+----+-------------+-------+------------+-------------+----------------+------+---------+------+------+----------+----------------------------------------------------+
| id | select_type | table | partitions | type        | possible_keys  | key  | key_len | ref  | rows | filtered | Extra                                              |
+----+-------------+-------+------------+-------------+----------------+------+---------+------+------+----------+----------------------------------------------------+
|  1 | SIMPLE      | users | NULL       | index_merge | PRIMARY,id,l,f | f,l  | 768,768 | NULL | 1346 |    50.00 | Using sort_union(f,l); Using where; Using filesort |
+----+-------------+-------+------------+-------------+----------------+------+---------+------+------+----------+----------------------------------------------------+
```

```json
{
  "query_block": {
    "select_id": 1,
    "cost_info": {
      "query_cost": "4317.85"
    },
    "ordering_operation": {
      "using_filesort": true,
      "table": {
        "table_name": "users",
        "access_type": "index_merge",
        "possible_keys": [
          "PRIMARY",
          "id",
          "f",
          "l"
        ],
        "key": "sort_union(f,l)",
        "key_length": "768,768",
        "rows_examined_per_scan": 1346,
        "rows_produced_per_join": 672,
        "filtered": "50.00",
        "cost_info": {
          "read_cost": "4183.25",
          "eval_cost": "134.60",
          "prefix_cost": "4317.85",
          "data_read_per_join": "3M"
        },
        "used_columns": [
          "id",
          "city",
          "first_name",
          "last_name"
        ],
        "attached_condition": "((`soc_db`.`users`.`id` > 22481) and ((`soc_db`.`users`.`first_name` like 'tom%') or (`soc_db`.`users`.`last_name` like 'tom%')))"
      }
    }
  }
}
```
[wrk tests без индекса](https://github.com/ios116/social/blob/master/social/assets/index_no) | [wrk tests с индексом](https://github.com/ios116/social/blob/master/social/assets/index_yes)

![latency](social/assets/img/latency.png)


![throughput](social/assets/img/throughput.png)

- Индекс выбран не составной, т.к. используется OR если бы AND то лучше работал бы составной.
- Выдача разбивается по страницам по последнему выданному Id + Limit,а не с помощью offset, т.к. чем больше offset тем больше планировщику приходится просчитывать отступ и тем медленне запрос. 
- Чем более селективнее запрос тем лучше работает индекс, т.е индекс при поисковом запросе "tom" будет работать лучше чем при "t".
- Очевидно, что производительность с индексом существенно выше. 
- При 1000 одновременных соединений увеличивается колличество долгих запросов без индекса 17% от 8624, c индексом 20% от 32735. 

### Task1
Tech stack:
- golang
- mysql 5.7
- session JWT
- css bootstrap

[sql](https://github.com/ios116/social/blob/master/social/sql/setup.sql)
