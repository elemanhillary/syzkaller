// AUTOGENERATED FROM executor/common_bsd.h
package csource

var commonHeaderFreebsd = `


#include <unistd.h>
#if defined(SYZ_EXECUTOR) || defined(SYZ_THREADED) || defined(SYZ_COLLIDE)
#include <pthread.h>
#include <stdlib.h>
#endif
#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT))
#include <errno.h>
#include <signal.h>
#include <stdarg.h>
#include <stdio.h>
#include <sys/time.h>
#include <sys/wait.h>
#include <time.h>
#endif
#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT) && defined(SYZ_USE_TMP_DIR))
#include <dirent.h>
#endif

#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT)) ||      \
    defined(SYZ_USE_TMP_DIR) || defined(SYZ_HANDLE_SEGV) || defined(SYZ_TUN_ENABLE) || \
    defined(SYZ_SANDBOX_NAMESPACE) || defined(SYZ_SANDBOX_SETUID) ||                   \
    defined(SYZ_SANDBOX_NONE) || defined(SYZ_FAULT_INJECTION) || defined(__NR_syz_kvm_setup_cpu)
__attribute__((noreturn)) static void doexit(int status)
{
	_exit(status);
	for (;;) {
	}
}
#endif



#include <stdint.h>
#include <string.h>
#if defined(SYZ_EXECUTOR) || defined(SYZ_USE_TMP_DIR)
#include <errno.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#endif
#if defined(SYZ_EXECUTOR) || defined(SYZ_HANDLE_SEGV)
#include <setjmp.h>
#include <signal.h>
#include <string.h>
#endif
#if defined(SYZ_EXECUTOR) || defined(SYZ_DEBUG)
#include <stdarg.h>
#include <stdio.h>
#endif

#if defined(SYZ_EXECUTOR)
#define exit vsnprintf
#define _exit vsnprintf

typedef unsigned long long uint64;
typedef unsigned int uint32;
typedef unsigned short uint16;
typedef unsigned char uint8;

#if defined(__GNUC__)
#define SYSCALLAPI
#define NORETURN __attribute__((noreturn))
#define ALIGNED(N) __attribute__((aligned(N)))
#define PRINTF __attribute__((format(printf, 1, 2)))
#else
#define SYSCALLAPI WINAPI
#define NORETURN __declspec(noreturn)
#define ALIGNED(N) __declspec(align(N))
#define PRINTF
#endif

typedef long(SYSCALLAPI* syscall_t)(long, long, long, long, long, long, long, long, long);

struct call_t {
	const char* name;
	int sys_nr;
	syscall_t call;
};

extern call_t syscalls[];
extern unsigned syscall_count;
#endif

#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT)) ||               \
    defined(SYZ_USE_TMP_DIR) || defined(SYZ_TUN_ENABLE) || defined(SYZ_SANDBOX_NAMESPACE) ||    \
    defined(SYZ_SANDBOX_NONE) || defined(SYZ_SANDBOX_SETUID) || defined(SYZ_FAULT_INJECTION) || \
    defined(__NR_syz_kvm_setup_cpu)
const int kFailStatus = 67;
const int kRetryStatus = 69;
#endif

#if defined(SYZ_EXECUTOR)
const int kErrorStatus = 68;
#endif

#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT)) ||                  \
    defined(SYZ_USE_TMP_DIR) || defined(SYZ_TUN_ENABLE) || defined(SYZ_SANDBOX_NAMESPACE) ||       \
    defined(SYZ_SANDBOX_NONE) || defined(SYZ_SANDBOX_SETUID) || defined(__NR_syz_kvm_setup_cpu) || \
    defined(__NR_syz_init_net_socket)
NORETURN PRINTF static void fail(const char* msg, ...)
{
	int e = errno;
	va_list args;
	va_start(args, msg);
	vfprintf(stderr, msg, args);
	va_end(args);
	fprintf(stderr, " (errno %d)\n", e);
	doexit((e == ENOMEM || e == EAGAIN) ? kRetryStatus : kFailStatus);
}
#endif

#if defined(SYZ_EXECUTOR)
NORETURN PRINTF static void error(const char* msg, ...)
{
	va_list args;
	va_start(args, msg);
	vfprintf(stderr, msg, args);
	va_end(args);
	fprintf(stderr, "\n");
	doexit(kErrorStatus);
}
#endif

#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT) && defined(SYZ_USE_TMP_DIR)) || defined(SYZ_FAULT_INJECTION)
NORETURN PRINTF static void exitf(const char* msg, ...)
{
	int e = errno;
	va_list args;
	va_start(args, msg);
	vfprintf(stderr, msg, args);
	va_end(args);
	fprintf(stderr, " (errno %d)\n", e);
	doexit(kRetryStatus);
}
#endif

#if defined(SYZ_EXECUTOR) || defined(SYZ_DEBUG)
static int flag_debug;

PRINTF static void debug(const char* msg, ...)
{
	if (!flag_debug)
		return;
	va_list args;
	va_start(args, msg);
	vfprintf(stderr, msg, args);
	va_end(args);
	fflush(stderr);
}
#endif

#if defined(SYZ_EXECUTOR) || defined(SYZ_USE_BITMASKS)
#define BITMASK_LEN(type, bf_len) (type)((1ull << (bf_len)) - 1)

#define BITMASK_LEN_OFF(type, bf_off, bf_len) (type)(BITMASK_LEN(type, (bf_len)) << (bf_off))

#define STORE_BY_BITMASK(type, addr, val, bf_off, bf_len)                         \
	if ((bf_off) == 0 && (bf_len) == 0) {                                     \
		*(type*)(addr) = (type)(val);                                     \
	} else {                                                                  \
		type new_val = *(type*)(addr);                                    \
		new_val &= ~BITMASK_LEN_OFF(type, (bf_off), (bf_len));            \
		new_val |= ((type)(val)&BITMASK_LEN(type, (bf_len))) << (bf_off); \
		*(type*)(addr) = new_val;                                         \
	}
#endif

#if defined(SYZ_EXECUTOR) || defined(SYZ_USE_CHECKSUMS)
struct csum_inet {
	uint32 acc;
};

static void csum_inet_init(struct csum_inet* csum)
{
	csum->acc = 0;
}

static void csum_inet_update(struct csum_inet* csum, const uint8* data, size_t length)
{
	if (length == 0)
		return;

	size_t i;
	for (i = 0; i < length - 1; i += 2)
		csum->acc += *(uint16*)&data[i];

	if (length & 1)
		csum->acc += (uint16)data[length - 1];

	while (csum->acc > 0xffff)
		csum->acc = (csum->acc & 0xffff) + (csum->acc >> 16);
}

static uint16 csum_inet_digest(struct csum_inet* csum)
{
	return ~csum->acc;
}
#endif

#if defined(SYZ_EXECUTOR) || defined(SYZ_HANDLE_SEGV)
static __thread int skip_segv;
static __thread jmp_buf segv_env;

static void segv_handler(int sig, siginfo_t* info, void* uctx)
{
	uintptr_t addr = (uintptr_t)info->si_addr;
	const uintptr_t prog_start = 1 << 20;
	const uintptr_t prog_end = 100 << 20;
	if (__atomic_load_n(&skip_segv, __ATOMIC_RELAXED) && (addr < prog_start || addr > prog_end)) {
		debug("SIGSEGV on %p, skipping\n", (void*)addr);
		_longjmp(segv_env, 1);
	}
	debug("SIGSEGV on %p, exiting\n", (void*)addr);
	doexit(sig);
}

static void install_segv_handler()
{
	struct sigaction sa;

	memset(&sa, 0, sizeof(sa));
	sa.sa_sigaction = segv_handler;
	sa.sa_flags = SA_NODEFER | SA_SIGINFO;
	sigaction(SIGSEGV, &sa, NULL);
	sigaction(SIGBUS, &sa, NULL);
}

#define NONFAILING(...)                                              \
	{                                                            \
		__atomic_fetch_add(&skip_segv, 1, __ATOMIC_SEQ_CST); \
		if (_setjmp(segv_env) == 0) {                        \
			__VA_ARGS__;                                 \
		}                                                    \
		__atomic_fetch_sub(&skip_segv, 1, __ATOMIC_SEQ_CST); \
	}
#endif

#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT))
static uint64 current_time_ms()
{
	struct timespec ts;

	if (clock_gettime(CLOCK_MONOTONIC, &ts))
		fail("clock_gettime failed");
	return (uint64)ts.tv_sec * 1000 + (uint64)ts.tv_nsec / 1000000;
}
#endif

#if defined(SYZ_EXECUTOR)
static void sleep_ms(uint64 ms)
{
	usleep(ms * 1000);
}
#endif

#if defined(SYZ_EXECUTOR) || (defined(SYZ_REPEAT) && defined(SYZ_WAIT_REPEAT) && defined(SYZ_USE_TMP_DIR))
static void remove_dir(const char* dir)
{
	DIR* dp;
	struct dirent* ep;
	int iter = 0;
retry:
	dp = opendir(dir);
	if (dp == NULL)
		return;
	while ((ep = readdir(dp))) {
		if (strcmp(ep->d_name, ".") == 0 || strcmp(ep->d_name, "..") == 0)
			continue;
		char filename[FILENAME_MAX];
		snprintf(filename, sizeof(filename), "%s/%s", dir, ep->d_name);
		struct stat st;
		if (lstat(filename, &st))
			return;
		if (S_ISDIR(st.st_mode)) {
			remove_dir(filename);
			continue;
		}
		int i;
		for (i = 0;; i++) {
			if (unlink(filename) == 0)
				break;
			if (errno == EROFS)
				break;
			if (errno != EBUSY || i > 100)
				return;
		}
	}
	closedir(dp);
	int i;
	for (i = 0;; i++) {
		if (rmdir(dir) == 0)
			break;
		if (i < 100) {
			if (errno == EROFS)
				break;
			if (errno == ENOTEMPTY) {
				if (iter < 100) {
					iter++;
					goto retry;
				}
			}
		}
		return;
	}
}
#endif

#if defined(SYZ_EXECUTOR) || defined(SYZ_FAULT_INJECTION)
static int inject_fault(int nth)
{
	return 0;
}

static int fault_injected(int fail_fd)
{
	return 0;
}
#endif
`
