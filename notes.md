Compiled
    Go tool can run file directory without VM
    Exectuables are differnet for OS

What
    System apps to web apps- cloud

Object Oriented
    yes & no
    What all you see is the code

Missing
    is it really missing
    no try catch
    Lexer does a lot of work

Question: What is Lexer?
Ans: A lexer, or lexical analyzer, is a program that breaks down a stream of characters into smaller units called tokens. It identifies and categorizes elements of a programming language, such as keywords and operators, to facilitate further analysis and interpretation by a parser.

###################################################################################################
The type of data that we use is always declared after the name
EX- a int, valOne int, name string
###################################################################################################

Type
    Case inisensitive; almost
    Variable type should be known in advance
    Everything is a type

There is no garbage value in int and string
the default value of int is 0 and string is " " empty

IF we use capital Letter at the time of declaration it is termed as public variable

Bufio: NewReader read until new line is hit and used as
    reader := bufio.NewReader(Os.stdin)
To read the string we use
    input := reader.ReadString('\n')

For Conversion we use stdconv
newRating, err := strconv.ParseFlow(strings,Trimspaces(input),64)




15:- Functions
    strat with func and name of function
        func main(){}
    Anonymus functions exit like func (){}
    To return the value we need to define the signature i.e. the type of data that will be returning
    EX- func adder (a int, b int) int{
	    return a+b
        }
    int is the signature
    ... is a veridatict function and can take any number of value
    
16:- Method
    When functions go in classes they are called methods
    method should contain struct as a part
    if a data is change the original data doesn't change thus pointer are used to give the actual reference

17:- Defer
    the execution of the text that is present after defer is done when the function is ending
    defer function invoked immediately before the surrounding function return int he reverse order they were defered (LIFO).

18-Files
    we need IOutils package and defer after when you are done reading and writing so needed to close the file
    Creation of file we use OS pacakge for the reading and writing or doing manupulation we use IOUTILITY package
    Data is always read in byte form so use DATABYTE OR DATA

19-webRequest
    it is always important to close the request once generated as if we switch to other function still the current request will be running in the background.
    defer response.Body.Close()
    hhtp Package is the package used to GET package is used to retrieve/get the website
    To read a file ioutils Package is used

20- urls
    To get the parmeters we use the value result.querry 
    in partOfUrl we give the value always in reference &url.URL

###- Creating server in golang
    -do npm start
    -routes and its usage 
    Thunderclient by ranga extenstion
    for post request we need to send some of the json data in body
    form need to be form-encode 

21- WebReqVerb
    Strings :- Builder:- a builder is used to efficiently build a string using Write method.
    responseString.String() this will convert the data that is holding in bytes to the 
    
    POST:- use /post
        USE jSON data and url encoded form in data for backend.
        NewReader syntax: it is a part of strings package and it has a thing that is if we put `` and write the data inside this it will be stored.
    
    PostFormRequest:- to make a request we use http.PostForm(myurl,data)

22-Bitmorejson
    to give tags we use this type [] strings  [] (data type) and the data inserted is in the form of []string{"abc","xyz"}
    We use JSON package
    we need to pass interface in marshal now what is Interface?
    its a word being borrowed and an alternative version of struct
    JSON is also byte data
    to represent a well managed data we use marshalindent and in prefix we dont need to give anything and to indent we need to write "\t"
    to give the field in a better way or to change the name of the field we use `json:name` this is the name that will be depicted and if we want to hide some fie we use `json : -` this will hide the field to be displayed in terminal
    whenever we use tags, omite we use a space after comma in this and most of the time it gets error

    when the data comes it is always in bytes
    if you are not sure about the data that is going to be entered then use interface

23- MOD
    indiret = any file in go is not using the library when once you start using it it will be changed to direct.
    ServeHome (just teh name) is a type of response in which we can write the response using 'w' and read the response and this is done tih the help of the gorilla mux package.
    basically mux is used to give the routing and a fucntion is used to call the system locally.
    Log fatal is used to check for the error

    What is go mod tidy?
     is removed it tidy all the package it remove all the unused package and brings back the package which are used and are  removed by mistake

    For Production case:- we use go mod vendor it is same as node module
