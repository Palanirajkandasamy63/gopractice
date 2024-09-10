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




go learning 

go is faster 
less memory usage 
go routines are simpler 

variables which stores values
var := 20
var confe int=20


datattypes
int 
uint
float complex 

fmt 
printf -->we can use whenever we need this  like %v we can substirute values and call 
sprintf(we can store it in a variable and call 
)

if else
if initial statement condition{}
else if condition
else

length:=getlength(email)
if lenth<1{
    fmt.Println("invalid email)
}
we can write as below 

if length:= getlength(email); length <1{
    fmt.Println("invalid email)
}


Functions 

func sub (x int , y int) int{
    return x-y
}

function signature -->what the function gets and what output it returns func sub (x int , y int) int

fun add(x,y int ) int {
    return x+y
}

fun createUSer(firstname string , lastname string , age int ) can we write it 
fun createuUSer(firstname lastname String, age int)

f func(func(int,int)int,int)int
a function named f takes  a function as a arugument and also int as a argument which returns int

go 
values are passed by value not by reference 

x:=5  --->memory 5 
x:2----->memory 2 
y:x----> memory new 2 
y=1---->memory =1  so x wont be unaffected in memory 


func main(){
    x :=5
    increment(x)
    fmt.println(x)

}
func increment (x int){
    x++
}

output is 5 because a new copy of x is create din function so it is unaffected 


func main(){
    x :=5
    x:= increment(x)
    fmt.println(x)

}
func increment (x int){
    x++
}
now the output will be 6 

ignoring a value to be unuse 

func getPoint()(x int , y int )int {
    return 3 ,4
}
to ignore y value 

x ,_ :=getPoint()


guard clause 
early return 


struct 

type bike struct{
    make string 
    model string 
    heigt int 
}

nested struct 

type bike struct{
    Make string 
    Model string 
    Heigt int 
    Frontwheel wheel
    Backwheel wheel
}

type wheel struct{
    Radius  nt 
    Material string
}

anaonyms struct 
if we are nto using more than one time so we use this 


embedded strunc

type car struct {
    make string 
    model string 
}

type truck struct{
    car 
    bedsize int 
}

accessing embedded struct 

lanesTruck :=truck{
    bedsize:10,
    car: car{
        make: "toyato",
        model:"camary",
    },

}

fmt.println(lanestruck.make)
fmt.println(lanes.truck.model)


func tests(s sender){
    fmt.println()
}

using a method on struct 
func (r rect ) area() int{
    return r.width * r .height
}



#interfaces 

type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}


loops 

package main

import "fmt"

func main(){
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }

   var temp ,sum int 

   for i:=0; i<5;i++{
	fmt.Scanf("%d\n",&temp)
	sum=sum+temp
   }
   fmt.Printf("%d",sum)


   //while loop like for loop 
   i:=0;
   for i<5{
	fmt.Scanf("%d\n",&temp)
	sum=sum+temp
	i++
   }
   fmt.Printf("%d",sum)
   }

   to exit infinite loop 

   we can use break


   arrays 

   var arr[5]int 
   it can only store 5 elemets 
   arr =[5]int{78,56,55,66,52} similar to arr =[...]int{78,56,55,66,52}
   for i:=0; i<len(arr);i++{
    fmt.printf("%d: %d \n",i arr[i])
   }

   for i , n := range arr{
    fmt.Printf(%d\n", n ,i , n )
   }

   
   for i , n := range arr{
    fmt.Printf(%d\n", n ,i , n )
   }


   slices 

   slies are like dynamic where 

   extracting slices form array 

   daysOfWeek :=[...]string {"sun","mon","tue","wed","thus","fri","sat'"}


    weekdays:= daysfweek[1:6]
    weekdays[2]="wensday"
    printslice(daysofweek[:])
    printslice(weekdays)

    func Printslice(slice []string){
        fmt.Printf("%v",slice)
    }


    var list[]int 
    list=append(list, 1,2,)

    capacity is additional on ein go 

    func printSlic(slice []int){
        slice ,len(slice),cap(slice)
    }