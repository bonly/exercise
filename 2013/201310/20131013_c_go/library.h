typedef struct IOContext {
    void *opaque;
    int (*read)(void *opaque, uint8_t *buf, int size);
    int64_t (*seek)(void *opaque, int64_t offset, int whence);
    int64_t (*size)(void *opaque);
} IOContext;