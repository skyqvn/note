#include "run_callback.h"

int run_callback(void* fn,int state) {
	//void* fn再转换是因为类型直接为callback_func时，CGO类型为*[0]byte
	return ((callback_func)fn)(state);
}
