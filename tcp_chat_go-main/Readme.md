# TCP chat in GO


### libs
<br />
Go net package and goroutines and channels

<br />

### tcp
<br />  
tcp or transmission  control protocol  <br />
provide transport mechanism for application layer  <br />
such as http smtp irc etc  <br />
<br />

### Commands

- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time. 
- `/rooms` - show list of available rooms to join.
- `/msg <msg>` - broadcast message  to everyone in a room.
- `/quite` - disconnects from the chat server.

### types

* client (name of the user current connection to a room)
* room   (list of member of the room)
* comment (comment of the client to server such as join etc)
* centralized server (responsible for incoming commands and storing rooms and ...)
* tcp server (for accepting connection)