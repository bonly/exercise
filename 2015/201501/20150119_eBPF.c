#include <errno.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include <linux/version.h>

#include <bcc/bpf_common.h>
#include <bcc/libbpf.h>

int main() {
	int map_fd, prog_fd, key=0, ret;
	long long value;
	char log_buf[8192];
	void *kprobe;

	/* Map size is 1 since we store only one value, the chown count */
	map_fd = bpf_create_map(BPF_MAP_TYPE_HASH, sizeof(key), sizeof(value), 1);
	if (map_fd < 0) {
		fprintf(stderr, "failed to create map: %s (ret %d)\n", strerror(errno), map_fd);
		return 1;
	}

	ret = bpf_update_elem(map_fd, &key, &value, 0);
	if (ret != 0) {
		fprintf(stderr, "failed to initialize map: %s (ret %d)\n", strerror(errno), ret);
		return 1;
	}

	struct bpf_insn prog[] = {
		/* Put 0 (the map key) on the stack */
		BPF_ST_MEM(BPF_W, BPF_REG_10, -4, 0),
		/* Put frame pointer into R2 */
		BPF_MOV64_REG(BPF_REG_2, BPF_REG_10),
		/* Decrement pointer by four */
		BPF_ALU64_IMM(BPF_ADD, BPF_REG_2, -4),
		/* Put map_fd into R1 */
		BPF_LD_MAP_FD(BPF_REG_1, map_fd),
		/* Load current count from map into R0 */
		BPF_RAW_INSN(BPF_JMP | BPF_CALL, 0, 0, 0,
			     BPF_FUNC_map_lookup_elem),
		/* If returned value NULL, skip two instructions and return */
		BPF_JMP_IMM(BPF_JEQ, BPF_REG_0, 0, 2),
		/* Put 1 into R1 */
		BPF_MOV64_IMM(BPF_REG_1, 1),
		/* Increment value by 1 */
		BPF_RAW_INSN(BPF_STX | BPF_XADD | BPF_DW, BPF_REG_0, BPF_REG_1, 0, 0),
		/* Return from program */
		BPF_EXIT_INSN(),
	};

	prog_fd = bpf_prog_load(BPF_PROG_TYPE_KPROBE, prog, sizeof(prog), "GPL", LINUX_VERSION_CODE, log_buf, sizeof(log_buf));
	if (prog_fd < 0) {
		fprintf(stderr, "failed to load prog: %s (ret %d)\ngot CAP_SYS_ADMIN?\n%s\n", strerror(errno), prog_fd, log_buf);
		return 1;
	}

	kprobe = bpf_attach_kprobe(prog_fd, "p_sys_fchownat", "p:kprobes/p_sys_fchownat sys_fchownat", -1, 0, -1, NULL, NULL);
	if (kprobe == NULL) {
		fprintf(stderr, "failed to attach kprobe: %s\n", strerror(errno));
		return 1;
	}

	for (;;) {
		ret = bpf_lookup_elem(map_fd, &key, &value);
		if (ret != 0) {
			fprintf(stderr, "failed to lookup element: %s (ret %d)\n", strerror(errno), ret);
		} else {
			printf("fchownat(2) count: %lld\n", value);
		}
		sleep(1);
	}

	return 0;
}
/*
gcc -I/usr/include/bcc/compat main.c -o chowncount -lbcc
*/
