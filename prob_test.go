package lda

import (
    "testing"
    "fmt"
    "math"
)


func TestProb(t *testing.T) {
    pr := NewProb()

    // set probability
    pr.AddValue(1,20)
    pr.AddValue(2,30)
    pr.AddValue(3,50)
    
    // count sample
    precision := 1e-2

    counter := NewProb()
    for i := 0; i < 10000; i++{
        sample_key := pr.Sample()
        counter.AddValue(sample_key, float64(1))        
    }

    for k, _ := range pr.Vector.Data {
        if math.Abs(pr.GetPr(k) - counter.GetPr(k)) > precision {
            fmt.Println(k)
            fmt.Println(pr.GetPr(k))
            fmt.Println(counter.GetPr(k))    
            t.Error("Sample error.")
        }
    }

}