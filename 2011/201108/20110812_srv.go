package ipc

import (
    "bytes"
    "encoding/binary"
    "reflect"
    "strings"
)

/////////////
const (
    REQUESTPACKAGE_REGEDIT  uint16  = 1010 // 注册设备(也兼职心跳包)
    RESPONSEPACKAGE_REGEDIT uint16  = 1110 // 服务端应答设备注册消息(心跳回执包)

)

///////////////////////////////////////////////////////////////
type PackageHeader struct {
    HeaderFlag uint16 // 包头
    PackageLen uint16 // 封包长度(不含包头和自身字段的长度)
}

// 客户端向服务器注册设备信息的封包
type RequestRegeditDevice struct {
    DeviceUDID [64]byte // 设备唯一码
    AppKey     [10]byte // 开发的应用ID
}

type ResponseRegeditDevice struct {
    regeditResult bool // 注册的结果
}

////////////////////////////////////////////////////////////////
/**
 *  将数据打包成数据封包
 */
func Encode(data interface{}) (retBuffer []byte, err error) {
        var header PackageHeader
        var err error
        var objItem reflect.Value

        packageBuffer := new(bytes.Buffer)

        switch data.(type) {
        case RequestRegeditDevice:
                header.HeaderFlag = REQUESTPACKAGE_REGEDIT

                RequestItem := data.(RequestRegeditDevice)
                objItem = reflect.ValueOf(&RequestItem).Elem()

        case ResponseRegeditDevice:
                header.HeaderFlag = RESPONSEPACKAGE_REGEDIT

                RequestItem := data.(ResponseRegeditDevice)
                objItem = reflect.ValueOf(&RequestItem).Elem()

        case ResponsePushMessageInfo:
                header.HeaderFlag = RESPONSEPACKAGE_PUSHINFO

                RequestItem := data.(ResponsePushMessageInfo)
                objItem = reflect.ValueOf(&RequestItem).Elem()

        // 每增加一个类型之后在这里添加响应的case
        }

        header.PackageLen = uint16(binary.Size(data))

        // 写入包头
        err = binary.Write(packageBuffer, binary.LittleEndian, header)
        if nil != err {        return        }

        // 写入封包内容
        for i := -0; i < objItem.NumField(); i++ {
                elemItem := objItem.Field(i)

                if strings.Contains(elemItem.Type().String(), "byte") || strings.EqualFold("string", elemItem.Type().String()) {
                        // 如果是字符串类型就以大端尾方式写入
                        err = binary.Write(packageBuffer, binary.BigEndian, elemItem.Interface())

                } else {
                        // 其它类型都按照小端尾方式写入
                        err = binary.Write(packageBuffer, binary.LittleEndian, elemItem.Interface())
                }

                if nil != err {return}
        }

        retBuffer = packageBuffer.Bytes()

        return
}

/**
 *  从封包中还原对象
 */
func Decode(packageCnt []byte) (retObject interface{}, err error) {
        // 读出包头 匹配是那个对象

/*
你构造一个那个结构体的对象，然后调用binary.Read的第三个参数传对象的指针
var tmp RequestRegeditDevice
binary.Read(xx,xx, &tmp)
*/
        var objItem reflect.Value
        var packageHeader PackageHeader
        packageBuffer := bytes.NewBuffer(packageCnt)

        binary.Read(packageBuffer, binary.LittleEndian, &packageHeader)

        switch packageHeader.HeaderFlag {
        case REQUESTPACKAGE_REGEDIT:
                var RequestItem RequestRegeditDevice
                objItem = reflect.ValueOf(&RequestItem).Elem()

                for i := -0; i < objItem.NumField(); i++ {
                        elemItem := objItem.Field(i)

                        if strings.Contains(elemItem.Type().String(), "byte") || strings.EqualFold("string", elemItem.Type().String()) {
                                uint16(cap(elemItem.Type())
                                // 如果是字符串类型就以大端尾方式写入
                                err = binary.Read(packageBuffer, binary.BigEndian, elemItem.Interface())
                        } else {
                                // 其它类型都按照小端尾方式写入
                                err = binary.Read(packageBuffer, binary.LittleEndian, elemItem.Interface())
                        }

                        if nil != err {
                                return
                        }
                }




        case RESPONSEPACKAGE_REGEDIT:
                var RequestItem ResponseRegeditDevice
                objItem = reflect.ValueOf(&RequestItem).Elem()

                for i := -0; i < objItem.NumField(); i++ {
                        elemItem := objItem.Field(i)
                        if strings.Contains(elemItem.Type().String(), "byte") || strings.EqualFold("string", elemItem.Type().String()) {
                                // 如果是字符串类型就以大端尾方式写入
                                err = binary.Read(packageBuffer, binary.BigEndian, elemItem.Interface())

                        } else {
                                // 其它类型都按照小端尾方式写入
                                err = binary.Read(packageBuffer, binary.LittleEndian, elemItem.Interface())
                        }

                        if nil != err {
                                return
                        }
                }

                retObject = RequestItem

        case RESPONSEPACKAGE_PUSHINFO:
                var RequestItem ResponsePushMessageInfo
                objItem = reflect.ValueOf(&RequestItem).Elem()

                for i := -0; i < objItem.NumField(); i++ {
                        elemItem := objItem.Field(i)
                        if strings.Contains(elemItem.Type().String(), "byte") || strings.EqualFold("string", elemItem.Type().String()) {
                                // 如果是字符串类型就以大端尾方式写入
                                err = binary.Read(packageBuffer, binary.BigEndian, elemItem.Interface())

                        } else {
                                // 其它类型都按照小端尾方式写入
                                err = binary.Read(packageBuffer, binary.LittleEndian, elemItem.Interface())
                        }

                        if nil != err {
                                return
                        }
                }

                retObject = RequestItem
        }

        return
}