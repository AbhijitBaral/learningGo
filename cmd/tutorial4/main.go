package main;

import "fmt";

func main(){
  var intArr[5] int32=[5]int32{0,1,2,3,4};
  fmt.Println(intArr[0]);
  //fmt.Println(intArr[2:3]);

  anArray:=[...]int32{1,2,4,1,42,4};
  fmt.Println(anArray);

  var intSlice []int32=[]int32{2,4,1};
  fmt.Println(intSlice);
  intSlice=append(intSlice,2);
  fmt.Println(intSlice);

  //maps are key-value pairs
  var myMap map[string]uint8=make(map[string]uint8);
  fmt.Println(myMap);

  var myMap2 = map[string]uint8{"Adam":23, "Sarah":32};
  fmt.Println(myMap2["Adam"]);
  var age, ok=myMap2["Sarah"];
  if (ok){
    fmt.Printf("The age is %v", age);
  }else{
    fmt.Println("Invalid Name");
  }

}
