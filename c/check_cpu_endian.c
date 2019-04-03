//
//  main.c
//  test
//
//  Created by 朱鸿达 on 2019/4/1.
//  Copyright © 2019 aaron. All rights reserved.
//

#include <stdio.h>

int main(int argc, const char * argv[]) {
    // insert code here...
    printf("%d\n", checkCPUendian());
    return 0;
}

//  由于联合体union的存放顺序是所有成员都从低地址开始存放，利用该特性就可以轻松地获得了CPU对内存采用Little- endian还是Big-endian模式读写。
int checkCPUendian(){
    union {
        unsigned int a;
        unsigned char b;
    }c;
    c.a = 1;
    return (c.b == 1);
}/*return 1 : little-endian, return 0:big-endian*/

