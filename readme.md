go mod init booking-app this initialize a mod file  describes the moduel with name/module path 
go mod init (path for moduel)
the module path is also a import match 
all code must be defined in package 

we need to tell the entrypoint where it needs to start excuting 
func main(){
    we need to tell so one applicatuion one main program 
}
print also needto be imported from fmt so import those 
a package is a collection of source files 

go run filename --compiles and runs

variable -->same like all progrlanga 
to create variables var conferanceName="go conferance" using var 
var conferancename int=50 
conferancename:=50 but here we cant explictly define type but wedont need to use var 

const values can be changed when used 

print f --->formated data 
fmt.Println("welcome to",conferanceName,"booking application ")
fmt.Printf("welcome to %v booking application\n ",conferanceName)
mt.Println("we have total of ",conferanceTickets,"tickets and",remainingTickets,"are still availbel")
fmt.Printf("we have total of %v tickets and %v are still availbel",conferanceTickets,remainingTickets)


to find the type of variable we can use %T

    fmt.Printf("confertickeis type %T",conferanceName)

we also have otger for boolean like  for boolean %t

fmt.scan(&username) to get user input 
&pointer 
variables are stored in memory find th evalue in memory so it needs address so we need a pointer so is a variable that points to the memory address of another variable 


data types  strings integers boolean maps arrays struct
string "textual ab +cd 
intege 22 , cal with integer 
byte ,float int long we have some for int so if valkue is not negative we can declare uint so that if we change negative value it wont work
floating point types for high precisons 

so go can now value when we assign imedoiatley so we need specifall define the type explicity 

typecast 
we need to change the type of user ticket or type cast 
 remainingTickets=remainingTickets-uint(userTicket)

 arays and slices in go is like array
to store multiple user ticket details we are using this one only same datatype

"nana ,"tom","nicole",Miachel"
arrays yntas 
var variable_name[size] varaible_type{"nana","nicol","peter"}
var variable_name[size] varaible_type{}

var bookings =[50] size of an arrays 
var bookings[50]string

arrays are index based 
bookings[0]="nana"
bookings[1}="raju"
len(bookings)to find len 
if we specify 52 it more than arrays size index out of bound 

slices
slices are are more flexible which is declared with variable 
var bookings[]string
appends 
to add value 
bookings =append(bookings,firstname+" "+lastname)

//var bookings[]string like this we cna declare 
var bookings =[]string{} like this or 
bookings :=[]strings{}

fmt.Printf("the whole slice %v\n",bookings)
	 fmt.Printf("the first element in  slice %v\n",bookings[0])
	 fmt.Printf("the type %T\n",bookings)
	 fmt.Printf("the  slice length  %v\n",len(bookings))



loops 
for {

}

specific loop 
for index, booking :=range bookings{
    strings.Fields(booking)
    var names =strings.Fields(booking)
    var firstname =names[0]

} -->syntax
1st iteration 0 'nj cam' string field does [nj cam]
2nd iter      1 'tg raj'
3 iter        2 "ram kumar"


string .fields splits the values 

_ are used to for using blankidentiifer 
	for index,booking :=range bookings{
		var names=strings.Fields(booking)
		firstNames = append(firstNames, names[0])

        remove index and use _

get firstname after booking is changed to list of maps 

bookings =["raj palani", "ram kumar"]
elemnt is a string 
"raj palani"
while looping 

now e have list of maps 
bookings =[{firstName:"raj",lastName:"palani}
{firstName:"ram",lastName:"kumar"}]
element is a map 

{firstname:"raj", lastname:"palan"}
so we need to access the key 


waitgroup in sync 
add
wait -->end 
done -->removes the thred from waiting list 