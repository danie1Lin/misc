//from https://hackmd.io/@sysprog/linux-macro-minmax?fbclid=IwAR0rifZ-jMc7FfN_w93pXrW3tGVjEsnknjdMmxMpWgBSvBzmAmKuyix2NHQ
// use gcc -E to check marco
#include <stdio.h>
#define max(a, b) ({ \
	typeof (a) _a = (a); \
	typeof (b) _b = (b); \
	(void) (&_a == &_b); \
	_a > _b ? _a : _b; \
})

void doOneTime() { printf("called doOneTime!\n"); }
int f1() { doOneTime(); return 0; }
int f2() { doOneTime(); return 1; }

int main() {
	printf("%i\n", max(f1(), f2()));

	printf("%i\n", max(1, "fdsafsd"));
}
