#include "_cgo_export.h"

#include <stdint.h>
#include <library.h>

void set_callbacks(IOContext *ioctx, void *opaque)
{
    ioctx->opaque        = opaque;
    ioctx->read          = ReadCallback;
    ioctx->seek          = SeekCallback;
    ioctx->size          = SizeCallback;
}