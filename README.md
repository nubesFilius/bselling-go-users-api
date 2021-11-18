# bselling-go-users-api

This is a microservice written in Golang using a MySQL DB.

It follows an MVC pattern where we implement the following components:
- App -> This is our main. Contains all the url routes pointing to the controller.
- Controller -> Will handle each operation to a Service.
- Service -> Where all our Business Logic is.
- Domain -> User(s) is our Domain. Here we interact with a Rest Provider or backend through a DAO(Data Access Objec), we define our main DTO(Data Transfer Object) and how we'll present the domain data to internal/external users.
- Datasource -> Our connection with the backend
- Logger -> All our logs config
- Utils -> Series of reusable pkgs to handle pwd encryption, date formatting, error handling and mysql error handling.

The microservice is used to get, create, fully update, partially update and delete user(s) with respective information stored in a backend db.