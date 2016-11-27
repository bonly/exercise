// #include <stdint.h>
import "C"

//export ReadCallback
func ReadCallback(r unsafe.Pointer, buf *C.uint8_t, size C.int) C.int {
    return someReadCallback((*package.Reader)(r), buf, size)
}

//export SeekCallback
func SeekCallback(r unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
    return someSeekCallback((*package.Reader)(r), offset, whence)
}

//export SizeCallback
func SizeCallback(r unsafe.Pointer) C.int64_t {
    sb   := (*package.Reader)(r)
    size := sb.Size()

    return C.int64_t(size)
}