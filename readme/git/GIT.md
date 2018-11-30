# GIT

## 只删除远程仓库文件

把要忽略的文件（文件夹）添加到 `.gitignore` 中，提交使 `.gitignore` 生效：
```shell
git rm -r --cached xxx  # -r 表示递归，当 xxx 是文件夹时有用
git add xxx             # 若 .gitignore 文件中已经忽略了 xxx 则可以不用执行此句
git commit -m "ignore xxx"
git push
```
某些时候删除失败（比如在 Github 上创建 .gitignore ，但是没有在执行命令之前将 .gitignore pull 到本地），注意看到这个结果可表示命令生效：
[git status 状态](/Users/aaron/go/src/collections/readme/src/git_status.png)
