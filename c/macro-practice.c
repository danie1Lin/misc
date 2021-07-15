//from https://hackmd.io/@sysprog/linux-macro-minmax?fbclid=IwAR0rifZ-jMc7FfN_w93pXrW3tGVjEsnknjdMmxMpWgBSvBzmAmKuyix2NHQ
#include <stdio.h>
#define max(x, y) (x) > (y) ? (x) : (y)


void doOneTime() { printf("called doOneTime!\n"); }
int f1() { doOneTime(); return 0; }
int f2() { doOneTime(); return 1; }

int main() {
	printf("%i\n", max(f1(), f2()));
}
