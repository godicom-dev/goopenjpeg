#ifndef GOOPENJPEG_J2K_DECODE_H
#define GOOPENJPEG_J2K_DECODE_H

#include <openjpeg.h>
#include <stddef.h>

typedef struct JPEG2000Parameters {
    OPJ_UINT32 columns;
    OPJ_UINT32 rows;
    OPJ_COLOR_SPACE colourspace;
    OPJ_UINT32 nr_components;
    OPJ_UINT32 precision;
    unsigned int is_signed;
    OPJ_UINT32 nr_tiles;
} j2k_parameters_t;

const char* j2k_openjpeg_version(void);

int j2k_get_parameters(
    const unsigned char* data,
    size_t data_len,
    int codec_format,
    j2k_parameters_t* output
);

int j2k_decode(
    const unsigned char* data,
    size_t data_len,
    unsigned char* out,
    int codec_format
);

#endif
