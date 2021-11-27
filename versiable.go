packeage server
import "fmt"

Max_bandwith=100000 //最大带宽

func Set_Max_bandwith(Max int) int{
  Max_bandwith=Max
}

func Get_Max_bandwith() int{
  return Max_bandwith
}
