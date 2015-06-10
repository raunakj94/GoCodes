package bolt
import (
	"testing"
	"log"
 	"math/rand"
    	"time"

       "github.com/boltdb/bolt"
)

var i = []uint8{1,2,6,3,4,5,6,7,8,9,2,3,4,34,6,7,42}

func benchmarkCalc(i []uint8, b *testing.B) {
        for n := 0; n < b.N; n++ {
                Calculate_Current_spend(i)
        }
}


func BenchmarkCalc1(b *testing.B){
j := i[0:2]
benchmarkCalc(j, b) 
}

func BenchmarkCalc2(b *testing.B){ 
j := i[0:5]
 benchmarkCalc(j, b) }

func BenchmarkCalc3(b *testing.B) {
j := i[0:5] 
benchmarkCalc(j, b) }

func BenchmarkCalc4(b *testing.B) {
j := i[0:15]
 benchmarkCalc(j, b) }

func BenchmarkCalc5(b *testing.B) {
j := i[0:17] 
benchmarkCalc(j, b) }




func benchmarkUpdate(q uint8,w int,b *testing.B) {
  db, err :=bolt.Open("abc.db",0666,nil)
		if err!=nil{
		log.Fatal(err)
  	     }

	 	            
        for n := 0; n < b.N; n++ {
                Update_Current_Spend(db,q,w)
        }
db.Close()
}
func BenchmarkUpdate1(b *testing.B){

r1:=random(1, 20)
benchmarkUpdate(16,r1,b)
}


func BenchmarkUpdate2(b *testing.B){

r2:=random(1, 20)
benchmarkUpdate(15,r2,b)
}

func BenchmarkUpdate3(b *testing.B){

r3 := random(1, 20)
benchmarkUpdate(10,r3,b)
}
func BenchmarkUpdate4(b *testing.B){

r4 := random(1, 20)
benchmarkUpdate(6,r4,b)
}

func BenchmarkUpdate5(b *testing.B){

r5 := random(1, 20)
benchmarkUpdate(4,r5,b)
}

func BenchmarkUpdate6(b *testing.B){

r6 := random(1, 20)
benchmarkUpdate(7,r6,b)
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}










