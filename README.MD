# Context
Created a POC for mimic the behaviour of any seat booking application

# Approach
1. Create a user table (id, name)
2. Create a seat table (id, name, user_id)
3. Try to acquire any free seat concurrently

# steps to run
1. docker-compose up 
2. go run setup.go
3. go run main.go
