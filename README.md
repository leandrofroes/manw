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

  manw - A multiplatform command line search engine for Windows OS info.
  
SYNOPSIS: 

  ./manw [OPTION...] [STRING]
          
OPTIONS:

  -f, --function  string  Search for a Windows API Function.
  -s, --structure string  Search for a Windows API Structure.    
  -k, --kernel    string  Search for a Windows Kernel Structure.
  -t, --type      string  Search for a Windows Data Type.
  -c, --cache     bool    Enable cache feature.
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
$ ./manw -s peb -c
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
typedef struct _TOKEN_CONTROL
{
     LUID TokenId;
     LUID AuthenticationId;
     LUID ModifiedId;
     TOKEN_SOURCE TokenSource;
} TOKEN_CONTROL, *PTOKEN_CONTROL;

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

## **Version 1.2**:

* General bug fix
* Now the cache directory is created only if you specify the -c flag

## **Version 1.3**:

* Now the cache feature is enabled by default and -c flag was removed.
* New -s flag for Windows API Structure search.
* -a renamed to -f.
* Fix flag number checking in order to allow only a single flag usage.
* General code updates.

## :warning: **Warning**

The scraper relies on the way the pages used by the project (e.g. google, MSDN, etc) are implemented so keep in mind that if it changes the search might not work. That being said always keep your manw up-to-date and please let me know if you find any issue.

## **Known Issues**

* Currently the kernel struct info search supports Windows Vista 32bits kernel only. I do have plans to support other versions in the future.

## **Special Thanks**

* [@merces](https://github.com/merces) for the core idea and all the support.

## **License**

The manw is published under the GPL v3 License. Please refer to the file named LICENSE for more information.
