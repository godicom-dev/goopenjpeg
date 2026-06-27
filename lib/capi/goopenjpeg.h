#ifndef GOOPENJPEG_H
#define GOOPENJPEG_H

#include <stddef.h>

#ifdef _WIN32
#define GOOPENJPEG_EXPORT __declspec(dllexport)
#else
#define GOOPENJPEG_EXPORT __attribute__((visibility("default")))
#endif

#ifdef __cplusplus
extern "C" {
#endif

#define GOOPENJPEG_OK          0
#define GOOPENJPEG_ERR_PARAM  -1
#define GOOPENJPEG_ERR_MEMORY -2
#define GOOPENJPEG_ERR_DECODE -3

#define GOOPENJPEG_CODEC_J2K 0
#define GOOPENJPEG_CODEC_JPT 1
#define GOOPENJPEG_CODEC_JP2 2

GOOPENJPEG_EXPORT int goopenjpeg_version(char* buf, int buf_len);

GOOPENJPEG_EXPORT int goopenjpeg_get_parameters(
    const unsigned char* data,
    int data_len,
    int codec,
    int* width,
    int* height,
    int* components,
    int* precision,
    int* is_signed,
    int* colourspace
);

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
    int* is_signed
);

GOOPENJPEG_EXPORT void goopenjpeg_free(unsigned char* p);

#ifdef __cplusplus
}
#endif

#endif
