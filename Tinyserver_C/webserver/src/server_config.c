#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <unix_wrap.h>

#include "variable.h"
#include "server_config.h"


int strkv(char *src, char *key, char *value)
{
    char *p,*q;

    if(*src == '#') return 0;

    p = strchr(src, '=');
    q = strchr(src, '\n');

    if (p != NULL && q != NULL)
    {
        *q = '\0';
        strncpy(key, src, p - src);
        strncpy(value, p+1, strlen(p+1));
        return 1;
    }
    else
        return 0;
}



void get_config(char *configFilePath, struct configItem * configVar, int configNum)
{
    int i;
    FILE *fd;
    char buf[100] = "";
    char key[20] = "";
    char value[80] = "";

    fd = Fopen(configFilePath, "r");

    while(fgets(buf, 100, fd))
    {
        if (strkv(buf, key, value))
        {
            for(i = 0; i< configNum; i++)
            {
                if (!strcmp(key, configVar[i].key))
                {
                    strncpy(configVar[i].value, value, strlen(value));
                }
            }
            memset(key, 0, strlen(key));
            memset(value, 0, strlen(value));
        }
    }

    fclose(fd);

}

