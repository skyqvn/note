int run_callback(void* fn,int state) {
	return ((int(*)(int))(fn))(state);
}
