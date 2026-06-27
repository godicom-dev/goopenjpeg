#include <stddef.h>

int goopenjpeg_version(char* buf, int buf_len)
{
    if (buf == NULL || buf_len <= 0) {
        return -1;
    }
    const char* version = "0.0.0-stub";
    for (int i = 0; version[i] != '\0' && i < buf_len - 1; ++i) {
        buf[i] = version[i];
    }
    buf[buf_len - 1] = '\0';
    return 0;
}

int goopenjpeg_get_parameters(
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
    (void)data;
    (void)data_len;
    (void)codec;
    (void)width;
    (void)height;
    (void)components;
    (void)precision;
    (void)is_signed;
    (void)colourspace;
    return -3;
}

int goopenjpeg_decode(
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
    (void)data;
    (void)data_len;
    (void)codec;
    (void)output;
    (void)output_len;
    (void)width;
    (void)height;
    (void)components;
    (void)precision;
    (void)is_signed;
    return -3;
}

void goopenjpeg_free(unsigned char* p)
{
    (void)p;
}
