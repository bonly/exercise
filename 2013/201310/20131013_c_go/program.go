// #include <stdint.h>
// #Include <library.h>
// extern void set_callbacks(IOContext *ioctx, void *opaque);

func RunSomething() {
    ioctx := C.alloc_io_context()
    // checks

    r := package.NewReader()
    C.set_callbacks(ioctx, unsafe.Pointer(r)

    // Stuff Happens

    // Works in 1.4, random failure in 1.5.
    C.process()

    // ...
}