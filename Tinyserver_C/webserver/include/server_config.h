#ifndef TINYSERVER_SERVER_CONFIG_H
#define TINYSERVER_SERVER_CONFIG_H

int strkv(char *src, char *key, char *value);
void get_config(char *configFilePath, struct configItem * configVar, int configNum);

#endif 
