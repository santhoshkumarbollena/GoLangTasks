package main
import ("fmt"
		"net/http")
func index_handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"its easy")
}
func about_handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"its me")
}
func pointers(){
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about",about_handler)
	http.ListenAndServe(":9090",nil)

	// acar :=car{2,25}
	// fmt.Println(acar.speed)
}
// type car struct{
// 	engine int
// 	speed float64
// }
func pointer(){
	x:=15
	a:=&x //Memeory address
	fmt.Println(x,a,*a)
	*a=5
	fmt.Println(x,a,*a)
}