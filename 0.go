// Created by EldersJavas(EldersJavas&gmail.com)

package main

import (
	"io/fs"
)

func init() {
	f, err := fs.Sub(resourceRootFS, "res")
	if err != nil {
		panic(err)
	}
	ResourceFS = f
}
