// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// commands/provider/assets/readme_template.md
package assets

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

var _commandsProviderAssetsReadme_templateMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xc1\x6e\x9c\x30\x10\x86\xef\x7e\x8a\x51\xc8\x09\x15\x1e\x20\x52\x6f\xe9\xa1\x8a\xd4\x48\x55\x1e\xc0\xee\xee\x6c\x6a\xb1\x01\x8a\x49\x24\xe4\xf8\xdd\xab\x19\x1b\x6c\x03\x4d\x57\x8a\xf6\x12\x86\xf1\x3f\x7f\x3c\xff\x7e\xa2\x00\x6b\xeb\x1f\xea\x05\x9d\x13\xc2\xda\xfa\x1e\xcd\x61\xd0\xfd\xa8\xbb\xd6\x39\x61\xad\x3e\x41\xfd\x13\xfb\xce\xe8\xb1\x1b\x26\xe7\x84\x94\xf2\x97\x32\xbf\xc5\x0b\x9a\xe7\xea\xd0\x0d\x08\x06\x87\x37\x7d\x40\x38\x62\x7f\xee\x26\xf2\x5b\x0f\x08\x6b\xb1\x3d\x3a\x27\x0a\xb8\xc7\x93\x6e\x35\xd9\x1b\x11\xec\xbf\xbd\x61\x3b\x1a\xe7\x0a\xf0\x95\xb0\x76\x50\xed\x33\xc2\x6d\x83\xd3\x17\xb8\x45\xea\xc2\xdd\xd7\xa8\x14\x05\x5d\x5b\x9f\x00\xff\x84\x63\x5e\x01\x6e\x6e\x9c\xb3\x96\xc6\xe8\x09\x78\x36\x08\x5c\xa5\x22\x6a\xf8\xdb\x08\xf6\x83\x06\xa7\x3b\x90\xf3\x9c\xa4\x6b\x05\x7d\x1e\x86\xbf\xee\x7c\xa4\x46\xe5\xdc\x3b\x94\x25\x99\x96\x25\x50\xf9\x80\x53\xa8\x9e\xa6\x7e\x6e\x26\x26\xd4\x11\xef\x50\x55\x15\xec\xfc\x8d\x7b\x1f\xd5\xa8\x1e\x78\x77\xaa\x68\xf5\xcd\x7f\x5d\xd6\x27\x45\xb6\x7d\x18\x5e\x27\x10\x75\x4b\x00\x7c\x41\x99\x8e\xc8\xa4\x51\xd3\x0e\xbe\xe5\x17\xe7\xe6\x23\x2f\xa2\xce\x34\x2d\xbb\xf0\x22\xcb\x12\x82\xe7\x3c\x9c\x45\xc7\xbb\x85\xe3\x9d\x47\xe0\xe0\x49\x99\x86\x31\xe0\x62\x4d\xc1\xa8\x4c\xc3\x10\x04\x59\xc6\x00\x1d\xfe\x0f\x81\xa8\x49\x08\x20\xb3\x5d\x00\x58\xbd\xfb\xfb\xf3\xc9\xf7\xb6\x7f\x65\x66\x8b\x02\x7c\x2d\xc4\x95\x60\xd0\x64\xef\x69\xe0\x92\x71\xc8\x2e\x91\xf1\xc0\x9a\x2c\x8b\xd9\x60\x1d\x48\xa2\xdc\x22\x11\x87\x64\xda\xd9\x40\xe1\xbb\x17\x50\xe1\x85\x1f\x63\x91\x26\xfc\xf8\x3a\xc6\x88\xc3\x8b\x88\xa9\x74\xdc\xf1\xb1\xf8\x3a\xe6\x92\x8e\x26\x90\x78\x59\x16\xcd\xe2\xb2\xce\x26\xd5\x26\xb8\x78\xe7\x08\x4c\x32\xef\xb1\x09\x73\x2b\x70\xae\x44\xc6\x3f\x32\x98\xef\xb0\xfd\x56\x7c\x3a\x81\x05\x8f\x74\xf1\xac\xb5\x01\x24\xb4\x2f\x20\x64\x2f\xbb\x88\xc8\xc7\x5f\x90\xbf\x01\x00\x00\xff\xff\x2c\x62\x9e\x8a\xc7\x06\x00\x00")

func commandsProviderAssetsReadme_templateMdBytes() ([]byte, error) {
	return bindataRead(
		_commandsProviderAssetsReadme_templateMd,
		"commands/provider/assets/readme_template.md",
	)
}

func commandsProviderAssetsReadme_templateMd() (*asset, error) {
	bytes, err := commandsProviderAssetsReadme_templateMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/provider/assets/readme_template.md", size: 1735, mode: os.FileMode(420), modTime: time.Unix(1543895910, 0)}
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
	"commands/provider/assets/readme_template.md": commandsProviderAssetsReadme_templateMd,
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
	"commands": &bintree{nil, map[string]*bintree{
		"provider": &bintree{nil, map[string]*bintree{
			"assets": &bintree{nil, map[string]*bintree{
				"readme_template.md": &bintree{commandsProviderAssetsReadme_templateMd, map[string]*bintree{}},
			}},
		}},
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
