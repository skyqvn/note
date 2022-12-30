extern "C"
{
#include "sum.h"
}

class sum
{
public:
    int f;
    int s;
    int getSum();
};

int sum::getSum()
{
    return this->f + this->s;
}

int s()
{
    sum s;
    s.f = 8;
    s.s = 9;
    return s.getSum();
}
