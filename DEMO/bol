package newton
import (
	"testing"
	"log"
 	"math/rand"
    	"time"

       "github.com/boltdb/bolt"
)


func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}


func benchmarkUpdate(q int,w uint64,b *testing.B) {
  db, err :=bolt.Open("f.db",0666,nil)
		if err!=nil{
		log.Fatal(err)
  	     }

	 	            
        for n := 0; n < b.N; n++ {
                Update_Current_spend(db,q,w)
        }
db.Close()
}


	func BenchmarkUpdateRandomBid1(b *testing.B){
					r1:=random(1, 20)
					benchmarkUpdate(r1,2,b)
				}


	func BenchmarkUpdateRandomBid2(b *testing.B){
					r2:=random(1, 20)
					benchmarkUpdate(r2,21,b)
				}

	func BenchmarkUpdateRandomBid3(b *testing.B){
					r3:=random(1, 20)
					benchmarkUpdate(r3,22,b)
				}

	func BenchmarkUpdateRandomBid4(b *testing.B){
					r4:=random(1, 20)
					benchmarkUpdate(r4,254,b)
				}

	func BenchmarkUpdateRandomBid5(b *testing.B){
					r5:=random(1, 20)
					benchmarkUpdate(r5,122,b)
				}
