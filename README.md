# kuberstore [K8s Exploration]
This is a exploration project for kubernetes in the context of a microservice architecture. 

### Actors
* [React TS] Simple client for adding a product
* [Go] Publisher microservice that generates a product id and pushes an event to a RabbitMQ deployment
* [Go] Catalog microservice that is subscribed to the product events and creates a new "listing" to a Mongo database when a product is added
* [Java Spring] Warehouse microservice that is subscribed to the product events and creates a new "stock item" to a Postgres database when a product is added
