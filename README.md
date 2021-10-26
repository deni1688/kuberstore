# kuberstore [K8s Exploration]
This is a exploration project for kubernetes in the context of a microservice architecture. 

### Services
* [React TS] Simple client for adding a product
* [Go] Publisher microservice that generates a product id and pushes an event to a RabbitMQ deployment
* [Go] Catalog microservice that is subscribed to the product events and creates a new "listing" in a Mongo database when a product is added
* [Java Spring] Warehouse microservice that is subscribed to the product events and creates a new "stock item" in a Postgres database when a product is added

The arrow signify the flow of data.

![image](https://user-images.githubusercontent.com/14905199/138826387-a759f5e0-886a-4102-b4b0-4e2a350cc616.png)

