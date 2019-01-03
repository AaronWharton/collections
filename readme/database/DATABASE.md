# DATABASE

## mysql

- 常见的 sql 操作：
    - 修改表中字段的类型：  
    修改表中字段的类型并赋默认值：  
    ```sql
    ALTER TABLE table_name MODIFY COLUMN field_name type
    # 类型转换实例：注意转换可能会导致数据（精度）丢失
    ALTER TABLE test_table MODIFY COLUMN test_field char(30)
    ALTER TABLE test_table MODIFY COLUMN test_field varchar(60)

    ALTER TABLE table_name MODIFY COLUMN field_name type DEFAULT default_value
    ```
    
    - 查看表结构：
    ```mysql
    DESCRIBE table_name;
    ```
    
    ```
    Example:
    mysql> describe users;
    +------------+------------------+------+-----+---------+----------------+
    | Field      | Type             | Null | Key | Default | Extra          |
    +------------+------------------+------+-----+---------+----------------+
    | id         | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
    | login_name | varchar(64)      | YES  | UNI | NULL    |                |
    | pwd        | text             | YES  |     | NULL    |                |
    +------------+------------------+------+-----+---------+----------------+
    ```
    
    - 查看表使用的存储引擎（可用来查询建表语句）：
    ```mysql
    # 查看系统支持的引擎
    SHOW ENGINES;
    # 输出建表的详细信息
    SHOW TABLE STATUS FROM db_name WHERE NAME='table_name';
    # 输出建表的简略信息
    SHOW CREATE TABLE table_name;
    ```
    
    ```
    Example:
    mysql> show create table users;
    +-------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    | Table | Create Table                                                                                                                                                                                                                       |
    +-------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    | users | CREATE TABLE `users` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `login_name` varchar(64) DEFAULT NULL,
      `pwd` text,
      PRIMARY KEY (`id`),
      UNIQUE KEY `login_name` (`login_name`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 |
    +-------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
    ```