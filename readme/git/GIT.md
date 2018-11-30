# GIT

## 只删除远程仓库文件

把要忽略的文件（文件夹）添加到 `.gitignore` 中，提交使 `.gitignore` 生效：
```shell
git rm -r --cached xxx  # -r 表示递归，当 xxx 是文件夹时有用
git commit -m "ignore xxx"
git push
```
