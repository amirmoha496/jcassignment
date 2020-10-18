# jcassignment

First Install this dependency 

#Run fallowing command
go get -u github.com/gorilla/mux

To run the application locally

#Run fallowing command
$ chmod +x run.sh
$ ./run.sh

Open the browser and try the application URLS at http://localhost:8080/
curl -d "password=angryMonkey" -X POST http://localhost:8080/hash
curl -X GET http://localhost:8080/hash/{id}
curl -X GET http://localhost:8080/stats


The application is also deployed on heroku.
The application URL for heroku is https://jcassignment.herokuapp.com/
If you hit the URL /shutdown the service will shutdown and then automatocally restarted.



