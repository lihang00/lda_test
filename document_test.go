package lda

import (
    "testing"
    "fmt"
)


func TestDocument(t *testing.T) {
    dic := NewWordDic()

    line1 := "the quick brown fox jumps over the lazy dog"
    doc1 := NewDoc(line1, dic)

    line2 := "dog fox"
    doc2 := NewDoc(line2, dic)

    if len(doc1.Words) != 9 || len(doc2.Words) != 2  {
        t.Error("Doc length error")
    }

    if doc2.Words[0] != 7 || doc2.Words[1] != 3 {
        fmt.Println(dic.GetId("dog"))
        fmt.Println(dic.GetId("fox"))
        fmt.Println(doc2.Words[0])
        fmt.Println(doc2.Words[1])
        t.Error("Word in doc mapping error")
    }

}