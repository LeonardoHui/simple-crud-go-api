This is a simple API created for learning purpose.

The original project belongs to [karanpratapsingh](https://github.com/karanpratapsingh) whose videos were followed to develop the basis of this one, please check out his videos:

-   [Build a REST API with Go - For Beginners (Part 1)](https://www.youtube.com/watch?v=bFYZrEuEDLE&t=0s&ab_channel=Karan)
-   [Build a REST API with Go - Connecting to PostgreSQL using GORM (Part 2)](https://www.youtube.com/watch?v=Yk5ZjKq4qDQ&ab_channel=Karan)

To connect to your local postgres

-   create a DB named "crud"
-   set the user to "postgress"
-   set the password to "root"

or change the line 12 on pkg > db > db.go to fit your postgres configuration

These were used to write unit testing

-   https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
-   https://github.com/gorilla/mux#testing-handlers

To run unit tests:

-   go test -v ./pkg/... -coverprofile cover.out
-   go tool cover -html="cover.out"
