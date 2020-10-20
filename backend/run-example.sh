#!/bin/sh
export DB_USER=ming
export DB_USER_PASSWD=password
export DB_ROOT_PASSWD=password
export DB_NAME=breakfaster

export PORT=80
export CHANNEL_SECRET=myLineBotChannelSecret
export ACCESS_TOKEN=myLineBotAccessToken
export BOT_VERSION=v1
export ORDER_PAGE_URI=https://liff.line.me/myLiffURI
export MAX_DB_IDLE_CONNS=10
export MAX_DB_OPEN_CONNS=100
export DB_DSN='ming:password@tcp(db:3306)/breakfaster?charset=utf8mb4&parseTime=True&loc=Local'
export GIN_MODE=debug
export LOG_PATH=server.log
export DEFAULT_CACHE_EXPIRATION=86400
export CLEAN_CACHE_INTERVAL=86400
export CLOVA_SECRET_KEY=myClovaSecretKey
export CLOVA_BUILDER_URL=https://myClovaChatBotEndpoint

docker-compose up