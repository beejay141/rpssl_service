#### R.P.S.S.L API Service

------------------------

#### Dependencies
  * Redis server
  * Internet connection (To make api call to the random number service)
  
  
#### Setup Config

- With Go command
  * Pull the code
  * run go command to sync packages
  * make sure you have redis run on default port on your local machine or 
  config the host using the `.env` file
  * run `go run main.go` to startup the app
  
- With docker-compose
 * pull the code
 * make sure you have docker installed on your local machine
 * run `docker-compose up` to build and start up the service
 
#### API documentation
 
> Choice Controller
- GET :`/choices`
  * returns all the available choices
- GET : `/choice`
  * returns a random choice
      
> Play Controller
- POST : `/play`
- BODY : `{player: 1 ,playerId:"bolaji"}`
  * player: player's choice, required
  * playerId: to track player's scores, optional
- RESULT : `{player:1,computer:2,results:"lose"}`
   
> Scoreboard Controller
- GET : `/score-board/:id`
  * param : playerId
- DELETE : `/score-board/reset/:id`
  * param : playerId
  
   