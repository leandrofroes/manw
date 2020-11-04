package config

import (
  "os"
  "strings"

  "manw/pkg/utils"
)

func Load() (cachePath string){
  cacheDir, err := os.UserCacheDir()
  utils.CheckError(err)
  cachePath = cacheDir + "/manw/"

  if _, err := os.Stat(cachePath); os.IsNotExist(err) {
    err := os.Mkdir(cachePath, 0700)
    utils.CheckError(err)
  }

  strings.ReplaceAll(cachePath, ":", "\\:")
  strings.ReplaceAll(cachePath, "\\", "\\\\")

  return cachePath
}
