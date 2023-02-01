package utils

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

func GeneratUUID() string {
	// V1 基于时间
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u1.String())

	// V4 基于随机数
	u4 := uuid.New()

	return u4.String() // a0d99f20-1dd1-459b-b516-dfeca4005203
}
