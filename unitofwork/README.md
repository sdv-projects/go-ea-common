# Unit of work

A unit of work is a behavioral pattern in software development. A Unit of Work keeps track of everything you do during a business transaction that can affect the database. When you're done, it figures out everything that needs to be done to alter the database as a result of your work. The pattern is used to improve performance by reducing transaction duration and blocking for the use case.

Martin Fowler, This pattern is part of Patterns of Enterprise Application Architecture:
https://martinfowler.com/eaaCatalog/unitOfWork.html

[<img src="https://martinfowler.com/eaaCatalog/index/unitOfWorkInterface.gif">](https://martinfowler.com/eaaCatalog/index/unitOfWorkInterface.gif)


# How to use

Example of use unitofwork.UnitOfWork 
```go
import (
     unitofwork "github.com/sdv-projects/go-ea-common/unitofwork"
     database "github.com/sdv-projects/go-ea-common/database"
)
 
var (
   tf database.TxFactory
   rf unitofwork.UoWRepositoryFactory
)
 
// Initialize TxFactory and UoWRepositoryFactory
...
 
// Create UnitOfWork instance
uow := unitofwork.NewUoW(tf, rf)
 
...
// Register changes
uow.RegisterNew(entity)     // adds and marks the entity as new
uow.RegisterDirty(entity)   // adds and marks the entity as updated
uow.RegisterDeleted(entity) // adds and marks the entity as deleted
 
...
// Commit applies changes to the repository in a single transaction.
err := uow.Commit(ctx)
if err != nil {
   ...
}
```