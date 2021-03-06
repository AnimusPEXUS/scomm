// Code generated by go-bindata.
// sources:
// ui/main.glade
// DO NOT EDIT!

package builtin_owntlscert

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

var _uiMainGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdf\x6f\xdb\x36\x10\x7e\xf7\x5f\x71\xe3\xeb\x6a\x2b\x6d\x31\x60\x0f\xb2\x0a\x74\x58\x8b\x02\x05\x56\xa0\x29\x86\x3d\x19\x27\xea\x2c\xdd\x4c\x93\x2a\x79\xf2\x8f\xff\x7e\x50\x9c\xc4\x49\x2c\xc6\xb2\x1c\x60\x09\x90\x37\x59\xbc\xef\xc8\xef\xee\xe3\x27\xd2\xe9\x87\xcd\xd2\xc0\x8a\x7c\x60\x67\xa7\xea\xed\xe4\x42\x01\x59\xed\x0a\xb6\xe5\x54\xfd\xb8\xfc\x34\xfe\x5d\x7d\xc8\x46\xe9\x2f\xe3\x31\x7c\x26\x4b\x1e\x85\x0a\x58\xb3\x54\x50\x1a\x2c\x08\xde\x4f\xde\x5d\x4c\x2e\x60\x3c\xce\x46\x29\x5b\x21\x3f\x47\x4d\xd9\x08\x20\xf5\xf4\xb3\x61\x4f\x01\x0c\xe7\x53\x55\xca\xe2\x57\xb5\x9f\xa8\x85\xa9\xe4\x2a\xce\xe5\xff\x92\x16\xd0\x06\x43\x98\xaa\xcf\xb2\xf8\x9b\x6d\xe1\xd6\x0a\xb8\x98\xaa\xf5\xee\xb9\x0d\x04\x48\x6b\xef\x6a\xf2\xb2\x05\x8b\x4b\x9a\x2a\x8d\x76\x36\x77\xba\x09\x2a\xfb\x84\x26\x50\x9a\xdc\x04\x74\xc7\xef\x92\xcd\x6a\x17\x58\xd8\x59\x95\x69\x6a\x57\x7c\x0c\x56\x7a\x5c\xb1\x6c\xa3\xe1\xba\x62\x53\xec\x9e\xbb\xe8\x7c\x74\x1b\x75\x33\x7a\x98\x7c\xc5\x81\x73\x43\x2a\xbb\xf4\xcd\x01\x81\x21\xa4\xbb\x30\x4b\xf4\x25\xdb\x99\xa1\xb9\xa8\xec\xb7\x13\x10\x9e\xcb\xea\x44\x88\xb8\xfa\x34\x40\xee\x44\xdc\xb2\x27\xc6\x79\x26\x2b\xb8\x6b\xdf\x8a\xbc\xb0\x46\xd3\x07\x18\x6a\xd4\x6c\xcb\xd8\x34\xf7\x7a\xd8\xdd\xc7\x2f\x4b\x2c\x49\xdd\x8d\x19\xd0\xcd\xa1\x1d\xed\xc2\x89\x73\x46\xb8\x9e\x09\x6d\x44\x81\x78\xb4\xc1\xa0\x60\x6e\x68\xaa\xb6\x14\x54\xf6\x97\x05\xa9\x38\x40\x8d\x25\xc1\x3f\xae\x01\x8d\x16\x96\x68\xaf\x7f\x7a\xb8\xfc\xfa\x1d\x74\x5b\xc4\x39\x6b\x14\x7a\x03\xeb\x8a\x75\x05\x1c\xa0\x09\x54\x80\xb8\xd6\x0b\xfc\xb6\x16\x40\x5b\x00\x59\xf1\x4d\x10\x68\x6c\x20\xdd\x78\x2a\x40\x3b\x6b\x49\xb7\xcd\x08\x6f\x20\x34\xba\x02\x0c\x70\xf9\xc7\xb7\xe4\xcb\xb7\xd1\xe8\x63\x23\xe2\xec\xad\x6b\x80\xf6\x84\x42\x01\x2c\xad\x1f\x4e\x3c\x19\x8d\x7e\x04\x82\xef\xb8\xa2\xe4\xab\xc3\x02\xf2\x1d\x56\x1c\x84\xf6\x9d\x69\xdf\xdd\x89\x07\x71\xc9\xdc\xbb\x25\x04\x71\x1e\x4b\x9a\xf4\x2d\x59\x85\x86\x4b\xab\xb2\x20\xe8\xa5\x2f\x28\x88\xd3\x0b\x95\x95\xb2\x18\x17\x8c\xc6\x95\xe3\x9f\x0d\x85\x96\x75\x77\x86\x34\xd9\x69\xe7\xde\xbb\x1a\xf5\x82\x6d\xf9\xf8\x4c\xb4\xa9\xd1\x16\x27\xca\x60\xce\xc6\x9c\x26\xb8\xbd\x01\x5e\xc4\x18\x1c\x2c\x37\x4d\x1e\x6c\x91\x3e\x5b\xa6\xb5\xbe\x2b\x1b\xcf\xdd\x66\x76\xa7\x7f\xcf\x69\x17\x3d\x6e\x0c\xdd\x4c\xe3\x6c\x1f\x46\x0d\x24\x77\x0e\xc1\x2e\x6c\x6f\xdb\x1c\x5e\xa5\x78\xa5\x22\xd5\xba\xda\xe2\xd7\xf2\xb8\x7a\x9e\x95\xd7\x56\x31\x73\x6b\xfb\x88\x5c\x62\x6b\x34\x98\x93\xe9\x34\xc2\x1b\x0f\x7a\x6c\xed\x67\xf5\xea\x48\xbf\x86\xc0\x3d\x69\xe2\x15\x85\x59\x41\x73\x6c\x8c\xf4\xc9\xd2\xe5\x3c\xfb\xfc\x5d\x0e\x14\x9b\xbd\x97\x13\xc5\xc0\xc7\x1d\x29\x86\x3c\xea\x4c\x7b\xa6\x51\x3a\x07\x4e\x75\x3b\x70\x9e\x3a\xa9\x60\x79\x72\x65\xfe\x59\x70\xf4\x3b\x14\xcb\xf6\xaa\xca\xff\x4d\x95\x6f\x9f\x9d\x2a\xdb\xf3\xd1\x93\xab\xb2\x3d\x88\xbd\xaa\xf2\xc5\xa8\xf2\xdd\xb3\x53\x65\x7b\x60\x7f\x72\x55\xb6\x37\x83\x57\x55\xbe\x18\x55\xbe\x7f\x6a\x55\xc6\x6a\x13\xaf\xcb\xe0\x9a\x0c\xab\xc7\xa0\xd3\x4c\xa4\x0e\x1d\x35\x38\xff\x5e\x79\xca\x7d\xea\x9c\x6b\x65\xe4\x33\x79\xfc\x5a\x79\x9f\xe2\xbd\xc1\x9d\x23\x81\x6c\x6b\x9a\x2a\x61\x31\x94\xa3\xbf\xf5\x94\xb4\x36\xa8\xa9\x72\xa6\x20\x9f\x1c\xa0\xf7\x69\xd3\xe4\xce\x9f\x94\xff\x05\x00\x00\xff\xff\x2b\x3e\x70\xc9\xfd\x14\x00\x00")

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

	info := bindataFileInfo{name: "ui/main.glade", size: 5373, mode: os.FileMode(420), modTime: time.Unix(1494419177, 0)}
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

