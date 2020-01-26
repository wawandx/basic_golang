package main
 
import (
	"fmt"
)
 
func main() {
	var name, course string
	var score, studentID int
	fmt.Println("Input Your Student ID : ")
	fmt.Scan(&studentID)
	fmt.Println("Input Your Name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Input Your Course: ")
	fmt.Scanf("%s", &course)
	fmt.Println("Input Your Score: ")
	fmt.Scanln(&score)
	
	fmt.Printf("Name (StudentID) : %s (%d)\n", name, studentID)
	fmt.Printf("Course (Score) : %s (%d)\n", course, score)
}