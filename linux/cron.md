# Crontab

感觉我好像从来没把crontab里五个设定时间参数的具体含义记住，每次想要设置cronjob都得现翻历史记录或者google一下，以防万一自己设置有误。这一次次的反复我自己都不胜其烦了，，今天一定好好来记一下。

| * | * | * | * | * |
| --- | --- | --- | --- | --- |
| 分钟 | 小时 | 日 | 月 | 星期 |

1. 没有网络的情况下，怎么查看crontab的帮助呢？`man crontab` 是最基础的，不过发现这个帮助页面里并没有关于具体字段的含义，但从see also里有个`crontab(5)`, 执行`man 5 crontab`之后就出来具体的字段含义了。下回没记住也不慌了。哈哈。（ps：man里5代表什么含义可以移步这里： https://pjf.name/blogs/374.html）

2. 如果还是没记住，在能上网的环境下，可以访问站点https://crontab.guru/，这是个非常不错的crontab在线工具，除了能输入不同的组合后提示相应的执行间隔含义，站主还在[example](https://crontab.guru/examples.html)页面提供了一些非常常用的cron schedule。从这个网站顺藤摸瓜还能发现有一个专门提供crontab状态监控的服务网站👉https://cronitor.io，不得不感慨，别把豆包不当干粮，小小功能也都能做出个生意。佩服佩服

_20191010_