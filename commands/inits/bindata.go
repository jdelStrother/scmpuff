// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/aliases.fish (131B)
// data/aliases.sh (131B)
// data/git_wrapper.fish (847B)
// data/git_wrapper.sh (662B)
// data/status_shortcuts.fish (866B)
// data/status_shortcuts.sh (1.288kB)

package inits

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _dataAliasesFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xc9\xc1\x0d\x83\x21\x08\x06\xd0\xbb\x53\x78\x63\x0a\x67\x69\x08\x88\x35\xb5\xb1\xf1\xc3\xfd\x7b\xe0\x0f\xd7\xf7\x78\x4d\x46\x1d\x68\x04\xf9\xfe\xae\xd9\x0b\xce\x7e\x41\xe5\x19\x6e\x34\xa6\x57\x56\x4d\xd2\x20\x9d\x66\x69\x2b\x6c\xed\x91\x24\x3b\x4c\xde\x5d\x3e\xfb\x7a\xc6\x41\xc4\xe9\xe8\x4e\xe5\x1f\x00\x00\xff\xff\xb3\xf6\xf4\x39\x83\x00\x00\x00")

func dataAliasesFishBytes() ([]byte, error) {
	return bindataRead(
		_dataAliasesFish,
		"data/aliases.fish",
	)
}

func dataAliasesFish() (*asset, error) {
	bytes, err := dataAliasesFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aliases.fish", size: 131, mode: os.FileMode(0644), modTime: time.Unix(1558658378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x33, 0x16, 0x65, 0xc3, 0xc7, 0xa3, 0xb5, 0x63, 0xe1, 0x80, 0x8, 0x9, 0xc3, 0xaf, 0x35, 0xe5, 0xc5, 0x17, 0x0, 0x9, 0x98, 0x7c, 0xe7, 0xc8, 0xd5, 0xbf, 0x19, 0xcb, 0x77, 0x79, 0x65}}
	return a, nil
}

var _dataAliasesSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xc9\xc1\x0d\x83\x21\x08\x06\xd0\xbb\x53\x78\x63\x0a\x67\x69\x08\x88\x35\xb5\xb1\xf1\xc3\xfd\x7b\xe0\x0f\xd7\xf7\x78\x4d\x46\x1d\x68\x04\xf9\xfe\xae\xd9\x0b\xce\x7e\x41\xe5\x19\x6e\x34\xa6\x57\x56\x4d\xd2\x20\x9d\x66\x69\x2b\x6c\xed\x91\x24\x3b\x4c\xde\x5d\x3e\xfb\x7a\xc6\x41\xc4\xe9\xe8\x4e\xe5\x1f\x00\x00\xff\xff\xb3\xf6\xf4\x39\x83\x00\x00\x00")

func dataAliasesShBytes() ([]byte, error) {
	return bindataRead(
		_dataAliasesSh,
		"data/aliases.sh",
	)
}

func dataAliasesSh() (*asset, error) {
	bytes, err := dataAliasesShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aliases.sh", size: 131, mode: os.FileMode(0644), modTime: time.Unix(1558658378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x33, 0x16, 0x65, 0xc3, 0xc7, 0xa3, 0xb5, 0x63, 0xe1, 0x80, 0x8, 0x9, 0xc3, 0xaf, 0x35, 0xe5, 0xc5, 0x17, 0x0, 0x9, 0x98, 0x7c, 0xe7, 0xc8, 0xd5, 0xbf, 0x19, 0xcb, 0x77, 0x79, 0x65}}
	return a, nil
}

var _dataGit_wrapperFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x51\x4d\x6f\x13\x31\x10\xbd\xef\xaf\x78\x4a\x23\x35\x41\x98\xa5\x12\x17\xe0\x46\x51\x11\x87\x4a\x48\xc0\x09\xa1\xc8\xeb\x9d\x5d\x5b\xf8\xab\xf6\x38\x09\xff\x1e\xad\x37\xa4\x69\x69\x24\xd4\x83\x0f\x9e\x99\xf7\x31\x6f\x2e\xf0\x41\x66\xea\x11\x3c\x34\x73\xcc\xef\xda\x76\x34\xac\x4b\xf7\x4a\x05\xd7\xca\xd4\x91\xe5\x76\x30\x59\x8b\x68\xcb\x68\xbc\xc8\xca\xc5\x32\x0c\x2f\x9b\x0b\xec\x0c\x6b\x1c\xfe\x82\xf6\xa4\x90\x4b\x8c\x21\x31\x56\x4f\x70\xb9\x14\x58\xb7\x87\xf1\x36\x16\x6b\xdb\x37\x6f\xd7\x4d\x26\x86\xd8\xe3\xeb\xf5\xed\x97\xef\x37\x37\x9b\x4f\x9f\xbf\x6d\xae\x6f\x3f\x62\xb5\xd3\x46\x69\x8c\x86\xd7\x4d\x63\x06\xf0\xef\x48\x10\x77\xd0\xa5\x6b\x00\xe0\x0c\x6c\xa1\x4b\xb7\x68\xc8\xf7\x4d\x45\xf9\xc0\x47\xe4\x41\xb9\xa2\x69\x6f\x18\x57\xf3\xdc\x50\xbc\x62\x13\x7c\x86\xa0\x49\xef\xbe\x52\x7f\xd3\xf8\x5f\x8a\xe5\x23\xb9\xf7\x08\xe9\x9c\x93\x07\x0b\x4c\x24\xd3\x12\x94\x19\x2b\x15\x8a\x67\x2c\x65\x1a\xb7\x6b\x08\xba\xc3\xeb\xda\xaf\xbe\xb6\xd2\xfe\xa3\x72\xec\x56\x25\x8b\x8c\x65\x66\xc9\x25\x1f\x1b\x89\xb8\x24\x8f\xe5\x5c\xa9\x5b\xd5\xf9\x9d\x61\xa5\x67\xa5\x1f\x57\x3f\x6b\x4d\xc9\x4c\x50\xc1\x39\xc3\xe8\xac\x74\x04\x1b\x46\x24\xea\xa6\xba\xa3\x34\xd2\xbd\xdc\x1c\x18\xea\x65\x85\xc0\xe2\xb1\xb1\xc5\x4c\x7d\xc2\xab\x49\xfd\x0a\x85\xd1\x9b\x61\x40\x72\x48\x94\x89\xcf\x11\x26\xb2\x92\xcd\x96\xfe\x8f\x5b\xf6\xfd\xf3\x9c\x9d\x00\x36\x27\xb9\x55\xd2\xcb\x17\x97\x0f\xb3\x9f\x92\x91\xbe\x7f\x8a\x71\x95\x39\x19\x3f\x82\xb2\x92\xb1\x9a\x9e\x4f\x78\xcc\x7c\x7a\x7f\x02\x00\x00\xff\xff\xfc\x8a\x95\xd2\x4f\x03\x00\x00")

func dataGit_wrapperFishBytes() ([]byte, error) {
	return bindataRead(
		_dataGit_wrapperFish,
		"data/git_wrapper.fish",
	)
}

func dataGit_wrapperFish() (*asset, error) {
	bytes, err := dataGit_wrapperFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/git_wrapper.fish", size: 847, mode: os.FileMode(0644), modTime: time.Unix(1558658378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9f, 0x8b, 0x9e, 0x1f, 0xbd, 0x42, 0xb8, 0x67, 0xc, 0x8, 0xfa, 0xfa, 0xd3, 0x77, 0x4, 0x9c, 0x1a, 0xcc, 0x6a, 0x7d, 0xb3, 0x4c, 0x4e, 0x3a, 0x20, 0x5c, 0xcb, 0xc7, 0x22, 0x97, 0xec, 0x85}}
	return a, nil
}

var _dataGit_wrapperSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\x51\x6b\x14\x31\x14\x85\xdf\xf3\x2b\x0e\xd3\xc5\x76\xc5\xa5\xd4\xd7\xc5\x22\x54\x2a\x3e\x14\x44\x2d\xbe\x08\x4b\x26\x73\xb3\xb9\x98\x49\x42\x72\x33\xbb\xc5\xf1\xbf\x4b\xa6\x56\xb0\x58\xc5\xd7\x7b\x4e\xbe\xef\xe4\x04\x1f\x68\x8c\x13\x41\x87\x3b\xd0\x91\x8b\x70\xd8\x63\xcf\x02\xed\x59\x17\xc4\x0c\x5b\x83\x11\x8e\x41\xd5\x70\x7f\x6b\xe9\x25\xce\x07\x9a\xce\x43\xf5\x1e\x2f\x2f\x9f\x5d\xa8\x1a\x0a\x09\x36\xf6\xcf\xa9\x3a\xc1\x6d\x21\x88\x23\xd8\x76\x4c\x5a\x1c\x24\x2e\x65\x89\xd0\x53\xe4\x01\x1c\x2c\x07\x16\x82\x8f\x31\xe1\xc0\xe2\x96\xfc\x97\x9f\x8e\x29\x66\xc1\xc7\xab\x9b\xf7\xb7\xd7\xd7\xbb\xb7\xef\x3e\xed\xae\x6e\xde\xbc\xea\x56\x67\x5f\x0e\x8e\xcd\xd2\x5e\x77\xcd\xf5\x39\xeb\xb4\xbc\x5d\x20\xcd\x7a\xea\x6a\x7f\xda\x4e\xae\xf6\x38\x64\x9d\x12\xe5\x17\x60\x0b\x0e\x45\xb4\xf7\x34\x28\xb6\x90\xbb\x44\x68\x8d\xc7\x1f\xd8\x36\x48\xc0\x53\x0b\x5c\xed\xbb\x2d\x2c\x2b\xf5\x30\xb6\xa9\xce\xd6\xf8\xa6\x00\xa3\x0b\x61\x75\x01\x0e\x0a\x00\x4c\x1c\x47\x96\xb9\xf7\x7a\xa4\xd9\xc7\xfd\x9c\xa9\xd7\x85\xe6\x91\xf2\x9e\xd6\x4b\x05\x28\x66\x4c\xd5\x5a\xd0\x91\x0c\x36\x1b\x74\xab\x47\xce\x0e\xdd\xea\x75\xb7\xdd\xde\x23\x1d\x99\xaf\xb1\xca\x3c\xb0\xb5\x73\x1e\xe7\x4c\x85\xe4\x09\x56\x26\xaf\x85\x27\xfa\x37\x56\x0f\xc3\x7f\xee\xf9\xbd\xbd\x2b\xa2\xa5\x96\x9f\xb4\xe7\x0f\xac\xbf\x49\xa9\x68\xa3\xbe\xab\x1f\x01\x00\x00\xff\xff\x0a\x57\xfa\x78\x96\x02\x00\x00")

func dataGit_wrapperShBytes() ([]byte, error) {
	return bindataRead(
		_dataGit_wrapperSh,
		"data/git_wrapper.sh",
	)
}

func dataGit_wrapperSh() (*asset, error) {
	bytes, err := dataGit_wrapperShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/git_wrapper.sh", size: 662, mode: os.FileMode(0644), modTime: time.Unix(1558658378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0x6c, 0xb3, 0x23, 0xc9, 0xa5, 0xd6, 0x4b, 0x8b, 0x7e, 0x93, 0xcd, 0xeb, 0x5f, 0x2f, 0x0, 0xed, 0x55, 0x9f, 0xc6, 0xd4, 0xc2, 0xfb, 0x3a, 0x74, 0xb7, 0x2e, 0xf6, 0xb2, 0x82, 0x1c, 0x1b}}
	return a, nil
}

var _dataStatus_shortcutsFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x8f\x9b\x30\x10\x85\xef\xfc\x8a\x27\xaf\x25\x82\x5a\xc7\x4d\xf7\xd6\x63\xa5\xfe\x8a\xdd\x6c\x44\x60\x00\xab\x8e\xa1\xf6\x98\x8d\xd4\xed\x7f\xaf\x4c\xb6\x04\x92\x4a\x55\x39\xce\xcc\x37\xef\xbd\x31\x0f\xf8\x5a\x06\xaa\xd1\x3b\x74\xcc\x43\xf8\xa2\x75\x6b\xb8\x8b\xc7\x6d\xd5\x9f\x74\xe9\x8f\x64\x59\x37\x26\x74\x6a\xb0\xb1\x35\x4e\x85\xea\x34\xc4\xa6\xf9\x98\x3d\xe0\xd5\x70\x87\xd4\x7b\x44\x63\xce\xff\xc7\xeb\x21\x5a\xab\x1f\xb3\x26\xba\x8a\x4d\xef\xf0\x5e\x3f\x04\x2e\x39\x86\x0c\xc0\x5c\xaa\x2c\x95\xfe\x30\x96\xfe\xbd\x4c\x0c\x65\xcf\x73\x9b\xdc\x78\xa8\xba\xd2\x43\x90\x58\x0c\xa0\x3a\xd5\x87\x3e\xf2\x10\x19\x1b\x1d\x83\xd7\x47\xe3\x34\xb9\xf1\x0f\x88\x8b\x14\x94\x6a\x8c\x25\x6b\x02\x43\x96\xbe\x1d\xf1\xa2\x6b\x1a\xb5\x8b\xd6\x16\xcb\x7d\x14\x20\xe4\x85\x11\xd9\xd4\x30\x0d\x98\x12\x46\x01\xca\x11\x3e\x4d\xd5\xf4\xb5\x86\xb1\x48\x92\x3e\x4f\x1c\xbd\x83\x5c\x94\xc9\xd5\xd9\x52\x21\xf9\x08\xd8\x04\xf6\xc6\xb5\x08\x83\x35\x8c\x67\x86\xbc\x26\x79\xda\xed\x8b\x95\xf4\xa6\xea\xa3\x63\xc8\x09\x2d\xa0\x5a\x5e\xb8\x68\x7a\x0f\x82\x71\xd8\x04\xfa\x71\x33\x5a\xcc\x53\xb3\x81\xf6\x9c\x02\xde\x5c\x55\x08\x49\x02\xe2\x42\x3d\x49\xda\x8b\x19\x4c\xf6\x57\x31\x92\x9e\x35\x6e\x92\x5c\x9a\xfe\xbc\xdd\xaa\xdd\xfe\xca\x55\x5d\x0f\x99\x06\x67\x7c\x5a\x71\xf7\x2f\xfc\xf5\xe1\xff\xf9\xee\xcb\x7e\x62\x53\x7a\x86\x3a\xe3\x0d\xe5\xeb\x77\xe4\x3f\x07\x6f\xd2\x1d\x76\xbf\x72\xbc\xa1\xf5\x34\x40\x7d\x43\xfe\x92\xdf\x65\xcf\x9f\xeb\x0f\x79\x71\xcd\x36\x4e\xc1\x6e\xf7\x67\xab\x23\x12\xe4\xb8\x8a\xf5\x3b\x00\x00\xff\xff\xed\x91\x51\xd2\x62\x03\x00\x00")

func dataStatus_shortcutsFishBytes() ([]byte, error) {
	return bindataRead(
		_dataStatus_shortcutsFish,
		"data/status_shortcuts.fish",
	)
}

func dataStatus_shortcutsFish() (*asset, error) {
	bytes, err := dataStatus_shortcutsFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/status_shortcuts.fish", size: 866, mode: os.FileMode(0644), modTime: time.Unix(1558658378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6c, 0x5c, 0xb0, 0x18, 0x5e, 0xd, 0xc7, 0xf6, 0xa9, 0x92, 0x4e, 0x2b, 0x68, 0xd2, 0x22, 0xcb, 0xf1, 0xe6, 0xdf, 0x51, 0x49, 0x57, 0x36, 0xd, 0xef, 0xb0, 0xd4, 0xf2, 0x51, 0xd7, 0xef, 0x86}}
	return a, nil
}

var _dataStatus_shortcutsSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\x41\x6f\xfb\x44\x10\xc5\xef\xfe\x14\x0f\xff\x2d\x35\x51\x12\x95\x72\xab\x82\x05\x12\x4d\x45\x2f\x14\x25\x85\x03\xb4\x4a\x37\xf6\xb8\x5e\xb1\xd9\xb5\x76\xc6\x71\x21\xe4\xbb\xa3\x5d\xc7\x49\x29\x54\xbd\x25\xda\xd9\x99\xf7\x7e\xf3\xd6\x5c\x6c\x9b\xb6\xaa\xd6\x2c\x4a\x5a\x1e\x8d\xb1\x4f\x00\xe3\x0a\x65\x30\x1c\x91\xdd\xad\x8b\x5a\xf9\x3c\xa5\x34\x49\x80\x2f\x58\x58\x6e\x3d\x81\xeb\xce\xf9\x92\x1b\xa3\x05\x9a\xe1\x2c\x2a\xe7\xf1\x17\xd7\x09\xa0\x2b\xfc\x8e\x99\x45\x9a\xfd\xb6\xfa\x71\xfd\xeb\x62\xb9\xba\xbb\xff\x29\xc5\xd3\x1c\x52\x93\x05\x93\xb8\x46\xde\x76\x98\xa3\xd2\xf3\xbe\xfd\xb2\xb5\xc3\x70\xf4\xba\xa6\x60\x71\x9e\xe0\x5a\x69\x5a\x89\x45\xa3\xe7\xa8\xf2\x19\x96\xa8\x64\x88\xc3\x86\x82\x04\x2d\x0c\xd7\x59\x18\x6d\x09\x4e\x6a\xf2\x9d\x66\x02\xbd\x6a\x41\xe1\x4a\x0a\x4a\xb9\x53\xc6\xb8\x8e\xca\xaf\xc6\x27\xb7\xc5\xb6\x5c\x9f\xda\x9f\xff\xe4\x69\x36\xba\x6c\xd9\x5f\x6e\xb4\xbd\x24\xbb\x7b\x27\x0c\xb3\x59\xa5\x0d\x19\xcd\x82\xec\xfb\xf1\x91\x8f\xae\x82\x4b\x4f\xe8\x14\x43\x59\x90\xf7\xce\x4f\x7b\x0d\x8d\xa7\xad\x92\xd6\x93\xf9\x73\x0a\x65\x4b\x34\x8a\x19\xca\x38\xfb\x12\x2e\x9d\x85\xf6\x2e\x57\x0f\x37\xf7\xbf\x3c\xc4\x46\x27\xd5\xd8\xb4\x02\xeb\x04\xab\x87\x9b\xc5\x72\x39\x05\x3b\xb4\x4c\x1e\x5c\xbb\xd6\x94\x60\xd1\xc6\x80\x89\xfa\xc1\xd8\xf2\xcb\xd9\x27\x71\x9e\x7d\x37\x2c\x28\x23\xc6\xcc\x12\xbe\x1e\x16\x93\x00\x80\x27\x69\xbd\x0d\x87\x09\x50\xe9\xde\xd3\x2d\x49\x51\x23\x1a\x75\x15\x82\x69\xc6\xa8\xf2\x6e\x8b\x4a\x7b\x96\x23\xef\x0a\x5c\x78\xdd\xc8\x71\x53\xe3\xd8\xc0\x10\x07\x8c\x54\xd4\x0e\x69\x76\x66\x9b\xe2\x6f\xd4\xa4\xca\x90\x93\xab\x81\xdd\xe2\xb5\x71\x5e\x60\xdb\xed\x86\x3c\x95\x08\xcc\x77\xca\x6b\xb5\x09\x13\x43\xc0\x48\x15\x75\xec\x9a\xe0\x94\xd1\xc2\x90\xf2\xeb\x9d\xf2\x41\xf2\xdd\xed\x2a\xcf\x2e\x1e\xe5\xe2\x6c\x3a\xbf\x0a\x4a\x9c\x8f\xf7\xa0\x2d\xb2\x28\x6b\x8e\xd2\x45\xcb\xd4\x4f\xcd\xde\x67\x3e\xa3\x3c\x8d\xa5\x69\x2c\x33\x24\xa0\xc9\x24\x01\x4a\x67\xe9\x34\x0a\x8f\xf2\x68\x2f\x7a\xfd\x3f\x7b\x6d\x65\x48\x47\x0f\x28\xa2\x91\xce\xc1\xd9\x4e\xf9\x32\x40\xf9\x5f\x16\xa2\xb4\x09\x2c\x26\xdf\x1c\xdf\x01\x31\x49\x78\x4f\x01\x82\xf6\xce\x6e\xc9\x4a\x48\x7a\x49\x95\x6a\x8d\x7c\xfe\xcc\x5a\xfb\xd1\x43\x3b\x24\x49\xf2\x05\x3f\x04\x6c\x1f\xb0\x4e\xfe\xcb\xf6\xf3\xcf\xc3\x70\x18\x43\x13\x78\x8f\x46\xd0\xf9\xd5\x1c\xfa\xdb\xfc\xfa\xfa\x7a\x0e\x3d\x99\x60\x3c\x3e\x71\x3f\xae\xc7\xee\x42\xff\xb5\xce\xb3\xfd\xfb\xb6\x87\x6c\xaf\x0f\xb1\x36\x78\x8d\x66\xb3\xfd\xe9\xc2\x01\x4f\x6f\x93\x8b\xde\xf1\xbf\x2a\xfa\xfd\x1a\xa6\x63\xc5\xc6\x93\xfa\x23\xfe\xae\xf4\xb0\xc8\x43\xf2\x4f\x00\x00\x00\xff\xff\xfa\x9a\x5d\x4e\x08\x05\x00\x00")

func dataStatus_shortcutsShBytes() ([]byte, error) {
	return bindataRead(
		_dataStatus_shortcutsSh,
		"data/status_shortcuts.sh",
	)
}

func dataStatus_shortcutsSh() (*asset, error) {
	bytes, err := dataStatus_shortcutsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/status_shortcuts.sh", size: 1288, mode: os.FileMode(0644), modTime: time.Unix(1558658378, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0xdf, 0xa8, 0x6c, 0x48, 0x6c, 0x1f, 0xe1, 0x63, 0xd1, 0x1e, 0x49, 0x78, 0x9c, 0x1b, 0x33, 0x83, 0xe4, 0xbb, 0x2, 0xe1, 0x18, 0x78, 0x8a, 0xc3, 0xae, 0x40, 0x78, 0xd8, 0xab, 0xe8, 0x1b}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"data/aliases.fish": dataAliasesFish,

	"data/aliases.sh": dataAliasesSh,

	"data/git_wrapper.fish": dataGit_wrapperFish,

	"data/git_wrapper.sh": dataGit_wrapperSh,

	"data/status_shortcuts.fish": dataStatus_shortcutsFish,

	"data/status_shortcuts.sh": dataStatus_shortcutsSh,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"data": &bintree{nil, map[string]*bintree{
		"aliases.fish":          &bintree{dataAliasesFish, map[string]*bintree{}},
		"aliases.sh":            &bintree{dataAliasesSh, map[string]*bintree{}},
		"git_wrapper.fish":      &bintree{dataGit_wrapperFish, map[string]*bintree{}},
		"git_wrapper.sh":        &bintree{dataGit_wrapperSh, map[string]*bintree{}},
		"status_shortcuts.fish": &bintree{dataStatus_shortcutsFish, map[string]*bintree{}},
		"status_shortcuts.sh":   &bintree{dataStatus_shortcutsSh, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
