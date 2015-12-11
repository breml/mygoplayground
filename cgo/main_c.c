#include <stdio.h>
#include <stdlib.h>
#include "_cgo_export.h"

void cCall() {
	const char *buffer = "test from c\n";
	fputs(buffer, stdout);
	goCall();
}