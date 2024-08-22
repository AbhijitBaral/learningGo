//~~~~~~~~~~~~STRUCTS in GoLang~~~~~~~~~~~~~~~~~~~~

package main;

import "fmt";

//Struct is defined with the "type" keyword
//Here "Person" is a struct with three fields
type Person struct{
  name string;
  age int;
  gender string;
};

//Below is a method that operates on the type Person.
/*Methods provide a more organized, modular, and idiomatic way to structure your code,
  making it easier to manage and scale. */
func(p Person) Greet(){
  fmt.Printf("Hello, my name is %s and I am %d years old.\n",p.name, p.age);
}

func greet(p Person){
  fmt.Printf("Hello, my name is %s amd I am %d years old.\n",p.name, p.age);
}

func main(){
  var person1 Person = Person{name:"Steve", age:32, gender:"male"};
  /*
  Shorthand: 
    person1 := Person{name:"Steve", age:"32", gender:"male"};
  */
  
  //Alternatively : 
  var person1Alt Person;
  person1Alt.name = "Steve";
  person1Alt.age = 32;
  person1Alt.gender = "male";

  fmt.Println(person1);
  fmt.Println(person1Alt);

  //Anonymous structs
  student := struct{
    name string
    age int
    gender string
  }{
    name: "Daisy",
    age: 21,
    gender: "female",
  }
  fmt.Println(student);

  //Pointers to structs
  studentPointer := &student;
  fmt.Println(*studentPointer);
  fmt.Println(studentPointer);

  //Structs with methods
  /*A method is a function that has a special receiver that allows it to
    access the fields of a structs  */

  person := Person{name: "Frank", age: 32, gender: "male"};
  person.Greet();
  
  greet(person);
}
