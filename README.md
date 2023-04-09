### kredit-api
- Built with domain design pattern and runs on **docker compose**
- Generate "nomor kontrak" automatically from the system when user create transaction
- Handling user transactions with rabbitmq aims to stay smooth while receiving high transaction traffic

### Getting started
- Go to the application project directory then run the command "docker-compose up --build -d"
- Wait a while for MySQL and RabbitMQ to be ready
- Import json data collection "credit-api.postman_collection.json" to test with postman
- Stop running application and remove container, run the command "docker-compose down"