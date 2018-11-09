;Text.asm	第九章：基址-变址操作数P292

.386
.model flat,stdcall
.stack 4096

ExitProcess PROTO, dwExitCode:DWORD

.data
array WORD 1000h, 2000h, 3000h

.code
main PROC
	mov ebx,OFFSET array
	mov esi,2
	mov ax,[ebx+esi]

	mov edi,OFFSET array
	mov ecx,4
	mov ax,[edi+ecx]

	mov ebp,OFFSET array
	mov esi,0
	mov ax,[ebp+esi]

	INVOKE ExitProcess, 0
main ENDP
END main
