// Code generated by go-bindata.
// sources:
// ui/key-cert-editor.glade
// ui/main.glade
// ui/plugin-acceptor.glade
// ui/storage-load.glade
// ui/systray-replacer.glade
// DO NOT EDIT!

package main

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

var _uiKeyCertEditorGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5a\x41\x6f\xe3\x36\x13\xbd\xfb\x57\xf0\xe3\xf5\x83\xe3\x64\x8b\x16\x3d\xd8\x5a\x74\x17\xc8\x76\x91\x6d\x11\x20\x69\x8b\x9e\x04\x8a\x1a\xcb\x5c\x53\x1c\x95\xa4\xec\xa8\x45\xff\x7b\x21\xab\x46\xec\x88\xb2\x44\x29\x71\x1c\xf7\xe6\xc0\x1c\x72\xe6\xbd\xc7\xe1\xcc\x38\xd3\xf7\x0f\xa9\x24\x2b\xd0\x46\xa0\x9a\xd1\xab\x8b\x4b\x4a\x40\x71\x8c\x85\x4a\x66\xf4\x97\xfb\xeb\xf1\xf7\xf4\x7d\x30\x9a\xfe\x6f\x3c\x26\x9f\x40\x81\x66\x16\x62\xb2\x16\x76\x41\x12\xc9\x62\x20\xdf\x5c\xbc\xbb\xbc\xb8\x24\xe3\x71\x30\x9a\x0a\x65\x41\xcf\x19\x87\x60\x44\xc8\x54\xc3\x1f\xb9\xd0\x60\x88\x14\xd1\x8c\x26\x76\xf9\x7f\xfa\x78\x50\x69\x46\x27\x9b\x75\x18\x7d\x05\x6e\x09\x97\xcc\x98\x19\xfd\x64\x97\x3f\x81\xca\x29\x11\xf1\x8c\xa6\xa0\xf2\x2b\x5a\xae\x22\x64\x9a\x69\xcc\x40\xdb\x82\x28\x96\xc2\x8c\xae\x84\x11\x91\x04\x1a\xdc\xeb\x1c\xa6\x93\xed\xb7\xee\xc5\x9c\xa9\x70\x8e\x3c\x37\x34\xb8\x66\xd2\xd4\xd7\xf3\x85\x90\x71\xf5\xd9\xe5\xd2\xc7\x05\xf0\x8d\x5f\x9f\x2d\xa4\x95\x6f\x3c\x15\x21\xc7\xac\xa0\x5b\x2b\x4f\x17\xfb\xb8\xe9\xb2\x51\x18\x9a\x05\xae\x43\x26\x65\xf7\x83\x24\x8b\x40\x52\x62\x35\x53\x46\x32\xcb\x22\x09\x33\x5a\x80\xa1\xc1\x0f\x52\xe2\x9a\x7c\xc4\xac\xe8\xb2\x4f\x6e\x20\xcc\x55\x0c\x5a\x0a\xd5\x18\xe7\x74\x52\xe1\xf9\x2f\xd4\x93\x1d\xac\x7b\xe2\x5e\x06\x7c\x66\xb8\xdf\x2d\x70\x4d\x6e\xc0\x81\xfa\x00\xf8\xee\x20\x63\x9a\x59\xd4\xfb\x10\x9a\x54\x84\x16\x93\x44\x82\x09\xcd\x76\xc9\x71\xf1\x1c\x10\xd4\x7e\x2c\xa9\x08\x25\xb2\xf8\xb5\xd5\xd0\x4c\xec\x17\x64\x31\x99\x6b\x4c\xc9\xb5\x90\x70\x71\x71\x0a\xd7\xaa\x06\xa1\x61\x2b\x38\x5d\x08\xef\xd8\x0a\x88\xc5\x13\x02\xb0\x7e\xb1\xde\xec\xed\xe1\xa0\xad\x45\x94\xe1\x12\x8a\x50\xa8\x39\xbe\xb6\x0e\x9e\x39\xb1\x6e\xe3\x23\xe3\xf1\x12\x8a\x71\x19\xe1\x29\x28\xa8\x99\x87\x2c\x8f\xce\x9f\x8a\x2a\xc8\x37\xc0\x46\xf9\x41\xcc\x05\x67\x16\xce\x9c\x92\x9d\x48\x8f\xc4\xcb\xee\x17\x75\x4e\x3e\xe0\x43\x45\x87\x46\xb4\x47\x69\x02\x9e\xac\x47\x2d\x40\x59\x66\x05\x2a\x1a\xac\x4a\x74\x38\x93\x6d\x46\x26\x63\x5c\xa8\x84\x06\xdf\xfa\xf6\x18\x65\xbc\xc7\x16\xd7\x9e\x4b\x6e\xb7\xbe\x54\xda\x29\x89\xd8\xc8\x28\xb4\x45\xb6\x5b\x2b\xf4\x72\xb6\xaf\xc3\x7e\xe2\xfe\xeb\x06\x0a\x72\x5f\x64\xf0\xb7\x7b\xaf\x7d\x61\x6e\x77\x67\x7c\x29\x54\x72\xf8\x4c\x78\xc8\x98\x8a\x3d\x1d\x9d\x8b\x43\x77\xd6\x65\x91\xa1\x11\x95\xfc\x2e\x9b\x22\xa8\xb9\xbb\x97\xf7\x7a\x93\x0c\x5a\xef\xb5\x04\x47\x66\xf9\x0d\x31\x73\xd5\x97\x99\xa7\x31\xd6\xe3\xf3\x8e\xcd\x2f\x2e\x2f\xb5\x3d\x89\xc7\xfb\x75\xfd\x90\x5b\x5b\x6e\xfd\x22\x09\xae\xab\x89\x06\x0e\x62\x05\x26\x8c\x61\xce\x72\x69\x7d\xb0\xc9\xf2\x8c\x06\x9b\x81\x54\xc7\x4c\x9a\x49\xc6\x61\x81\x32\x06\x3d\x79\x8b\x22\x70\xc4\x39\x44\x04\x3f\xa3\x85\x08\x71\x59\x65\x19\x15\x1d\x5b\x0a\x1d\xb2\xe0\x53\x17\xc3\x79\x5e\x42\xf7\xa2\x49\xd0\xc7\x6c\x53\xf8\x59\x16\xf9\xbe\x90\x1b\xbb\x08\x75\x0c\xba\xcd\xb2\x86\x92\x1b\xa9\x3b\xae\x51\x4a\x88\x7f\x13\x2a\xde\x1b\xc5\x0d\x82\x69\x00\x54\x2e\xd3\x85\xd9\x78\x19\x31\x1d\x66\x28\x05\x2f\x68\xc0\xe4\x9a\x15\xc6\x67\x93\x55\xaf\x4d\x9c\x30\xba\xa1\xbc\x87\x07\xfb\xab\x80\x75\x25\x3a\xbb\x76\x89\x6e\x30\xaa\x03\x91\x75\x99\x43\x2c\x36\xb5\x56\x8b\xa2\x9a\xcc\x53\x54\x58\xd6\xca\x9d\x9c\x77\x95\x03\xdb\x6f\xdc\x8a\x75\x1a\x38\x17\x57\x5c\x91\xb2\xaa\x9d\x51\xcb\xa2\x1a\xf4\x4d\xd5\xd2\x71\x44\xdf\x0a\xae\xdf\x88\x59\x35\x6f\xd5\x04\xb2\xbb\xee\x72\x1d\x6d\x59\x14\x56\xaf\x4d\x8b\xd7\x8e\xe2\xa8\x85\x9f\x76\x52\xf6\xdb\xa7\x93\xa2\x24\x65\x3a\x11\x2a\x94\x30\xb7\x34\xb8\x6a\x28\xe5\x0f\x9a\x6a\x91\x2c\xfa\xda\x5a\xcc\x7a\x5a\x46\x68\x2d\xa6\xbe\xc6\x9d\xda\xe6\x76\x86\xdd\x2c\x7f\x4e\x59\xf2\xb4\xfb\x1c\xcc\xf5\x50\xbe\x5d\xf6\x05\x93\x22\x69\x6e\x0f\x9a\x0d\x8d\x45\xbe\xa4\x41\x62\x97\x8d\x53\x98\x43\xe6\x82\xa3\x0a\x8d\xf8\x13\x68\xf0\x5d\xdf\xa4\x7a\xe8\xce\xbb\xce\xdc\x96\xa4\x7d\x70\x6e\x6f\xbd\x9a\x2c\x5b\x9b\xe3\xc7\x40\x9b\x33\x98\xfb\xf9\xf0\xd3\x64\xd3\x73\xe0\x72\xfa\x55\x35\xd9\xfc\x34\x58\x24\x65\x0d\x42\x6e\xb5\x58\x31\x0b\xe4\x06\x8a\x11\x7d\xfc\x29\x98\x12\xa6\x62\x42\xb7\xbf\x51\xd2\x91\x50\xa4\x6c\x85\x08\x8b\x70\x05\x64\x34\x42\x25\x0b\xf2\x3b\xe6\xc4\x94\x0d\x0f\x59\x2a\x5c\x97\x7f\xea\xdd\x0d\x7d\xbd\xfd\x9a\x1b\x2b\xe6\x05\x0d\x38\x28\x0b\xba\xef\x05\x6c\x15\xc7\xb9\xdc\x82\x96\x4c\xe3\x7f\x0b\x86\x17\x04\x9d\x9c\xf3\x2f\x06\xfe\x0b\xc5\xda\x8f\x22\x8e\xe1\x45\xab\xb5\x8e\xca\x39\x42\x91\x57\x8f\xa5\x61\x88\xd9\x40\x7c\x37\xd2\x4f\x74\x48\x7d\x9d\x4b\xd9\x9c\x1e\xfb\x0f\x42\x3b\xf1\xf4\x5c\xf3\xe3\x83\xf3\x80\xd7\x1c\x99\x78\xcf\x00\x06\xf7\xff\x5d\xfb\x95\x5a\xdf\x9f\x31\x6d\x4f\x74\x92\xe2\xd1\xeb\xf7\xee\xf3\x3b\xb7\xec\xfd\xaf\xc4\x33\x0c\xed\xcf\x39\x13\xdd\xe6\x91\x14\x9c\xdc\x32\x6d\x9f\x3b\x19\x75\x78\x6b\x5e\x28\x7f\xf5\x1f\xb0\x77\x1d\x3e\xf7\x9d\xaf\xbf\xf3\x99\xaf\x3f\x86\x31\x9d\xec\xfc\x9b\xed\x3f\x01\x00\x00\xff\xff\x38\x8a\x33\xb3\xbf\x2b\x00\x00")

func uiKeyCertEditorGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiKeyCertEditorGlade,
		"ui/key-cert-editor.glade",
	)
}

func uiKeyCertEditorGlade() (*asset, error) {
	bytes, err := uiKeyCertEditorGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/key-cert-editor.glade", size: 11199, mode: os.FileMode(420), modTime: time.Unix(1489591472, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiMainGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5f\x6f\xdb\x36\x10\x7f\xef\xa7\xd0\xf4\x3a\x38\x76\x92\x6e\xeb\x00\xdb\x45\xd7\xad\x41\xb1\xa2\x28\x96\x6c\x7d\x14\x28\xe9\x2c\x71\xa6\x49\x8d\xa4\xec\xa4\x9f\x7e\xb0\x64\x3b\x8e\xcd\x7f\x92\x15\x47\xae\x95\xa7\x08\xe6\x9d\x78\x77\xbf\x3b\x1e\x4f\x3c\x0e\xdf\xde\xcf\x88\x37\x07\x2e\x30\xa3\x23\xff\xf2\x62\xe0\x7b\x40\x23\x16\x63\x9a\x8c\xfc\xbf\xef\x3e\xf4\xde\xf8\x6f\xc7\xaf\x86\x3f\xf4\x7a\xde\x0d\x50\xe0\x48\x42\xec\x2d\xb0\x4c\xbd\x84\xa0\x18\xbc\xeb\x8b\xab\xab\x8b\x4b\xaf\xd7\x1b\xbf\x1a\x62\x2a\x81\x4f\x50\x04\xe3\x57\x9e\x37\xe4\xf0\x5f\x8e\x39\x08\x8f\xe0\x70\xe4\x27\x72\xfa\xa3\xff\xf8\xa2\xeb\x8b\xab\x81\xdf\x2f\xc6\xb1\xf0\x5f\x88\xa4\x17\x11\x24\xc4\xc8\xbf\x91\xd3\xaf\x98\xc6\x6c\xe1\x7b\x38\x1e\xf9\x8b\xf2\xff\xe5\x40\xcf\x1b\x66\x9c\x65\xc0\xe5\x83\x47\xd1\x0c\x46\x7e\x84\x68\x30\x61\x51\x2e\xfc\xf1\x07\x44\x04\x0c\xfb\xeb\x01\xab\xf1\x51\x8a\x49\x5c\xfe\xbf\xa4\x26\x28\x82\x94\x91\x18\x78\x7f\x35\xa0\xbf\x35\x62\x67\xf4\xde\xb4\x7e\x63\xf7\xfe\xfa\xd7\xfd\xb9\xcc\xb1\xc0\x21\x01\x7f\x7c\xc7\xf3\xbd\x89\xd4\x99\xbc\x8a\x86\x71\x0c\x54\x22\x89\x19\xf5\xc7\x73\xe0\x12\x47\x88\x28\x09\x9f\xc8\xa2\x96\xe7\x33\x93\x10\x32\x36\x2d\x15\x4d\x57\x4f\xc1\x0c\x61\xea\x6f\x53\xd6\x90\xd5\x22\x6f\x15\x32\x91\xb2\x45\x10\x32\x1e\x03\x37\x28\x4a\x49\x19\x71\x46\x08\xb2\x4f\x74\x4f\x57\x6a\x7d\x7d\x41\x14\x62\x7f\x77\x5c\x4d\xed\x1c\xa0\x21\x15\xe9\x0c\xf1\x04\xd3\x80\xc0\x44\xfa\xe3\x9f\x6a\x50\x72\x9c\xa4\x35\x49\x25\xcb\xea\x11\x86\x4c\x4a\x36\xb3\xd1\x2a\xad\xa3\xb6\xd0\x6d\x61\x71\x88\xbf\x6e\x07\x8d\xc6\xcc\x75\xa0\xc9\x54\xe4\x22\x45\x31\x5b\x04\xf2\x21\x03\x7f\x8c\xa9\x95\x5c\xab\x0b\xb5\x3e\x3e\x46\x8c\xfe\x83\x41\xa7\x89\x83\xb5\xd1\x80\x46\x54\x2c\x4a\x78\xf8\xe3\x9f\xed\xc4\xc3\x7e\x29\xb4\x46\x5f\x7d\x3d\x78\x0c\x84\xc3\x0c\x45\x53\x4c\x13\x37\x1b\x72\x10\xf8\x1b\x58\x62\x93\x8e\x58\xa4\x1c\xd3\xa9\x8b\xaa\x86\x7d\xed\xac\xb4\x62\x76\xbe\xe3\xa2\x0b\xb5\x3e\xee\x38\xc0\x29\xf8\x4e\x21\x96\x57\x64\x5e\x14\x91\x5e\xf1\x38\xf2\x05\x10\x88\xca\x34\x41\x4b\xaa\x93\xfa\x76\x43\xdb\x37\xbc\x57\xef\x5a\x5e\xbb\xfc\xb2\x1e\xb6\x9e\xcf\x2d\xd5\x12\x2a\x07\xaf\x8c\xbb\x84\xf8\xc8\x97\x28\xdc\x33\xe6\xbe\x01\x3f\xa1\x10\xc8\x91\x92\x14\x6b\xc4\xdb\xa5\x25\xc5\xe4\x3c\xc9\x11\x15\x04\xc9\x65\x6a\x36\xf2\x1f\x40\xf8\xe3\x77\x59\x46\x70\x54\x24\xb6\x42\xcf\x51\x07\x0e\x3d\x30\x76\x67\x20\x51\x18\x4c\x30\x21\xd6\xc9\x6b\xec\x6a\x30\x93\xdd\x36\xcb\x0d\x44\x91\x6b\x87\xec\x3e\x98\x30\x1e\x10\x96\xb4\xd5\x54\xce\x3b\x0d\xb3\x12\x3c\xf5\xbe\xcb\xaa\xd2\x26\x6c\x9d\x31\x81\xcb\xf9\xbf\x6e\xd2\xce\xe7\xe0\x8e\x9f\x58\xf2\x9c\x5e\xe8\x64\x99\xb6\x39\x6f\x2e\x25\xa3\x4f\x6b\x00\xad\xb2\x66\x75\x8f\xdd\x83\xc3\x03\xcb\x65\x20\xe4\xc3\x72\xea\x22\xe3\x80\xe2\xa6\x76\x66\xa5\xf6\x56\xd1\xaf\xf8\x3f\xe0\x30\x85\x07\xc7\x3c\x53\x8f\xd4\x82\x4b\xd5\xe5\xfd\x25\xd3\x56\x0e\x11\xe0\x39\x88\x20\x86\x09\xca\x89\x74\x4b\x32\x9a\x4a\x8a\xe0\x3e\x43\x34\xae\x37\xf3\xd2\xf5\xea\x50\x3e\xfa\xfb\xe0\x38\xd9\x54\xbd\x60\x74\xd5\x2d\x13\x2a\x5a\xbd\xf3\xdd\x4a\xc6\x51\x02\xde\x7b\x46\x25\x67\x86\x60\x73\x24\x2b\xb5\x6d\xc9\x68\xef\x62\xf1\x1d\xd7\x0b\x9b\x5e\x16\x45\x86\x22\x4c\x93\x06\xcb\x94\x37\x1c\xab\xea\xc8\xaa\x97\x37\xb5\x52\xd5\x2a\x4e\x71\xb6\x08\xdc\xa4\x57\xbe\x9d\x91\x7c\x46\xab\x30\xa8\x58\xa1\x31\x65\x64\x8d\xe8\xb3\x09\x9d\xaa\x78\x54\x86\xa4\x8e\x91\xbb\x6e\x55\xd4\x3b\xf9\x9e\x44\x5c\xba\x57\x9c\xaa\x95\x95\x14\xe9\x1f\x8a\xa2\x00\x65\x59\x00\xb4\xfc\x42\x63\xe0\x57\x65\x49\xfa\xa3\x60\xe7\x22\x87\x8a\x6f\x65\x78\xa8\x98\xd4\xa8\xe2\xa9\xd8\xd4\x49\x14\x37\xbc\x8c\xe5\xb7\xf5\xfb\x4c\x89\xa3\x6e\x56\x55\x12\x48\x1d\x0f\xf7\x44\x52\xc7\xc1\x39\xa1\x7c\x54\x88\x55\x5a\x4b\x45\xb3\x71\xe0\xc7\x58\x34\x8a\xfc\xdf\x4b\x7e\x1d\xf4\xcf\x05\xfa\x97\xed\x86\xfe\x2d\x64\x88\x23\xc9\x78\x89\x7e\xb1\x7e\xbc\xac\x88\xf8\xa6\x81\xe9\xbc\x6c\x9f\x13\x96\x8c\x9b\x2a\x13\x0b\xca\x68\x90\xb2\x19\x4b\x80\x02\xab\xe6\xf7\x2d\x8c\xc9\x39\xe5\x90\x60\x21\x81\x37\x16\x96\xff\x82\x19\x9b\x77\x51\xf9\x6c\x3c\xe9\x97\x93\x05\xbf\xa9\x20\xab\x93\xdb\x84\xfb\x3f\xed\x85\x59\x1d\xdb\x0e\xf6\xb6\x59\xb5\x0d\xf6\x6f\xda\x0d\x7b\x5d\x32\x72\xdd\x25\x23\xed\xc3\xd2\xaf\x5d\x32\xb2\x8c\xc7\xc5\xc9\xd7\x85\xe9\x38\x96\x4e\x0f\x86\x8a\x7d\xca\x16\x5d\x50\x3e\x17\x47\xba\x3c\x5a\x75\xc4\xa6\x53\xbb\x3e\xf7\x30\x0c\x13\x19\x20\x29\x51\x94\x3a\x56\x79\xf6\xbe\xfd\xb0\x6c\xc3\xc0\x61\xaf\x6c\x51\x82\x51\x01\x5d\xe9\xfa\xac\x4a\xd7\x73\x14\x70\x98\x70\x10\x69\x83\xe9\x72\xc1\xaf\x8b\xcd\xe7\x12\x9b\x4f\xb0\x70\x3d\x47\xcb\xdc\x04\x32\xd9\x18\xea\xdf\x15\xec\x3a\xd0\x9f\x0b\xe8\x8f\x56\xb2\x6e\x7b\x3e\x72\xdd\xae\x7c\x44\x77\xfe\x48\x27\x4a\x2b\x72\x91\x14\x11\x9c\xd0\x2a\xcb\x7f\xc5\xb8\x04\xb1\xf7\x85\xe4\x09\x36\x1d\x49\xdf\xe2\x8e\xa4\xe4\x38\xcc\x25\x08\x23\xb6\x37\xc3\x56\xb3\x58\x40\x71\x94\xc6\x9b\x23\x92\xc3\xc8\x0f\x19\x89\xcd\x8d\x17\xf6\xf7\xb4\x1d\xfd\xb5\x18\x2c\x70\x2c\x53\xa7\x2f\x15\x9d\xe7\x58\x78\x3c\xa7\xe7\xcc\x11\x2e\x7a\x6e\x3b\xd7\x79\x16\xd7\x71\xfa\x50\x77\x1a\xae\xe3\xd4\x71\xa8\x12\xa7\x0d\x3d\xaa\x69\xd5\xd4\x6a\x4f\x88\x43\x19\x54\xec\x7f\xf4\x6a\x6e\x04\x36\x7d\x90\xc5\x56\x40\x2e\x96\x5b\x80\x17\x2e\x9a\x1f\xc2\xa6\x3c\x6e\x16\x24\x1c\xc7\x01\xc1\x74\x19\xb2\x42\x26\xdd\xf7\xdd\x87\xb4\x5b\xea\x15\xec\xd4\x72\xe9\xb9\xec\xf5\x5c\x76\x13\x2f\x9b\x3b\x3b\xed\x01\x4e\xa8\x96\xd7\x85\xb1\xd3\x0c\x63\x73\xd4\x85\xb1\x2e\x8c\xbd\x5c\x18\x7b\xd6\x12\x40\xd7\x31\x76\x94\x8e\x31\x83\x0d\xbb\x8e\x31\xd5\xf6\xd0\xba\x29\x3c\x92\x75\x14\xae\xd9\x75\x8a\x29\x69\x4f\xa4\x69\xab\x42\xfb\xd5\xc7\x19\x4a\x74\xc7\xee\x5b\xd5\x7f\x35\x5f\x95\x69\x80\x1a\xfb\xb0\x55\xa4\x42\xb2\x68\xea\x8f\x13\x39\xed\x51\x90\x0b\xc6\xa7\x55\x39\xe0\x88\xd1\xa0\xbc\x41\xc5\x72\x0f\x53\xb7\xd6\x34\x74\x05\x93\xa9\x96\xd8\x2a\x60\x1a\xce\x17\xbd\xff\x6c\x65\xe6\x50\x84\xdb\x2d\xf4\x15\x38\x5c\x97\xf9\xae\x06\x83\xc1\x40\x97\x61\xda\x6a\x7c\xcd\x83\xb5\x96\x06\x9b\x40\xab\x25\x03\xed\xd0\x6a\x47\x2b\x44\x39\x07\xef\x3d\x9b\xcd\x72\xba\xbe\x7e\xc8\xfb\xec\x18\x30\xeb\xe0\xb8\x38\xb0\xb2\x06\x32\x96\x88\xe0\x48\xbf\x57\xda\x27\x8f\xd0\x16\xf9\xe0\xe2\xcd\xe0\xe9\xdf\xeb\xce\x2b\x6c\x45\xee\xce\x2b\x76\x92\x0b\xa7\x6f\x40\xee\x2e\x35\x1f\x5c\x58\xbf\x71\x1c\xbc\x02\x5c\x5e\xb5\x6a\x05\x78\xa9\x74\xc5\x52\x41\x78\xd9\xad\x71\x77\xe7\x96\x92\xd6\xf0\xe5\x34\x64\xb9\xc1\x11\xbf\xa7\x5b\xb7\x54\xb2\xa8\xe5\xa8\xe3\x72\xd5\x5d\xad\x72\x8e\xa5\x10\x75\x47\xcc\xa7\x22\x6e\xfd\xb8\xfd\x83\xe2\x63\x02\xfe\x06\x37\x9c\xe5\xd9\xfa\xc6\xf3\x05\x8e\x13\x90\x9b\x40\xb6\x7a\x5e\xcd\x53\x7d\x95\x42\xdf\x69\xf0\xba\xfd\x5c\x33\x7a\xab\x5d\xd7\x8d\xdf\x56\xeb\xa4\x1b\x41\xd9\x6e\x66\x7b\xfd\xb5\x23\xb7\xed\x66\x09\x0b\xc5\xd6\xe9\x5d\xfb\xc8\xd5\x81\xc7\xcd\x7d\xf1\x5b\xf6\x78\xb4\xe4\xb0\xbf\x75\xf3\xfd\xff\x01\x00\x00\xff\xff\xf9\x0f\x41\xf6\x52\x5f\x00\x00")

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

	info := bindataFileInfo{name: "ui/main.glade", size: 24402, mode: os.FileMode(420), modTime: time.Unix(1541660915, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiPluginAcceptorGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x6d\x6f\xdb\xb6\x13\x7f\xef\x4f\x71\x7f\xfd\x5f\xe4\x61\x8e\x9d\xa6\xdd\xd6\x02\xb6\x8b\xa5\x6b\x82\x02\x43\x57\x20\x19\x86\xbd\x32\x68\xf2\x24\x71\xa1\x48\x8d\x3c\xd9\xf1\xb7\x1f\x28\xd9\x89\x13\x4b\xb6\xe5\xc4\x89\xb2\xa6\xed\x0b\x37\xe4\x1d\x79\x0f\xbf\x7b\xa0\x2f\xbd\x8f\xd7\x89\x82\x31\x5a\x27\x8d\xee\x07\x6f\x3a\xc7\x01\xa0\xe6\x46\x48\x1d\xf5\x83\x3f\x2e\xcf\x8e\xde\x07\x1f\x07\xad\xde\xff\x8e\x8e\xe0\x1c\x35\x5a\x46\x28\x60\x22\x29\x86\x48\x31\x81\xf0\xb6\x73\x72\xd2\x79\x03\x47\x47\x83\x56\x4f\x6a\x42\x1b\x32\x8e\x83\x16\x40\xcf\xe2\x3f\x99\xb4\xe8\x40\xc9\x51\x3f\x88\xe8\xea\x87\xe0\xf6\xa0\xb7\x9d\x93\xe3\xa0\x9b\xef\x33\xa3\xbf\x91\x13\x70\xc5\x9c\xeb\x07\xe7\x74\xf5\xa7\xd4\xc2\x4c\x02\x90\xa2\x1f\x4c\x8a\xcf\x7e\x23\x40\x2f\xb5\x26\x45\x4b\x53\xd0\x2c\xc1\x7e\xc0\x99\x1e\x86\x86\x67\x2e\x18\x9c\x31\xe5\xb0\xd7\x9d\x6f\x28\xdf\x4f\x92\x14\x06\x40\x96\x69\xa7\x18\xb1\x91\xc2\x7e\x30\x45\x17\x0c\x7e\xe1\x1c\x53\x92\x3a\x82\x51\x26\x15\x49\x0d\xc6\x02\x5e\x13\x5a\xcd\x14\x24\x46\x64\x6a\x99\x3b\x8f\xa5\x12\xc5\x67\x7f\x96\x62\x1c\x63\xa3\x04\xda\xee\x6c\x43\x77\x61\xc7\xbd\xdd\x4b\x42\x9f\x5b\x29\x82\xf9\xf2\xf2\xd5\xc7\xd2\xc9\x91\xc2\x60\x70\x69\xb3\xa5\x9b\x6c\xa3\x9b\x32\x9a\x84\xd9\x48\xea\xa1\xc2\x90\x82\xc1\x8f\x35\x28\xac\x8c\xe2\x9a\x24\x64\xd2\x7a\x04\x23\x43\x64\x92\x0d\x69\xac\x99\x0c\x5d\xca\xb8\xd4\xd1\x86\x14\xdc\xa8\x2c\xd1\xeb\x88\xee\x18\xb1\xdc\x90\xbf\xb1\x11\xaa\x60\x71\xcf\x16\xd6\xdc\xd6\xa2\x65\x74\x31\x53\x32\xd2\xc1\xc0\x11\xb3\xb4\x29\x91\xca\xa5\x28\x83\xca\x37\x95\x45\x52\x43\x81\x18\xa6\x39\xc2\x17\x1d\x9a\x4d\xf9\x4e\x2c\x4b\xd7\x88\xcd\x88\xac\x1c\x65\x84\xee\xee\xc2\xe2\xd2\x9c\x1b\xe6\x8e\x07\x63\xa6\x32\xec\x07\x23\xa3\x44\xd0\xbd\xc7\xae\x5b\xce\xaf\xd7\x2d\x2c\x77\xe7\x67\x29\xe3\x57\x52\x47\x6b\x54\x83\x21\x0d\x19\x11\xe3\x71\x30\x38\xde\x54\x70\x32\x69\x7d\xa2\x89\x14\x14\x07\x83\x93\xf2\xfd\xbd\xee\xd2\x7d\xef\x84\x1c\x78\xe9\x1e\x8b\x5a\x3c\xdc\x5f\x4f\x8b\x80\xfe\xb1\x4a\x87\xcf\xe7\x06\x6f\x9e\xc2\xac\x79\x16\xcd\xb5\x33\x9c\xa5\xb6\x26\x5a\xfa\x91\x62\x53\xbe\xb2\x53\x43\x57\xd8\xac\x01\x86\x6e\xa0\x55\x1f\x05\xbf\x5f\x59\x82\x87\xcd\x03\xef\x93\xc4\xe4\x05\xf0\xfa\x1b\x34\xd1\xc6\xff\x6d\xe4\xbe\x66\xde\x07\xd8\xf4\x53\x8c\xfc\xca\x65\xc9\x61\x03\xe1\xfb\xf6\xd5\xb0\xdb\x1b\xf6\xd2\xf7\xd1\x0d\x34\x6a\x69\xcf\xf6\x6a\xd4\xcd\x8c\xfa\x2b\x3a\x6e\x65\x4a\xd2\xe8\x06\x9a\xf6\xa7\x27\x4e\xb7\xc5\x53\x51\x03\x6d\xbd\xfb\x7c\x5b\xbf\x69\x7f\xce\x0c\xfd\x24\x98\x5f\x70\x0c\x71\x0b\x93\x57\xf7\x68\xbe\x7b\x3c\x49\xdc\x68\xa0\x1f\x3c\x4a\x4a\x38\x93\x0a\xe1\x1b\xa3\xb8\x89\x15\xdc\xbb\xdd\x5a\xf6\xd4\x5c\x37\xc9\xae\xab\xdf\xa5\xcb\x85\x2a\x17\x6c\x29\xa2\x85\x52\x61\xca\x28\x0e\xee\xd3\x6e\x29\xf2\x43\xc4\x2e\xa3\xdd\x28\xb4\x95\x11\x6e\x19\xde\xca\x58\xad\x7f\xb6\xae\xa2\x1a\x26\x46\x60\x30\xe0\x31\xb3\xd5\xa4\x65\x38\x82\x4a\x2c\x95\x1d\x85\xd7\x29\xd3\x62\x0b\xed\x86\x52\xa9\xfa\x92\xa5\xc6\xc9\x22\x05\x56\x62\xb7\x14\x8a\x50\x06\x47\xa8\xe1\xbd\xa7\x19\x91\xd1\x85\xfb\x8e\xf2\xcf\x43\x8b\x0e\x99\xe5\x9b\xf8\x6f\xb5\x47\x58\xbc\xc8\x99\xd4\x51\xc2\x63\xa0\xa1\x2e\xa9\x45\x8e\x72\x8c\x6e\x28\x30\x64\x99\xa2\x75\x1c\x5e\xb0\x67\x55\xe6\xfb\x3a\x9e\xf5\x9c\xb5\xc7\x8e\x33\xd4\x52\x20\xe7\xb3\xf7\x96\x26\xe5\xad\xef\xe3\x99\x70\xc7\xaf\x49\x25\x41\xcf\xa4\xa8\x87\xf3\x01\x85\x35\x06\xaf\x56\xe7\xef\x29\x6a\xf8\x3c\x1f\x73\x28\xbe\xd9\xdd\x54\xfe\x87\xba\x51\x1d\xb2\x7a\x51\x6f\xc9\x07\xe7\x21\x6c\xe7\x2d\xca\x56\x85\xec\x87\xd7\x6f\x83\x9f\x32\x84\x1c\xc2\x11\x7c\x09\x21\x2d\xc6\x18\xa4\xbb\x1d\xf3\xd9\xd7\x86\xe6\xf3\x3f\x07\x6d\xa0\x18\x35\x50\x2c\x1d\x28\xa9\xd1\xef\xbc\xd2\x66\xa2\xc1\x68\x35\x05\x16\x12\xda\x39\x13\x5f\x41\x83\x47\xa4\xd4\x51\xeb\xd0\x1f\x70\xb9\x48\x96\x20\xf3\x2b\x0a\x9d\x83\xd0\xd8\x9b\x11\xa3\x82\xda\x3d\xcf\xe8\x84\xe3\x4c\xe1\xcd\xe4\xc4\x71\xe7\xfd\xf1\xdd\x3f\xef\x5e\xc2\x20\xc5\xcf\xcd\x82\x4e\x11\xa7\x5f\x5e\xe7\xb8\x84\x9d\xa9\xc9\x68\xe8\x68\xea\xef\xb8\xea\x25\x61\xfb\xa2\x9d\xe5\xe3\x43\xc3\x62\xb2\xee\x41\x95\x7b\x31\x88\xb4\x26\x79\x6d\x6d\x85\x35\x96\x78\x39\xe5\x7b\xdd\x9b\xbe\xfc\xbe\x90\x33\xcd\x97\x32\x59\xd9\xa5\x57\x7c\x83\x99\xb3\x78\x75\xaa\xa6\x38\x55\x93\x5a\xc2\xed\x46\xbe\x1a\x36\xfa\xf7\x39\x37\x24\xda\x1d\x27\xac\x47\xae\xd6\xeb\x3e\x76\x7e\x3f\x0f\x9b\x7f\x99\x0c\x98\x45\x88\x8c\xd4\x11\x90\x81\x22\xcf\x82\xc6\xc9\xac\xdc\x6c\xb5\x4e\xef\x96\x9f\x20\x0c\xf8\xc2\x77\x36\xae\x0f\x4c\x88\x1c\x6d\x4c\x41\x6a\x91\xb3\xcc\xff\xc7\x01\xd3\x02\x38\xd3\x30\xc2\x19\x4f\x14\xc0\x1c\x48\xd7\x69\xb5\x6e\x3a\xc8\x19\xcb\x36\xc4\x66\x82\x63\xb4\x6d\x70\xb1\xc9\x94\xf0\x54\xc2\x64\x23\x85\x90\x3f\x52\xa0\x68\xc3\x08\x43\x63\x11\xc8\x4e\xa5\xf6\x37\xc5\x31\xea\xbc\x88\x86\x7d\x65\x98\x38\xf0\xf5\x77\xd2\xf6\x67\xcc\x2a\x6b\x38\x37\x8a\xe9\x68\x5e\x74\xa7\x16\x5d\x96\xa0\x03\x49\x7b\x0e\xa4\x96\xb4\x7f\x00\x61\xa6\xb9\xbf\x2f\xe0\x35\xf2\xfc\xe6\x9d\x56\xeb\xc2\xdc\x9c\xc6\x95\xcc\x61\x03\x7b\x9e\xe9\x1e\x14\xa9\xc2\x2f\x03\x6a\x97\x59\x14\x40\x31\xa3\xa2\xec\xf7\xff\xb4\x40\x14\xf3\x23\xbd\x76\x63\x36\xf6\xea\x25\x08\xad\x49\x8a\xb3\x2d\xba\x14\x39\xc9\x31\x02\xb7\xc8\xc8\xd8\x5c\x5b\x7e\x37\xd9\xcc\x51\xb1\x8b\x1b\x81\x9d\x56\xeb\x4b\x98\x2f\x4c\xa4\x52\x73\xe3\x14\xdc\xdb\xfe\xe7\x16\x2e\x3e\x99\x24\x81\xd4\x9a\xc8\xb2\xa4\xd8\x96\x37\x22\x2c\x23\x93\x30\x92\x9c\x29\x35\x2d\xb4\x78\xdb\xb9\xcc\xf4\xee\x4f\x99\xbd\x01\x01\x32\x1e\x03\xc9\x04\xe7\xa2\xcf\xb5\x28\xa9\x9d\xab\x3b\x17\x17\x24\x79\x31\x2d\xe6\x5c\x29\x46\x70\x2c\xc1\xa2\xa1\x99\xc4\x92\xc7\xfe\x52\x7b\xe3\x5b\x9b\x77\x76\xf0\x6a\x5e\xdd\xb5\xc0\xa3\x74\x2e\xb0\xa2\x7b\x81\xca\x7c\xb8\xa2\x1c\x01\x9a\xa6\x37\x30\x5c\x1f\x84\xce\xec\xf2\xec\x5e\x99\xba\x9e\x25\x08\x15\x4f\x87\xd7\xb3\x50\xb4\xa2\x6a\x2b\xe9\x2c\x62\x26\xcc\x64\xe8\x95\x11\x0c\xd6\xd4\xde\xa5\xf1\xba\x5c\x5d\x55\x31\xfb\x41\x2a\x7b\xa8\xda\x2a\x55\x57\x1a\x86\x2f\x90\x67\x56\xd2\x14\xbe\x1a\x92\x1c\x61\x3f\x55\xc8\x1c\xb6\x3d\xd0\x44\x8e\xc0\x7b\xa0\xbc\x87\xe4\x83\xba\x57\xd9\x0c\x69\xb0\x1e\x6d\xf0\x68\x88\xab\xe0\xe5\xc5\x8e\xac\xc9\xb4\xb8\x61\xf8\x7f\x0c\x31\x3c\xf9\x50\xfc\xad\x62\xb5\x0a\xc0\xb0\xa2\xa8\xad\x00\x32\x94\x82\x79\x28\x09\xef\x3f\x9f\xcf\x15\xbe\xf4\xeb\x5d\x1b\x9c\xb1\x71\x64\x79\xce\x9a\xf4\xfd\x53\x95\xa4\x77\x85\x5c\x58\xbc\x5d\xe8\x75\x17\x7e\x87\xf0\xdf\x00\x00\x00\xff\xff\xa7\xe5\xb2\xfa\x9c\x38\x00\x00")

func uiPluginAcceptorGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiPluginAcceptorGlade,
		"ui/plugin-acceptor.glade",
	)
}

func uiPluginAcceptorGlade() (*asset, error) {
	bytes, err := uiPluginAcceptorGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/plugin-acceptor.glade", size: 14492, mode: os.FileMode(420), modTime: time.Unix(1541671899, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiStorageLoadGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4d\x6f\xe3\x36\x10\xbd\xe7\x57\xb0\x3c\xb5\x28\x1c\x3b\xd9\x16\xe8\x41\xd6\x02\xbb\x68\x72\x68\xd0\x16\xd8\x2d\x8a\x9e\x04\x9a\x1a\xcb\xd3\xd0\x1c\x95\x1c\xd9\x56\x7f\x7d\x21\x6b\xd3\xd8\x31\x65\x7d\xd8\x0d\x8c\xa0\x37\x21\x9c\xc7\x79\x9c\x99\xc7\xe1\x38\xd1\xfb\xcd\xd2\x88\x15\x38\x8f\x64\xa7\xf2\xe6\x7a\x22\x05\x58\x4d\x29\xda\x6c\x2a\x7f\xfb\x7c\x37\xfa\x41\xbe\x8f\xaf\xa2\xaf\x46\x23\x71\x0f\x16\x9c\x62\x48\xc5\x1a\x79\x21\x32\xa3\x52\x10\xef\xae\x6f\x6f\xaf\x6f\xc4\x68\x14\x5f\x45\x68\x19\xdc\x5c\x69\x88\xaf\x84\x88\x1c\xfc\x55\xa0\x03\x2f\x0c\xce\xa6\x32\xe3\xc7\x6f\xe5\xb3\xa3\x77\xd7\xb7\x13\x39\xde\xda\xd1\xec\x4f\xd0\x2c\xb4\x51\xde\x4f\xe5\x3d\x3f\xfe\x8e\x36\xa5\xb5\x14\x98\x4e\xe5\xba\xfe\xae\x0c\x85\x88\x72\x47\x39\x38\x2e\x85\x55\x4b\x98\x4a\xad\x6c\x32\x27\x5d\x78\x19\xdf\x29\xe3\x21\x1a\x3f\x19\x84\xed\x19\xd9\x80\x14\xec\x94\xf5\x46\xb1\x9a\x19\x98\xca\x12\xbc\x8c\x1f\x48\xa5\xe2\x13\x93\x53\x59\xeb\x2e\x35\xa5\x24\x27\x8f\x8c\x64\x65\xac\xa1\x3a\x77\xab\xf3\x32\x87\x64\x81\x96\x65\x9c\xa2\x32\x94\xb5\x01\x32\xa7\x56\xc8\x65\xe3\xfe\x7a\x81\x26\xad\xbf\x2b\xb0\x51\x1a\x16\x64\x52\x70\xe3\x2f\x06\xe3\x1d\x8b\x17\xd6\x07\x31\xff\x40\x1b\xf9\xb4\x7a\x48\x65\x85\x1e\x67\x06\x64\xfc\xd9\x15\x07\xf1\x19\x92\x99\x10\x66\xa9\x5c\x86\x36\x31\x30\x67\x19\x7f\xdf\x03\xe1\x30\x5b\xf4\x84\x30\xe5\xfd\x00\x33\x62\xa6\x65\x47\x0c\x39\x04\xcb\xaa\xae\x8e\x15\x38\x46\xad\x4c\x17\xa0\xcf\x95\x46\x9b\x35\xb9\xd9\xcb\x61\x38\x8f\x0f\x6a\x06\x46\xee\xda\x0c\xc8\xe6\xd0\x8c\x86\x70\x66\x4b\x28\xa4\xb9\x5f\x72\xb0\xe2\x0f\x2a\x9c\xf8\xf4\x91\x96\xcb\x26\xf9\x35\x6d\xbc\x51\x06\x33\x2b\xe3\x49\x23\x40\x31\x3b\x9c\x15\x0c\x7e\x7f\x61\x77\xe9\x49\xd3\xb0\xad\x21\xb1\x52\xa6\x80\xa9\x9c\x91\x49\xe5\xb8\x15\xe5\xf1\x6f\xf8\x17\x73\xf3\xdd\x64\x32\x79\x09\x8a\xc6\x61\x12\xd1\xb8\xce\xdc\xde\xdf\x72\xa5\x1f\xd1\x66\xc7\x8f\x0d\x9b\x5c\xd9\xb4\x67\x12\xe6\x68\x4c\xbf\x74\x3f\xdf\x6e\x0d\xf1\x8d\xc6\x07\x74\xf7\x2e\x1c\xd1\xb1\x60\xef\x1d\xa6\x97\x54\xaf\x8e\xd6\xc9\x71\x19\x06\xbd\x91\x29\x96\xb6\x0b\xf0\x20\x24\xe1\xb0\x84\x74\x3c\x38\x36\xa7\xc4\x27\x84\x6d\xd6\xf4\x17\x0d\x8b\x9f\xd5\xb2\xd7\x8e\xad\x62\x0e\x4b\x46\x34\xca\x26\x48\x1b\xe6\x9c\x28\x66\xa5\x17\x47\x3d\x05\x7a\x37\xe5\x5d\x80\x01\x4d\x88\x90\x2e\xc4\xdb\x2f\x84\x9f\xa0\x14\x5f\xff\xaa\xbc\x5f\x93\x4b\xbf\x79\xa3\xa5\x70\xf3\xba\xa5\xf0\xa3\x65\x57\xd6\xcf\x62\xa8\x3e\x93\x8a\xd1\x2b\x95\xc6\x10\x68\x0a\x73\x55\x18\xee\x0f\x66\x22\xc3\x98\x27\x0c\x1b\x0e\x5f\x33\x60\xaa\xd0\x58\xb2\x23\xd8\xa0\x67\xb4\x99\x60\x12\xda\x81\xaa\x9a\x33\xac\xfb\x78\x5b\x3c\xb5\xd4\xbe\x34\x73\x87\x4b\xe5\xca\x04\x35\xd9\x44\x69\xc6\x55\x4d\x72\x80\x90\x3c\x68\xb2\xe9\x99\x36\xdb\xe3\xe5\xc1\x56\x9d\x7c\x75\x06\x56\xa7\x6c\xb5\xc7\xa9\x6b\x7e\xef\xd0\x1c\x71\x72\xe6\x5b\xe0\x88\x98\x2f\xaa\x21\x1c\xdc\x02\x8f\x50\x5e\xe8\x25\x30\x58\x59\x5b\xb2\x68\xb6\xd3\x6f\x4b\xa9\x5d\x4a\x15\x9c\xa5\x17\x9c\x3e\x16\xf4\x79\x22\xd7\x53\x41\xaf\xd7\xf1\xf3\x58\xd0\x70\xde\x73\x8d\x05\x1f\x0a\x66\xb2\xfb\xbf\x4a\x84\x08\xbd\xf2\x2c\x5b\x52\xc1\x89\xe7\xb2\xf2\xe8\x73\x07\x2a\x3d\xf5\xa1\x5f\x1f\xb4\xd6\xf3\x6c\xfb\x9d\x18\x52\x2f\x47\xa2\x30\x9b\xa6\xc7\xd7\x03\x35\xf3\x1a\x1c\xc4\x96\x40\xf6\x85\x3a\xd0\x80\x2b\xf0\x5d\x5f\x09\xa7\xeb\xbc\xd3\xe4\x1c\x02\xb6\x4f\xcf\x21\x54\xeb\x04\x2d\xfe\xa3\x06\x11\xa8\x28\x6d\xc8\x77\x79\x29\x36\x97\xd4\xc7\x6a\x87\xff\x6b\x6a\xc8\xb5\x1b\xc2\x9d\x5a\x52\x17\xd2\x6d\x5e\xed\x47\xa8\xdb\xa1\xdd\x66\xff\x8c\x3b\x8b\xcf\x0b\xd1\x78\xe7\xff\x16\xff\x04\x00\x00\xff\xff\x48\x93\x6f\x0c\x10\x19\x00\x00")

func uiStorageLoadGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiStorageLoadGlade,
		"ui/storage-load.glade",
	)
}

func uiStorageLoadGlade() (*asset, error) {
	bytes, err := uiStorageLoadGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/storage-load.glade", size: 6416, mode: os.FileMode(420), modTime: time.Unix(1541606376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiSystrayReplacerGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\xb1\x6e\xc2\x30\x10\x86\x77\x9e\xe2\x7a\x6b\xe5\x18\xc2\xd2\xc1\x31\x1b\xa8\x3b\x55\xc7\xca\xb1\x8f\xc4\x8d\x63\x53\xdb\x81\xf6\xed\x2b\x0a\x08\x54\x77\xea\x76\xd2\xff\xfd\xba\xef\x74\x62\xf5\x39\x3a\x38\x50\x4c\x36\xf8\x06\x17\xd5\x1c\x81\xbc\x0e\xc6\xfa\xae\xc1\x97\xed\x9a\x3d\xe1\x4a\xce\xc4\x03\x63\xb0\x21\x4f\x51\x65\x32\x70\xb4\xb9\x87\xce\x29\x43\xb0\xac\xea\xba\x5a\x00\x63\x72\x26\xac\xcf\x14\x77\x4a\x93\x9c\x01\x88\x48\x1f\x93\x8d\x94\xc0\xd9\xb6\xc1\x2e\x0f\x8f\x78\x5b\xb4\xac\xea\x39\xf2\x1f\x2e\xb4\xef\xa4\x33\x68\xa7\x52\x6a\x70\x93\x87\x57\xeb\x4d\x38\xe2\x29\x04\x10\xfb\x18\xf6\x14\xf3\x17\x78\x35\x52\x83\x5a\xf9\xb7\x5d\xd0\x53\x42\xb9\x56\x2e\x91\xe0\x57\xe0\xc2\xeb\xde\x3a\x73\x9e\x4f\x6d\xa7\x34\xf5\xc1\x19\x8a\xfc\x02\xf0\x3b\xe2\x17\x5d\xa8\x3c\x8f\xaa\x23\xbc\xe6\xa5\xcd\xc1\x26\xdb\x3a\x42\xb9\x8d\x53\xa1\xf2\x1f\xfd\xbf\x3a\x29\x07\x3d\xa0\xec\xf2\xc0\x54\x1b\xa6\x5c\x76\x04\x3f\x9b\x17\x17\xde\x02\xc1\xef\xbe\xf3\x1d\x00\x00\xff\xff\xf5\xed\x48\x1a\xf6\x01\x00\x00")

func uiSystrayReplacerGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiSystrayReplacerGlade,
		"ui/systray-replacer.glade",
	)
}

func uiSystrayReplacerGlade() (*asset, error) {
	bytes, err := uiSystrayReplacerGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/systray-replacer.glade", size: 502, mode: os.FileMode(420), modTime: time.Unix(1541518671, 0)}
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
	"ui/key-cert-editor.glade": uiKeyCertEditorGlade,
	"ui/main.glade": uiMainGlade,
	"ui/plugin-acceptor.glade": uiPluginAcceptorGlade,
	"ui/storage-load.glade": uiStorageLoadGlade,
	"ui/systray-replacer.glade": uiSystrayReplacerGlade,
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
		"key-cert-editor.glade": &bintree{uiKeyCertEditorGlade, map[string]*bintree{}},
		"main.glade": &bintree{uiMainGlade, map[string]*bintree{}},
		"plugin-acceptor.glade": &bintree{uiPluginAcceptorGlade, map[string]*bintree{}},
		"storage-load.glade": &bintree{uiStorageLoadGlade, map[string]*bintree{}},
		"systray-replacer.glade": &bintree{uiSystrayReplacerGlade, map[string]*bintree{}},
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

