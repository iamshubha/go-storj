Create a .env file in the root of the directory 

ACCESS_TOKEN, HOST_PORT & R_MODE

R_MODE belongs to any of this - debug/test/release

HOST_PORT = specify host port 

ACCESS_TOKEN = storj access token


To build the application in docker and run fire this 2 command 

1 - docker build -t storj-app .      
2 - docker run -d --name storj-container -p 8080:8080 storj-app
