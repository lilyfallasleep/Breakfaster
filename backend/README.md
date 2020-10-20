# Breakfaster - A Breakfast Ordering System @LINE
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
For development, you can start the DB container only by executing:
```bash
docker-compose up db
```

The database will be serving on `127.0.0.1:3306`. You can then set env `DB_DSN=<user>:<password>@tcp(127.0.0.1:3306)/breakfaster?charset=utf8mb4&parseTime=True&loc=Local` and connect to the database.
