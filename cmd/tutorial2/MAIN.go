package main;

import "fmt";
import "unicode/utf8";

func Main(){
  var intNum int=32767;
  intNum=intNum+1;
  fmt.Println(intNum);

  var floatNum float64=12345678.9;
  fmt.Println(floatNum);

  var floatNum32 float32=10.1;
  var intNum32 int32=2;

  //var result float32=floatNum32+intNum32; Invalid operation
    var result float32=floatNum32+float32(intNum32);
  fmt.Println(result);
  
  var intNum1 int=3;
  var intNum2 int=2;
  fmt.Println(intNum1/intNum2);
  fmt.Println(intNum1%intNum2);

  var myString string="Hello"+" "+"World";
  //The declaration formats below are not recommended
  /*var myString1="The variable type is inferrred";
  myString2 :="Shorthand declaration";
  */
  
  var1,var2:=1,2;
  fmt.Println(var1, var2);

  fmt.Println(myString);

  fmt.Println(utf8.RuneCountInString(myString));

  //const types cannot be only declared. They need to be initialized explicitly.
  const myConst string="constant string"; 
  fmt.Println(myConst);

}
