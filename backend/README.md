# Breakfaster - A Breakfast Ordering System @LINE
[![Build Status](http://morris.csie.ntu.edu.tw:5601/api/badges/minghsu0107/Breakfaster/status.svg)](http://morris.csie.ntu.edu.tw:5601/minghsu0107/Breakfaster)
## Usage
First, make sure that Docker and Docker-Compose are installed on your machine.

Then, edit environmental variables in `run-example.sh`. Please refer to `.env.example` for more details.

Next, launch the application by executing:
```bash
chmod +x run-example.sh
./run-example.sh
```

Wait for a while and the Breakfaster will be up and running!

You can visit `http://<host-ip:port>/api/v1/doc/index.html` for API documentation. Note that the API documentation is only available when `GIN_MODE=debug`.
## Model Training
On the Clova ChatBot dashboard, create a numbers of question-answer pairs. Current answers should include: `問題回報`, `取消訂單`, `點餐紀錄`, `規則`. See more details in [Chat Bot Custom API Documentation](https://apidocs.ncloud.com/en/ai-application-service/chatbot/chatbot/).
## Development
For development, you can start the DB and Redis containers only by executing:
```bash
export DB_USER=ming
export DB_USER_PASSWD=password
export DB_ROOT_PASSWD=password
export DB_NAME=breakfaster
export REDIS_CLUSTER_IP=$(ifconfig | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3}" \
    | grep -v 127.0.0.1 | awk '{ print $2 }' | cut -f2 -d: | head -n1)
export REDIS_PASSWD=pass.123

docker-compose up db redis-cluster-creator \
    redis-node1 redis-node2 redis-node3 redis-node4 redis-node5 redis-node6
```

MySQL database will be serving on `127.0.0.1:3306`. You can then set env `DB_DSN=<user>:<password>@tcp(127.0.0.1:3306)/breakfaster?charset=utf8mb4&parseTime=True&loc=Local` and connect to MySQL. Also, the Redis cluster will be serving on `REDIS_CLUSTER_IP:7000` to `REDIS_CLUSTER_IP:7005`. You can connect to the cluster via any one of the 6 Redis nodes.
