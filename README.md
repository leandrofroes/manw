# manw

manw is a very simple and fast command line search engine for Windows API functions written Go. The idea is basically scrapy MSDN and retrieve useful informations about a specific Windows API function.

## Why?

On Linux systems we are able to search for the documentation of a function using man command, but what about Windows functions? Open the browser and read the full documentation is the best idea to completely understand a function but sometimes the only thing we need is an overview of it and not the full documentation.

## **Installation**

```
git clone https://github.com/leandrofroes/manw
cd manw
make
```

## **Usage**

```
./manw <function_name>
```

## **Examples**

```
$ ./manw createprocessa
CreateProcessA function (processthreadsapi.h) - Win32 apps

Creates a new process and its primary thread. The new process runs in the security context of the calling process.

BOOL CreateProcessA(
  LPCSTR                lpApplicationName,
  LPSTR                 lpCommandLine,
  LPSECURITY_ATTRIBUTES lpProcessAttributes,
  LPSECURITY_ATTRIBUTES lpThreadAttributes,
  BOOL                  bInheritHandles,
  DWORD                 dwCreationFlags,
  LPVOID                lpEnvironment,
  LPCSTR                lpCurrentDirectory,
  LPSTARTUPINFOA        lpStartupInfo,
  LPPROCESS_INFORMATION lpProcessInformation
);


Return value: If the function succeeds, the return value is nonzero. If the function fails, the return value is zero. Note that the function returns before the process has finished initialization. If a required DLL cannot be located or fails to initialize, the process is terminated.

Example code:

        LPTSTR szCmdline = _tcsdup(TEXT("C:\\Program Files\\MyApp -L -S"));
        CreateProcess(NULL, szCmdline, /* ... */);

Source: https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessa
```

```
$ ./manw createfile
CreateFileA function (fileapi.h) - Win32 apps

Creates or opens a file or I/O device. The most commonly used I/O devices are as follows:\_file, file stream, directory, physical disk, volume, console buffer, tape drive, communications resource, mailslot, and pipe.

HANDLE CreateFileA(
  LPCSTR                lpFileName,
  DWORD                 dwDesiredAccess,
  DWORD                 dwShareMode,
  LPSECURITY_ATTRIBUTES lpSecurityAttributes,
  DWORD                 dwCreationDisposition,
  DWORD                 dwFlagsAndAttributes,
  HANDLE                hTemplateFile
);


Return value: If the function succeeds, the return value is an open handle to the specified file, device, named pipe, or mail slot. If the function fails, the return value is INVALID_HANDLE_VALUE.

Source: https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilea
```

You can search for Windows Structures too:

```
$ ./manw peb
PEB (winternl.h) - Win32 apps

Contains process information.

typedef struct _PEB {
  BYTE                          Reserved1[2];
  BYTE                          BeingDebugged;
  BYTE                          Reserved2[1];
  PVOID                         Reserved3[2];
  PPEB_LDR_DATA                 Ldr;
  PRTL_USER_PROCESS_PARAMETERS  ProcessParameters;
  PVOID                         Reserved4[3];
  PVOID                         AtlThunkSListPtr;
  PVOID                         Reserved5;
  ULONG                         Reserved6;
  PVOID                         Reserved7;
  ULONG                         Reserved8;
  ULONG                         AtlThunkSListPtr32;
  PVOID                         Reserved9[45];
  BYTE                          Reserved10[96];
  PPS_POST_PROCESS_INIT_ROUTINE PostProcessInitRoutine;
  BYTE                          Reserved11[128];
  PVOID                         Reserved12[1];
  ULONG                         SessionId;
} PEB, *PPEB;


Example code:

typedef struct _PEB {
    BYTE Reserved1[2];
    BYTE BeingDebugged;
    BYTE Reserved2[21];
    PPEB_LDR_DATA LoaderData;
    PRTL_USER_PROCESS_PARAMETERS ProcessParameters;
    BYTE Reserved3[520];
    PPS_POST_PROCESS_INIT_ROUTINE PostProcessInitRoutine;
    BYTE Reserved4[136];
    ULONG SessionId;
} PEB;

Source: https://docs.microsoft.com/en-us/windows/win32/api/winternl/ns-winternl-peb
```

## **License**

The manw is published under the GPL v3 License. Please refer to the file named LICENSE for more information.
