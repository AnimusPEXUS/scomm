// Code generated by go-bindata.
// sources:
// ui/db.go
// ui/main.glade
// DO NOT EDIT!

package builtin_ownkeypair

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

var _uiDbGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2a\xcd\xcc\x29\xc9\xcc\x8b\xcf\x2f\xcf\xcb\x4e\xad\x2c\x48\xcc\x2c\xe2\xe2\x02\x04\x00\x00\xff\xff\xf5\xc3\x8c\x78\x1c\x00\x00\x00")

func uiDbGoBytes() ([]byte, error) {
	return bindataRead(
		_uiDbGo,
		"ui/db.go",
	)
}

func uiDbGo() (*asset, error) {
	bytes, err := uiDbGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/db.go", size: 28, mode: os.FileMode(420), modTime: time.Unix(1494405570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiMainGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4d\x6f\xdb\x46\x10\xbd\xeb\x57\x4c\xf7\xd2\x43\x23\xc9\x49\x50\xa0\x07\x8a\x01\x8a\xd6\x41\x91\xa2\x35\x90\x04\x45\x4f\xc2\x70\x39\xa4\xa6\x5a\xed\x32\xbb\x43\x7d\xfc\xfb\x62\x45\x3b\xb2\x2c\xd2\x92\x28\x1f\x6c\xc0\x37\x8a\x3b\x6f\x3e\xdf\xbe\xe5\x2a\xf9\xb0\x5e\x18\x58\x92\x0f\xec\xec\x44\xbd\x1d\x5d\x29\x20\xab\x5d\xce\xb6\x9c\xa8\xaf\x5f\xae\x87\xbf\xa8\x0f\xe9\x20\xf9\x61\x38\x84\x8f\x64\xc9\xa3\x50\x0e\x2b\x96\x19\x94\x06\x73\x82\xf7\xa3\x77\x57\xa3\x2b\x18\x0e\xd3\x41\xc2\x56\xc8\x17\xa8\x29\x1d\x00\x24\x9e\xbe\xd5\xec\x29\x80\xe1\x6c\xa2\x4a\x99\xff\xa4\x76\x81\x22\x4c\x8d\xb7\x76\x2e\xfb\x8f\xb4\x80\x36\x18\xc2\x44\x7d\x94\xf9\x3f\x6c\x73\xb7\x52\xc0\xf9\x44\xad\x9a\xe7\x68\x08\x90\x54\xde\x55\xe4\x65\x03\x16\x17\x34\x51\x1a\xed\xb4\x70\xba\x0e\x2a\xbd\x46\x13\x28\x19\xdf\x19\xb4\xdb\x37\xce\xa6\x95\x0b\x2c\xec\xac\x4a\x35\xc5\x8c\x8f\xc1\x4a\x8f\x4b\x96\x4d\xa7\xb9\x9e\xb1\xc9\x9b\xe7\xb6\x72\x7e\x75\x6b\x75\xb7\x7a\xe8\x7c\xc9\x81\x33\x43\x2a\xfd\xe2\xeb\x83\x02\xfa\x14\xdd\x86\x59\xa0\x2f\xd9\x4e\x0d\x15\xa2\xd2\x9f\xcf\x40\x78\x2e\x67\x67\x42\xc4\x55\xe7\x01\x32\x27\xe2\x16\x27\x62\x9c\x67\xb2\x82\xcd\xf8\x96\xe4\x85\x35\x9a\x53\x80\xa1\x42\xcd\xb6\xec\x0a\xb3\x37\xc3\xf6\x39\xfe\xb1\xc0\x92\xd4\x7d\x9b\x1e\xd3\xec\x3b\xd1\x36\x9c\x38\x67\x84\xab\xa9\xd0\x5a\x14\x88\x47\x1b\x0c\x0a\x66\x86\x26\x6a\x43\x41\xa5\x7f\x5b\x90\x19\x07\xa8\xb0\x24\xf8\xd7\xd5\xa0\xd1\xc2\x02\xed\xed\x4f\x0f\x37\x9e\x97\x28\x04\x9f\x68\x33\x1a\x0c\x7e\xfb\x8b\x04\x30\x84\x7a\x41\x01\xea\x10\xad\x5c\x01\xd5\xad\xcd\x9c\x36\xe1\x0d\xac\x66\xac\x67\x80\xc6\x13\xe6\x1b\xd0\xce\x0a\xb2\x85\xd5\x8c\x36\x1e\xaa\x3a\x33\xac\xa1\x42\x2f\xa3\xc1\xe0\xb3\x83\xad\xc3\x92\x44\xd8\x96\x70\xd3\xac\x7e\xa2\x0d\x14\xde\x2d\xee\xc7\x8e\x6e\xc9\x02\xcb\x8f\x01\x2c\x51\x4e\xf9\x68\x30\xf8\x1a\x08\xb2\x5a\xc4\xd9\xef\xc2\xf3\x06\xc4\x81\xf6\x14\x61\x96\x56\xfb\x2e\x9c\x87\xca\x53\x08\xf0\x7b\xce\x02\x68\x73\x90\xe8\xd4\x38\xcc\xef\x1b\x36\xc1\x83\x5b\x10\x14\x6c\xa8\xc1\x61\x10\x02\x16\x60\x0b\xd7\xb5\x31\x5b\x3b\xca\x79\x9b\x77\xc1\x64\x62\x3e\x37\x5b\xe7\x9f\x71\xd9\xa4\x11\x70\x49\x7b\x7e\xd9\xc6\xb7\xe2\x7c\xec\x1b\x16\x42\x7e\xdb\xf2\xf2\x36\xf7\x18\x27\xba\x04\x96\xd1\xa9\x03\x9e\xa1\xe1\xd2\xaa\x34\x08\x7a\x39\x15\x14\xc4\xe9\xb9\x4a\x4b\x99\x0f\x73\x46\xe3\xca\xe1\xb7\x9a\x42\xdc\x30\xed\x1e\x92\x71\xc3\xf4\xbd\x77\x15\xea\x39\xdb\xf2\xf1\x48\xb4\xae\xd0\xe6\x67\x92\xb6\x60\x63\xce\xdb\x1e\x3b\xb9\xbe\xea\xaa\xe0\x20\xdd\x64\xfc\x60\x43\x9f\xb2\xc1\xa3\x50\x6f\x0f\x9d\xcc\xad\xa7\x91\xee\xcf\x69\xb3\x3f\xae\x5f\xed\x25\x76\x97\xf9\xd0\xaa\x67\x71\x97\x14\xd8\x86\x3d\x59\xdd\xfb\x77\xa9\xbb\x53\x1d\xdd\xda\x4a\xd0\x2d\x2f\xb6\xcf\xd3\xbb\x2d\x3d\x75\x2b\x1b\x79\x32\xad\x90\x7d\x4b\x4b\xdb\x12\x34\x98\x91\x69\x15\xeb\x3b\x91\x7b\x2c\xf1\x8b\x06\x75\x64\x58\x7d\xe0\x9e\x34\xf1\x92\xc2\x34\xa7\x02\x6b\x23\xa7\x78\x69\xd3\x9b\x9d\xff\x36\xdd\xe9\x8a\x7e\x92\xfe\x74\x81\x8f\xeb\x50\x17\xf2\xa8\x1e\xed\x2a\xed\x2c\xe7\x40\x9f\xbe\x2f\x5c\x46\xcd\x78\xc2\x3c\x2d\x2d\xe3\x69\xfa\x4a\xc9\x17\x43\xc9\xb7\xcf\x8e\x92\xf1\x43\xe9\x69\x29\x19\xbf\xc1\x5e\x29\xf9\x62\x28\xf9\xee\xd9\x51\x32\xde\x09\x9e\x96\x92\x7f\x3a\xcc\x5f\x29\xf9\x62\x28\xf9\xfe\xa9\x29\xd9\xd5\x9b\xee\xbe\xf4\xee\x49\xbf\x7e\xf4\xfa\x88\xe9\xe8\x43\x4b\x0f\x2e\xbf\x44\x9e\x73\x87\xba\xe4\x0e\xd9\x71\x40\x1e\xbf\x43\xee\x97\xb8\xb7\xd8\xc8\x11\xc8\xa6\xa2\x89\x12\x16\x43\x19\xee\x34\x25\xa9\x0c\x6a\x9a\x39\x93\x93\x1f\x1f\xa0\x77\x6e\x93\xf1\xbd\xff\x4f\xff\x0f\x00\x00\xff\xff\x71\xf9\x9a\x25\x98\x15\x00\x00")

func uiMainGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiMainGlade,
		"ui/main.glade",
	)
}

func uiMainGlade() (*asset, error) {
	bytes, err := uiMainGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/main.glade", size: 5528, mode: os.FileMode(420), modTime: time.Unix(1494317550, 0)}
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
	"ui/db.go": uiDbGo,
	"ui/main.glade": uiMainGlade,
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
	"ui": &bintree{nil, map[string]*bintree{
		"db.go": &bintree{uiDbGo, map[string]*bintree{}},
		"main.glade": &bintree{uiMainGlade, map[string]*bintree{}},
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

