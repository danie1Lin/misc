#include <stdio.h>
#define ___PASTE(a , b) a##b
#define __PASTE(a , b) ___PASTE(a, b)

// __COUNTER__ is defined in https://gcc.gnu.org/onlinedocs/cpp/Common-Predefined-Macros.html
#define __UNIQUE_ID(prefix) \
    __PASTE(__PASTE(__UNIQUE_ID_, prefix), __COUNTER__)

#define __max(t1, t2, max1, max2, x, y) ({		\
	t1 max1 = (x);					\
	t2 max2 = (y);					\
	(void) (&max1 == &max2);			\
	max1 > max2 ? max1 : max2; })

#define max(x, y)					\
	__max(typeof(x), typeof(y),			\
	      __UNIQUE_ID(max1_), __UNIQUE_ID(max2_),	\
	      x, y)

void doOneTime() { printf("called doOneTime!\n"); }
int f1() { doOneTime(); return 0; }
int f2() { doOneTime(); return 1; }

int main() {
	printf("%i\n", max(f1(), f2()));

	printf("%i\n", max(1, "fdsafsd"));
}
