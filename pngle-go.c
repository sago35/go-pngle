#include "miniz.h"
#include "pngle.h"
#include "pngle-go.h"

void on_draw(pngle_t *pngle, uint32_t x, uint32_t y, uint32_t w, uint32_t h, uint8_t rgba[4])
{
    uint8_t r = rgba[0]; // 0 - 255
    uint8_t g = rgba[1]; // 0 - 255
    uint8_t b = rgba[2]; // 0 - 255
    uint8_t a = rgba[3]; // 0: fully transparent, 255: fully opaque

    //printf("put pixel at (%d, %d) (w%d h%d) with color #%02x%02x%02x\n", x, y, w, h, r, g, b);
    callbackFromPngle(x, y, w, h, r, g, b, a);
}

int decodeFromBytes(unsigned char *b, int length, int scale)
{
    pngle_t *pngle = pngle_new();

    pngle_set_draw_callback(pngle, on_draw);

    // Feed data to pngle
    int fed = pngle_feed(pngle, b, length);
    if (fed < 0) {
        return -1;
    }

    pngle_destroy(pngle);

    return 0;
}
