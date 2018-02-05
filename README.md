# simple_db

## Installation 
In order to clone and run the project you must have a valid Go-installation on your computer. If you have not installed Go on your computer, a guide can be found [here.](https://golang.org/doc/install)

If Go is installed you can simply just clone the project and execute the *Aflevering_1* file with matching arguments for the file. 
If you're on a Windows computer you might have to type *go install* or *go build* in order to make a new executable that Windows can read. Both commands will build a new executeable, but in different locations. *go build* should create a new executable in the root of the project. 

## Arguments
In order to run the program you must specify what operation you wish to use on the database. The project supports *get* and *set* on the database.
### get
```
./Aflevering_1 get 2
```
This will get you the entry with id 2. You can only query on id's, and the id has to exist - otherwise the program will fail. 
### set
```
./Aflevering_1 set "this is a test"
```
This will create a new entry in the database. It is important to write the second argument in quotes if you wish to insert a line of text.

## Indexes
When the user inserts a new entry in the database a byteOffset for that entry is stored in a hashmap. This way when the user tries to query from the database the hashmap will supply at what byteOffset the program should look for in the database-text-file. This way, the program doesn't look through the entire file but at a certain offset. The program will then read all bytes until a newline is found. Afterwards the entry the user was looking for will be returned. 

## Notes
The file *index.txt* should never be completely empty if you wish to reset the database. The file must at least contain a pair of brackets "{}" in order to work. 
