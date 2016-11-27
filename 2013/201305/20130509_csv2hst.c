#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <time.h>

#pragma pack(1)
struct RateInfo
{
    time_t ctm;           // current time in seconds since 01.01.1970
    double open;
    double low;
    double high;
    double close;
    int    vol;
};
#pragma pack(8)
struct HistoryHeader
{
    int    version;
    char   copyright[64];
    char   symbol[8];
    int    period;
    int    unused[16];
};
//+------------------------------------------------------------------+
//| usage 1:  csv2hst SSSSSSSP.csv                                   |
//| usage 2:  csv2hst file.csv SSSSSSSP.hst                          |
//+------------------------------------------------------------------+
int main(int argc, char* argv[])
{
    FILE *in, *out;
    char *cp;
    int   i, period;
    char  inputfilename[256] = "";
    char  outputfilename[256] = "";
    char  symbol[8];
    char  buffer[256];
    char *delimiter = ",";
    tm            tt;
    RateInfo      ri;
    HistoryHeader hdr;

    if (argc < 2) return -1;

    strcpy(inputfilename,argv[1]);
    strcpy(outputfilename,argv[1]);
    if ((cp = strstr(outputfilename, ".csv")) == NULL) return -2;

    if (argc > 2)
        strcpy(outputfilename,argv[2]);
    else
    {
        *cp=0;
        strcat(outputfilename, ".hst");
    }
    if (stricmp(inputfilename, outputfilename) == 0)   return -3;
    if ((cp = strstr(outputfilename, ".hst")) == NULL) return -4;

    if ((in = fopen(inputfilename, "ri")) == NULL)     return -5;
    if ((out = fopen(outputfilename, "wb")) == NULL)
    {
        fclose(in);
        return -6;
    }

    // take period and symbol from output file name
    *cp=0;
    period=0;
    for (i=strlen(outputfilename)-1; i>=0; i--)
    {
        if ((outputfilename[i]<'0' || outputfilename[i]>'9') && period == 0)
        {
            period = atoi(outputfilename+i+1);
            outputfilename[i+1]=0;
        }
        if (outputfilename[i]=='\\' || outputfilename[i]=='/')
            break;
    }
    strncpy(symbol, outputfilename+i+1, 7);
    symbol[7] = 0;
    // prepare and write hst header
    memset(&hdr, 0, sizeof(HistoryHeader));
    hdr.version = 1;
    hdr.period = period;
    strcpy(hdr.symbol, symbol);
    strcpy(hdr.copyright, "Copyright © 2003, MetaQuotes Software Corp.");
    fwrite(&hdr, sizeof(HistoryHeader), 1, out);
    // prepare and write RateInfo records
    i=0;
    while (fgets(buffer, 255, in) != NULL)
    {
        buffer[255] = 0;
        cp = buffer;
        //---- date DD.MM.YYYY
        tt.tm_mday = atoi(cp);
        if (tt.tm_mday < 1 || tt.tm_mday > 31)        continue;
        if ((cp = strstr(cp, ".")) == NULL)           continue;
        tt.tm_mon = atoi(cp+1) - 1;
        if (tt.tm_mon < 0  || tt.tm_mon > 11)         continue;
        if ((cp = strstr(cp+1, ".")) == NULL)         continue;
        tt.tm_year = atoi(cp+1);
        if (tt.tm_year > 1900) tt.tm_year -= 1900;
        if (tt.tm_year < 50)   tt.tm_year += 100;
        //---- time ,HH:MI
        if ((cp = strstr(buffer, delimiter)) == NULL) continue;
        tt.tm_hour = atoi(cp+1);
        if (tt.tm_hour < 0 || tt.tm_hour > 23)        continue;
        if ((cp = strstr(cp+1, ":")) == NULL)         continue;
        tt.tm_min = atoi(cp+1);
        if (tt.tm_min < 0  || tt.tm_min > 59)         continue;
        if ((ri.ctm = mktime(&tt)) == 0)              continue;
        ri.ctm -= _timezone;
        //---- open,high,low,close,volume
        if ((cp = strstr(cp, delimiter)) == NULL)     continue;
        if ((ri.open = atof(cp+1)) <= 0)              continue;
        if ((cp = strstr(cp+1, delimiter)) == NULL)   continue;
        if ((ri.high = atof(cp+1)) <= 0)              continue;
        if ((cp = strstr(cp+1, delimiter)) == NULL)   continue;
        if ((ri.low = atof(cp+1)) <= 0)               continue;
        if ((cp = strstr(cp+1, delimiter)) == NULL)   continue;
        if ((ri.close = atof(cp+1)) <= 0)             continue;
        if ((cp = strstr(cp+1, delimiter)) == NULL)   continue;
        if ((ri.vol = atoi(cp+1)) < 0)                continue;
        //----- check prices
        if (ri.low > ri.open  || ri.low > ri.close  || ri.low > ri.high) continue;
        if (ri.high < ri.open || ri.high < ri.close || ri.high < ri.low) continue;
        if (gmtime(&ri.ctm) == NULL)                  continue;
        i++;
        fwrite(&ri,sizeof(RateInfo),1,out);
    }
    fclose(in);
    fclose(out);
    printf("%d records written\n", i);
    return 0;
}
