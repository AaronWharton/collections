# OS

## macOS

- 使用 `homebrew` 安装 mysql
    - 启动/关闭一个后台 mysql server ： `brew services start/stop mysql`
    - 启动/关闭一个前台 mysql server ： `mysql.server start/stop`
    - 连接 mysql ： `mysql -uroot`
    - 启动 mysql 出错 `mysqld_safe Directory '/usr/local/mysql/tmp' for UNIX socket file don't exists.` ：删除或者修改 `/etc/my.cnf` ，此时 mysql 去读取其他位置的配置文件。因为 `/etc/my.cnf` 里的脚本路径不存在。