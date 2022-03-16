### 介绍

`gitlab-custom-hooks`为go实现的gitlab自定义脚本，目前已实现以下钩子：

- pre-receive：遵循`Angular`的commit规范检测。

### 使用

需要获取Gitlab的项目ID，步骤为：`设置-通用-项目ID`，接着使用`to-repo.sh`将钩子复制到gitlab对应项目的`custom_hooks`内。
#### 脚本参数
- GITLAB_HOME：gitlab数据的目录。
- RULES_HOME：规则存放的目录，及编译后脚本存放的位置。
> 此脚本是为了方便多个单项目要使用此规则来操作，若全部仓库要使用则开启全局钩子即可。
#### 例子
- 如`id=20`，则为`./to-repo.sh 20`

### 来源

- [pre-receive](https://github.com/mritd/pre-receive) （项目的`pre-receive`是在此基础上增加改进）