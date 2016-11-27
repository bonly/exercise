package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _res_texture_png = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x02\xd2\xfc\x20\xcc\xc1\x06\x24\xad\xaf\x4d\xf5\x02\x52\x6c\x49\xde\xee\x2e\x0c\xff\x41\x70\xc1\xde\xe5\x93\x81\x22\x9c\x05\x1e\x91\xc5\x0c\x0c\xdc\xc2\x20\xcc\xc8\x30\x6b\x8e\x04\x50\x90\xbd\xc4\xd3\xd7\x95\xfd\x3e\x0b\x2f\xaf\xb1\xe8\x4e\x8b\x8c\xf3\x40\x21\xd9\xcc\x90\x88\x12\xe7\xfc\xdc\xdc\xd4\xbc\x12\x06\x10\x70\x2e\x4a\x4d\x2c\x49\x4d\x51\x28\xcf\x2c\xc9\x50\x70\xf7\xf4\x0d\x48\xd1\x4b\x65\x07\x8a\x7b\x7a\xba\x38\x86\x68\x9c\x7f\x3b\xc9\x90\xeb\x80\x01\x0f\xb3\x7f\x6d\xd0\xbf\xff\xdb\x4f\x5f\x28\xeb\xb2\x6d\xb9\xb4\xd2\xd6\x40\x4e\x25\x71\x83\xec\xd4\x2e\x95\x55\x2a\x8e\xbe\x4f\x2d\x2b\x0e\x88\x1f\x66\x7c\x73\xa4\xfa\x76\xee\x8d\x77\x61\xf9\xfe\x6e\xb1\x5c\x12\xcb\x97\x84\xeb\xd7\xfe\xb4\xcf\xe9\xe7\x54\x69\x70\x17\xd7\xd4\xb4\xff\x0a\xb2\xcd\xd3\xd5\xcf\x65\x9d\x53\x42\x13\x20\x00\x00\xff\xff\x0f\x5f\xd2\x1f\xe5\x00\x00\x00")

func res_texture_png_bytes() ([]byte, error) {
	return bindata_read(
		_res_texture_png,
		"res_texture.png",
	)
}

func res_texture_png() (*asset, error) {
	bytes, err := res_texture_png_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "res_texture.png", size: 229, mode: os.FileMode(420), modTime: time.Unix(1428933081, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"res_texture.png": res_texture_png,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"res_texture.png": &_bintree_t{res_texture_png, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

