package lda

import (
    "hector/core"
    "math/rand"
)

type Prod struct {
    Vector *core.Vector
    Sum float64
}

func NewProd() *Prod {
    pr := Prod{}
    pr.Vector = core.NewVector()
    pr.Sum = 0
    return &pr
}

func (pr *Prod) AddValue(key int64, value float64) {
    pr.Vector.AddValue(key, value)
    pr.Sum += value;
}

func (pr *Prod) Sample() int64 {
    ss := pr.Sum * rand.Float64()

    var ff int64
    for ff := int64(0) ; ff < int64(len(pr.Vector.Data)) ; ff++ {
        ss -= pr.Vector.GetValue(ff)
        if ss < 0 {
            break
        }
    }

    return ff
}