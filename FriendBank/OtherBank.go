package FriendBank

import (
  "math/rand"
  "time"
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)
type Julian struct{
  Balance int
}
// Friend transfer money
func OtherBank() int{
  bank:=Make_obj()
  rand.Seed(time.Now().UTC().UnixNano())
  bank.Balance += rand.Intn(10000)
  fmt.Println("Your friend has ", bank.Balance, "in his/her account")
  money_out:=bufio.NewReader(os.Stdin)
  fmt.Println("How much is your friend sending you? We will trust your judgement :)")
  direct_depo,err:=money_out.ReadString('\n')
  if err!=nil{
    fmt.Println("There was an error at first err...")
    panic(err)
  }
  direct_depo= strings.Replace(direct_depo,"\n","", -1)
  intDeposit, err:=strconv.Atoi(direct_depo)
  if err!=nil{
  fmt.Println("There was an error at second err...")
  panic(err)
  }
  if (intDeposit> bank.Balance) {
    fmt.Println("Error you get nothing.")
    intDeposit =0
    return intDeposit
  }

  return intDeposit

}
func Make_obj()*Julian{
  friend:=&Julian{Balance:0}
  return friend
}
