# manw

manw is a very simple and fast command line search engine for Windows API written in Go.

## Why?

On Linux systems we are able to search for the documentation of a function using man command, but what about Windows? Open the browser and read the full documentation is always a good option  but sometimes the only thing we need is a high level overview and not the full documentation.

## **Installation**

```
git clone https://github.com/leandrofroes/manw
cd manw
make
```

## **Usage**

```
NAME

  manw - A multiplatform command line search engine for Windows API.
  
SYNOPSIS: 

  ./manw [OPTION]... [STRING]
          
OPTIONS:

  -a, --api     string  Search for a Windows API Function/Structure.
  -c, --cache           Enable caching feature.
  -k, --kernel  string  Search for a Windows Kernel Structure.
  -t, --type    string  Search for a Windows Data Type.
```

## **Examples**

```
$ ./manw -a createfilew -c
CreateFileW function (fileapi.h) - Win32 apps - Kernel32.dll

Creates or opens a file or I/O device. The most commonly used I/O devices are as follows:\_file, file stream, directory, physical disk, volume, console buffer, tape drive, communications resource, mailslot, and pipe.

HANDLE CreateFileW(
  LPCWSTR               lpFileName,
  DWORD                 dwDesiredAccess,
  DWORD                 dwShareMode,
  LPSECURITY_ATTRIBUTES lpSecurityAttributes,
  DWORD                 dwCreationDisposition,
  DWORD                 dwFlagsAndAttributes,
  HANDLE                hTemplateFile
);


Return value: If the function succeeds, the return value is an open handle to the specified file, device, named pipe, or mail slot. If the function fails, the return value is INVALID_HANDLE_VALUE. 

Source: https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew

```

```
$ ./manw -a peb
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

```
$ ./manw -t callback

Data Type: CALLBACK

The calling convention for callback functions. This type is declared in WinDef.h as follows: #define CALLBACK __stdcall CALLBACK, WINAPI, and APIENTRY are all used to define functions with the __stdcall calling convention. Most functions in the Windows API are declared using WINAPI. You may wish to use CALLBACK for the callback functions that you implement to help identify the function as a callback function.

```

```
$ ./manw -k _token_control
//0x28 bytes (sizeof)
struct _TOKEN_CONTROL
{
    struct _LUID TokenId;                                                   //0x0
    struct _LUID AuthenticationId;                                          //0x8
    struct _LUID ModifiedId;                                                //0x10
    struct _TOKEN_SOURCE TokenSource;                                       //0x18
}; 

Used in_SECURITY_CLIENT_CONTEXT

```

## **Version 1.0**:

* DLL dependency added to Windows API module.
* New Command Line flags support.
* New Caching feature for offline usage.
* New Windows Data Type search module.
* New Windows Kernel Structure search module.
* Now the project is modular.

## **Version 1.1**:

* Fix v0.1 compatibility. If no parameter is passed manw is going to run the API Search module by default.
* Remove caching path configuration requirement. Now the caching path is created by manw itself.
* Now both Kernel Structure and Data Type module supports caching feature.
* Some other code improvements.

## **Special Thanks**

* [@merces](https://github.com/merces) for the core idea and all the support.

## **License**

The manw is published under the GPL v3 License. Please refer to the file named LICENSE for more information.
