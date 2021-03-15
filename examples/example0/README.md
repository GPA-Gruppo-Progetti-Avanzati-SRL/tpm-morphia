### Example - 0

Toy example to set the stage. 

TestFind: setup the collection (setup.go with the const for the collection names and mongo url) and query the records 
by searching for an author with firstName == "John" or an author with lastName == "Ward" and city == "Naples".
The find is carried out twice with different cities: in the first case two records are matched whereas the second find returns only one record.

TestUpdate: setup the collection and update the recvord for firstName == "Susan" and lastName = "Red".
