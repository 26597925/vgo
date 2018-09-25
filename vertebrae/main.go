package main

import(
"fmt"
)

type MyApp struct{}

func (myApp *MyApp) TestService(param int) int {
	return 12 + param
}

func main() {
	myApp := &MyApp{}
	container := new(Vertebrae)
	
	container.Add("my.service", myApp)
	container.Add("my.string", "Testing 123")
	
	myService, ok := container.Get("my.service")
	
	if !ok {
		panic("Service 'my.service' not found!")    
	}
	
	myString, ok := container.Get("my.string")
	
	if !ok {
		panic("Service 'my.string' not found!")    
	}
	
	fmt.Println(myService.(*MyApp).TestService(23))
	fmt.Println(myString.(string))
}