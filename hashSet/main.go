package main;

import "fmt";

func main(){
  //Creation of a hashset
  hashSet:=make(map[string]struct{});

  //Addition of elements to the set
  hashSet["apple"]=struct{}{};
  hashSet["banana"]=struct{}{};
  hashSet["cherry"]=struct{}{};

  //check if an element exists
  if_,exists:=hashSet["banana"];exists{
    fmt.Println("banana exists in the set");
  }
  else{
    fmt.Println("banana doesn't exist in the set");
  }

  //Delete an element from the set
  delete(hashSet, "banana");

  //check again after deletion
  if _,exists:=hashSet["banana"];exists{
    fmt.Println("banana exists in the set");
  }
  else{
    fmt.Println("banana does not exist int the set");
  }

  //Iterate over the set
  for item := range hasSet{
    fmt.Println(item);
  }
}
