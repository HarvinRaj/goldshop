//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package avx2

var _text_vstring = []byte{
	// .p2align 5, 0x00
	// LCPI0_0
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, // QUAD $0x2222222222222222; QUAD $0x2222222222222222  // .space 16, '""""""""""""""""'
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, //0x00000010 QUAD $0x2222222222222222; QUAD $0x2222222222222222  // .space 16, '""""""""""""""""'
	//0x00000020 LCPI0_1
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000020 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000030 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	//0x00000040 LCPI0_2
	0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, //0x00000040 QUAD $0x1f1f1f1f1f1f1f1f; QUAD $0x1f1f1f1f1f1f1f1f  // .space 16, '\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f'
	0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, //0x00000050 QUAD $0x1f1f1f1f1f1f1f1f; QUAD $0x1f1f1f1f1f1f1f1f  // .space 16, '\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f\x1f'
	//0x00000060 .p2align 4, 0x90
	//0x00000060 _vstring
	0x55,             //0x00000060 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000061 movq         %rsp, %rbp
	0x41, 0x57, //0x00000064 pushq        %r15
	0x41, 0x56, //0x00000066 pushq        %r14
	0x41, 0x55, //0x00000068 pushq        %r13
	0x41, 0x54, //0x0000006a pushq        %r12
	0x53,                   //0x0000006c pushq        %rbx
	0x48, 0x83, 0xec, 0x18, //0x0000006d subq         $24, %rsp
	0x4c, 0x8b, 0x16, //0x00000071 movq         (%rsi), %r10
	0xf6, 0xc1, 0x20, //0x00000074 testb        $32, %cl
	0x0f, 0x85, 0x2b, 0x01, 0x00, 0x00, //0x00000077 jne          LBB0_12
	0x4c, 0x8b, 0x6f, 0x08, //0x0000007d movq         $8(%rdi), %r13
	0x4c, 0x89, 0x6d, 0xc8, //0x00000081 movq         %r13, $-56(%rbp)
	0x4d, 0x29, 0xd5, //0x00000085 subq         %r10, %r13
	0x0f, 0x84, 0x8a, 0x03, 0x00, 0x00, //0x00000088 je           LBB0_41
	0x4c, 0x8b, 0x1f, //0x0000008e movq         (%rdi), %r11
	0x49, 0x83, 0xfd, 0x40, //0x00000091 cmpq         $64, %r13
	0x0f, 0x82, 0x89, 0x03, 0x00, 0x00, //0x00000095 jb           LBB0_42
	0x4c, 0x89, 0xd3, //0x0000009b movq         %r10, %rbx
	0x48, 0xf7, 0xd3, //0x0000009e notq         %rbx
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x000000a1 movq         $-1, $-48(%rbp)
	0x45, 0x31, 0xe4, //0x000000a9 xorl         %r12d, %r12d
	0xc5, 0xfe, 0x6f, 0x05, 0x4c, 0xff, 0xff, 0xff, //0x000000ac vmovdqu      $-180(%rip), %ymm0  /* LCPI0_0+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x0d, 0x64, 0xff, 0xff, 0xff, //0x000000b4 vmovdqu      $-156(%rip), %ymm1  /* LCPI0_1+0(%rip) */
	0x49, 0xbf, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, //0x000000bc movabsq      $6148914691236517205, %r15
	0x4d, 0x89, 0xd0, //0x000000c6 movq         %r10, %r8
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000000c9 .p2align 4, 0x90
	//0x000000d0 LBB0_4
	0xc4, 0x81, 0x7e, 0x6f, 0x14, 0x03, //0x000000d0 vmovdqu      (%r11,%r8), %ymm2
	0xc4, 0x81, 0x7e, 0x6f, 0x5c, 0x03, 0x20, //0x000000d6 vmovdqu      $32(%r11,%r8), %ymm3
	0xc5, 0xed, 0x74, 0xe0, //0x000000dd vpcmpeqb     %ymm0, %ymm2, %ymm4
	0xc5, 0x7d, 0xd7, 0xcc, //0x000000e1 vpmovmskb    %ymm4, %r9d
	0xc5, 0xe5, 0x74, 0xe0, //0x000000e5 vpcmpeqb     %ymm0, %ymm3, %ymm4
	0xc5, 0xfd, 0xd7, 0xcc, //0x000000e9 vpmovmskb    %ymm4, %ecx
	0xc5, 0xed, 0x74, 0xd1, //0x000000ed vpcmpeqb     %ymm1, %ymm2, %ymm2
	0xc5, 0xfd, 0xd7, 0xc2, //0x000000f1 vpmovmskb    %ymm2, %eax
	0xc5, 0xe5, 0x74, 0xd1, //0x000000f5 vpcmpeqb     %ymm1, %ymm3, %ymm2
	0xc5, 0xfd, 0xd7, 0xfa, //0x000000f9 vpmovmskb    %ymm2, %edi
	0x48, 0xc1, 0xe1, 0x20, //0x000000fd shlq         $32, %rcx
	0x49, 0x09, 0xc9, //0x00000101 orq          %rcx, %r9
	0x48, 0xc1, 0xe7, 0x20, //0x00000104 shlq         $32, %rdi
	0x48, 0x09, 0xf8, //0x00000108 orq          %rdi, %rax
	0x0f, 0x85, 0x30, 0x00, 0x00, 0x00, //0x0000010b jne          LBB0_8
	0x4d, 0x85, 0xe4, //0x00000111 testq        %r12, %r12
	0x0f, 0x85, 0x3d, 0x00, 0x00, 0x00, //0x00000114 jne          LBB0_10
	0x45, 0x31, 0xe4, //0x0000011a xorl         %r12d, %r12d
	0x4d, 0x85, 0xc9, //0x0000011d testq        %r9, %r9
	0x0f, 0x85, 0x79, 0x00, 0x00, 0x00, //0x00000120 jne          LBB0_11
	//0x00000126 LBB0_7
	0x49, 0x83, 0xc5, 0xc0, //0x00000126 addq         $-64, %r13
	0x48, 0x83, 0xc3, 0xc0, //0x0000012a addq         $-64, %rbx
	0x49, 0x83, 0xc0, 0x40, //0x0000012e addq         $64, %r8
	0x49, 0x83, 0xfd, 0x3f, //0x00000132 cmpq         $63, %r13
	0x0f, 0x87, 0x94, 0xff, 0xff, 0xff, //0x00000136 ja           LBB0_4
	0xe9, 0x20, 0x02, 0x00, 0x00, //0x0000013c jmp          LBB0_30
	//0x00000141 LBB0_8
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x00000141 cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x0b, 0x00, 0x00, 0x00, //0x00000146 jne          LBB0_10
	0x48, 0x0f, 0xbc, 0xc8, //0x0000014c bsfq         %rax, %rcx
	0x4c, 0x01, 0xc1, //0x00000150 addq         %r8, %rcx
	0x48, 0x89, 0x4d, 0xd0, //0x00000153 movq         %rcx, $-48(%rbp)
	//0x00000157 LBB0_10
	0x4c, 0x89, 0xe1, //0x00000157 movq         %r12, %rcx
	0x48, 0xf7, 0xd1, //0x0000015a notq         %rcx
	0x48, 0x21, 0xc1, //0x0000015d andq         %rax, %rcx
	0x4c, 0x8d, 0x34, 0x09, //0x00000160 leaq         (%rcx,%rcx), %r14
	0x4d, 0x09, 0xe6, //0x00000164 orq          %r12, %r14
	0x4c, 0x89, 0xf7, //0x00000167 movq         %r14, %rdi
	0x48, 0xf7, 0xd7, //0x0000016a notq         %rdi
	0x48, 0x21, 0xc7, //0x0000016d andq         %rax, %rdi
	0x48, 0xb8, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, //0x00000170 movabsq      $-6148914691236517206, %rax
	0x48, 0x21, 0xc7, //0x0000017a andq         %rax, %rdi
	0x45, 0x31, 0xe4, //0x0000017d xorl         %r12d, %r12d
	0x48, 0x01, 0xcf, //0x00000180 addq         %rcx, %rdi
	0x41, 0x0f, 0x92, 0xc4, //0x00000183 setb         %r12b
	0x48, 0x01, 0xff, //0x00000187 addq         %rdi, %rdi
	0x4c, 0x31, 0xff, //0x0000018a xorq         %r15, %rdi
	0x4c, 0x21, 0xf7, //0x0000018d andq         %r14, %rdi
	0x48, 0xf7, 0xd7, //0x00000190 notq         %rdi
	0x49, 0x21, 0xf9, //0x00000193 andq         %rdi, %r9
	0x4d, 0x85, 0xc9, //0x00000196 testq        %r9, %r9
	0x0f, 0x84, 0x87, 0xff, 0xff, 0xff, //0x00000199 je           LBB0_7
	//0x0000019f LBB0_11
	0x4d, 0x0f, 0xbc, 0xf1, //0x0000019f bsfq         %r9, %r14
	0xe9, 0x84, 0x01, 0x00, 0x00, //0x000001a3 jmp          LBB0_27
	//0x000001a8 LBB0_12
	0x4c, 0x8b, 0x6f, 0x08, //0x000001a8 movq         $8(%rdi), %r13
	0x4c, 0x89, 0x6d, 0xc8, //0x000001ac movq         %r13, $-56(%rbp)
	0x4d, 0x29, 0xd5, //0x000001b0 subq         %r10, %r13
	0x0f, 0x84, 0x5f, 0x02, 0x00, 0x00, //0x000001b3 je           LBB0_41
	0x4c, 0x8b, 0x1f, //0x000001b9 movq         (%rdi), %r11
	0x49, 0x83, 0xfd, 0x40, //0x000001bc cmpq         $64, %r13
	0x0f, 0x82, 0x7c, 0x02, 0x00, 0x00, //0x000001c0 jb           LBB0_43
	0x4c, 0x89, 0xd3, //0x000001c6 movq         %r10, %rbx
	0x48, 0xf7, 0xd3, //0x000001c9 notq         %rbx
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x000001cc movq         $-1, $-48(%rbp)
	0x45, 0x31, 0xe4, //0x000001d4 xorl         %r12d, %r12d
	0xc5, 0xfe, 0x6f, 0x05, 0x21, 0xfe, 0xff, 0xff, //0x000001d7 vmovdqu      $-479(%rip), %ymm0  /* LCPI0_0+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x0d, 0x39, 0xfe, 0xff, 0xff, //0x000001df vmovdqu      $-455(%rip), %ymm1  /* LCPI0_1+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x15, 0x51, 0xfe, 0xff, 0xff, //0x000001e7 vmovdqu      $-431(%rip), %ymm2  /* LCPI0_2+0(%rip) */
	0x4d, 0x89, 0xd1, //0x000001ef movq         %r10, %r9
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000001f2 .p2align 4, 0x90
	//0x00000200 LBB0_15
	0xc4, 0x81, 0x7e, 0x6f, 0x1c, 0x0b, //0x00000200 vmovdqu      (%r11,%r9), %ymm3
	0xc4, 0x81, 0x7e, 0x6f, 0x64, 0x0b, 0x20, //0x00000206 vmovdqu      $32(%r11,%r9), %ymm4
	0xc5, 0xe5, 0x74, 0xe8, //0x0000020d vpcmpeqb     %ymm0, %ymm3, %ymm5
	0xc5, 0xfd, 0xd7, 0xcd, //0x00000211 vpmovmskb    %ymm5, %ecx
	0xc5, 0xdd, 0x74, 0xe8, //0x00000215 vpcmpeqb     %ymm0, %ymm4, %ymm5
	0xc5, 0x7d, 0xd7, 0xfd, //0x00000219 vpmovmskb    %ymm5, %r15d
	0xc5, 0xe5, 0x74, 0xe9, //0x0000021d vpcmpeqb     %ymm1, %ymm3, %ymm5
	0xc5, 0x7d, 0xd7, 0xf5, //0x00000221 vpmovmskb    %ymm5, %r14d
	0xc5, 0xdd, 0x74, 0xe9, //0x00000225 vpcmpeqb     %ymm1, %ymm4, %ymm5
	0xc5, 0x7d, 0xd7, 0xc5, //0x00000229 vpmovmskb    %ymm5, %r8d
	0xc5, 0xdd, 0xda, 0xea, //0x0000022d vpminub      %ymm2, %ymm4, %ymm5
	0xc5, 0xdd, 0x74, 0xe5, //0x00000231 vpcmpeqb     %ymm5, %ymm4, %ymm4
	0xc5, 0xfd, 0xd7, 0xc4, //0x00000235 vpmovmskb    %ymm4, %eax
	0x49, 0xc1, 0xe7, 0x20, //0x00000239 shlq         $32, %r15
	0x4c, 0x09, 0xf9, //0x0000023d orq          %r15, %rcx
	0x49, 0xc1, 0xe0, 0x20, //0x00000240 shlq         $32, %r8
	0x48, 0xc1, 0xe0, 0x20, //0x00000244 shlq         $32, %rax
	0x4d, 0x09, 0xc6, //0x00000248 orq          %r8, %r14
	0x0f, 0x85, 0x48, 0x00, 0x00, 0x00, //0x0000024b jne          LBB0_21
	0x4d, 0x85, 0xe4, //0x00000251 testq        %r12, %r12
	0x0f, 0x85, 0x55, 0x00, 0x00, 0x00, //0x00000254 jne          LBB0_23
	0x45, 0x31, 0xe4, //0x0000025a xorl         %r12d, %r12d
	//0x0000025d LBB0_18
	0xc5, 0xe5, 0xda, 0xe2, //0x0000025d vpminub      %ymm2, %ymm3, %ymm4
	0xc5, 0xe5, 0x74, 0xdc, //0x00000261 vpcmpeqb     %ymm4, %ymm3, %ymm3
	0xc5, 0xfd, 0xd7, 0xfb, //0x00000265 vpmovmskb    %ymm3, %edi
	0x48, 0x09, 0xf8, //0x00000269 orq          %rdi, %rax
	0x48, 0x85, 0xc9, //0x0000026c testq        %rcx, %rcx
	0x0f, 0x85, 0x8a, 0x00, 0x00, 0x00, //0x0000026f jne          LBB0_24
	0x48, 0x85, 0xc0, //0x00000275 testq        %rax, %rax
	0x0f, 0x85, 0x2c, 0x04, 0x00, 0x00, //0x00000278 jne          LBB0_80
	0x49, 0x83, 0xc5, 0xc0, //0x0000027e addq         $-64, %r13
	0x48, 0x83, 0xc3, 0xc0, //0x00000282 addq         $-64, %rbx
	0x49, 0x83, 0xc1, 0x40, //0x00000286 addq         $64, %r9
	0x49, 0x83, 0xfd, 0x3f, //0x0000028a cmpq         $63, %r13
	0x0f, 0x87, 0x6c, 0xff, 0xff, 0xff, //0x0000028e ja           LBB0_15
	0xe9, 0x23, 0x01, 0x00, 0x00, //0x00000294 jmp          LBB0_35
	//0x00000299 LBB0_21
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x00000299 cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x0b, 0x00, 0x00, 0x00, //0x0000029e jne          LBB0_23
	0x49, 0x0f, 0xbc, 0xfe, //0x000002a4 bsfq         %r14, %rdi
	0x4c, 0x01, 0xcf, //0x000002a8 addq         %r9, %rdi
	0x48, 0x89, 0x7d, 0xd0, //0x000002ab movq         %rdi, $-48(%rbp)
	//0x000002af LBB0_23
	0x4d, 0x89, 0xe0, //0x000002af movq         %r12, %r8
	0x49, 0xf7, 0xd0, //0x000002b2 notq         %r8
	0x4d, 0x21, 0xf0, //0x000002b5 andq         %r14, %r8
	0x4f, 0x8d, 0x3c, 0x00, //0x000002b8 leaq         (%r8,%r8), %r15
	0x4d, 0x09, 0xe7, //0x000002bc orq          %r12, %r15
	0x4c, 0x89, 0x7d, 0xc0, //0x000002bf movq         %r15, $-64(%rbp)
	0x49, 0xf7, 0xd7, //0x000002c3 notq         %r15
	0x4d, 0x21, 0xf7, //0x000002c6 andq         %r14, %r15
	0x48, 0xbf, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, //0x000002c9 movabsq      $-6148914691236517206, %rdi
	0x49, 0x21, 0xff, //0x000002d3 andq         %rdi, %r15
	0x45, 0x31, 0xe4, //0x000002d6 xorl         %r12d, %r12d
	0x4d, 0x01, 0xc7, //0x000002d9 addq         %r8, %r15
	0x41, 0x0f, 0x92, 0xc4, //0x000002dc setb         %r12b
	0x4d, 0x01, 0xff, //0x000002e0 addq         %r15, %r15
	0x48, 0xbf, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, //0x000002e3 movabsq      $6148914691236517205, %rdi
	0x49, 0x31, 0xff, //0x000002ed xorq         %rdi, %r15
	0x4c, 0x23, 0x7d, 0xc0, //0x000002f0 andq         $-64(%rbp), %r15
	0x49, 0xf7, 0xd7, //0x000002f4 notq         %r15
	0x4c, 0x21, 0xf9, //0x000002f7 andq         %r15, %rcx
	0xe9, 0x5e, 0xff, 0xff, 0xff, //0x000002fa jmp          LBB0_18
	//0x000002ff LBB0_24
	0x4c, 0x0f, 0xbc, 0xf1, //0x000002ff bsfq         %rcx, %r14
	0x48, 0x85, 0xc0, //0x00000303 testq        %rax, %rax
	0x0f, 0x84, 0x12, 0x00, 0x00, 0x00, //0x00000306 je           LBB0_26
	0x48, 0x0f, 0xbc, 0xc0, //0x0000030c bsfq         %rax, %rax
	0x4c, 0x39, 0xf0, //0x00000310 cmpq         %r14, %rax
	0x0f, 0x83, 0x13, 0x00, 0x00, 0x00, //0x00000313 jae          LBB0_27
	0xe9, 0x8c, 0x03, 0x00, 0x00, //0x00000319 jmp          LBB0_80
	//0x0000031e LBB0_26
	0xb8, 0x40, 0x00, 0x00, 0x00, //0x0000031e movl         $64, %eax
	0x4c, 0x39, 0xf0, //0x00000323 cmpq         %r14, %rax
	0x0f, 0x82, 0x7e, 0x03, 0x00, 0x00, //0x00000326 jb           LBB0_80
	//0x0000032c LBB0_27
	0x49, 0x29, 0xde, //0x0000032c subq         %rbx, %r14
	//0x0000032f LBB0_28
	0x4d, 0x85, 0xf6, //0x0000032f testq        %r14, %r14
	0x0f, 0x88, 0x79, 0x03, 0x00, 0x00, //0x00000332 js           LBB0_81
	0x4c, 0x89, 0x36, //0x00000338 movq         %r14, (%rsi)
	0x4c, 0x89, 0x52, 0x10, //0x0000033b movq         %r10, $16(%rdx)
	0x48, 0xc7, 0x02, 0x07, 0x00, 0x00, 0x00, //0x0000033f movq         $7, (%rdx)
	0x48, 0x8b, 0x4d, 0xd0, //0x00000346 movq         $-48(%rbp), %rcx
	0x4c, 0x39, 0xf1, //0x0000034a cmpq         %r14, %rcx
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x0000034d movq         $-1, %rax
	0x48, 0x0f, 0x4c, 0xc1, //0x00000354 cmovlq       %rcx, %rax
	0x48, 0x89, 0x42, 0x18, //0x00000358 movq         %rax, $24(%rdx)
	0xe9, 0x5a, 0x03, 0x00, 0x00, //0x0000035c jmp          LBB0_83
	//0x00000361 LBB0_30
	0x4d, 0x01, 0xd8, //0x00000361 addq         %r11, %r8
	0x49, 0x83, 0xfd, 0x20, //0x00000364 cmpq         $32, %r13
	0x0f, 0x82, 0x4f, 0x01, 0x00, 0x00, //0x00000368 jb           LBB0_48
	//0x0000036e LBB0_31
	0xc4, 0xc1, 0x7e, 0x6f, 0x00, //0x0000036e vmovdqu      (%r8), %ymm0
	0xc5, 0xfd, 0x74, 0x0d, 0x85, 0xfc, 0xff, 0xff, //0x00000373 vpcmpeqb     $-891(%rip), %ymm0, %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0xfd, 0xd7, 0xf9, //0x0000037b vpmovmskb    %ymm1, %edi
	0xc5, 0xfd, 0x74, 0x05, 0x99, 0xfc, 0xff, 0xff, //0x0000037f vpcmpeqb     $-871(%rip), %ymm0, %ymm0  /* LCPI0_1+0(%rip) */
	0xc5, 0xfd, 0xd7, 0xc0, //0x00000387 vpmovmskb    %ymm0, %eax
	0x85, 0xc0, //0x0000038b testl        %eax, %eax
	0x0f, 0x85, 0xcd, 0x00, 0x00, 0x00, //0x0000038d jne          LBB0_44
	0x4d, 0x85, 0xe4, //0x00000393 testq        %r12, %r12
	0x0f, 0x85, 0xe0, 0x00, 0x00, 0x00, //0x00000396 jne          LBB0_46
	0x45, 0x31, 0xe4, //0x0000039c xorl         %r12d, %r12d
	0x48, 0x85, 0xff, //0x0000039f testq        %rdi, %rdi
	0x0f, 0x84, 0x0d, 0x01, 0x00, 0x00, //0x000003a2 je           LBB0_47
	//0x000003a8 LBB0_34
	0x48, 0x0f, 0xbc, 0xc7, //0x000003a8 bsfq         %rdi, %rax
	0x4d, 0x29, 0xd8, //0x000003ac subq         %r11, %r8
	0x4d, 0x8d, 0x34, 0x00, //0x000003af leaq         (%r8,%rax), %r14
	0x49, 0x83, 0xc6, 0x01, //0x000003b3 addq         $1, %r14
	0xe9, 0x73, 0xff, 0xff, 0xff, //0x000003b7 jmp          LBB0_28
	//0x000003bc LBB0_35
	0x4d, 0x01, 0xd9, //0x000003bc addq         %r11, %r9
	0x49, 0x83, 0xfd, 0x20, //0x000003bf cmpq         $32, %r13
	0x0f, 0x82, 0x59, 0x02, 0x00, 0x00, //0x000003c3 jb           LBB0_70
	//0x000003c9 LBB0_36
	0xc4, 0xc1, 0x7e, 0x6f, 0x01, //0x000003c9 vmovdqu      (%r9), %ymm0
	0xc5, 0xfd, 0x74, 0x0d, 0x2a, 0xfc, 0xff, 0xff, //0x000003ce vpcmpeqb     $-982(%rip), %ymm0, %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0xfd, 0xd7, 0xc1, //0x000003d6 vpmovmskb    %ymm1, %eax
	0xc5, 0xfd, 0x74, 0x0d, 0x3e, 0xfc, 0xff, 0xff, //0x000003da vpcmpeqb     $-962(%rip), %ymm0, %ymm1  /* LCPI0_1+0(%rip) */
	0xc5, 0xfd, 0xd7, 0xc9, //0x000003e2 vpmovmskb    %ymm1, %ecx
	0xc5, 0xfd, 0xda, 0x0d, 0x52, 0xfc, 0xff, 0xff, //0x000003e6 vpminub      $-942(%rip), %ymm0, %ymm1  /* LCPI0_2+0(%rip) */
	0x85, 0xc9, //0x000003ee testl        %ecx, %ecx
	0x0f, 0x85, 0x8d, 0x01, 0x00, 0x00, //0x000003f0 jne          LBB0_61
	0x4d, 0x85, 0xe4, //0x000003f6 testq        %r12, %r12
	0x0f, 0x85, 0xa0, 0x01, 0x00, 0x00, //0x000003f9 jne          LBB0_63
	0x45, 0x31, 0xe4, //0x000003ff xorl         %r12d, %r12d
	0xc5, 0xfd, 0x74, 0xc1, //0x00000402 vpcmpeqb     %ymm1, %ymm0, %ymm0
	0x48, 0x85, 0xc0, //0x00000406 testq        %rax, %rax
	0x0f, 0x84, 0xcd, 0x01, 0x00, 0x00, //0x00000409 je           LBB0_64
	//0x0000040f LBB0_39
	0x48, 0x0f, 0xbc, 0xc8, //0x0000040f bsfq         %rax, %rcx
	0xe9, 0xc9, 0x01, 0x00, 0x00, //0x00000413 jmp          LBB0_65
	//0x00000418 LBB0_41
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x00000418 movq         $-1, %r14
	0xe9, 0x91, 0x02, 0x00, 0x00, //0x0000041f jmp          LBB0_82
	//0x00000424 LBB0_42
	0x4f, 0x8d, 0x04, 0x13, //0x00000424 leaq         (%r11,%r10), %r8
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x00000428 movq         $-1, $-48(%rbp)
	0x45, 0x31, 0xe4, //0x00000430 xorl         %r12d, %r12d
	0x49, 0x83, 0xfd, 0x20, //0x00000433 cmpq         $32, %r13
	0x0f, 0x83, 0x31, 0xff, 0xff, 0xff, //0x00000437 jae          LBB0_31
	0xe9, 0x7b, 0x00, 0x00, 0x00, //0x0000043d jmp          LBB0_48
	//0x00000442 LBB0_43
	0x4f, 0x8d, 0x0c, 0x13, //0x00000442 leaq         (%r11,%r10), %r9
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x00000446 movq         $-1, $-48(%rbp)
	0x45, 0x31, 0xe4, //0x0000044e xorl         %r12d, %r12d
	0x49, 0x83, 0xfd, 0x20, //0x00000451 cmpq         $32, %r13
	0x0f, 0x83, 0x6e, 0xff, 0xff, 0xff, //0x00000455 jae          LBB0_36
	0xe9, 0xc2, 0x01, 0x00, 0x00, //0x0000045b jmp          LBB0_70
	//0x00000460 LBB0_44
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x00000460 cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x11, 0x00, 0x00, 0x00, //0x00000465 jne          LBB0_46
	0x4c, 0x89, 0xc1, //0x0000046b movq         %r8, %rcx
	0x4c, 0x29, 0xd9, //0x0000046e subq         %r11, %rcx
	0x48, 0x0f, 0xbc, 0xd8, //0x00000471 bsfq         %rax, %rbx
	0x48, 0x01, 0xcb, //0x00000475 addq         %rcx, %rbx
	0x48, 0x89, 0x5d, 0xd0, //0x00000478 movq         %rbx, $-48(%rbp)
	//0x0000047c LBB0_46
	0x44, 0x89, 0xe1, //0x0000047c movl         %r12d, %ecx
	0xf7, 0xd1, //0x0000047f notl         %ecx
	0x21, 0xc1, //0x00000481 andl         %eax, %ecx
	0x8d, 0x1c, 0x09, //0x00000483 leal         (%rcx,%rcx), %ebx
	0x45, 0x8d, 0x0c, 0x4c, //0x00000486 leal         (%r12,%rcx,2), %r9d
	0xf7, 0xd3, //0x0000048a notl         %ebx
	0x21, 0xc3, //0x0000048c andl         %eax, %ebx
	0x81, 0xe3, 0xaa, 0xaa, 0xaa, 0xaa, //0x0000048e andl         $-1431655766, %ebx
	0x45, 0x31, 0xe4, //0x00000494 xorl         %r12d, %r12d
	0x01, 0xcb, //0x00000497 addl         %ecx, %ebx
	0x41, 0x0f, 0x92, 0xc4, //0x00000499 setb         %r12b
	0x01, 0xdb, //0x0000049d addl         %ebx, %ebx
	0x81, 0xf3, 0x55, 0x55, 0x55, 0x55, //0x0000049f xorl         $1431655765, %ebx
	0x44, 0x21, 0xcb, //0x000004a5 andl         %r9d, %ebx
	0xf7, 0xd3, //0x000004a8 notl         %ebx
	0x21, 0xdf, //0x000004aa andl         %ebx, %edi
	0x48, 0x85, 0xff, //0x000004ac testq        %rdi, %rdi
	0x0f, 0x85, 0xf3, 0xfe, 0xff, 0xff, //0x000004af jne          LBB0_34
	//0x000004b5 LBB0_47
	0x49, 0x83, 0xc0, 0x20, //0x000004b5 addq         $32, %r8
	0x49, 0x83, 0xc5, 0xe0, //0x000004b9 addq         $-32, %r13
	//0x000004bd LBB0_48
	0x4d, 0x85, 0xe4, //0x000004bd testq        %r12, %r12
	0x0f, 0x85, 0x16, 0x02, 0x00, 0x00, //0x000004c0 jne          LBB0_85
	0x4c, 0x8b, 0x7d, 0xd0, //0x000004c6 movq         $-48(%rbp), %r15
	0x4d, 0x85, 0xed, //0x000004ca testq        %r13, %r13
	0x0f, 0x84, 0x8d, 0x00, 0x00, 0x00, //0x000004cd je           LBB0_58
	//0x000004d3 LBB0_50
	0x4c, 0x89, 0xdf, //0x000004d3 movq         %r11, %rdi
	0x48, 0xf7, 0xdf, //0x000004d6 negq         %rdi
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x000004d9 movq         $-1, %r14
	//0x000004e0 LBB0_51
	0x31, 0xc0, //0x000004e0 xorl         %eax, %eax
	//0x000004e2 LBB0_52
	0x41, 0x0f, 0xb6, 0x1c, 0x00, //0x000004e2 movzbl       (%r8,%rax), %ebx
	0x80, 0xfb, 0x22, //0x000004e7 cmpb         $34, %bl
	0x0f, 0x84, 0x69, 0x00, 0x00, 0x00, //0x000004ea je           LBB0_57
	0x80, 0xfb, 0x5c, //0x000004f0 cmpb         $92, %bl
	0x0f, 0x84, 0x12, 0x00, 0x00, 0x00, //0x000004f3 je           LBB0_55
	0x48, 0x83, 0xc0, 0x01, //0x000004f9 addq         $1, %rax
	0x49, 0x39, 0xc5, //0x000004fd cmpq         %rax, %r13
	0x0f, 0x85, 0xdc, 0xff, 0xff, 0xff, //0x00000500 jne          LBB0_52
	0xe9, 0x60, 0x00, 0x00, 0x00, //0x00000506 jmp          LBB0_59
	//0x0000050b LBB0_55
	0x49, 0x8d, 0x4d, 0xff, //0x0000050b leaq         $-1(%r13), %rcx
	0x48, 0x39, 0xc1, //0x0000050f cmpq         %rax, %rcx
	0x0f, 0x84, 0x99, 0x01, 0x00, 0x00, //0x00000512 je           LBB0_81
	0x4a, 0x8d, 0x0c, 0x07, //0x00000518 leaq         (%rdi,%r8), %rcx
	0x48, 0x01, 0xc1, //0x0000051c addq         %rax, %rcx
	0x49, 0x83, 0xff, 0xff, //0x0000051f cmpq         $-1, %r15
	0x48, 0x8b, 0x5d, 0xd0, //0x00000523 movq         $-48(%rbp), %rbx
	0x48, 0x0f, 0x44, 0xd9, //0x00000527 cmoveq       %rcx, %rbx
	0x48, 0x89, 0x5d, 0xd0, //0x0000052b movq         %rbx, $-48(%rbp)
	0x4c, 0x0f, 0x44, 0xf9, //0x0000052f cmoveq       %rcx, %r15
	0x49, 0x01, 0xc0, //0x00000533 addq         %rax, %r8
	0x49, 0x83, 0xc0, 0x02, //0x00000536 addq         $2, %r8
	0x4c, 0x89, 0xe9, //0x0000053a movq         %r13, %rcx
	0x48, 0x29, 0xc1, //0x0000053d subq         %rax, %rcx
	0x48, 0x83, 0xc1, 0xfe, //0x00000540 addq         $-2, %rcx
	0x49, 0x83, 0xc5, 0xfe, //0x00000544 addq         $-2, %r13
	0x49, 0x39, 0xc5, //0x00000548 cmpq         %rax, %r13
	0x49, 0x89, 0xcd, //0x0000054b movq         %rcx, %r13
	0x0f, 0x85, 0x8c, 0xff, 0xff, 0xff, //0x0000054e jne          LBB0_51
	0xe9, 0x58, 0x01, 0x00, 0x00, //0x00000554 jmp          LBB0_81
	//0x00000559 LBB0_57
	0x49, 0x01, 0xc0, //0x00000559 addq         %rax, %r8
	0x49, 0x83, 0xc0, 0x01, //0x0000055c addq         $1, %r8
	//0x00000560 LBB0_58
	0x4d, 0x29, 0xd8, //0x00000560 subq         %r11, %r8
	0x4d, 0x89, 0xc6, //0x00000563 movq         %r8, %r14
	0xe9, 0xc4, 0xfd, 0xff, 0xff, //0x00000566 jmp          LBB0_28
	//0x0000056b LBB0_59
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x0000056b movq         $-1, %r14
	0x80, 0xfb, 0x22, //0x00000572 cmpb         $34, %bl
	0x0f, 0x85, 0x36, 0x01, 0x00, 0x00, //0x00000575 jne          LBB0_81
	0x4d, 0x01, 0xe8, //0x0000057b addq         %r13, %r8
	0xe9, 0xdd, 0xff, 0xff, 0xff, //0x0000057e jmp          LBB0_58
	//0x00000583 LBB0_61
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x00000583 cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x11, 0x00, 0x00, 0x00, //0x00000588 jne          LBB0_63
	0x4c, 0x89, 0xcf, //0x0000058e movq         %r9, %rdi
	0x4c, 0x29, 0xdf, //0x00000591 subq         %r11, %rdi
	0x48, 0x0f, 0xbc, 0xd9, //0x00000594 bsfq         %rcx, %rbx
	0x48, 0x01, 0xfb, //0x00000598 addq         %rdi, %rbx
	0x48, 0x89, 0x5d, 0xd0, //0x0000059b movq         %rbx, $-48(%rbp)
	//0x0000059f LBB0_63
	0x44, 0x89, 0xe7, //0x0000059f movl         %r12d, %edi
	0xf7, 0xd7, //0x000005a2 notl         %edi
	0x21, 0xcf, //0x000005a4 andl         %ecx, %edi
	0x8d, 0x1c, 0x3f, //0x000005a6 leal         (%rdi,%rdi), %ebx
	0x45, 0x8d, 0x04, 0x7c, //0x000005a9 leal         (%r12,%rdi,2), %r8d
	0xf7, 0xd3, //0x000005ad notl         %ebx
	0x21, 0xcb, //0x000005af andl         %ecx, %ebx
	0x81, 0xe3, 0xaa, 0xaa, 0xaa, 0xaa, //0x000005b1 andl         $-1431655766, %ebx
	0x45, 0x31, 0xe4, //0x000005b7 xorl         %r12d, %r12d
	0x01, 0xfb, //0x000005ba addl         %edi, %ebx
	0x41, 0x0f, 0x92, 0xc4, //0x000005bc setb         %r12b
	0x01, 0xdb, //0x000005c0 addl         %ebx, %ebx
	0x81, 0xf3, 0x55, 0x55, 0x55, 0x55, //0x000005c2 xorl         $1431655765, %ebx
	0x44, 0x21, 0xc3, //0x000005c8 andl         %r8d, %ebx
	0xf7, 0xd3, //0x000005cb notl         %ebx
	0x21, 0xd8, //0x000005cd andl         %ebx, %eax
	0xc5, 0xfd, 0x74, 0xc1, //0x000005cf vpcmpeqb     %ymm1, %ymm0, %ymm0
	0x48, 0x85, 0xc0, //0x000005d3 testq        %rax, %rax
	0x0f, 0x85, 0x33, 0xfe, 0xff, 0xff, //0x000005d6 jne          LBB0_39
	//0x000005dc LBB0_64
	0xb9, 0x40, 0x00, 0x00, 0x00, //0x000005dc movl         $64, %ecx
	//0x000005e1 LBB0_65
	0xc5, 0xfd, 0xd7, 0xd8, //0x000005e1 vpmovmskb    %ymm0, %ebx
	0x48, 0x85, 0xc0, //0x000005e5 testq        %rax, %rax
	0x0f, 0x84, 0x24, 0x00, 0x00, 0x00, //0x000005e8 je           LBB0_68
	0x0f, 0xbc, 0xc3, //0x000005ee bsfl         %ebx, %eax
	0xbf, 0x40, 0x00, 0x00, 0x00, //0x000005f1 movl         $64, %edi
	0x0f, 0x45, 0xf8, //0x000005f6 cmovnel      %eax, %edi
	0x48, 0x39, 0xf9, //0x000005f9 cmpq         %rdi, %rcx
	0x0f, 0x87, 0xa8, 0x00, 0x00, 0x00, //0x000005fc ja           LBB0_80
	0x4d, 0x29, 0xd9, //0x00000602 subq         %r11, %r9
	0x4d, 0x8d, 0x34, 0x09, //0x00000605 leaq         (%r9,%rcx), %r14
	0x49, 0x83, 0xc6, 0x01, //0x00000609 addq         $1, %r14
	0xe9, 0x1d, 0xfd, 0xff, 0xff, //0x0000060d jmp          LBB0_28
	//0x00000612 LBB0_68
	0x85, 0xdb, //0x00000612 testl        %ebx, %ebx
	0x0f, 0x85, 0x90, 0x00, 0x00, 0x00, //0x00000614 jne          LBB0_80
	0x49, 0x83, 0xc1, 0x20, //0x0000061a addq         $32, %r9
	0x49, 0x83, 0xc5, 0xe0, //0x0000061e addq         $-32, %r13
	//0x00000622 LBB0_70
	0x4d, 0x85, 0xe4, //0x00000622 testq        %r12, %r12
	0x0f, 0x85, 0xf0, 0x00, 0x00, 0x00, //0x00000625 jne          LBB0_87
	0x48, 0x8b, 0x45, 0xd0, //0x0000062b movq         $-48(%rbp), %rax
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x0000062f movq         $-1, %r14
	0x4d, 0x85, 0xed, //0x00000636 testq        %r13, %r13
	0x0f, 0x84, 0x72, 0x00, 0x00, 0x00, //0x00000639 je           LBB0_81
	//0x0000063f LBB0_72
	0x41, 0x0f, 0xb6, 0x09, //0x0000063f movzbl       (%r9), %ecx
	0x80, 0xf9, 0x22, //0x00000643 cmpb         $34, %cl
	0x0f, 0x84, 0x81, 0x00, 0x00, 0x00, //0x00000646 je           LBB0_84
	0x80, 0xf9, 0x5c, //0x0000064c cmpb         $92, %cl
	0x0f, 0x84, 0x26, 0x00, 0x00, 0x00, //0x0000064f je           LBB0_77
	0x80, 0xf9, 0x20, //0x00000655 cmpb         $32, %cl
	0x0f, 0x82, 0x4c, 0x00, 0x00, 0x00, //0x00000658 jb           LBB0_80
	0x48, 0xc7, 0xc1, 0xff, 0xff, 0xff, 0xff, //0x0000065e movq         $-1, %rcx
	0xbb, 0x01, 0x00, 0x00, 0x00, //0x00000665 movl         $1, %ebx
	//0x0000066a LBB0_76
	0x49, 0x01, 0xd9, //0x0000066a addq         %rbx, %r9
	0x49, 0x01, 0xcd, //0x0000066d addq         %rcx, %r13
	0x0f, 0x85, 0xc9, 0xff, 0xff, 0xff, //0x00000670 jne          LBB0_72
	0xe9, 0x36, 0x00, 0x00, 0x00, //0x00000676 jmp          LBB0_81
	//0x0000067b LBB0_77
	0x49, 0x83, 0xfd, 0x01, //0x0000067b cmpq         $1, %r13
	0x0f, 0x84, 0x2c, 0x00, 0x00, 0x00, //0x0000067f je           LBB0_81
	0x48, 0xc7, 0xc1, 0xfe, 0xff, 0xff, 0xff, //0x00000685 movq         $-2, %rcx
	0xbb, 0x02, 0x00, 0x00, 0x00, //0x0000068c movl         $2, %ebx
	0x48, 0x83, 0xf8, 0xff, //0x00000691 cmpq         $-1, %rax
	0x0f, 0x85, 0xcf, 0xff, 0xff, 0xff, //0x00000695 jne          LBB0_76
	0x4c, 0x89, 0xc8, //0x0000069b movq         %r9, %rax
	0x4c, 0x29, 0xd8, //0x0000069e subq         %r11, %rax
	0x48, 0x89, 0x45, 0xd0, //0x000006a1 movq         %rax, $-48(%rbp)
	0xe9, 0xc0, 0xff, 0xff, 0xff, //0x000006a5 jmp          LBB0_76
	//0x000006aa LBB0_80
	0x49, 0xc7, 0xc6, 0xfe, 0xff, 0xff, 0xff, //0x000006aa movq         $-2, %r14
	//0x000006b1 LBB0_81
	0x4c, 0x8b, 0x55, 0xc8, //0x000006b1 movq         $-56(%rbp), %r10
	//0x000006b5 LBB0_82
	0x4c, 0x89, 0x16, //0x000006b5 movq         %r10, (%rsi)
	0x4c, 0x89, 0x32, //0x000006b8 movq         %r14, (%rdx)
	//0x000006bb LBB0_83
	0x48, 0x83, 0xc4, 0x18, //0x000006bb addq         $24, %rsp
	0x5b,       //0x000006bf popq         %rbx
	0x41, 0x5c, //0x000006c0 popq         %r12
	0x41, 0x5d, //0x000006c2 popq         %r13
	0x41, 0x5e, //0x000006c4 popq         %r14
	0x41, 0x5f, //0x000006c6 popq         %r15
	0x5d,             //0x000006c8 popq         %rbp
	0xc5, 0xf8, 0x77, //0x000006c9 vzeroupper
	0xc3, //0x000006cc retq
	//0x000006cd LBB0_84
	0x4d, 0x29, 0xd9, //0x000006cd subq         %r11, %r9
	0x49, 0x83, 0xc1, 0x01, //0x000006d0 addq         $1, %r9
	0x4d, 0x89, 0xce, //0x000006d4 movq         %r9, %r14
	0xe9, 0x53, 0xfc, 0xff, 0xff, //0x000006d7 jmp          LBB0_28
	//0x000006dc LBB0_85
	0x4d, 0x85, 0xed, //0x000006dc testq        %r13, %r13
	0x0f, 0x84, 0x7c, 0x00, 0x00, 0x00, //0x000006df je           LBB0_89
	0x4d, 0x89, 0xdf, //0x000006e5 movq         %r11, %r15
	0x49, 0xf7, 0xd7, //0x000006e8 notq         %r15
	0x4d, 0x01, 0xc7, //0x000006eb addq         %r8, %r15
	0x48, 0x8b, 0x4d, 0xd0, //0x000006ee movq         $-48(%rbp), %rcx
	0x48, 0x83, 0xf9, 0xff, //0x000006f2 cmpq         $-1, %rcx
	0x48, 0x89, 0xc8, //0x000006f6 movq         %rcx, %rax
	0x49, 0x0f, 0x44, 0xc7, //0x000006f9 cmoveq       %r15, %rax
	0x4c, 0x0f, 0x45, 0xf9, //0x000006fd cmovneq      %rcx, %r15
	0x49, 0x83, 0xc0, 0x01, //0x00000701 addq         $1, %r8
	0x49, 0x83, 0xc5, 0xff, //0x00000705 addq         $-1, %r13
	0x48, 0x89, 0x45, 0xd0, //0x00000709 movq         %rax, $-48(%rbp)
	0x4d, 0x85, 0xed, //0x0000070d testq        %r13, %r13
	0x0f, 0x85, 0xbd, 0xfd, 0xff, 0xff, //0x00000710 jne          LBB0_50
	0xe9, 0x45, 0xfe, 0xff, 0xff, //0x00000716 jmp          LBB0_58
	//0x0000071b LBB0_87
	0x4d, 0x85, 0xed, //0x0000071b testq        %r13, %r13
	0x0f, 0x84, 0x3d, 0x00, 0x00, 0x00, //0x0000071e je           LBB0_89
	0x4c, 0x89, 0xd8, //0x00000724 movq         %r11, %rax
	0x48, 0xf7, 0xd0, //0x00000727 notq         %rax
	0x4c, 0x01, 0xc8, //0x0000072a addq         %r9, %rax
	0x48, 0x8b, 0x7d, 0xd0, //0x0000072d movq         $-48(%rbp), %rdi
	0x48, 0x83, 0xff, 0xff, //0x00000731 cmpq         $-1, %rdi
	0x48, 0x89, 0xf9, //0x00000735 movq         %rdi, %rcx
	0x48, 0x0f, 0x44, 0xc8, //0x00000738 cmoveq       %rax, %rcx
	0x48, 0x0f, 0x45, 0xc7, //0x0000073c cmovneq      %rdi, %rax
	0x49, 0x83, 0xc1, 0x01, //0x00000740 addq         $1, %r9
	0x49, 0x83, 0xc5, 0xff, //0x00000744 addq         $-1, %r13
	0x48, 0x89, 0x4d, 0xd0, //0x00000748 movq         %rcx, $-48(%rbp)
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x0000074c movq         $-1, %r14
	0x4d, 0x85, 0xed, //0x00000753 testq        %r13, %r13
	0x0f, 0x85, 0xe3, 0xfe, 0xff, 0xff, //0x00000756 jne          LBB0_72
	0xe9, 0x50, 0xff, 0xff, 0xff, //0x0000075c jmp          LBB0_81
	//0x00000761 LBB0_89
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x00000761 movq         $-1, %r14
	0xe9, 0x44, 0xff, 0xff, 0xff, //0x00000768 jmp          LBB0_81
	0x00, 0x00, 0x00, //0x0000076d .p2align 2, 0x00
	//0x00000770 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x00000770 .long 2
}
