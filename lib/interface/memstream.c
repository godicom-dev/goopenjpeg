#include "memstream.h"

#include <string.h>

memstream_t memstream_from_buffer(const unsigned char* data, size_t size)
{
    memstream_t stream = {0};
    stream.data = data;
    stream.size = (OPJ_UINT64)size;
    stream.offset = 0;
    return stream;
}

OPJ_SIZE_T mem_read(void* dest, OPJ_SIZE_T nr_bytes, void* user_data)
{
    memstream_t* stream = (memstream_t*)user_data;
    if (stream == NULL || dest == NULL) {
        return (OPJ_SIZE_T)-1;
    }

    if (stream->offset >= stream->size) {
        return (OPJ_SIZE_T)-1;
    }

    OPJ_UINT64 remaining = stream->size - stream->offset;
    OPJ_SIZE_T to_read = nr_bytes;
    if ((OPJ_UINT64)to_read > remaining) {
        to_read = (OPJ_SIZE_T)remaining;
    }

    memcpy(dest, stream->data + stream->offset, to_read);
    stream->offset += to_read;
    return to_read;
}

OPJ_BOOL mem_seek(OPJ_OFF_T offset, void* user_data)
{
    memstream_t* stream = (memstream_t*)user_data;
    if (stream == NULL) {
        return OPJ_FALSE;
    }
    if (offset < 0 || (OPJ_UINT64)offset > stream->size) {
        return OPJ_FALSE;
    }
    stream->offset = (OPJ_UINT64)offset;
    return OPJ_TRUE;
}

OPJ_OFF_T mem_skip(OPJ_OFF_T offset, void* user_data)
{
    memstream_t* stream = (memstream_t*)user_data;
    if (stream == NULL) {
        return -1;
    }

    OPJ_INT64 new_offset = (OPJ_INT64)stream->offset + offset;
    if (new_offset < 0 || (OPJ_UINT64)new_offset > stream->size) {
        return -1;
    }

    stream->offset = (OPJ_UINT64)new_offset;
    return offset;
}

OPJ_UINT64 mem_length(void* user_data)
{
    memstream_t* stream = (memstream_t*)user_data;
    if (stream == NULL) {
        return 0;
    }
    return stream->size;
}
