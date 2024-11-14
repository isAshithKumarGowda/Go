// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

// Output: Hello World

// package main

// import "fmt"

// // Define a function that matches the type of the first parameter in 'f'
// func multiply(a, b int) int {
//     return a * b
// }

// func main() {
//     // Declare 'f' as a variable of function type
//     var f func(func(int, int) int, int) int

//     // Assign a function to 'f'
//     f = func(op func(int, int) int, x int) int {
//         return op(x, 2) // Call the passed function with x and 2
//     }io

//     // Use 'f' with 'multiply' and 5
//     result := f(multiply, 5)
//     fmt.Println(result)
// }

// Output: 10

// package main

// import "fmt"

// func main() {
// 	sendsSoFar := 430
// 	const sendsToAdd = 25
// 	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
// 	fmt.Println("you've sent",sendsSoFar, "messages")
// }

// func incrementSends(sendsSoFar, sendsToAdd int)int {
// 	sendsSoFar = sendsSoFar + sendsToAdd
// 	return sendsSoFar
// }

// Output: you've sent 455 messages

// package main

// import "fmt"

// func main() {
// 	firstName, _:=getNames()
// 	fmt.Println("Welcome to Textio,", firstName)
// }

// func getNames()(string, string) {
// 	return "John", "Doe"
// }

// Output: Welcome to Textio, John

// package main

// import "fmt"

// func yearsUntilEvents(age int) (
// 	yearsUntilAdult int,
// 	yearsUntilDrinking int,
// 	yearsUntilCarRental int,
// ) {
// 	yearsUntilAdult = 18 - age
// 	if yearsUntilAdult <0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrinking = 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}
// 	yearsUntilCarRental = 25 - age
// 	if yearsUntilCarRental < 0{
// 		yearsUntilCarRental = 0
// 	}
// 	return
// }

// func main(){
// 	yearsUntilAdult,yearsUntilDrinking,yearsUntilCarRental := yearsUntilEvents(5)
// 	fmt.Println("Age until u are adult ", yearsUntilAdult)
// 	fmt.Println("Age until you can drink ",yearsUntilDrinking)
// 	fmt.Println("Age until you can rent a car ",yearsUntilCarRental)

// }

// package main

// import "fmt"

// type messagesToSend struct {
// 	phoneNumber int
// 	message     string
// }

// func test(m messagesToSend) {
// 	fmt.Printf("Sending message : '%s' to :%v \n", m.message, m.phoneNumber)
// 	fmt.Println("===========================================================")
// }

// func main() {
// 	test(messagesToSend{
// 		phoneNumber: 8951169155,
// 		message:     "Thank you for signing up",
// 	})
// }

// Output: Sending message : 'Thank you for signing up' to :8951169155
// 		   ===========================================================

// package main

// import "fmt"

// type messagesToSend struct {
// 	message string
// 	sender user
// 	recipient user
// }

// type user struct {
// 	name string
// 	number int
// }

// func canSendMessage(mToSend messagesToSend) bool {
// 	if mToSend.sender.name == "" {
// 		return false
// 	}
// 	if mToSend.recipient.name == "" {
// 		return false
// 	}
// 	if mToSend.sender.number == 0 {
// 		return false
// 	}
// 	if mToSend.recipient.number == 0 {
// 		return false
// 	}
// 	return true
// }

// func main() {
// 	if canSendMessage(messagesToSend{
// 		message: "hello world",
// 		sender: user{
// 			name: "Ashith",
// 			number: 8951169155,
// 		},
// 		recipient: user{
// 			name: "Anikethan",
// 			number: 8147920866,
// 		},
// 	}){
// 		fmt.Println("User Details are correct ")
// 	} else {
// 		fmt.Println("User Details are not correct ")
// 	}
// }

// Output: User Details are correct

// package main

// import "fmt"

// type user struct {
//     name string
//     number int
// }

// type sender struct {
//     rateLimit int
//     user
// }

// func test(s sender) {
//     fmt.Println("Sender name: ",s.name)
//     fmt.Println("Sender number: ",s.number)
//     fmt.Println("Sender Rate Limit: ",s.rateLimit)
// }

// func main() {
//     test(sender{ rateLimit: 10, user: user{ name: "Ashith", number: 8951169155,}});
// }

// Output: Sender name:  Ashith
//         Sender number:  8951169155
//         Sender Rate Limit:  10

// package main

// import "fmt"

// type rectangle struct {
//     width int
//     height int
// }

// func (r rectangle) area() int {
//     return r.width * r.height
// }

// func main() {
//     r := rectangle{
//         width: 10,
//         height: 5,
//     }
//     fmt.Printf("Area of the rectangle is : %d\n",r.area())
// }

// Output: Area of the rectangle is : 50

// package main

// import (
// 	"fmt"
	
// )

// type authenticationInfo struct {
//     userName string
//     password string
// }

// func (authI authenticationInfo) getBasicAuth() string {
//     return fmt.Sprintf(
//         "Authentication: Basic %s: %s",
//         authI.userName,
//         authI.password,
//     )
// }


// func main() {
//     userInfo := authenticationInfo{
//         userName: "Ashith",
//         password: "123456",
//     }
//     fmt.Printf("%s",userInfo.getBasicAuth())
// }

// Output: Authentication: Basic Ashith: 123456


// package main 

// import (
//     "fmt"
//     "time"
// )

// func sendMessage(msg message) {
//     fmt.Println(msg.getMessage())
// }

// type message interface {
//     getMessage() string
// }

// type birthdayMessage struct {
//     name string
//     birthdayTime time.Time
// }

// type sendingReport struct {
//     reportName string
//     numberOfSends int
// }

// func (bm birthdayMessage) getMessage() string {
//     return fmt.Sprintf("Hi %s, it's your BIRTHDAY on %v",bm.name,bm.birthdayTime)
// }

// func (sr sendingReport) getMessage() string {
//     return fmt.Sprintf("Your \" %s \" report is ready. You've sent %d messages.",sr.reportName,sr.numberOfSends)
// }

// func test(m message) {
//     sendMessage(m)
//     fmt.Println("=================================================io==============")
// }


// func main () {
//     test(sendingReport{
//         reportName: "First Report",
//         numberOfSends: 10,
//     })
//     test(birthdayMessage{
//         name: "Ashith",
//         birthdayTime: time.Date(2005, 05, 03, 0, 34, 0, 0, time.UTC),
//     })
// }

// Output: Your " First Report " report is ready. You've sent 10 messages.
//         ===============================================================
//         Hi Ashith, it's your BIRTHDAY on 2005-05-03 00:34:00 +0000 UTC
//         ===============================================================


// package main

// import "fmt"
// TYPE ASSERTION 

// func  main()  {
//     var i interface {} = "Hello, Go"
//     s, ok := i.(string)
//     if ok {
//         fmt.Println(s)
//     } else {
//         fmt.Println("i is not a string")
//     }
 

//     n, ok := i.(int)
//     if ok {
//         fmt.Println(n)
//     } else {
//         fmt.Println("I am not a int")
//     }
// }io

// Output: Hello, Go
//         I am not a int


// package main

// import "fmt"

// func main() {
//     user, err := getUser()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     // use user here
//     profile, err := getUserProfile(user.ID)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
// }


// func getUser() (User, error){
//     //DO SOME GETUSER LOGIC
//     return user , nil
// }

// Output: complete the code


// package main

// import "fmt"


// func main() {
//     totalCost := 0.0
//     for i := 0; i < 10; i++ {io
//         totalCost += 1.0 + (0.01 * float64(i))
//     }
//     fmt.Println(totalCost)
// }
// Output: 10.45

// package main

// import "fmt"

// func main () {
//     for i := 1; i <= 100; i++ {
//         if i % 3 == 0 && i % 5 == 0 {
//             fmt.Printf("fizzbuzz\t")
//         }else if i % 3 == 0 {
//             fmt.Printf("fizz\t")
//         }else if i % 5 == 0 {
//             fmt.Printf("buzz\t")
//         }else {
//             fmt.Printf("%d\t",i)
//         }
//     }
// }

// Output:1       2       fizz    4       buzz    fizz    7       8       fizz    buzz    11      fizz    13      14      fizzbuzz        16      17      fizz    19 buzz     fizz    22      23      fizz    buzz    26      fizz    28      29      fizzbuzz        31      32      fizz    34      buzz    fizz    37      38 fizz     buzz    41      fizz    43      44      fizzbuzz        46      47      fizz    49      buzz    fizz    52      53      fizz    buzz    56      fizz58      59      fizzbuzz        61      62      fizz    64      buzz    fizz    67      68      fizz    buzz    71      fizz    73      74      fizzbuzz   76       77      fizz    79      buzz    fizz    82      83      fizz    buzz    86      fizz    88      89      fizzbuzz        91      92      fizz    94 buzz     fizz    97      98      fizz    buzz 



// package main

// import "fmt"

// func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
//     actualCostInPennies := 1.0
//     maxMessagesToSend := 0
//     for maxCostInPennies > int(actualCostInPennies) {
//         // actualCostInPennies <= float64(maxCostInPennies)
//         maxMessagesToSend++
//         actualCostInPennies *= costMultiplier
//     }
//     return maxMessagesToSend
// }

// func main (){
//     cost := getMaxMessagesToSend(5, 250)
//     fmt.Println(cost)
// }

// Output: 4


// package main 

// import "fmt"

// func sum(nums ...int) []int {
//     for i := 0; i < len(nums); i++ {
//         fmt.Println(nums[i])
//     }
//     return nums
// }


// func main() {
//     total := sum(1, 2, 3)
//     fmt.Println(total)
// }

// Output: 
// 1
// 2
// 3
// [1 2 3]

package main 

import "fmt"

func printStrings(strings ...string) {
    for i := 0; i < len(strings); i++ {
        fmt.Println(strings[i])
    }
}

func main() {
    names := [] string {"Ashith", "Anikethan", "Deekshitha"}
    printStrings(names...)
}