package main

import(
  "bufio"
  "os"
  "fmt"
  "strconv"
  "strings"
)

type ba struct{
  Balance int
  Payments[] string
  Name string


}
func login()*ba{
  for{
    log_in:=bufio.NewReader(os.Stdin)
    fmt.Println("Please Log In With Your Name: or Control C to exit ")
    id,err:=log_in.ReadString('\n')

    if err != nil{
      fmt.Println("There was an error, ", err)
      continue
    }
    if id=="\n" {
      continue
    }
    newName:= &ba{Name:id}
    fmt.Println(newName.Name)
    return newName
  }

}
func deposit(bank *ba, money int) int{
  bank.Balance +=money
  return bank.Balance
}

func withdrawal(bank *ba, money int)int{
  bank.Balance-= money
  return bank.Balance
}

func add_subtract(object *ba, addSubtract string){
  add:=bufio.NewReader(os.Stdin)
  fmt.Println("How much money should we ", addSubtract)
  depo,err:=add.ReadString('\n')
  if err != nil{
    fmt.Println("There was an error, ", err)
  }
  // Have to get rid of new line char from user input.
  depo = strings.Replace(depo,"\n","", -1)
  new_depo,err:= strconv.Atoi(depo)
  fmt.Println("Money",addSubtract, "ed: ", new_depo)
  if addSubtract =="add"{
    fmt.Println("Total Money in Account: ", deposit(object,new_depo))
  }else{
    fmt.Println("Money left in account: ", withdrawal(object,new_depo))
    }

}

func main(){
    obj:=login()
    fmt.Println("Welcome to Bank of America")
    for{
      choices:=bufio.NewReader(os.Stdin)
      fmt.Println("What would you like to: \n", "1. Make a deposit\n", "2. Withdrawal Cash\n", "3. View Account Info\n", "4. Pay a Bill\n", "5. Dispute a Transaction\n", "6. Logout\n")
      decision, err:=choices.ReadString('\n')
        if err != nil{
      fmt.Println("There was an error, ", err)
      continue
        }

      // switch statement, cases correspond to the previous print.
      switch decision{
      default:
        os.Exit(0)

      case "1\n":
        adder:="add"
        add_subtract(obj,adder)
      case "2\n":
        subtracter:="subtract"
        add_subtract(obj,subtracter)

      }
    }

}
