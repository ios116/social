## Social network 

### Task2

**Без индекса**

```mysql
explain format=json select id, first_name, last_name from users WHERE first_name LIKE 'to%' OR last_name LIKE 'to%' ORDER BY id limit 100;
```

```json
 {
   "query_block": {
     "select_id": 1,
     "cost_info": {
       "query_cost": "197049.60"
     },
     "ordering_operation": {
       "using_filesort": false,
       "table": {
         "table_name": "users",
         "access_type": "index",
         "key": "PRIMARY",
         "used_key_parts": [
           "id"
         ],
         "key_length": "8",
         "rows_examined_per_scan": 100,
         "rows_produced_per_join": 188976,
         "filtered": "20.99",
         "cost_info": {
           "read_cost": "159254.27",
           "eval_cost": "37795.33",
           "prefix_cost": "197049.60",
           "data_read_per_join": "974M"
         },
         "used_columns": [
           "id",
           "first_name",
           "last_name"
         ],
         "attached_condition": "((`soc_db`.`users`.`first_name` like 'to%') or (`soc_db`.`users`.`last_name` like 'to%'))"
       }
     }
   }
 }
```
**С индексом**

```mysql
create index f on users(first_name);
create index l on users(last_name);
explain format=json select id, first_name, last_name from users WHERE first_name LIKE 'to%' OR last_name LIKE 'to%' ORDER BY id limit 100;
```

```json
 {
   "query_block": {
     "select_id": 1,
     "cost_info": {
       "query_cost": "53897.79"
     },
     "ordering_operation": {
       "using_filesort": true,
       "table": {
         "table_name": "users",
         "access_type": "index_merge",
         "possible_keys": [
           "l",
           "f"
         ],
         "key": "sort_union(f,l)",
         "key_length": "768,768",
         "rows_examined_per_scan": 19637,
         "rows_produced_per_join": 19637,
         "filtered": "100.00",
         "cost_info": {
           "read_cost": "49970.39",
           "eval_cost": "3927.40",
           "prefix_cost": "53897.79",
           "data_read_per_join": "101M"
         },
         "used_columns": [
           "id",
           "first_name",
           "last_name"
         ],
         "attached_condition": "((`soc_db`.`users`.`first_name` like 'to%') or (`soc_db`.`users`.`last_name` like 'to%'))"
       }
     }
   }
 }
```

### Task1
Tech stack:
- golang
- mysql 5.7
- session JWT
- css bootstrap

[sql](https://github.com/ios116/social/blob/master/social/sql/setup.sql)
