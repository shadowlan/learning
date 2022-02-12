# git-fork

在github里fork一个repo的动作是很简单的，点个按钮就能做到，不过怎么和上游的改动保持同步之前没做过，今天操作了一下，发现还有个小坑。具体的更新和操作步骤记录如下：

1. 先查看本地仓库的remote信息：`git remote -v`
2. 添加远程上游仓库： `git remote add upstream https://github.com/ORIGINAL_OWNER/ORIGINAL_REPOSITORY.git`
3. 此时再用`git remote -v`查看本地目录信息会发现除了origin的fetch/pull远程路径外，还有upstream的远程路径
4. 运行`git fetch upstream`
5. 运行`git checkout master`
6. 将来自upstream/master的更改合并到本地master分支中。 这会使复刻的master分支与上游仓库同步，而不会丢失本地更改。
7. 如果没有任何冲突需要解决，git会执行快速合并，如果有冲突，需要解决冲突后检入。

在这个过程中更新和合并远程代码都没有问题，但是当我在当前仓库下继续执行`hub pull-request`的时候却出了状况，新建的PR直接提到了被fork的主仓库中去，难道hub命令在创建PR的时候会优先使用upstream定义的远程仓库路径？不过当我执行`git remote remove upstream`后再尝试建PR的时候，就从我fork后的远程仓库创建PR了。看来upstream的确是hub会优先选择的。

hub的pull-request代码在[这里](https://github.com/github/hub/blob/b45ad4b7611c53ffd1c85a9acaceb8b36de89999/commands/pull_request.go),看起来可能除了删除upstream的定义外，应该还可以通过传入参数-b来解决这个问题。

传送门：[upstream vs origin](https://blog.csdn.net/liuchaoxuan/article/details/80656145)