# screen-wechat-bot
给我一个URL，我能截图发给企业微信机器人

docker build -t wechat-screenshot .

docker run --rm wechat-screenshot -u https://baidu.com -e "#s_lg_img" -k 1200 -g 800 -b "你的企业微信机器人key"
