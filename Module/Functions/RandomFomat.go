package functions

import (
    "fmt"
    "math/rand"
)

func RandomFomat(name string) string {

    formats := []string{
        "HI, %v, welcome",
        "xin chào, %v, welcome",
        "Ni hao, %v, laile",
    }

    return fmt.Sprintf(
        formats[rand.Intn(len(formats))],
        name,
    )
}