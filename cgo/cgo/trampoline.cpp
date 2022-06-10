#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <pthread.h>
#include <stdio.h>
#include <time.h>
#include <unistd.h>
#include <thread>
#include <chrono>
#include "trampoline.h"

void pidato_stub() {}

//typedef void pidato_trampoline_handler(size_t arg0, size_t arg1);

void pidato_trampoline(size_t fn, size_t arg0, size_t arg1) {
	((pidato_trampoline_handler*)fn)(arg0, arg1);
}

void pidato_sleep(size_t arg0, size_t arg1) {
	std::this_thread::sleep_for((std::chrono::nanoseconds)arg0);
}