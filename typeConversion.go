package main;

import "fmt";

func main(){
  var temperatureFloat float32 = 44.432;
  //Explicit type conversion of 'temperatureInt' from integer to float
  var temperatureInt int32 = int32(temperatureFloat);
  fmt.Printf("temperatureFloat: %f, temperatureInt: %v\n",temperatureFloat,temperatureInt);

  temperatureInt=0;
  temperatureInt=temperatureFloat;
  fmt.Printf("temperatureFloat: %f, temperatureInt: %v\n",temperatureFloat,temperatureInt);
}
