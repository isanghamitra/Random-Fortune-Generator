package main 

import "fmt"
// Import the "math/rand" and "time" packages
import "math/rand"
import "time"

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().Unix())
    
    // Add a for loop
    for i:=1; i<=3; i++ {
        fmt.Printf("Fortune number %d: ",i)
        number := rand.Intn(10)
        switch number {
            case 0:
                fmt.Println("Change the world by being yourself. – Amy Poehler")
            case 1:
                fmt.Println("Every moment is a fresh beginning. – T.S Eliot")
            case 2:
                fmt.Println("Aspire to inspire before we expire. – Unknown")
            case 3:
                fmt.Println("What we think, we become. – Buddha")
            case 4:
                fmt.Println("Whatever you do, do it well. – Walt Disney")
            case 5:
                fmt.Println("Dream as if you’ll live forever, live as if you’ll die today. – James Dean")
            case 6:
                fmt.Println("Yesterday you said tomorrow. Just do it. – Nike")
            case 7:
                fmt.Println("Happiness depends upon ourselves. – Aristotle")
            case 8:
                fmt.Println("If the world was blind how many people would you impress? – Boonaa Mohammed")
            case 9:
                fmt.Println("I have nothing to lose but something to gain. – Eminem")
            case 10:
                fmt.Println("When words fail, music speaks. – Shakespeare")
        }
    }
}
