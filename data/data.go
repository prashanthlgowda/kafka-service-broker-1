// Code generated by go-bindata.
// sources:
// data/assets/catalog.json
// DO NOT EDIT!

package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsCatalogJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x93\x4f\x8f\xd3\x30\x10\xc5\xef\xf9\x14\x23\x4b\xe5\x44\x6a\x37\xd9\x6d\x9a\xdc\x60\xf7\x82\x90\x96\xc3\x82\x38\xa0\x3d\x4c\xec\x49\x6a\x35\xb5\x2d\xdb\x6d\x55\x55\xfd\xee\x28\x29\xfd\x13\x76\x81\x1b\x97\x48\x9e\x3c\xdb\xef\xf7\xf2\x72\x48\x00\x58\x20\xbf\xd5\x92\x02\xab\x7e\x24\x00\x00\x87\xe1\x09\xc0\xb4\x62\x15\xbb\xc3\x52\xe5\x34\xc3\x34\x9f\x0b\x4c\x67\x33\x2a\xd2\xfa\xfe\xae\x48\xb3\x7c\x2e\xa5\x14\x6a\xde\x60\xcd\xde\x9f\xb7\x18\x5c\x13\xab\x58\x88\xe8\x57\x68\xd4\x0e\xf7\x86\xd2\x15\x36\x2b\xbc\x6a\x14\x05\xe9\xb5\x8b\xda\x1a\x56\xb1\x0f\x0e\xe5\x92\xe0\xf3\x58\x53\x6b\xa3\xb0\xee\x88\x55\x10\xfd\x86\x2e\xf3\x88\xed\xc5\xe7\x30\x38\x9d\xfd\x6b\xfd\x72\xd1\xad\x29\xa2\xc2\x88\xac\x3a\x5c\xb5\x4a\x07\xd7\xe1\xfe\xe9\xe4\xf1\xb9\xf7\x08\xef\xe0\x7b\xef\xf1\xb7\xfb\x7b\xf8\x35\xb6\xf4\xcd\x77\xac\x62\xcb\x18\x5d\xa8\x38\x0f\x5b\x33\xc5\xc1\xee\xd4\xfa\x96\x7b\x72\x36\x70\x0c\x0d\x1f\x4c\xf0\xa0\x23\xf1\xce\xb6\x36\x70\xeb\x75\xab\x0d\x76\x81\x3b\xd3\xf2\x4f\x0f\x5f\x9e\x26\x99\x48\x27\x99\xf8\xd8\xa1\x5c\x4d\x32\x61\xcd\x24\x13\x5f\x3d\x9a\xe0\xd0\x93\x89\x53\x67\xda\xdb\xdb\x3b\x6b\xda\xc7\x7f\xe7\x04\xc0\x9c\xb7\x5b\xad\xc8\x3f\xfe\x19\xee\x56\xae\xac\xdc\xac\xc9\x44\xec\x8f\x1d\xe3\xed\x76\xbb\xa9\xd2\xa6\xb5\x51\x77\x14\xa6\xd2\xae\xf9\xb0\x4c\x47\x19\x1f\x2f\x19\xbb\x0e\xcd\xe8\x63\x5c\xa3\xbe\x94\x67\x91\x09\x95\xe5\xf2\xa6\x3c\x65\x5f\x1e\x55\x2c\x8a\x42\x60\x9e\xcb\xfb\xfa\xc6\xdd\xb5\x40\xd1\x3a\x2d\xc7\x6f\xc6\xb5\x79\x5e\xa2\x27\x40\x08\xda\xb4\x1d\xc1\xa0\x07\x6b\x20\xf4\x73\xf5\x2a\xa6\x51\x25\x46\x46\x01\x98\xb4\x21\xb2\x0a\xc4\xcd\xf4\x38\xda\xda\x78\x3a\x37\x31\x79\x43\xf1\x9a\x5b\x64\x8d\x2a\x33\xb9\x48\x65\x59\x16\x27\xee\x85\x14\x59\x5a\x17\x72\x21\x55\x39\xab\x9b\xd9\xdd\x9b\xdc\x27\xfb\x7f\x03\x7f\xf0\x84\x91\x60\x6f\x37\x1e\xec\xce\x9c\xc8\xc3\xff\x45\x3f\xff\x6e\xc9\x79\xf5\x92\x1c\x93\x9f\x01\x00\x00\xff\xff\x84\x1d\xbf\xd5\x4c\x04\x00\x00")

func assetsCatalogJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsCatalogJson,
		"assets/catalog.json",
	)
}

func assetsCatalogJson() (*asset, error) {
	bytes, err := assetsCatalogJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/catalog.json", size: 1100, mode: os.FileMode(420), modTime: time.Unix(1510703633, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/catalog.json": assetsCatalogJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"catalog.json": &bintree{assetsCatalogJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

