# Welcome!

This is Alex Chow's internship assignment for the Summer 2019 Topos back-end engineer internship.

## Remarks and Rationale

This was actually my first time working with Go, which has proved to be a lot of fun to learn. The process has been challenging, but overall enjoyable, and I am grateful for the opportunity to learn something new.

Previous work in back-end engineering has taught me that some priorities are (in no particular order):

- Ease of Use
- Power
- Configurability / Adaptability

I sought to fulfill these as best as I can, which informed my design philosophy. I present 3 runnable programs, 2 relating to the ETL and 1 relating to the API. createAndImport creates a new database then imports some starter data. justImport skips the database check and is useful after initial setup, only importing the the data specified by the user. API is the API, which allows for http GET requests to retrieve data from the database and returning them in JSON form. I am also including a config.json file, where you can set the database username and password to be used, set the database name you want created, the table name in the database, the host of the database (here just localhost for local functions), and the port the API listens to.

I present two functions because while one function works just fine thanks to how mySQL handles database creation, it speeds it up ever so slightly, and more importantly. aims to reduce confusion by making it clear that this function works fine when there is a database that already exists. I also use a config.json file because of its human-readability and ease of changing some values that a user would want to customize without having these in the source code. It also means you don't have to keep rebuilding the functions when you just want to change the user in the database.

## Dependencies

I did my best to build them into runnable binaries that include the dependencies. If they don't work, or you want to build them from source, here are the external libraries I used:

- [mux](github.com/gorilla/mux)
- [Go-MySQL-Driver](github.com/go-sql-driver/mysql)

Everything else should be included in Go.

## Instructions

Okay, enough talk. Time to run the program! I'm assuming the latest versions of both Go and mySQL are being used.

1. Install and configure mySQL Server how you like it. You can find mySQL server [here.](https://dev.mysql.com/downloads/mysql/) Instructions for installing mySQL can be found [here.](https://dev.mysql.com/doc/refman/8.0/en/installing.html)
2. Setup config.json. You can use the default values just fine, the main things that need to be changed are the username and password for the SQL server.
3. I included a testbuilding.csv file with 30 rows from the larger dataset (of over 1 million records). Feel free to use this or download the full dataset and use some (or all) of that data from [here.](https://data.cityofnewyork.us/Housing-Development/Building-Footprints/nqwf-w8eh)
4. Install the dependencies above and then build from source using go build [function].go from the working directory.
5. The first program that must be run is createAndImport, as that creates the database (assuming you set your username and password correctly in config,json!). A command prompt will pop up, asking for a file. Type the full filepath of the file you want imported to the database and hit enter. If you entered a valid path it should return a success message.
6. To retrieve data from the API, run the API binary, and either use a browser or a program such as Postman to call http://localhost:port, where port is what you set your port to in the config file, with the default being 8086, followed by the endpoint, listed below.

## API endpoints
- /getData/ALL -- Returns all the data in the table. Of questionable usage for larger tables so there's...
- /getData/BIN/{BIN} -- Returns all the data about a given BIN. For rows with the unassigned BIN (X000000) this will return all info about unassigned BIN buildings.
- /getData/Random/{Count} -- Returns a given number of random rows from the table. This is more useful than just one.
- /getData/Borough/{Borough Number} -- Returns all the data about buildings from a Borough using the NYC's own codes, listed below.
- /getData/Type/{LSTSTATTYPE} -- Returns information about buildings with a given last status type.
- /getData/Year/{Year} -- Returns the buildings constructed in a certain year.
- /getData/Feature/{Feature Code} -- Returns all buildings matching a certain feature code.
- /aggregate/{Operation}/{Column} -- Allows for any of the mySQL aggregation functions on a column, see below for a list of tested aggregation functions.

## Borough Codes
- 1 = Manhattan 
- 2 = The Bronx
- 3 = Brooklyn
- 4 = Queens
- 5 = Staten Island

## Status Types
- Demolition
- Alteration
- Geometry
- Initialization
- Correction
- Marked for Construction
- Marked For Demolition
- Constructed

## Feature Codes
- 2100 = Building
- 5100 = Building Under Construction
- 5110 = Garage
- 2110 = Skybridge
- 1001 = Gas Station Canopy
- 1002 = Storage Tank
- 1003 = Placeholder (triangle for permitted bldg)
- 1004 = Auxiliary Structure (non-addressable, not garage)
- 1005 = Temporary Structure (e.g. construction trailer)

## Aggregation Functions
- AVG to compute the average value of column_name
- STD to compute the standard deviation of column_name
- SUM to compute the sum
- MIN to find the minimum value
- MAX to fund the maximum value
- COUNT to count the total number of entries having a not null entry in column

