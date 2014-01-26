package lda

import (
    "hector/core"
    "math/rand"
)

type Prob struct {
    Vector *core.Vector
    Sum float64
    Keys []int64
}

func NewProb() *Prob {
    pr := Prob{}
    pr.Vector = core.NewVector()
    pr.Sum = 0
    pr.Keys = []int64{}
    return &pr
}

func (pr *Prob) AddValue(key int64, value float64) {
    if _,ok := pr.Vector.Data[key]; !ok {
        pr.Keys = append(pr.Keys, key)
    }
    pr.Vector.AddValue(key, value)
    pr.Sum += value;
}

func (pr *Prob) GetPr(key int64) float64{
    value, ok := pr.Vector.Data[key]
    if !ok {
        return 0.0
    } else {
        return value / pr.Sum
    }
}

func (pr *Prob) Sample() int64 {
    ss := pr.Sum * rand.Float64()

    var ff int64
    for _, ff = range pr.Keys {
        ss -= pr.Vector.GetValue(ff)
        if ss < 0 {
            break
        }
    }

    return ff
}

func (pr *Prob) GetMaxPrIndex() int64 {
    ff := int64(0)
    maxPr := float64(0)

    for _, idx := range pr.Keys {
        value := pr.Vector.GetValue(idx)
        if  value > maxPr {
            ff = idx
            maxPr = value
        }
    }

    return ff
}