package main;

import "fmt";

func main(){
  /*var x string = "Hello World";
  printMe(x);*/
  var quotient,remainder int=intDivision(12,3);
  fmt.Printf("The quotient of the division is: %v and remainder is: %v",quotient,remainder);
}

func printMe(printVale string){
  fmt.Println(printVale);
}

func intDivision(numerator int, denominator int)(int,int){
  var quotient int=numerator/denominator;
  var remainder int=numerator%denominator;
  return quotient, remainder;
}
