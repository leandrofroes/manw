# manw

manw is a very simple and fast command line search engine for Windows API written in Go.

## Why?

On Linux systems we are able to search for the documentation of a function using man command, but what about Windows? Open the browser and read the full documentation is always a good option  but sometimes the only thing we need is a high level overview and not the full documentation.

## **Installation**

You can either download the latest version from the [releases](https://github.com/leandrofroes/manw/releases) page or build it manually:

```
git clone https://github.com/leandrofroes/manw
cd manw
make
```

OR

```
go get https://github.com/leandrofroes/manw
```

NOTE: Tested on Linux and Windows.

## **Usage**

```
NAME

  manw - A multiplatform command line search engine for Windows API.
  
SYNOPSIS: 

  ./manw [OPTION...] [STRING]
          
OPTIONS:

  -f, --function  string  Search for a Windows API Function.
  -s, --structure string  Search for a Windows API Structure.    
  -k, --kernel    string  Search for a Windows Kernel Structure.
  -t, --type      string  Search for a Windows Data Type.
  -a, --arch      string  Specify the architecture you are looking for.
  -n, --syscall   string  Search for a Windows Syscall ID. If you don't use -a the default value is "x86".
  -c, --no-cache  bool    Disable the caching feature.
```

## **Examples**

```
$ ./manw -f createprocess
CreateProcessA function (processthreadsapi.h) - Win32 apps

Exported by: Kernel32.dll

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
$ ./manw -s peb
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

```
$ ./manw -n NtAllocateVirtualMemory -a x64
Windows 7
	- SP0: 21
	- SP1: 21
Windows Server 2012
	- SP0: 22
	- R2: 23
Windows 8
	- 8.1: 23
	- 8.0: 22
Windows 10
	- 1803: 24
	- 1809: 24
	- 1903: 24
	- 1909: 24
	- 1507: 24
	- 1703: 24
	- 1709: 24
	- 2004: 24
	- 20H2: 24
	- 1511: 24
	- 1607: 24
Windows XP
	- SP1: 21
	- SP2: 21
Windows Server 2003
	- R2 SP2: 21
	- SP0: 21
	- SP2: 21
	- R2: 21
Windows Vista
	- SP0: 21
	- SP1: 21
	- SP2: 21
Windows Server 2008
	- SP0: 21
	- SP2: 21
	- R2: 21
	- R2 SP1: 21

```

If no parameter is specified manw is going to use -f flag by default.

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

## **Version 1.2**:

* General bug fix
* Now the cache directory is created only if you specify the -c flag

## **Version 1.3**:

* Now the cache feature is enabled by default and -c flag was removed.
* New -s flag for Windows API Structure search.
* -a renamed to -f.
* Fix flag number checking in order to allow only a single flag usage.
* General code updates.

## **Version 2.0**:

* Update Google search to meet with the new google page format.
* Improve Google search to be generic and to avoid declare a new one for each scrapy function.
* Add Windows Syscall ID searching through -n flag.
* Add -a flag to specify the architecture along with -n flag. If you use -n with no architecture flag the default value is x86.
* Add -c flag to disable caching feature.
* Change module usage to use go get.
* General code improvement and bug fixing.

## :warning: **Warning**

The scraper relies on the way the pages used by the project (e.g. google, MSDN, etc) are implemented so keep in mind that if it changes the search might not work. That being said always keep your manw up-to-date and please let me know if you find any issue.

## **Known issues:**

* Data type search might show a very weird output depending the data you search.
* Currently the kernel struct info search supports Windows Vista 32bits kernel only. I do have plans to support other versions in the future.
* I'm also always trying to improve the code (e.g. performance, best practices, etc.) since I started it as a study project.

## **Special Thanks**

* [@merces](https://github.com/merces) for the core idea and all the support.

## **License**

The manw is published under the GPL v3 License. Please refer to the file named LICENSE for more information.
