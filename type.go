package main
import (
	"fmt"
)

type Students interface {
	Greet()
	LearnMore()
}

type Student struct{
	name string
	age int
}

func (s Student ) Greet() {
	fmt.Printf("Hi, I am %s\n", s.name)
}
//A Whiptail method
func (s Student) LearnMore() {
	fmt.Printf("[Name: %s, Age: %d]\n", s.name,s.age)
}

type Whiptail struct {
	Student
	//embedded type

}


type CoolPerson struct {
	Whiptail //type embedding for composition
	coolBecause []string
}
//overrides LearnMore
func (s CoolPerson) LearnMore() {
	//Call Whiptail LearnMore
	s.Whiptail.LearnMore()
	fmt.Println("CoolPerson is cool because:")
	for _, value := range s.coolBecause {
		fmt.Println(value)
	}
}

type Boss struct {
	Whiptail //type embedding for composition
	bossySayings []string
}
//overrides LearnMore
func (b Boss) LearnMore() {
	//Call Whiptail LearnMore
	b.Whiptail.LearnMore()
	fmt.Println("Boss says these things:")
	for _, value := range b.bossySayings {
		fmt.Println(value)
	}
}

type UltraBoss struct {
	Student //type embedding for composition
	bossySayings []string
}
//overrides LearnMore
func (u UltraBoss) LearnMore() {
	//Call Whiptail LearnMore
	u.Student.LearnMore()
	fmt.Println("UltraBoss says these things:")
	for _, value := range u.bossySayings {
		fmt.Println(value)
	}
}

type Presentation struct
{
	location string
	s []Students
}
func (m Presentation) PresentationPeople(){
	for _, v := range m.s {
		v.Greet()
		v.LearnMore()
	}
}

func main() {
	karina := CoolPerson{Whiptail{Student{"Karina", 20}},
		[]string{"Learned how to bike","Eats ramen", "Plays frisbee"}}
	vivienne := Boss{Whiptail{Student{"Vivienne", 34}},
		[]string{"Ummm","Yikes."}}
	emery := Boss{Whiptail{Student{"Emery", 15}},
		[]string{"Go do it","This is a goroutine."}}
		ben := UltraBoss{Student{"Ben", 100},
			[]string{"yay PLs","final projects are fun"}}
		team:=Presentation {
				"Ben's Office",
				[]Students{karina, emery, vivienne, ben},
			}


			team.PresentationPeople()
}
