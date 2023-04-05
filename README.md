# kredit-api
- generate nomor_kontrak automatically from the system when user make transaction
- handling user transactions with rabbitmq aims to stay smooth while receiving high transaction traffic

# Getting started
- create a database with the name "credits"
- import Query.sql database to mysql database
- go to the application project directory then run the command "docker-compose up"
- run the application by using the command "make start"
- import json data collection "credit-api.postman_collection.json" to test with postman