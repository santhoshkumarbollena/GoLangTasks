package main
import ("fmt"
		)
const usixteenbitmax float64=65535
const kmhmultiple float64=1.60934
func main(){
	acar :=car{gas_pedal:22341,
		break_pedal :0,
		steering_wheel:12561,
		top_spead_kmh :225.0 }
	fmt.Println(acar.gas_pedal)
	fmt.Println(acar.kmp())
	fmt.Println(acar.mph)
}
func (c car) kmp()float64{
	return float64(c.gas_pedal)*(c.top_spead_kmh/usixteenbitmax)
}
func (c car) mph()float64{
	return float64(c.gas_pedal)*(c.top_spead_kmh/usixteenbitmax/kmhmultiple)
}
type car struct{
	gas_pedal uint16 //min 0 max 65535
	break_pedal uint16 
	steering_wheel int16
	top_spead_kmh float64
}
