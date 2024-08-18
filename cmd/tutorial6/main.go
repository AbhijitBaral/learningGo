package main;
import "fmt";

type gasEngine struct{
  mpg uint8;      //Field 1
  gallon uint8;   //Field 2
  ownerInfo owner;//FIeld 3
}

/* The same struct can be defined with only types instead of variables in the fields
type gasEngine struct{
  uint8;      //Field 1
  uint8;      //Field 2
  owner;      //FIeld 3
}
*/

type owner struct{
  name string;
}


func (e gasEngine) milesLeft() uint8{
  return (e.gallon*e.mpg);
}

func main(){
  var Owner owner=owner{"Owner name"};
  
  var Engine gasEngine= gasEngine{25, 21, Owner};
  /*
  //One way to to field assignement
  var myEngine1 gasEngine = gasEngine{mpg:23, gallon:32, name:NULL };

  //Another way to assign values to the fields
  var myEngine2 gasEngine = gasEngine{23,32}  //The values are assigned in order of their description in the struct

  //Another way
  var myEngine3 gasEngine;
  myEngine3.mpg=23;
  myEngine3.gallon=32;
  */

  func (e gasEngine) milesLeft() uint8{
    return e.gallons*e.mpg;
  }

  func canMakeIt(e gasEngine, Mile uint8){
    if(miles<=e.milesLeft()){
      fmt.Println("You can make it there !");
    }
    else{
      fmt.Println("need to fuel up first! ");
    }
  }

  //I don't understand interfaces in go...
}
