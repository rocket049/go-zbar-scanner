/*
Author: Fu Huizhong <fuhuizn@163.com>
2025-06-02
*/
#include "zbar.h"
#include <string.h>
#include <stdlib.h>
#include <stdio.h>

char *zbar_scan(unsigned char *image_data,int width,int height, char *fmt)
{
    zbar_image_t *zbar_image = zbar_image_create();
    zbar_image_set_format(zbar_image,*(unsigned int *)fmt);
    zbar_image_set_size(zbar_image,width,height);
    uint8_t *qrcodedata = (uint8_t *)calloc(width*height,sizeof(uint8_t));
    memcpy(qrcodedata,image_data,width*height);

    zbar_image_set_data(zbar_image,qrcodedata,width * height,zbar_image_free_data);
    //创建图像扫描器
    zbar_image_scanner_t *scanner = zbar_image_scanner_create();
    //配置扫描器（可选）
    zbar_image_scanner_set_config(scanner,ZBAR_QRCODE,ZBAR_CFG_ENABLE,1);
    //扫描图像
    int n = zbar_scan_image(scanner,zbar_image);
    char *ret = NULL;
    if (n > 0){
        //处理扫描结果
        const zbar_symbol_set_t *symbols = zbar_image_scanner_get_results(scanner);
        const zbar_symbol_t *symbol = zbar_symbol_set_first_symbol(symbols);
        if (symbol != NULL)
        {
            //打印二维码内容
            ret = (char *)zbar_symbol_get_data(symbol);
        }
    }
    char *res = malloc(strlen(ret)+1);
    memcpy(res, ret,strlen(ret)+1);
    //清理资源
    zbar_image_destroy(zbar_image);
    zbar_image_scanner_destroy(scanner);
    
    return res;
}
