# Test Project

Want to include:
- Event Driven Architecture?
- Event Sourcing
- CQRS
- Protobuf
- Websockets

## CQRS

Command Query Responsibility Segregation.
Create a materialized view from the published events.


### Command

Very specific - i.e. UserAddedToCart, AddedItemToOrder, etc..


## Event Sourcing 

Use a log of events rather than a Database/Table.
Want to make it swappable as possible, for instance:
- SQL Table
- Mongo Table
- Google PubSub etc.
- Kafka
- RabbitMQ
- My Own :) 


### What Happens on a Query?


### Who and how do events get sent


### What happens when an event is received


## Aggregates

Effectively the object

## ProtoBuf

`protoc --go_out=. todo/todo.proto`

## Good Resources

- https://victoramartinez.com/posts/event-sourcing-in-go/