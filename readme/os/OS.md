# OS

## macOS

- 使用 `homebrew` 安装 mysql ： `brew install mysql`
    - 启动/关闭一个后台 mysql server ： `brew services start/stop mysql`
    - 启动/关闭一个前台 mysql server ： `mysql.server start/stop`
    - 连接 mysql ： `mysql -uroot`
    - 启动 mysql 出错 `mysqld_safe Directory '/usr/local/mysql/tmp' for UNIX socket file don't exists.` ：
        - 配置文件内容出错，删除或者修改 `/etc/my.cnf` ，此时 mysql 去读取其他位置的配置文件。
    - 启动 mysql 出错 `Starting MySQL... ERROR! The server quit without updating PID file` ：
        - 系统已经运行了一个 mysql 进程，杀死所有 mysql 进程重新启动即可：
            ```shell
            ➜  ~ killall -9 mysql mysqld
            ➜  ~ mysql.server start
            
            ➜  ~ mysql_secure_installation  # 设置 mysql 密码
            ```

## 跨平台

- golang 跨平台编译：虽然 `go build` 可以让可执行文件跨平台运行，但是 build 时要指定系统类型才能让可执行文件在相应的平台上运行。类 Unix 系统可使用 `uname -a` 查看系统类型，编译最简单的方式是指定系统类型（架构相对复杂，一般不指定也可运行）可解决 `cannot execute binary file` ：
    ```
    GOOS=windows go build xxx.go    // 生成 xxx.exe
    GOOS=linux   go build xxx.go    // 生成 xxx
    ```