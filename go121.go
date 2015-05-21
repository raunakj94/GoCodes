package main
import "fmt" 

func main(){
sum:=0
i:=0
 A :=[5]int{1,2,3,4,5}
for i<5{
        sum=sum+A[i]
	i=i+1
      }
avg:=sum/5
fmt.Println(sum)
fmt.Println(avg)
for k:=0 ; k<20; k++{
	              if k%2==0{
                                 fmt.Println("Even")
} else{ fmt.Println("Odd")}


}


