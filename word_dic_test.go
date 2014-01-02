package lda

import (
    "testing"
    "fmt"
)

func TestWordDic(t *testing.T) {
    dic := NewWordDic()

    word1 := "abc"
    word2 := "def"
    word3 := "gh"
    
    id1 := dic.AddWord(word1)
    id2 := dic.AddWord(word2)
    id3 := dic.AddWord(word3)

    if dic.AddWord(word1) != id1 {
        fmt.Println(id1)
        fmt.Println(dic.AddWord(word1))
        t.Error("Add word wrong.")
    }

    if dic.GetId(word2) != id2 {
        t.Error("Get id wrong")
    }

    if dic.GetWord(id3) != word3 {
        t.Error("Get word wrong")
    }
}