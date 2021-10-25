# kuberstore [K8s Exploration]
This is a exploration project for kubernetes on the context of a microservice architecture. 

### Actors
* Simple React client for adding a product
* Publisher microservice that generates a prodcut id and pushes an event to a RabbitMQ deployment
* Catalog microservice that is subscribed to the product events and creates a new "listing" to a Mongo database when a product is added
* Warehouse microservice that is subscribed to the product events and creates a new "stock item" to a Postgres database when a product is added
