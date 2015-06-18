
package newton

import (
       "log"
       	"github.com/boltdb/bolt"
        "fmt"
	"encoding/binary"
)

func Display(db *bolt.DB){
	db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("ADSERVER"))
    c := b.Cursor()

    for k, v := c.First(); k != nil; k, v = c.Next() {
                n:=decode(v)
        fmt.Printf("key=%s, value=%d\n", k, n)
    }

    return nil
})
}

//testSuite:=make([]string,20,20)

 var testSuite = [] string
 {"abcdefgh01","abcdefgh02","abcdefgh03","abcdefgh04","abcdefgh05","abcdefgh06",
 "abcdefgh07","abcdefgh08","abcdefgh09","abcdefgh10","abcdefgh11","abcdefgh12",
 "abcdefgh13","abcdefgh14","abcdefgh15","abcdefgh16","abcdefgh17","abcdefgh18","abcdefgh19","abcdefgh20"}


func encode(n uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	return buf
}

func decode(buf []byte) uint64 {
	if buf == nil {
		return 0
	}
	return binary.BigEndian.Uint64(buf)
}


//Function to Calculate a Current_spend


	func Update_Current_spend(db *bolt.DB,subline_key int,current_bid uint64,){
          var bid uint64
			db.View(func(tx *bolt.Tx) error {
							var temp string
							temp=testSuite[subline_key]
  					  		b := tx.Bucket([]byte("ADSERVER"))
        						bid = decode(b.Get([]byte(temp)))
							   return nil
})
			db.Update(func(tx *bolt.Tx) error {
							var temp string
							temp=testSuite[subline_key]
 						   bucket := tx.Bucket([]byte("ADSERVER"))
							c_bid:=current_bid+bid
   						 err := bucket.Put([]byte(temp), encode(c_bid))
   						 return err
					})
			

}

  func main(){
	//Database Access Codes

	  	db, err :=bolt.Open("f.db",0666,nil)
			if err!=nil{
			log.Fatal(err)
	  	     }              
		defer db.Close()
	
db.Update(func(tx *bolt.Tx) error {
   			 _, err := tx.CreateBucket([]byte("ADSERVER"))
    	
		if err != nil {
        			return fmt.Errorf("create bucket: %s", err)
    				}
   	 	return nil
})
 

		db.Update(func(tx *bolt.Tx) error {
	   		 bucket := tx.Bucket([]byte("ADSERVER"))
			count:=0
				for count<20{
						var temp string
							temp=testSuite[count]
	     				bucket.Put([]byte(temp),encode(0))
	                		count++
					}
    return err
	})

Display(db)
}
