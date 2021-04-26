package service

import (
    "context"
    proto "go-micro-demos/circuitbreaker/hystrixdo/proto/house"
    "strconv"
    "time"
)

type HouseService struct {

}

func newHouse(id int32, name string, floor int32) *proto.RequestData {
    return &proto.RequestData{Name:name, Id:id, Floor:floor}
}

func (*HouseService) GetHouse(ctx context.Context, req *proto.RequestData, resp *proto.ResponseMsg) error {
    // hystrix 超时测试
    time.Sleep(time.Second * 5)

    reqdata := make([]*proto.RequestData, 0)
    for i := 1; i < 6;i++ {
        reqdata = append(reqdata, newHouse(int32(i), "name"+strconv.Itoa(i), int32(i)))
    }

    return nil
}

func (*HouseService) Build(ctx context.Context, req *proto.RequestData, resp *proto.ResponseMsg) error {

    return nil
}