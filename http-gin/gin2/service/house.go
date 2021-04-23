package service

import "strconv"

const (
    HouseMaxNum = 1000
    HouseMinNum = 1
    HouseDefaultNum = 2
)

type House struct {
    ID int
    Name string
}

type Req struct {
    Num int `form:"num"`
}

// 初始化结构体
func New(id int, name string) *House {
    return &House{
        ID:id,
        Name:name,
    }
}

// 新建多少栋房子
func BuildHouse(n int) []*House {
    sliceHouse := make([]*House, 0)
    for i := 0; i < n; i++{
        sliceHouse = append(sliceHouse, New(10 + i, "house name "+strconv.Itoa(10+i)))
    }
    return sliceHouse
}