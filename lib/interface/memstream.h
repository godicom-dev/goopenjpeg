#ifndef GOOPENJPEG_MEMSTREAM_H
#define GOOPENJPEG_MEMSTREAM_H

#include <openjpeg.h>
#include <stddef.h>

typedef struct memstream {
    const unsigned char* data;
    OPJ_UINT64 size;
    OPJ_UINT64 offset;
} memstream_t;

memstream_t memstream_from_buffer(const unsigned char* data, size_t size);

OPJ_SIZE_T mem_read(void* dest, OPJ_SIZE_T nr_bytes, void* user_data);

OPJ_BOOL mem_seek(OPJ_OFF_T offset, void* user_data);

OPJ_OFF_T mem_skip(OPJ_OFF_T offset, void* user_data);

OPJ_UINT64 mem_length(void* user_data);

#endif
