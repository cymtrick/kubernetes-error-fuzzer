// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/include /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build ppc64le && linux
// +build ppc64le,linux

package unix

const (
	SYS_RESTART_SYSCALL         = 0
	SYS_EXIT                    = 1
	SYS_FORK                    = 2
	SYS_READ                    = 3
	SYS_WRITE                   = 4
	SYS_OPEN                    = 5
	SYS_CLOSE                   = 6
	SYS_WAITPID                 = 7
	SYS_CREAT                   = 8
	SYS_LINK                    = 9
	SYS_UNLINK                  = 10
	SYS_EXECVE                  = 11
	SYS_CHDIR                   = 12
	SYS_TIME                    = 13
	SYS_MKNOD                   = 14
	SYS_CHMOD                   = 15
	SYS_LCHOWN                  = 16
	SYS_BREAK                   = 17
	SYS_OLDSTAT                 = 18
	SYS_LSEEK                   = 19
	SYS_GETPID                  = 20
	SYS_MOUNT                   = 21
	SYS_UMOUNT                  = 22
	SYS_SETUID                  = 23
	SYS_GETUID                  = 24
	SYS_STIME                   = 25
	SYS_PTRACE                  = 26
	SYS_ALARM                   = 27
	SYS_OLDFSTAT                = 28
	SYS_PAUSE                   = 29
	SYS_UTIME                   = 30
	SYS_STTY                    = 31
	SYS_GTTY                    = 32
	SYS_ACCESS                  = 33
	SYS_NICE                    = 34
	SYS_FTIME                   = 35
	SYS_SYNC                    = 36
	SYS_KILL                    = 37
	SYS_RENAME                  = 38
	SYS_MKDIR                   = 39
	SYS_RMDIR                   = 40
	SYS_DUP                     = 41
	SYS_PIPE                    = 42
	SYS_TIMES                   = 43
	SYS_PROF                    = 44
	SYS_BRK                     = 45
	SYS_SETGID                  = 46
	SYS_GETGID                  = 47
	SYS_SIGNAL                  = 48
	SYS_GETEUID                 = 49
	SYS_GETEGID                 = 50
	SYS_ACCT                    = 51
	SYS_UMOUNT2                 = 52
	SYS_LOCK                    = 53
	SYS_IOCTL                   = 54
	SYS_FCNTL                   = 55
	SYS_MPX                     = 56
	SYS_SETPGID                 = 57
	SYS_ULIMIT                  = 58
	SYS_OLDOLDUNAME             = 59
	SYS_UMASK                   = 60
	SYS_CHROOT                  = 61
	SYS_USTAT                   = 62
	SYS_DUP2                    = 63
	SYS_GETPPID                 = 64
	SYS_GETPGRP                 = 65
	SYS_SETSID                  = 66
	SYS_SIGACTION               = 67
	SYS_SGETMASK                = 68
	SYS_SSETMASK                = 69
	SYS_SETREUID                = 70
	SYS_SETREGID                = 71
	SYS_SIGSUSPEND              = 72
	SYS_SIGPENDING              = 73
	SYS_SETHOSTNAME             = 74
	SYS_SETRLIMIT               = 75
	SYS_GETRLIMIT               = 76
	SYS_GETRUSAGE               = 77
	SYS_GETTIMEOFDAY            = 78
	SYS_SETTIMEOFDAY            = 79
	SYS_GETGROUPS               = 80
	SYS_SETGROUPS               = 81
	SYS_SELECT                  = 82
	SYS_SYMLINK                 = 83
	SYS_OLDLSTAT                = 84
	SYS_READLINK                = 85
	SYS_USELIB                  = 86
	SYS_SWAPON                  = 87
	SYS_REBOOT                  = 88
	SYS_READDIR                 = 89
	SYS_MMAP                    = 90
	SYS_MUNMAP                  = 91
	SYS_TRUNCATE                = 92
	SYS_FTRUNCATE               = 93
	SYS_FCHMOD                  = 94
	SYS_FCHOWN                  = 95
	SYS_GETPRIORITY             = 96
	SYS_SETPRIORITY             = 97
	SYS_PROFIL                  = 98
	SYS_STATFS                  = 99
	SYS_FSTATFS                 = 100
	SYS_IOPERM                  = 101
	SYS_SOCKETCALL              = 102
	SYS_SYSLOG                  = 103
	SYS_SETITIMER               = 104
	SYS_GETITIMER               = 105
	SYS_STAT                    = 106
	SYS_LSTAT                   = 107
	SYS_FSTAT                   = 108
	SYS_OLDUNAME                = 109
	SYS_IOPL                    = 110
	SYS_VHANGUP                 = 111
	SYS_IDLE                    = 112
	SYS_VM86                    = 113
	SYS_WAIT4                   = 114
	SYS_SWAPOFF                 = 115
	SYS_SYSINFO                 = 116
	SYS_IPC                     = 117
	SYS_FSYNC                   = 118
	SYS_SIGRETURN               = 119
	SYS_CLONE                   = 120
	SYS_SETDOMAINNAME           = 121
	SYS_UNAME                   = 122
	SYS_MODIFY_LDT              = 123
	SYS_ADJTIMEX                = 124
	SYS_MPROTECT                = 125
	SYS_SIGPROCMASK             = 126
	SYS_CREATE_MODULE           = 127
	SYS_INIT_MODULE             = 128
	SYS_DELETE_MODULE           = 129
	SYS_GET_KERNEL_SYMS         = 130
	SYS_QUOTACTL                = 131
	SYS_GETPGID                 = 132
	SYS_FCHDIR                  = 133
	SYS_BDFLUSH                 = 134
	SYS_SYSFS                   = 135
	SYS_PERSONALITY             = 136
	SYS_AFS_SYSCALL             = 137
	SYS_SETFSUID                = 138
	SYS_SETFSGID                = 139
	SYS__LLSEEK                 = 140
	SYS_GETDENTS                = 141
	SYS__NEWSELECT              = 142
	SYS_FLOCK                   = 143
	SYS_MSYNC                   = 144
	SYS_READV                   = 145
	SYS_WRITEV                  = 146
	SYS_GETSID                  = 147
	SYS_FDATASYNC               = 148
	SYS__SYSCTL                 = 149
	SYS_MLOCK                   = 150
	SYS_MUNLOCK                 = 151
	SYS_MLOCKALL                = 152
	SYS_MUNLOCKALL              = 153
	SYS_SCHED_SETPARAM          = 154
	SYS_SCHED_GETPARAM          = 155
	SYS_SCHED_SETSCHEDULER      = 156
	SYS_SCHED_GETSCHEDULER      = 157
	SYS_SCHED_YIELD             = 158
	SYS_SCHED_GET_PRIORITY_MAX  = 159
	SYS_SCHED_GET_PRIORITY_MIN  = 160
	SYS_SCHED_RR_GET_INTERVAL   = 161
	SYS_NANOSLEEP               = 162
	SYS_MREMAP                  = 163
	SYS_SETRESUID               = 164
	SYS_GETRESUID               = 165
	SYS_QUERY_MODULE            = 166
	SYS_POLL                    = 167
	SYS_NFSSERVCTL              = 168
	SYS_SETRESGID               = 169
	SYS_GETRESGID               = 170
	SYS_PRCTL                   = 171
	SYS_RT_SIGRETURN            = 172
	SYS_RT_SIGACTION            = 173
	SYS_RT_SIGPROCMASK          = 174
	SYS_RT_SIGPENDING           = 175
	SYS_RT_SIGTIMEDWAIT         = 176
	SYS_RT_SIGQUEUEINFO         = 177
	SYS_RT_SIGSUSPEND           = 178
	SYS_PREAD64                 = 179
	SYS_PWRITE64                = 180
	SYS_CHOWN                   = 181
	SYS_GETCWD                  = 182
	SYS_CAPGET                  = 183
	SYS_CAPSET                  = 184
	SYS_SIGALTSTACK             = 185
	SYS_SENDFILE                = 186
	SYS_GETPMSG                 = 187
	SYS_PUTPMSG                 = 188
	SYS_VFORK                   = 189
	SYS_UGETRLIMIT              = 190
	SYS_READAHEAD               = 191
	SYS_PCICONFIG_READ          = 198
	SYS_PCICONFIG_WRITE         = 199
	SYS_PCICONFIG_IOBASE        = 200
	SYS_MULTIPLEXER             = 201
	SYS_GETDENTS64              = 202
	SYS_PIVOT_ROOT              = 203
	SYS_MADVISE                 = 205
	SYS_MINCORE                 = 206
	SYS_GETTID                  = 207
	SYS_TKILL                   = 208
	SYS_SETXATTR                = 209
	SYS_LSETXATTR               = 210
	SYS_FSETXATTR               = 211
	SYS_GETXATTR                = 212
	SYS_LGETXATTR               = 213
	SYS_FGETXATTR               = 214
	SYS_LISTXATTR               = 215
	SYS_LLISTXATTR              = 216
	SYS_FLISTXATTR              = 217
	SYS_REMOVEXATTR             = 218
	SYS_LREMOVEXATTR            = 219
	SYS_FREMOVEXATTR            = 220
	SYS_FUTEX                   = 221
	SYS_SCHED_SETAFFINITY       = 222
	SYS_SCHED_GETAFFINITY       = 223
	SYS_TUXCALL                 = 225
	SYS_IO_SETUP                = 227
	SYS_IO_DESTROY              = 228
	SYS_IO_GETEVENTS            = 229
	SYS_IO_SUBMIT               = 230
	SYS_IO_CANCEL               = 231
	SYS_SET_TID_ADDRESS         = 232
	SYS_FADVISE64               = 233
	SYS_EXIT_GROUP              = 234
	SYS_LOOKUP_DCOOKIE          = 235
	SYS_EPOLL_CREATE            = 236
	SYS_EPOLL_CTL               = 237
	SYS_EPOLL_WAIT              = 238
	SYS_REMAP_FILE_PAGES        = 239
	SYS_TIMER_CREATE            = 240
	SYS_TIMER_SETTIME           = 241
	SYS_TIMER_GETTIME           = 242
	SYS_TIMER_GETOVERRUN        = 243
	SYS_TIMER_DELETE            = 244
	SYS_CLOCK_SETTIME           = 245
	SYS_CLOCK_GETTIME           = 246
	SYS_CLOCK_GETRES            = 247
	SYS_CLOCK_NANOSLEEP         = 248
	SYS_SWAPCONTEXT             = 249
	SYS_TGKILL                  = 250
	SYS_UTIMES                  = 251
	SYS_STATFS64                = 252
	SYS_FSTATFS64               = 253
	SYS_RTAS                    = 255
	SYS_SYS_DEBUG_SETCONTEXT    = 256
	SYS_MIGRATE_PAGES           = 258
	SYS_MBIND                   = 259
	SYS_GET_MEMPOLICY           = 260
	SYS_SET_MEMPOLICY           = 261
	SYS_MQ_OPEN                 = 262
	SYS_MQ_UNLINK               = 263
	SYS_MQ_TIMEDSEND            = 264
	SYS_MQ_TIMEDRECEIVE         = 265
	SYS_MQ_NOTIFY               = 266
	SYS_MQ_GETSETATTR           = 267
	SYS_KEXEC_LOAD              = 268
	SYS_ADD_KEY                 = 269
	SYS_REQUEST_KEY             = 270
	SYS_KEYCTL                  = 271
	SYS_WAITID                  = 272
	SYS_IOPRIO_SET              = 273
	SYS_IOPRIO_GET              = 274
	SYS_INOTIFY_INIT            = 275
	SYS_INOTIFY_ADD_WATCH       = 276
	SYS_INOTIFY_RM_WATCH        = 277
	SYS_SPU_RUN                 = 278
	SYS_SPU_CREATE              = 279
	SYS_PSELECT6                = 280
	SYS_PPOLL                   = 281
	SYS_UNSHARE                 = 282
	SYS_SPLICE                  = 283
	SYS_TEE                     = 284
	SYS_VMSPLICE                = 285
	SYS_OPENAT                  = 286
	SYS_MKDIRAT                 = 287
	SYS_MKNODAT                 = 288
	SYS_FCHOWNAT                = 289
	SYS_FUTIMESAT               = 290
	SYS_NEWFSTATAT              = 291
	SYS_UNLINKAT                = 292
	SYS_RENAMEAT                = 293
	SYS_LINKAT                  = 294
	SYS_SYMLINKAT               = 295
	SYS_READLINKAT              = 296
	SYS_FCHMODAT                = 297
	SYS_FACCESSAT               = 298
	SYS_GET_ROBUST_LIST         = 299
	SYS_SET_ROBUST_LIST         = 300
	SYS_MOVE_PAGES              = 301
	SYS_GETCPU                  = 302
	SYS_EPOLL_PWAIT             = 303
	SYS_UTIMENSAT               = 304
	SYS_SIGNALFD                = 305
	SYS_TIMERFD_CREATE          = 306
	SYS_EVENTFD                 = 307
	SYS_SYNC_FILE_RANGE2        = 308
	SYS_FALLOCATE               = 309
	SYS_SUBPAGE_PROT            = 310
	SYS_TIMERFD_SETTIME         = 311
	SYS_TIMERFD_GETTIME         = 312
	SYS_SIGNALFD4               = 313
	SYS_EVENTFD2                = 314
	SYS_EPOLL_CREATE1           = 315
	SYS_DUP3                    = 316
	SYS_PIPE2                   = 317
	SYS_INOTIFY_INIT1           = 318
	SYS_PERF_EVENT_OPEN         = 319
	SYS_PREADV                  = 320
	SYS_PWRITEV                 = 321
	SYS_RT_TGSIGQUEUEINFO       = 322
	SYS_FANOTIFY_INIT           = 323
	SYS_FANOTIFY_MARK           = 324
	SYS_PRLIMIT64               = 325
	SYS_SOCKET                  = 326
	SYS_BIND                    = 327
	SYS_CONNECT                 = 328
	SYS_LISTEN                  = 329
	SYS_ACCEPT                  = 330
	SYS_GETSOCKNAME             = 331
	SYS_GETPEERNAME             = 332
	SYS_SOCKETPAIR              = 333
	SYS_SEND                    = 334
	SYS_SENDTO                  = 335
	SYS_RECV                    = 336
	SYS_RECVFROM                = 337
	SYS_SHUTDOWN                = 338
	SYS_SETSOCKOPT              = 339
	SYS_GETSOCKOPT              = 340
	SYS_SENDMSG                 = 341
	SYS_RECVMSG                 = 342
	SYS_RECVMMSG                = 343
	SYS_ACCEPT4                 = 344
	SYS_NAME_TO_HANDLE_AT       = 345
	SYS_OPEN_BY_HANDLE_AT       = 346
	SYS_CLOCK_ADJTIME           = 347
	SYS_SYNCFS                  = 348
	SYS_SENDMMSG                = 349
	SYS_SETNS                   = 350
	SYS_PROCESS_VM_READV        = 351
	SYS_PROCESS_VM_WRITEV       = 352
	SYS_FINIT_MODULE            = 353
	SYS_KCMP                    = 354
	SYS_SCHED_SETATTR           = 355
	SYS_SCHED_GETATTR           = 356
	SYS_RENAMEAT2               = 357
	SYS_SECCOMP                 = 358
	SYS_GETRANDOM               = 359
	SYS_MEMFD_CREATE            = 360
	SYS_BPF                     = 361
	SYS_EXECVEAT                = 362
	SYS_SWITCH_ENDIAN           = 363
	SYS_USERFAULTFD             = 364
	SYS_MEMBARRIER              = 365
	SYS_MLOCK2                  = 378
	SYS_COPY_FILE_RANGE         = 379
	SYS_PREADV2                 = 380
	SYS_PWRITEV2                = 381
	SYS_KEXEC_FILE_LOAD         = 382
	SYS_STATX                   = 383
	SYS_PKEY_ALLOC              = 384
	SYS_PKEY_FREE               = 385
	SYS_PKEY_MPROTECT           = 386
	SYS_RSEQ                    = 387
	SYS_IO_PGETEVENTS           = 388
	SYS_SEMTIMEDOP              = 392
	SYS_SEMGET                  = 393
	SYS_SEMCTL                  = 394
	SYS_SHMGET                  = 395
	SYS_SHMCTL                  = 396
	SYS_SHMAT                   = 397
	SYS_SHMDT                   = 398
	SYS_MSGGET                  = 399
	SYS_MSGSND                  = 400
	SYS_MSGRCV                  = 401
	SYS_MSGCTL                  = 402
	SYS_PIDFD_SEND_SIGNAL       = 424
	SYS_IO_URING_SETUP          = 425
	SYS_IO_URING_ENTER          = 426
	SYS_IO_URING_REGISTER       = 427
	SYS_OPEN_TREE               = 428
	SYS_MOVE_MOUNT              = 429
	SYS_FSOPEN                  = 430
	SYS_FSCONFIG                = 431
	SYS_FSMOUNT                 = 432
	SYS_FSPICK                  = 433
	SYS_PIDFD_OPEN              = 434
	SYS_CLONE3                  = 435
	SYS_CLOSE_RANGE             = 436
	SYS_OPENAT2                 = 437
	SYS_PIDFD_GETFD             = 438
	SYS_FACCESSAT2              = 439
	SYS_PROCESS_MADVISE         = 440
	SYS_EPOLL_PWAIT2            = 441
	SYS_MOUNT_SETATTR           = 442
	SYS_QUOTACTL_FD             = 443
	SYS_LANDLOCK_CREATE_RULESET = 444
	SYS_LANDLOCK_ADD_RULE       = 445
	SYS_LANDLOCK_RESTRICT_SELF  = 446
)
