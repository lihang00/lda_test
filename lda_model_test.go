package lda

import (
    "testing"
    "fmt"
    //"math"
)

func TestLDAModel(t *testing.T) {
    param := LDAParams{}
    param.Alpha = 0.1
    param.Beta = 0.1
    param.Ntopics = 10
    model := InitModel(&param)

    model_string := "a\t0\t5\na\t1\t2\nb\t2\t3\nc\t4\t5\nc\t0\t2\n"

    word_id := model.word_dic.AddWord("a")
    model.AddNWZ(word_id, int64(0), int64(5))
    model.AddNWZ(word_id, int64(1), int64(2))
    word_id = model.word_dic.AddWord("b")
    model.AddNWZ(word_id, int64(2), int64(3))
    word_id = model.word_dic.AddWord("c")
    model.AddNWZ(word_id, int64(4), int64(5))
    model.AddNWZ(word_id, int64(0), int64(2))

    if model_string != model.SaveModel("") {
        fmt.Println(model.SaveModel(""))
        t.Error("Model dump file error.")
    }
    
}