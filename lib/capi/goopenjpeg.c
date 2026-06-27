#include "goopenjpeg.h"

#include "../interface/j2k_decode.h"

#include <stdlib.h>
#include <string.h>

static int bytes_per_sample(int precision)
{
    if (precision <= 8) {
        return 1;
    }
    if (precision <= 16) {
        return 2;
    }
    if (precision <= 24) {
        return 3;
    }
    return 4;
}

GOOPENJPEG_EXPORT int goopenjpeg_version(char* buf, int buf_len)
{
    if (buf == NULL || buf_len <= 0) {
        return GOOPENJPEG_ERR_PARAM;
    }

    const char* version = j2k_openjpeg_version();
    if (version == NULL) {
        buf[0] = '\0';
        return GOOPENJPEG_OK;
    }

    strncpy(buf, version, (size_t)buf_len - 1);
    buf[buf_len - 1] = '\0';
    return GOOPENJPEG_OK;
}

GOOPENJPEG_EXPORT int goopenjpeg_get_parameters(
    const unsigned char* data,
    int data_len,
    int codec,
    int* width,
    int* height,
    int* components,
    int* precision,
    int* is_signed,
    int* colourspace)
{
    if (data == NULL || data_len <= 0 || width == NULL || height == NULL ||
        components == NULL || precision == NULL || is_signed == NULL || colourspace == NULL) {
        return GOOPENJPEG_ERR_PARAM;
    }

    j2k_parameters_t param = {0};
    const int code = j2k_get_parameters(data, (size_t)data_len, codec, &param);
    if (code != 0) {
        return code > 0 ? -code : GOOPENJPEG_ERR_DECODE;
    }

    *width = (int)param.columns;
    *height = (int)param.rows;
    *components = (int)param.nr_components;
    *precision = (int)param.precision;
    *is_signed = (int)param.is_signed;
    *colourspace = (int)param.colourspace;
    return GOOPENJPEG_OK;
}

GOOPENJPEG_EXPORT int goopenjpeg_decode(
    const unsigned char* data,
    int data_len,
    int codec,
    unsigned char** output,
    int* output_len,
    int* width,
    int* height,
    int* components,
    int* precision,
    int* is_signed)
{
    if (data == NULL || data_len <= 0 || output == NULL || output_len == NULL ||
        width == NULL || height == NULL || components == NULL || precision == NULL || is_signed == NULL) {
        return GOOPENJPEG_ERR_PARAM;
    }

    *output = NULL;
    *output_len = 0;

    j2k_parameters_t param = {0};
    int code = j2k_get_parameters(data, (size_t)data_len, codec, &param);
    if (code != 0) {
        return code > 0 ? -code : GOOPENJPEG_ERR_DECODE;
    }

    const int bpp = bytes_per_sample((int)param.precision);
    const int out_len = (int)(param.columns * param.rows * param.nr_components * bpp);
    if (out_len <= 0) {
        return GOOPENJPEG_ERR_PARAM;
    }

    unsigned char* out_buf = (unsigned char*)malloc((size_t)out_len);
    if (out_buf == NULL) {
        return GOOPENJPEG_ERR_MEMORY;
    }

    code = j2k_decode(data, (size_t)data_len, out_buf, codec);
    if (code != 0) {
        free(out_buf);
        return code > 0 ? -code : GOOPENJPEG_ERR_DECODE;
    }

    *output = out_buf;
    *output_len = out_len;
    *width = (int)param.columns;
    *height = (int)param.rows;
    *components = (int)param.nr_components;
    *precision = (int)param.precision;
    *is_signed = (int)param.is_signed;
    return GOOPENJPEG_OK;
}

GOOPENJPEG_EXPORT void goopenjpeg_free(unsigned char* p)
{
    free(p);
}
