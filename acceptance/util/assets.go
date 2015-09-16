// +build acceptance

package util

import (
	"fmt"
	"path"
	"os"
)

// Create a map type to reference.
// The key will be an enumeration and the value will be a string to represent a path (either relative or absolute)
type assetsMap map[AssetId]string

// Create the key for the assetsMap. Essentially an enumeration.
type AssetId int

// This enumeration contains all they keys for
// When adding more assets, step 1 of 2 is add another entry here to represent all the entries.
const (
	RootDirAssets AssetId = iota
	TestUser01Assets
)

// assetsRelPaths contains all the unverified relative paths to tests assets.
// When adding more assets, step 2 of 2 is add another entry here to represent the relative location of the asset with respect to the root test asset folder
var assets = assetsMap{
	RootDirAssets:    "",
	TestUser01Assets: path.Join("users", "testuser01"),
}

// FindAssets will verify and create a map of the absolute location of the assets.
func FindAssets() assetsMap {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dir = path.Join(dir, "test_assets")
	loadedAssets := make(assetsMap)
	for key, relPath := range assets {
		absPath := path.Join(dir, relPath)
		if _, err := os.Stat(absPath); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		loadedAssets[key] = absPath

	}
	return loadedAssets
}
