package FindNativeMounts

// Filesystem magic numbers defined as per /usr/include/linux/magic.h.
const (
	ADFS_SUPER_MAGIC      = 0xadf5
	AFFS_SUPER_MAGIC      = 0xadff
	AFS_SUPER_MAGIC       = 0x5346414F
	AUTOFS_SUPER_MAGIC    = 0x0187
	CODA_SUPER_MAGIC      = 0x73757245
	CRAMFS_MAGIC          = 0x28cd3d45 /* some random number */
	CRAMFS_MAGIC_WEND     = 0x453dcd28 /* magic number with the wrong endianess */
	DEBUGFS_MAGIC         = 0x64626720
	SECURITYFS_MAGIC      = 0x73636673
	SELINUX_MAGIC         = 0xf97cff8c
	SMACK_MAGIC           = 0x43415d53 /* "SMAC" */
	RAMFS_MAGIC           = 0x858458f6 /* some random number */
	TMPFS_MAGIC           = 0x01021994
	HUGETLBFS_MAGIC       = 0x958458f6 /* some random number */
	SQUASHFS_MAGIC        = 0x73717368
	ECRYPTFS_SUPER_MAGIC  = 0xf15f
	EFS_SUPER_MAGIC       = 0x414A53
	EXT2_SUPER_MAGIC      = 0xEF53
	EXT3_SUPER_MAGIC      = 0xEF53
	XENFS_SUPER_MAGIC     = 0xabba1974
	EXT4_SUPER_MAGIC      = 0xEF53
	BTRFS_SUPER_MAGIC     = 0x9123683E
	NILFS_SUPER_MAGIC     = 0x3434
	F2FS_SUPER_MAGIC      = 0xF2F52010
	HPFS_SUPER_MAGIC      = 0xf995e849
	ISOFS_SUPER_MAGIC     = 0x9660
	JFFS2_SUPER_MAGIC     = 0x72b6
	PSTOREFS_MAGIC        = 0x6165676C
	EFIVARFS_MAGIC        = 0xde5e81e4
	HOSTFS_SUPER_MAGIC    = 0x00c0ffee
	OVERLAYFS_SUPER_MAGIC = 0x794c7630

	MINIX_SUPER_MAGIC   = 0x137F /* minix v1 fs, 14 char names */
	MINIX_SUPER_MAGIC2  = 0x138F /* minix v1 fs, 30 char names */
	MINIX2_SUPER_MAGIC  = 0x2468 /* minix v2 fs, 14 char names */
	MINIX2_SUPER_MAGIC2 = 0x2478 /* minix v2 fs, 30 char names */
	MINIX3_SUPER_MAGIC  = 0x4d5a /* minix v3 fs, 60 char names */

	MSDOS_SUPER_MAGIC    = 0x4d44 /* MD */
	NCP_SUPER_MAGIC      = 0x564c /* Guess, what 0x564c is :-) */
	NFS_SUPER_MAGIC      = 0x6969
	OPENPROM_SUPER_MAGIC = 0x9fa1
	QNX4_SUPER_MAGIC     = 0x002f     /* qnx4 fs detection */
	QNX6_SUPER_MAGIC     = 0x68191122 /* qnx6 fs detection */

	REISERFS_SUPER_MAGIC = 0x52654973 /* used by gcc */
	/* used by file system utilities that
	   look at the superblock, etc.  */
	REISERFS_SUPER_MAGIC_STRING     = "ReIsErFs"
	REISER2FS_SUPER_MAGIC_STRING    = "ReIsEr2Fs"
	REISER2FS_JR_SUPER_MAGIC_STRING = "ReIsEr3Fs"

	SMB_SUPER_MAGIC     = 0x517B
	CGROUP_SUPER_MAGIC  = 0x27e0eb
	CGROUP2_SUPER_MAGIC = 0x63677270

	STACK_END_MAGIC = 0x57AC6E9D

	TRACEFS_MAGIC = 0x74726163

	V9FS_MAGIC = 0x01021997

	BDEVFS_MAGIC          = 0x62646576
	BINFMTFS_MAGIC        = 0x42494e4d
	DEVPTS_SUPER_MAGIC    = 0x1cd1
	FUTEXFS_SUPER_MAGIC   = 0xBAD1DEA
	PIPEFS_MAGIC          = 0x50495045
	PROC_SUPER_MAGIC      = 0x9fa0
	SOCKFS_MAGIC          = 0x534F434B
	SYSFS_MAGIC           = 0x62656572
	USBDEVICE_SUPER_MAGIC = 0x9fa2
	MTD_INODE_FS_MAGIC    = 0x11307854
	ANON_INODE_FS_MAGIC   = 0x09041934
	BTRFS_TEST_MAGIC      = 0x73727279
	NSFS_MAGIC            = 0x6e736673
	BPF_FS_MAGIC          = 0xcafe4a11
)
