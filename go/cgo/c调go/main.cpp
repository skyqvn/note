#include <iostream>
#include "sum.h"

using namespace std;

int main(){
    char c[12]="hello world";  
    GoString s={
        c,
        sizeof(c)
    };
    printString(s);

    int i=5;
    printNum(i);
}
