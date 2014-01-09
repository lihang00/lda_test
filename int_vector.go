package lda

type IntVector struct {
    data map[int64]int64
}

func NewIntVector() *IntVector {
    v := IntVector{}
    v.data = make(map[int64]int64)
    return &v
}

func (v *IntVector) AddValue(key int64, value int64) {
    _, ok := v.data[key]
    if ok{
        v.data[key] += value
    } else {
        v.data[key] = value
    }
}

func (v *IntVector) GetValue(key int64) int64{
    value, ok := v.data[key]
    if !ok {
        return 0
    } else {
        return value
    }
}
func (v *IntVector) SetValue(key int64, value int64) {
    v.data[key] = value
}

func (v *IntVector) AddVector(v2 *IntVector, alpha int64) {
    for key, value := range v2.data {
        if alpha != 1 {
            v.AddValue(key, value * alpha)
        } else {
            v.AddValue(key, value)
        }

    }
}
