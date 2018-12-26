# DATABASE

## mysql

- 常见的 sql 操作：  
    修改表中字段的类型：  
    修改表中字段的类型并赋默认值：  
    ```sql
    ALTER TABLE table_name MODIFY COLUMN field_name type
    # 类型转换实例：注意转换可能会导致数据丢失
    ALTER TABLE test_table MODIFY COLUMN test_field char(30)
    ALTER TABLE test_table MODIFY COLUMN test_field varchar(60)

    ALTER TABLE table_name MODIFY COLUMN field_name type DEFAULT default_value
    ```