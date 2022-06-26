# Test Project

Want to include:
- Event Driven Architecture?
- Event Sourcing
- CQRS
- Protobuff
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
- Kafka
- RabbitMQ
- My Own :) 


## Good Resources

- https://victoramartinez.com/posts/event-sourcing-in-go/