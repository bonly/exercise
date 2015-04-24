#include <stdio.h>
#include <stdlib.h>
#include <iconv.h>
#include <wchar.h>

#define UTF8_SEQUENCE_MAXLEN 6
/* #define UTF8_SEQUENCE_MAXLEN 16 */

int
main(int argc, char **argv)
{
    wchar_t *wcs = L"A";
    signed char utf8[(1 /* wcslen(wcs) */ + 1 /* L'\0' */) * UTF8_SEQUENCE_MAXLEN];
    char *iconv_in = (char *) wcs;
    char *iconv_out = (char *) &utf8[0];
    size_t iconv_in_bytes = (wcslen(wcs) + 1 /* L'\0' */) * sizeof(wchar_t);
    size_t iconv_out_bytes = sizeof(utf8);
    size_t ret;
    iconv_t cd;

    cd = iconv_open("UTF-8", "WCHAR_T");
    if ((iconv_t) -1 == cd) {
        perror("iconv_open");
        return EXIT_FAILURE;
    }

    ret = iconv(cd, &iconv_in, &iconv_in_bytes, &iconv_out, &iconv_out_bytes);
    if ((size_t) -1 == ret) {
        perror("iconv");
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}
