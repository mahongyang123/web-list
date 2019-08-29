
/********************************************************
LINUX NAMESPACE

主要是涉及三个系统调用
clone() –    实现线程的系统调用，用来创建一个新的进程，并可以通过设计上述参数达到隔离。
unshare() –  使某进程脱离某个namespace
setns() –    把某进程加入到某个namespace
********************************************************/
#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>

/* 定义一个给 clone 用的栈，栈大小1M */
#define STACK_SIZE (1024 * 1024)
static char container_stack[STACK_SIZE];

char* const container_args[] = {
    "/bin/bash",
    NULL
};

int container_main(void* arg)
{
    printf("Container - inside the container!\n");
    /* 直接执行一个shell，以便我们观察这个进程空间里的资源是否被隔离了 */
    execv(container_args[0], container_args); 
    printf("Something's wrong!\n");
    return 1;
}

int main()
{
    printf("Parent - start a container!\n");
    /* 调用clone函数，其中传出一个函数，还有一个栈空间的（为什么传尾指针，因为栈是反着的） */
    int container_pid = clone(container_main, container_stack+STACK_SIZE, SIGCHLD, NULL);
    /* 等待子进程结束 */
    waitpid(container_pid, NULL, 0);
    printf("Parent - container stopped!\n");
    return 0;
}


//UTS Namespace
int container_main(void* arg)
{
    printf("Container - inside the container!\n");
    sethostname("container",10);    /* 设置hostname */
    execv(container_args[0], container_args);
    printf("Something's wrong!\n");
    return 1;
}

int main()
{
    printf("Parent - start a container!\n");
    int container_pid = clone(container_main, container_stack+STACK_SIZE, 
            CLONE_NEWUTS | SIGCHLD, NULL); /*启用CLONE_NEWUTS Namespace隔离 */
    waitpid(container_pid, NULL, 0);
    printf("Parent - container stopped!\n");
    return 0;
}

//IPC Namespace
int container_pid = clone(container_main, container_stack+STACK_SIZE, 
	CLONE_NEWUTS | CLONE_NEWIPC | SIGCHLD, NULL);


//PID Namespace
int container_main(void* arg)
{
    /* 查看子进程的PID，我们可以看到其输出子进程的 pid 为 1 */
    printf("Container [%5d] - inside the container!\n", getpid());
    sethostname("container",10);
    execv(container_args[0], container_args);
    printf("Something's wrong!\n");
    return 1;
}

int main()
{
    printf("Parent [%5d] - start a container!\n", getpid());
    /*启用PID namespace - CLONE_NEWPID*/
    int container_pid = clone(container_main, container_stack+STACK_SIZE, 
            CLONE_NEWUTS | CLONE_NEWPID | SIGCHLD, NULL); 
    waitpid(container_pid, NULL, 0);
    printf("Parent - container stopped!\n");
    return 0;
}

//Mount Namespace
int container_main(void* arg)
{
    printf("Container [%5d] - inside the container!\n", getpid());
    sethostname("container",10);
    /* 重新mount proc文件系统到 /proc下 */
    system("mount -t proc proc /proc");
    execv(container_args[0], container_args);
    printf("Something's wrong!\n");
    return 1;
}

int main()
{
    printf("Parent [%5d] - start a container!\n", getpid());
    /* 启用Mount Namespace - 增加CLONE_NEWNS参数 */
    int container_pid = clone(container_main, container_stack+STACK_SIZE, 
            CLONE_NEWUTS | CLONE_NEWPID | CLONE_NEWNS | SIGCHLD, NULL);
    waitpid(container_pid, NULL, 0);
    printf("Parent - container stopped!\n");
    return 0;
}

//容器配置
#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/mount.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>

#define STACK_SIZE (1024 * 1024)

static char container_stack[STACK_SIZE];
char* const container_args[] = {
	"/bin/bash",
	"-l",
	NULL
};

int container_main(void* arg)
{
	printf("Container [%5d] - inside the container!\n", getpid());

	//set hostname
	sethostname("container",10);

	//remount "/proc" to make sure the "top" and "ps" show container's information
	if (mount("proc", "rootfs/proc", "proc", 0, NULL) !=0 ) {
		perror("proc");
	}
	if (mount("sysfs", "rootfs/sys", "sysfs", 0, NULL)!=0) {
		perror("sys");
	}
	if (mount("none", "rootfs/tmp", "tmpfs", 0, NULL)!=0) {
		perror("tmp");
	}
	if (mount("udev", "rootfs/dev", "devtmpfs", 0, NULL)!=0) {
		perror("dev");
	}
	if (mount("devpts", "rootfs/dev/pts", "devpts", 0, NULL)!=0) {
		perror("dev/pts");
	}
	if (mount("shm", "rootfs/dev/shm", "tmpfs", 0, NULL)!=0) {
		perror("dev/shm");
	}
	if (mount("tmpfs", "rootfs/run", "tmpfs", 0, NULL)!=0) {
		perror("run");
	}
	/*
		* 模仿Docker的从外向容器里mount相关的配置文件
		* 你可以查看：/var/lib/docker/containers/<container_id>/目录，
		* 你会看到docker的这些文件的。
		*/
	if (mount("conf/hosts", "rootfs/etc/hosts", "none", MS_BIND, NULL)!=0 ||
			mount("conf/hostname", "rootfs/etc/hostname", "none", MS_BIND, NULL)!=0 ||
			mount("conf/resolv.conf", "rootfs/etc/resolv.conf", "none", MS_BIND, NULL)!=0 ) {
		perror("conf");
	}
	/* 模仿docker run命令中的 -v, --volume=[] 参数干的事 */
	if (mount("/tmp/t1", "rootfs/mnt", "none", MS_BIND, NULL)!=0) {
		perror("mnt");
	}

	/* chroot 隔离目录 */
	if ( chdir("./rootfs") != 0 || chroot("./") != 0 ){
		perror("chdir/chroot");
	}

	execv(container_args[0], container_args);
	perror("exec");
	printf("Something's wrong!\n");
	return 1;
}

int main()
{
	printf("Parent [%5d] - start a container!\n", getpid());
	int container_pid = clone(container_main, container_stack+STACK_SIZE, 
			CLONE_NEWUTS | CLONE_NEWIPC | CLONE_NEWPID | CLONE_NEWNS | SIGCHLD, NULL);
	waitpid(container_pid, NULL, 0);
	printf("Parent - container stopped!\n");
	return 0;
}