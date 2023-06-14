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
    if a data is change the original data doesn't change thus pointer are used to give the actual refrence

17:- Defer
    the execution of the text that is present after defer is done when the function is cd

