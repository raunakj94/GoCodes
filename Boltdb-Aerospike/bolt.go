package main

import (
       "log"
       "github.com/boltdb/bolt"
        "fmt"
)
//testSuite:=make([]string,20,20)
	var testSuite = [] string{"abcdefgh01","abcdefgh02","abcdefgh03","abcdefgh04","abcdefgh05","abcdefgh06","abcdefgh07","abcdefgh08","abcdefgh09","abcdefgh10",
"abcdefgh11","abcdefgh12","abcdefgh13","abcdefgh14","abcdefgh15","abcdefgh16","abcdefgh17","abcdefgh18","abcdefgh19","abcdefgh20"}
//string slice of values i.e.timestamp
	var time_stamp = [] string{"1433834829474","1433834829475","1433834829476","1433834829477","1433834829478",
"1433834829479","1433834829480","1433834829481","1433834829482","1433834829483","1433834829484",
"1433834829485","1433834829486","1433834829487","1433834829489","1433834829490","1433834829491","1433834829492",
"1433834829493","1433834829494","1433834829495",}

//uint slice of keys i.e. bids 
    var bids = [] uint8{20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1}

//Function to Calculate a Current_spend

func Calculate_Current_spend(current_spend []uint8) uint8 {
	var current_spent uint8
	for u:=0 ; u<len(current_spend) ; u++ {
		current_spent=current_spent+current_spend[u]
	}

	//UpdateCurrentSpend(c_s,q)
	return current_spent
}


func Update_Current_Spend(currentSpend uint8,subline_key int){
	var temp_spend []uint8
	temp_spend=append(temp_spend,currentSpend)
	

	var temp string
	temp=testSuite[subline_key]
 	db, err :=bolt.Open("time.db",0666,nil)
	if err!=nil{
		log.Fatal(err)
       	}

//Each Update() waits for disk to commit the writes. This overhead can be minimized by combining multiple updates with the Batch operation.There will multiple goroutines calling Batch.
/*Each Update() waits for disk to commit the writes. This overhead can be minimized by 
combining multiple updates with the Batch operation.
There will multiple goroutines calling Batch.*/
	
	//Creating  a new bucket for storing the subline_id and current_spend
	
	db.Update(func(tx *bolt.Tx) error {
   			 _, err := tx.CreateBucket([]byte("ADSERVER"))
    	
		if err != nil {
        			return fmt.Errorf("create bucket: %s", err)
    				}
   	 	return nil
})
 
     db.Update(func(tx *bolt.Tx) error {
 			   bucket := tx.Bucket([]byte("ADSERVER"))
   			 err := bucket.Put([]byte(temp), []byte(temp_spend[0:1]))
   			 return err
	})
		
	db.View(func(tx *bolt.Tx) error {
   				 bucket := tx.Bucket([]byte("ADSERVER"))
   				 temp_value := bucket.Get([]byte(temp))
   				 fmt.Printf("For subline_id %s current_spend is: %d\n",temp,temp_value)
   				 return nil
	})

	db.Close()
}

	func main(){
//Database Access Codes

  db, err :=bolt.Open("time.db",0666,nil)
	if err!=nil{
	log.Fatal(err)
       }              

  subline_count:=0
 for subline_count<20{
     db.Update(func(tx *bolt.Tx) error {
   					 _, err := tx.CreateBucketIfNotExists([]byte(testSuite[subline_count]))
   						 if err != nil {
    								    return err
   										 }
	return nil
	})
subline_count++
	}
bid_count:=0
	for bid_count<20{
		iterator:=1

		db.Update(func(tx *bolt.Tx) error {
	   		 bucket := tx.Bucket([]byte(testSuite[bid_count]))
	   			 for iterator<=20{
						var temp1 string
						temp1=time_stamp[iterator]
	     				bucket.Put([]byte(temp1),bids[iterator-1:iterator])
	                		iterator++
					}
    return err
	})
bid_count++
}

db.Close()

display_count:=0
for display_count<20{
	db, err :=bolt.Open("time.db",0666,nil)
		if err!=nil{
		log.Fatal(err)
       		}
         temp_array:=make([]uint8,20)
	 db.View(func(tx *bolt.Tx) error {
     	 j:=tx.Bucket([]byte(testSuite[display_count]))
      	c := j.Cursor()
    	for k, val := c.First(); k != nil; k, val = c.Next() {
        	 						// fmt.Println(k, val)
								
                                                                  temp_array=append(temp_array,val[0])
}
return nil
})
var get_Current_Spend uint8
get_Current_Spend=Calculate_Current_spend(temp_array)
db.Close()
 Update_Current_Spend(get_Current_Spend,display_count)
	display_count++
       }
}
