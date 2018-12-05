# GIT

## 只删除远程仓库文件

把要忽略的文件（文件夹）添加到 `.gitignore` 中，提交使 `.gitignore` 生效。某些时候删除失败，是因为添加规则之前，远程仓库就提交了要被忽略的文件，此时需要删除远程仓库的文件：
```shell
git rm -r --cached xxx  # -r 表示递归，当 xxx 是文件夹时有用
git add xxx             # 若 .gitignore 文件中已经忽略了 xxx 则可以不用执行此句
git commit -m "ignore xxx"
git push
```
