package service

import "strconv"

const (
    HouseDefaultNum = 2
)

type House struct {
    ID int `json:"id" form:"id"`
    Name string `json:"name" form:"name"`
    Floor int `json:"floor" form:"floor"`
}

// 新建房子的数量
type Num struct {
    Num int `json:"num" form:"num"`
}

// 初始化结构体
func New(id int, name string, floor int) *House {
    return &House{
        ID: id,
        Name: name,
        Floor: floor,
    }
}

// 新建多少栋房子
func BuildHouse(num int) []*House {
    sliceHouse := make([]*House, 0)
    for i := 1; i <= num; i++{
        // 给房子简单命名
        sliceHouse = append(sliceHouse, New(i, "house name "+strconv.Itoa(i), 10))
    }
    return sliceHouse
}