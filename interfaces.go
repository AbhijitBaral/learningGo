//~~~~~~~~~~~~~~Interfaces in GoLang~~~~~~~~~~~~~~~~~

package main;

import "fmt";

//Struc "Rectangle"
type rectangle struct{
  length float32;
  width float32;
}
//method for rectangle
func (r rectangle)area() float64{
  return r.width * r.length;
}
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

//Struct "Circle"
type circle struct{
  radius float32; 
}
//method for type circle
func (c circle)area() float64{
  return Math.Pi * c.radius * c.radius;
}

//Interface
type Shape interface{
  area() float64;
}
