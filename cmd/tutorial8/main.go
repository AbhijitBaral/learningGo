//~~~~~~~~~~~~~~~~~~~~Go Routines~~~~~~~~~~~~~~~~~~~~~~~~~

package main;

import(
  "fmt"
  "math/rand"
  "time"
  "sync"
);

var wg = sync.WaitGroup{};
var dbData=[]string{"id1","id2","id3","id4","id5"};
var results = []string{};

func main(){
  t0:=time.Now();
  for(i:=0; i<len(dbData); i++){
    wg.Add(1);
    go dbCall(i);
  }

  wg.Wait();
  fmt.Printf("\nTotal execution time: %v", time.Since(t0));
  fmt.Printf("\nThe results are %v", results);
}

func dbCall(i int){
  //Simulate dbCall delay
  var delay float32 = 2000;
  time.Sleep(time.Duration(delay)*time.Millisecond);
  fmt.Println("The result from the dataBase is: %s",dbData[i]);
  results = append(results, dbData[i]);
  wg.Done();

  // I don't understand go routines
}
