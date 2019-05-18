package main
import (
    "fmt"
    "reflect"
)

func main() {
    x := struct{Foo string 
         Bar int 
         i int}{"foo", 2,9}

    v := reflect.ValueOf(x)

    values := make([]interface{}, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        values[i] = v.Field(i).Interface()
    }

    fmt.Println(values)
}
// package main

// import ("fmt"
// 		"strconv"
// 		"strings"
// 		)
// type date struct{
// 	day int64
// 	month int64
// 	year int64
// }

// type member struct {
// 			name string
// 			enrollmentStartDate date
// 			enrollmentEndDate date
// 			code string
// 			year  int64
// 			next *member
// }

// func addNodeEnd(newPerson, personList *member) *member {
// 	if personList == nil {
// 	  return personList
// 	}
// 	for p := personList; p != nil; p = p.next {
// 	  if p.next == nil {
// 		p.next = newPerson
// 		return personList
// 	  }
// 	 }
// 	return personList
//   }
// func Logic(personList *member)*member{
// 	var resultList *member
// 	//fmt.Println("Enter date int day-month-year format: ")
// 	var inputy string="2-2-1015"
// 	//var resultList *member
// 	//fmt.Scanln(&inputy)
// 	s := strings.Split(inputy, "-")
// 	day , err:=strconv.ParseInt(s[0],10,64)
// 	month , err:=strconv.ParseInt(s[1],10,64)
// 	year , err:=strconv.ParseInt(s[2],10,64)
// 	if err == nil {	}
// 	for p := personList; p != nil; p = p.next {
// 	//fmt.Println(p)
// 	//fmt.Print(p.enrollmentEndDate.day);fmt.Print(day);fmt.Println(p.enrollmentStartDate.day);
// 	//fmt.Println(month)
// 	//fmt.Println(year)
// 	if(year<p.enrollmentEndDate.year&&year>p.enrollmentStartDate.year){
// 			// result[k]=mems[i];
// 			// k++;
// 		//fmt.Println(mems[i].year)
// 		fmt.Println(p)
// 		resultList = p
// 		resultList = addNodeEnd(p, resultList)
// 	 }
// 	if(year==p.enrollmentEndDate.year&&year>p.enrollmentStartDate.year){
// 		if(month<p.enrollmentEndDate.month&&month>p.enrollmentStartDate.month){
// 			//fmt.Println(p)
// 			resultList := p
// 			resultList = addNodeEnd(p, resultList)
// 		}
// 		if(month==p.enrollmentEndDate.month&&month>p.enrollmentStartDate.month){
// 			if(day<p.enrollmentEndDate.day&&day>p.enrollmentStartDate.day){
// 			//fmt.Println(p)
// 			resultList := p
// 			resultList = addNodeEnd(p, resultList)
// 			}
// 		}
// 	 }
	 
// 	}
// 	return resultList
// 	}
// func printList(personList *member) {
// 		for p := personList; p != nil; p = p.next {
// 		 fmt.Println(p)
// 		}
// 	  }
// func main(){
// 	var memberList *member
// 	var demores *member
// 	s := &member{name : "santhosh",enrollmentStartDate : date{day :2,month :4,year :2015},enrollmentEndDate : date{day :1,month :1,year :2019},code : "b15cs067",year : 2019,next :nil}
// 	memberList = s
// 	p := &member{name : "santhosh kumar",enrollmentStartDate : date{day :2,month :4, year:1000},enrollmentEndDate : date{day :1,month :1,year :1050},code : "b15cs067",year : 1000,next :nil}
// 	memberList = addNodeEnd(p, memberList)
// 	z := &member{name : "santhosh bollena",enrollmentStartDate : date{day :2,month :4, year:1010},enrollmentEndDate : date{day :1,month :1,year :1020},code : "b15cs067",year : 1000,next :nil}
// 	memberList = addNodeEnd(z, memberList)
// 	demores = Logic(memberList)
// 	printList(demores)
// 	//fmt.Println(memberList);
// 	//var  members = []member {s,p}
// 	// members := make(chan member)
// 	// members <- s
// 	// members <- p
// 	//fmt.Println(members)
// 	//demo :=logic(members)
// 	//fmt.Println(demo);
// 	// var i int
// 	// for i = 0; i < len(demo);i++ {
// 	// 	fmt.Println(demo[i])
// 	// }

// 	//fmt.Println(p);
// }
// // func logic(mems []Channel)[numberOdMems]member{
// // 	fmt.Print("Enter year: ")
// //     var inputy string
// //     fmt.Scanln(&inputy)
// // 	var i int
// // 	for i = 0; i < len(mems);i++ {
// // 		inp , err:=strconv.ParseInt(inputy,10,64)
// // 		if err == nil {
// // 			//fmt.Println(inp)
// // 		}
// // 		if(inp<mems[i].year){
// // 			result[k]=mems[i];
// // 			k++;
// // 		//fmt.Println(mems[i].year)
// // 	 }
// // 	}
// // 	//fmt.Println(result)
// // 	return result
// // }