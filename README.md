
# screen-wechat-bot
给我一个URL，我能截图发给企业微信机器人
构建容器镜像
docker build -t wechat-screenshot .
容器
程序在主机上运行的时候，如果检查到运行环境不满足需求，会自动下载相关依赖，如果你想来去无痕地运行程序，可以使用如下程序运行：
docker run --rm wechat-screenshot -u https://baidu.com -e "#s_lg_img" -k 1200 -g 800 -b "你的企业微信机器人key"
