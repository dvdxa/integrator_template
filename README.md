# integrator_template
Template for integration service between two banks(agent and vendor).

## Architecture
The app was designed in layered architecture. The architecture combines Clean Architecture and Hexagonal Architecture.

## Adapter Layer
The adapter layer is implemented using adapter, which is responsible for interaction with vendor, it provides an 
abstraction over it so other parts of
application that call adapter don't care about the origin and decoupled  from specific implementations used.

## Domain Layer
The domain layer is implemented using service. It depends on adapter get necessary data from vendor
for further processing. Service not coupled to specific adapter implementation and can be reused if we add more adapters.

## Handler Layer
The handler layer depends on the domain layer and interacts with agent. We defined the route that can be called by agent
to initiate request to vendor.