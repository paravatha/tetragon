// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

//go:build arm64 && linux
// +build arm64,linux

package syscallinfo

import (
	"golang.org/x/sys/unix"
)

var syscallNames = map[int]string{
	unix.SYS_READ:                    "sys_read",
	unix.SYS_WRITE:                   "sys_write",
	unix.SYS_CLOSE:                   "sys_close",
	unix.SYS_FSTAT:                   "sys_fstat",
	unix.SYS_LSEEK:                   "sys_lseek",
	unix.SYS_MMAP:                    "sys_mmap",
	unix.SYS_MPROTECT:                "sys_mprotect",
	unix.SYS_MUNMAP:                  "sys_munmap",
	unix.SYS_BRK:                     "sys_brk",
	unix.SYS_RT_SIGACTION:            "sys_rt_sigaction",
	unix.SYS_RT_SIGPROCMASK:          "sys_rt_sigprocmask",
	unix.SYS_RT_SIGRETURN:            "sys_rt_sigreturn",
	unix.SYS_IOCTL:                   "sys_ioctl",
	unix.SYS_PREAD64:                 "sys_pread64",
	unix.SYS_PWRITE64:                "sys_pwrite64",
	unix.SYS_READV:                   "sys_readv",
	unix.SYS_WRITEV:                  "sys_writev",
	unix.SYS_SCHED_YIELD:             "sys_sched_yield",
	unix.SYS_MREMAP:                  "sys_mremap",
	unix.SYS_MSYNC:                   "sys_msync",
	unix.SYS_MINCORE:                 "sys_mincore",
	unix.SYS_MADVISE:                 "sys_madvise",
	unix.SYS_SHMGET:                  "sys_shmget",
	unix.SYS_SHMAT:                   "sys_shmat",
	unix.SYS_SHMCTL:                  "sys_shmctl",
	unix.SYS_DUP:                     "sys_dup",
	unix.SYS_NANOSLEEP:               "sys_nanosleep",
	unix.SYS_GETITIMER:               "sys_getitimer",
	unix.SYS_SETITIMER:               "sys_setitimer",
	unix.SYS_GETPID:                  "sys_getpid",
	unix.SYS_SENDFILE:                "sys_sendfile",
	unix.SYS_SOCKET:                  "sys_socket",
	unix.SYS_CONNECT:                 "sys_connect",
	unix.SYS_ACCEPT:                  "sys_accept",
	unix.SYS_SENDTO:                  "sys_sendto",
	unix.SYS_RECVFROM:                "sys_recvfrom",
	unix.SYS_SENDMSG:                 "sys_sendmsg",
	unix.SYS_RECVMSG:                 "sys_recvmsg",
	unix.SYS_SHUTDOWN:                "sys_shutdown",
	unix.SYS_BIND:                    "sys_bind",
	unix.SYS_LISTEN:                  "sys_listen",
	unix.SYS_GETSOCKNAME:             "sys_getsockname",
	unix.SYS_GETPEERNAME:             "sys_getpeername",
	unix.SYS_SOCKETPAIR:              "sys_socketpair",
	unix.SYS_SETSOCKOPT:              "sys_setsockopt",
	unix.SYS_GETSOCKOPT:              "sys_getsockopt",
	unix.SYS_CLONE:                   "sys_clone",
	unix.SYS_EXECVE:                  "sys_execve",
	unix.SYS_EXIT:                    "sys_exit",
	unix.SYS_WAIT4:                   "sys_wait4",
	unix.SYS_KILL:                    "sys_kill",
	unix.SYS_UNAME:                   "sys_uname",
	unix.SYS_SEMGET:                  "sys_semget",
	unix.SYS_SEMOP:                   "sys_semop",
	unix.SYS_SEMCTL:                  "sys_semctl",
	unix.SYS_SHMDT:                   "sys_shmdt",
	unix.SYS_MSGGET:                  "sys_msgget",
	unix.SYS_MSGSND:                  "sys_msgsnd",
	unix.SYS_MSGRCV:                  "sys_msgrcv",
	unix.SYS_MSGCTL:                  "sys_msgctl",
	unix.SYS_FCNTL:                   "sys_fcntl",
	unix.SYS_FLOCK:                   "sys_flock",
	unix.SYS_FSYNC:                   "sys_fsync",
	unix.SYS_FDATASYNC:               "sys_fdatasync",
	unix.SYS_TRUNCATE:                "sys_truncate",
	unix.SYS_FTRUNCATE:               "sys_ftruncate",
	unix.SYS_GETCWD:                  "sys_getcwd",
	unix.SYS_CHDIR:                   "sys_chdir",
	unix.SYS_FCHDIR:                  "sys_fchdir",
	unix.SYS_FCHMOD:                  "sys_fchmod",
	unix.SYS_FCHOWN:                  "sys_fchown",
	unix.SYS_UMASK:                   "sys_umask",
	unix.SYS_GETTIMEOFDAY:            "sys_gettimeofday",
	unix.SYS_GETRLIMIT:               "sys_getrlimit",
	unix.SYS_GETRUSAGE:               "sys_getrusage",
	unix.SYS_SYSINFO:                 "sys_sysinfo",
	unix.SYS_TIMES:                   "sys_times",
	unix.SYS_PTRACE:                  "sys_ptrace",
	unix.SYS_GETUID:                  "sys_getuid",
	unix.SYS_SYSLOG:                  "sys_syslog",
	unix.SYS_GETGID:                  "sys_getgid",
	unix.SYS_SETUID:                  "sys_setuid",
	unix.SYS_SETGID:                  "sys_setgid",
	unix.SYS_GETEUID:                 "sys_geteuid",
	unix.SYS_GETEGID:                 "sys_getegid",
	unix.SYS_SETPGID:                 "sys_setpgid",
	unix.SYS_GETPPID:                 "sys_getppid",
	unix.SYS_SETSID:                  "sys_setsid",
	unix.SYS_SETREUID:                "sys_setreuid",
	unix.SYS_SETREGID:                "sys_setregid",
	unix.SYS_GETGROUPS:               "sys_getgroups",
	unix.SYS_SETGROUPS:               "sys_setgroups",
	unix.SYS_SETRESUID:               "sys_setresuid",
	unix.SYS_GETRESUID:               "sys_getresuid",
	unix.SYS_SETRESGID:               "sys_setresgid",
	unix.SYS_GETRESGID:               "sys_getresgid",
	unix.SYS_GETPGID:                 "sys_getpgid",
	unix.SYS_SETFSUID:                "sys_setfsuid",
	unix.SYS_SETFSGID:                "sys_setfsgid",
	unix.SYS_GETSID:                  "sys_getsid",
	unix.SYS_CAPGET:                  "sys_capget",
	unix.SYS_CAPSET:                  "sys_capset",
	unix.SYS_RT_SIGPENDING:           "sys_rt_sigpending",
	unix.SYS_RT_SIGTIMEDWAIT:         "sys_rt_sigtimedwait",
	unix.SYS_RT_SIGQUEUEINFO:         "sys_rt_sigqueueinfo",
	unix.SYS_RT_SIGSUSPEND:           "sys_rt_sigsuspend",
	unix.SYS_SIGALTSTACK:             "sys_sigaltstack",
	unix.SYS_PERSONALITY:             "sys_personality",
	unix.SYS_STATFS:                  "sys_statfs",
	unix.SYS_FSTATFS:                 "sys_fstatfs",
	unix.SYS_GETPRIORITY:             "sys_getpriority",
	unix.SYS_SETPRIORITY:             "sys_setpriority",
	unix.SYS_SCHED_SETPARAM:          "sys_sched_setparam",
	unix.SYS_SCHED_GETPARAM:          "sys_sched_getparam",
	unix.SYS_SCHED_SETSCHEDULER:      "sys_sched_setscheduler",
	unix.SYS_SCHED_GETSCHEDULER:      "sys_sched_getscheduler",
	unix.SYS_SCHED_GET_PRIORITY_MAX:  "sys_sched_get_priority_max",
	unix.SYS_SCHED_GET_PRIORITY_MIN:  "sys_sched_get_priority_min",
	unix.SYS_SCHED_RR_GET_INTERVAL:   "sys_sched_rr_get_interval",
	unix.SYS_MLOCK:                   "sys_mlock",
	unix.SYS_MUNLOCK:                 "sys_munlock",
	unix.SYS_MLOCKALL:                "sys_mlockall",
	unix.SYS_MUNLOCKALL:              "sys_munlockall",
	unix.SYS_VHANGUP:                 "sys_vhangup",
	unix.SYS_PIVOT_ROOT:              "sys_pivot_root",
	unix.SYS_PRCTL:                   "sys_prctl",
	unix.SYS_ADJTIMEX:                "sys_adjtimex",
	unix.SYS_SETRLIMIT:               "sys_setrlimit",
	unix.SYS_CHROOT:                  "sys_chroot",
	unix.SYS_SYNC:                    "sys_sync",
	unix.SYS_ACCT:                    "sys_acct",
	unix.SYS_SETTIMEOFDAY:            "sys_settimeofday",
	unix.SYS_MOUNT:                   "sys_mount",
	unix.SYS_UMOUNT2:                 "sys_umount2",
	unix.SYS_SWAPON:                  "sys_swapon",
	unix.SYS_SWAPOFF:                 "sys_swapoff",
	unix.SYS_REBOOT:                  "sys_reboot",
	unix.SYS_SETHOSTNAME:             "sys_sethostname",
	unix.SYS_SETDOMAINNAME:           "sys_setdomainname",
	unix.SYS_INIT_MODULE:             "sys_init_module",
	unix.SYS_DELETE_MODULE:           "sys_delete_module",
	unix.SYS_QUOTACTL:                "sys_quotactl",
	unix.SYS_NFSSERVCTL:              "sys_nfsservctl",
	unix.SYS_GETTID:                  "sys_gettid",
	unix.SYS_READAHEAD:               "sys_readahead",
	unix.SYS_SETXATTR:                "sys_setxattr",
	unix.SYS_LSETXATTR:               "sys_lsetxattr",
	unix.SYS_FSETXATTR:               "sys_fsetxattr",
	unix.SYS_GETXATTR:                "sys_getxattr",
	unix.SYS_LGETXATTR:               "sys_lgetxattr",
	unix.SYS_FGETXATTR:               "sys_fgetxattr",
	unix.SYS_LISTXATTR:               "sys_listxattr",
	unix.SYS_LLISTXATTR:              "sys_llistxattr",
	unix.SYS_FLISTXATTR:              "sys_flistxattr",
	unix.SYS_REMOVEXATTR:             "sys_removexattr",
	unix.SYS_LREMOVEXATTR:            "sys_lremovexattr",
	unix.SYS_FREMOVEXATTR:            "sys_fremovexattr",
	unix.SYS_TKILL:                   "sys_tkill",
	unix.SYS_FUTEX:                   "sys_futex",
	unix.SYS_SCHED_SETAFFINITY:       "sys_sched_setaffinity",
	unix.SYS_SCHED_GETAFFINITY:       "sys_sched_getaffinity",
	unix.SYS_IO_SETUP:                "sys_io_setup",
	unix.SYS_IO_DESTROY:              "sys_io_destroy",
	unix.SYS_IO_GETEVENTS:            "sys_io_getevents",
	unix.SYS_IO_SUBMIT:               "sys_io_submit",
	unix.SYS_IO_CANCEL:               "sys_io_cancel",
	unix.SYS_LOOKUP_DCOOKIE:          "sys_lookup_dcookie",
	unix.SYS_REMAP_FILE_PAGES:        "sys_remap_file_pages",
	unix.SYS_GETDENTS64:              "sys_getdents64",
	unix.SYS_SET_TID_ADDRESS:         "sys_set_tid_address",
	unix.SYS_RESTART_SYSCALL:         "sys_restart_syscall",
	unix.SYS_SEMTIMEDOP:              "sys_semtimedop",
	unix.SYS_FADVISE64:               "sys_fadvise64",
	unix.SYS_TIMER_CREATE:            "sys_timer_create",
	unix.SYS_TIMER_SETTIME:           "sys_timer_settime",
	unix.SYS_TIMER_GETTIME:           "sys_timer_gettime",
	unix.SYS_TIMER_GETOVERRUN:        "sys_timer_getoverrun",
	unix.SYS_TIMER_DELETE:            "sys_timer_delete",
	unix.SYS_CLOCK_SETTIME:           "sys_clock_settime",
	unix.SYS_CLOCK_GETTIME:           "sys_clock_gettime",
	unix.SYS_CLOCK_GETRES:            "sys_clock_getres",
	unix.SYS_CLOCK_NANOSLEEP:         "sys_clock_nanosleep",
	unix.SYS_EXIT_GROUP:              "sys_exit_group",
	unix.SYS_EPOLL_CTL:               "sys_epoll_ctl",
	unix.SYS_TGKILL:                  "sys_tgkill",
	unix.SYS_MBIND:                   "sys_mbind",
	unix.SYS_SET_MEMPOLICY:           "sys_set_mempolicy",
	unix.SYS_GET_MEMPOLICY:           "sys_get_mempolicy",
	unix.SYS_MQ_OPEN:                 "sys_mq_open",
	unix.SYS_MQ_UNLINK:               "sys_mq_unlink",
	unix.SYS_MQ_TIMEDSEND:            "sys_mq_timedsend",
	unix.SYS_MQ_TIMEDRECEIVE:         "sys_mq_timedreceive",
	unix.SYS_MQ_NOTIFY:               "sys_mq_notify",
	unix.SYS_MQ_GETSETATTR:           "sys_mq_getsetattr",
	unix.SYS_KEXEC_LOAD:              "sys_kexec_load",
	unix.SYS_WAITID:                  "sys_waitid",
	unix.SYS_ADD_KEY:                 "sys_add_key",
	unix.SYS_REQUEST_KEY:             "sys_request_key",
	unix.SYS_KEYCTL:                  "sys_keyctl",
	unix.SYS_IOPRIO_SET:              "sys_ioprio_set",
	unix.SYS_IOPRIO_GET:              "sys_ioprio_get",
	unix.SYS_INOTIFY_ADD_WATCH:       "sys_inotify_add_watch",
	unix.SYS_INOTIFY_RM_WATCH:        "sys_inotify_rm_watch",
	unix.SYS_MIGRATE_PAGES:           "sys_migrate_pages",
	unix.SYS_OPENAT:                  "sys_openat",
	unix.SYS_MKDIRAT:                 "sys_mkdirat",
	unix.SYS_MKNODAT:                 "sys_mknodat",
	unix.SYS_FCHOWNAT:                "sys_fchownat",
	unix.SYS_UNLINKAT:                "sys_unlinkat",
	unix.SYS_RENAMEAT:                "sys_renameat",
	unix.SYS_LINKAT:                  "sys_linkat",
	unix.SYS_SYMLINKAT:               "sys_symlinkat",
	unix.SYS_READLINKAT:              "sys_readlinkat",
	unix.SYS_FCHMODAT:                "sys_fchmodat",
	unix.SYS_FACCESSAT:               "sys_faccessat",
	unix.SYS_PSELECT6:                "sys_pselect6",
	unix.SYS_PPOLL:                   "sys_ppoll",
	unix.SYS_UNSHARE:                 "sys_unshare",
	unix.SYS_SET_ROBUST_LIST:         "sys_set_robust_list",
	unix.SYS_GET_ROBUST_LIST:         "sys_get_robust_list",
	unix.SYS_SPLICE:                  "sys_splice",
	unix.SYS_TEE:                     "sys_tee",
	unix.SYS_SYNC_FILE_RANGE:         "sys_sync_file_range",
	unix.SYS_VMSPLICE:                "sys_vmsplice",
	unix.SYS_MOVE_PAGES:              "sys_move_pages",
	unix.SYS_UTIMENSAT:               "sys_utimensat",
	unix.SYS_EPOLL_PWAIT:             "sys_epoll_pwait",
	unix.SYS_TIMERFD_CREATE:          "sys_timerfd_create",
	unix.SYS_FALLOCATE:               "sys_fallocate",
	unix.SYS_TIMERFD_SETTIME:         "sys_timerfd_settime",
	unix.SYS_TIMERFD_GETTIME:         "sys_timerfd_gettime",
	unix.SYS_ACCEPT4:                 "sys_accept4",
	unix.SYS_SIGNALFD4:               "sys_signalfd4",
	unix.SYS_EVENTFD2:                "sys_eventfd2",
	unix.SYS_EPOLL_CREATE1:           "sys_epoll_create1",
	unix.SYS_DUP3:                    "sys_dup3",
	unix.SYS_PIPE2:                   "sys_pipe2",
	unix.SYS_INOTIFY_INIT1:           "sys_inotify_init1",
	unix.SYS_PREADV:                  "sys_preadv",
	unix.SYS_PWRITEV:                 "sys_pwritev",
	unix.SYS_RT_TGSIGQUEUEINFO:       "sys_rt_tgsigqueueinfo",
	unix.SYS_PERF_EVENT_OPEN:         "sys_perf_event_open",
	unix.SYS_RECVMMSG:                "sys_recvmmsg",
	unix.SYS_FANOTIFY_INIT:           "sys_fanotify_init",
	unix.SYS_FANOTIFY_MARK:           "sys_fanotify_mark",
	unix.SYS_PRLIMIT64:               "sys_prlimit64",
	unix.SYS_NAME_TO_HANDLE_AT:       "sys_name_to_handle_at",
	unix.SYS_OPEN_BY_HANDLE_AT:       "sys_open_by_handle_at",
	unix.SYS_CLOCK_ADJTIME:           "sys_clock_adjtime",
	unix.SYS_SYNCFS:                  "sys_syncfs",
	unix.SYS_SENDMMSG:                "sys_sendmmsg",
	unix.SYS_SETNS:                   "sys_setns",
	unix.SYS_GETCPU:                  "sys_getcpu",
	unix.SYS_PROCESS_VM_READV:        "sys_process_vm_readv",
	unix.SYS_PROCESS_VM_WRITEV:       "sys_process_vm_writev",
	unix.SYS_KCMP:                    "sys_kcmp",
	unix.SYS_FINIT_MODULE:            "sys_finit_module",
	unix.SYS_SCHED_SETATTR:           "sys_sched_setattr",
	unix.SYS_SCHED_GETATTR:           "sys_sched_getattr",
	unix.SYS_RENAMEAT2:               "sys_renameat2",
	unix.SYS_SECCOMP:                 "sys_seccomp",
	unix.SYS_GETRANDOM:               "sys_getrandom",
	unix.SYS_MEMFD_CREATE:            "sys_memfd_create",
	unix.SYS_KEXEC_FILE_LOAD:         "sys_kexec_file_load",
	unix.SYS_BPF:                     "sys_bpf",
	unix.SYS_EXECVEAT:                "sys_execveat",
	unix.SYS_USERFAULTFD:             "sys_userfaultfd",
	unix.SYS_MEMBARRIER:              "sys_membarrier",
	unix.SYS_MLOCK2:                  "sys_mlock2",
	unix.SYS_COPY_FILE_RANGE:         "sys_copy_file_range",
	unix.SYS_PREADV2:                 "sys_preadv2",
	unix.SYS_PWRITEV2:                "sys_pwritev2",
	unix.SYS_PKEY_MPROTECT:           "sys_pkey_mprotect",
	unix.SYS_PKEY_ALLOC:              "sys_pkey_alloc",
	unix.SYS_PKEY_FREE:               "sys_pkey_free",
	unix.SYS_STATX:                   "sys_statx",
	unix.SYS_IO_PGETEVENTS:           "sys_io_pgetevents",
	unix.SYS_RSEQ:                    "sys_rseq",
	unix.SYS_PIDFD_SEND_SIGNAL:       "sys_pidfd_send_signal",
	unix.SYS_IO_URING_SETUP:          "sys_io_uring_setup",
	unix.SYS_IO_URING_ENTER:          "sys_io_uring_enter",
	unix.SYS_IO_URING_REGISTER:       "sys_io_uring_register",
	unix.SYS_OPEN_TREE:               "sys_open_tree",
	unix.SYS_MOVE_MOUNT:              "sys_move_mount",
	unix.SYS_FSOPEN:                  "sys_fsopen",
	unix.SYS_FSCONFIG:                "sys_fsconfig",
	unix.SYS_FSMOUNT:                 "sys_fsmount",
	unix.SYS_FSPICK:                  "sys_fspick",
	unix.SYS_PIDFD_OPEN:              "sys_pidfd_open",
	unix.SYS_CLONE3:                  "sys_clone3",
	unix.SYS_CLOSE_RANGE:             "sys_close_range",
	unix.SYS_OPENAT2:                 "sys_openat2",
	unix.SYS_PIDFD_GETFD:             "sys_pidfd_getfd",
	unix.SYS_FACCESSAT2:              "sys_faccessat2",
	unix.SYS_PROCESS_MADVISE:         "sys_process_madvise",
	unix.SYS_EPOLL_PWAIT2:            "sys_epoll_pwait2",
	unix.SYS_MOUNT_SETATTR:           "sys_mount_setattr",
	unix.SYS_QUOTACTL_FD:             "sys_quotactl_fd",
	unix.SYS_LANDLOCK_CREATE_RULESET: "sys_landlock_create_ruleset",
	unix.SYS_LANDLOCK_ADD_RULE:       "sys_landlock_add_rule",
	unix.SYS_LANDLOCK_RESTRICT_SELF:  "sys_landlock_restrict_self",
	unix.SYS_MEMFD_SECRET:            "sys_memfd_secret",
	unix.SYS_PROCESS_MRELEASE:        "sys_process_mrelease",
	// unix.SYS_FUTEX_WAITV:             "sys_futex_waitv",
	// unix.SYS_SET_MEMPOLICY_HOME_NODE: "sys_set_mempolicy_home_node",
}
