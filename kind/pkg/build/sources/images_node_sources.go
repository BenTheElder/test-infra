// Code generated by go-bindata.
// sources:
// ../../../images/node/20-kubeadm.conf
// ../../../images/node/Dockerfile
// ../../../images/node/debug-node.service
// ../../../images/node/entrypoint/main.go
// ../../../images/node/journalctl-to-tty.service
// ../../../images/node/kubeadm.sh
// DO NOT EDIT!

package sources

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

var _imagesNode20KubeadmConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x48\xcb\xac\x50\xc8\x2e\x4d\x4a\xcd\x49\x2d\x51\x48\x4b\xcc\xcc\xc9\xcc\x4b\x57\xc8\xcf\x53\x28\x2e\x4f\x2c\x50\x48\xcd\x4b\x4c\xca\x49\x4d\xe1\x8a\x0e\x4e\x2d\x2a\xcb\x4c\x4e\x8d\xe5\x72\xcd\x2b\xcb\x2c\xca\xcf\xcb\x4d\xcd\x2b\xb1\x55\xf2\x0e\x75\x72\xf5\x71\x0d\x89\x77\x8d\x08\x09\x72\x8c\x77\x0c\x72\x0f\xb6\xd5\xd5\x05\x99\xa1\x0b\xd2\xad\x9b\x9f\x67\x9b\x96\x98\x53\x9c\xaa\xc4\x05\x08\x00\x00\xff\xff\xb9\x28\xc4\x46\x66\x00\x00\x00")

func imagesNode20KubeadmConfBytes() ([]byte, error) {
	return bindataRead(
		_imagesNode20KubeadmConf,
		"images/node/20-kubeadm.conf",
	)
}

func imagesNode20KubeadmConf() (*asset, error) {
	bytes, err := imagesNode20KubeadmConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/node/20-kubeadm.conf", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesNodeDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x7b\x6f\xdb\x38\xb6\xff\xdf\x9f\xe2\x5c\x39\x68\x93\x19\x4b\x4a\x9c\xb6\x93\xc9\x20\xb8\xd7\x4d\xdc\x8c\xd1\xd6\x2e\x6c\x77\x8a\x62\xee\x20\xa0\xa8\x63\x89\x6b\x8a\xd4\x92\x94\x5d\xb7\xdb\xf9\xec\x8b\x43\x49\x7e\x24\xe9\xee\xec\xce\x6e\x51\x20\x16\x1f\x87\xe7\xf9\x3b\x3f\xb2\x0b\xd7\xba\xdc\x18\x91\xe5\x0e\xfa\xa7\x67\x17\x30\xcf\x11\x5e\x57\x09\x1a\x85\x0e\x2d\x0c\x2a\x97\x6b\x63\xa3\x4e\xb7\xd3\x85\x37\x82\xa3\xb2\x98\x42\xa5\x52\x34\xe0\x72\x84\x41\xc9\x78\x8e\xed\x4c\x0f\x7e\x41\x63\x85\x56\xd0\x8f\x4e\xe1\x98\x16\x04\xcd\x54\x70\xf2\x53\xa7\x0b\x1b\x5d\x41\xc1\x36\xa0\xb4\x83\xca\x22\xb8\x5c\x58\x58\x08\x89\x80\x9f\x38\x96\x0e\x84\x02\xae\x8b\x52\x0a\xa6\x38\xc2\x5a\xb8\xdc\x1f\xd3\x08\x89\x3a\x5d\xf8\xd8\x88\xd0\x89\x63\x42\x01\x03\xae\xcb\x0d\xe8\xc5\xfe\x3a\x60\xce\x2b\x4c\xff\x72\xe7\xca\xcb\x38\x5e\xaf\xd7\x11\xf3\xca\x46\xda\x64\xb1\xac\x17\xda\xf8\xcd\xe8\x7a\x38\x9e\x0d\xc3\x7e\x74\xea\xb7\xbc\x57\x12\xad\x05\x83\x7f\xad\x84\xc1\x14\x92\x0d\xb0\xb2\x94\x82\xb3\x44\x22\x48\xb6\x06\x6d\x80\x65\x06\x31\x05\xa7\x49\xdf\xb5\x11\x4e\xa8\xac\x07\x56\x2f\xdc\x9a\x19\xec\x74\x21\x15\xd6\x19\x91\x54\xee\xc0\x59\xad\x76\xc2\x1e\x2c\xd0\x0a\x98\x82\x60\x30\x83\xd1\x2c\x80\x97\x83\xd9\x68\xd6\xeb\x74\xe1\xc3\x68\xfe\xf3\xe4\xfd\x1c\x3e\x0c\xa6\xd3\xc1\x78\x3e\x1a\xce\x60\x32\x85\xeb\xc9\xf8\x66\x34\x1f\x4d\xc6\x33\x98\xbc\x82\xc1\xf8\x23\xbc\x1e\x8d\x6f\x7a\x80\xc2\xe5\x68\x00\x3f\x95\x86\xf4\xd7\x06\x04\xb9\x11\x53\xf2\xd9\x0c\xf1\x40\x81\x85\xae\x15\xb2\x25\x72\xb1\x10\x1c\x24\x53\x59\xc5\x32\x84\x4c\xaf\xd0\x28\xa1\x32\x28\xd1\x14\xc2\x52\x30\x2d\x30\x95\x76\xba\x20\x45\x21\x1c\x73\x7e\xe4\x81\x51\x51\xa7\xd3\x85\xa5\x50\x29\x70\x59\x59\x87\x06\x12\x46\xa6\x16\x2c\xc3\x1e\x24\x95\x90\x8e\x0c\xdd\xa5\xd6\x53\x0b\x29\x26\x82\xa9\x90\x16\x76\xba\xb0\xce\x05\xcf\xc9\x37\x09\xb3\x82\x33\x29\x37\xc0\x60\x85\x66\x03\xe4\xaa\xb2\xc4\x14\x52\xbd\x56\xcd\xae\x5a\x32\x1c\xff\xfe\xec\xfc\xed\xcb\x93\x4e\x17\x5c\xce\x1c\xd8\x5c\x94\xb6\x4e\x1b\x06\xb6\x60\x52\x42\xe5\x84\x14\x6e\x03\x4f\xb9\x44\xa6\x42\xa1\xac\x63\x52\x3e\xdd\x1d\x47\xb2\x52\x64\x12\x52\xa3\x4b\x0a\xe8\x42\x9b\x4e\x17\x58\xe9\xc2\x0c\x1d\x54\x65\xca\x1c\xc2\x93\x27\xdb\x91\x46\x04\x1c\xfd\x1f\x74\xba\x60\x11\x7d\x8e\xd9\xcb\x38\xce\x84\xcb\xab\x24\xe2\xba\x88\x97\x5b\x43\xf7\x7f\x3a\x83\x18\x17\x8c\xfc\x13\x93\x4f\xd2\xf8\xc0\x07\x9d\x2e\xcc\x75\x5d\x14\x6b\x04\x96\xa6\x60\x37\xd6\x61\x91\xf6\xe0\x7a\x3c\xea\x51\x1c\x40\xfb\x38\x3b\xad\xa5\x05\x85\x98\xd6\x79\x68\xaa\xda\xb7\x2c\x2d\xbc\x98\x57\xda\xb4\x7b\xe1\x7b\x48\x35\x5f\xa2\x01\xae\xd5\x42\x64\x95\xf1\x31\xa4\xfa\x4b\x21\x41\xa9\xd7\x3d\x6f\x04\x05\x73\xa1\xa5\xd4\x6b\x0a\xbf\xc1\x05\x1a\x54\x1c\xed\x65\xa7\xbb\x35\x90\xaa\x68\x41\xb9\x8f\x76\xe9\x74\xe9\x4b\x69\x2d\x96\x22\x9e\x35\xb9\x1f\x37\x87\xc6\xd7\x5a\x51\x7d\xa2\x19\x29\x87\x66\xc1\x38\xc6\x7b\x72\x52\x5c\xa1\xd4\x25\x1a\x1b\x19\x4c\x73\xe6\xbc\xcf\x12\xa9\xb3\xb8\x7f\x7a\xf6\x2c\x3e\x7d\x4e\xff\x4d\xa5\x28\x15\xc3\x46\x66\x48\x71\x15\x2a\xac\xad\x09\x79\x7b\xc2\x1f\x17\xfc\x22\x3e\xfd\x31\x3e\x3b\x7f\x20\x58\xa8\x90\x85\x4a\xab\xb0\x34\x62\x25\x24\x66\x98\xee\x8b\xa7\xb0\x4c\x6e\x26\xc7\x09\x2a\x97\x23\xca\x14\xcd\xc9\x25\x95\x9a\xd4\x06\xc1\x56\x65\xa9\x0d\x41\x40\x13\x1b\x66\x78\x8e\xb6\x07\xa3\x61\x3b\x57\x43\xe5\xf4\xfa\x67\x60\x26\x8b\x9a\xf0\x28\xbd\xae\x43\x2d\x2c\x58\x27\xa4\xa4\x4c\x1c\x4c\x6f\xc1\x6a\x8a\x3e\x67\x0a\x0c\x12\x48\x0a\xda\x6f\x74\x95\xe5\xba\xaa\x45\xf9\xd4\xa1\x92\x11\xce\x8f\x91\xbd\xc2\x7a\x15\x68\xb1\xf6\x8b\xb6\x06\x44\x1d\x92\x4a\xc7\x5f\x05\xac\x48\x5f\x3c\x0b\xfc\xc0\xcb\xc1\x6c\x78\x37\x7a\x3b\xb8\x1d\xde\xfd\x32\x9c\xce\x46\x93\xf1\x55\x70\x1a\x9d\xdf\x9f\xbc\x0a\x96\x17\x36\xca\xb8\x89\x84\xde\xcf\xd6\xf0\xe8\x0b\x89\xfc\x7a\x79\xf4\xe5\xa1\xa4\xaf\x41\xe7\xd5\x74\xf2\x16\xf6\xe7\xbe\x92\x1f\xc7\x93\xf9\xf0\xb2\x76\x45\x51\x59\x07\x09\x42\x8a\x0b\xa1\x30\x05\x96\x79\x3c\x5f\x10\x76\xd0\xe6\xfd\xa8\x6a\x6e\xa3\x3a\xec\x3e\x9e\xa8\x32\xa1\x30\xde\xe6\x68\x5d\x4b\x68\xe2\xae\xe5\xba\xc4\x87\xf6\xfa\x3a\x75\xde\x41\x37\xc3\x97\xa3\xc1\xf8\xee\xd5\x74\x32\x9e\x0f\xc7\x37\x57\x4a\x2b\x41\x29\xca\xb8\x13\x2b\x04\xeb\x74\x69\xc1\xea\x02\xa9\xde\x61\xcd\x3c\x1a\xda\xde\x36\x54\xd4\xba\x08\x1e\xc0\x10\x66\x30\x93\x55\x05\x2a\xd7\x83\x35\x3e\x35\x08\xc7\x2c\x39\xa9\x2c\x9d\x43\x3a\x38\x0d\x04\x18\x0c\x1c\x16\xa5\x36\xcc\x6c\x60\x38\xfe\xa5\xb6\xb4\x8e\xca\x3f\xd6\xa6\xd3\x85\x5b\x74\x90\x62\x89\x2a\x45\xc5\x05\x5a\xca\x45\xca\x80\x2d\xb8\x02\x93\x06\x59\xba\x81\x9c\xd9\x4b\xb0\x36\xef\x91\xe2\x3d\xb0\x8a\x95\xa9\x5f\x5d\xeb\x9d\x18\xbd\x44\x55\x43\xa8\x50\x4e\x53\x59\x87\x50\x32\xbe\x64\x19\x12\x94\x70\xb4\x96\x34\xa4\xe6\xd0\x40\x1c\xd9\x51\xbb\xfd\xfe\xe2\x7d\xdc\xb1\x68\x56\x82\xa3\x85\xe3\xa6\x9e\x4e\x1e\x5d\x4d\x72\x1b\x20\x8a\x21\xdf\x94\x68\x08\x18\x21\x86\x1d\x3e\xfa\xde\xaf\x15\x2a\x67\xbd\x88\x42\x58\xbe\x93\x73\x5c\x43\xb9\x40\x8f\x8e\xbe\x32\x14\xe8\xca\x00\x99\x44\x80\x28\x54\x76\x52\xfb\x47\xf9\x0a\x22\xc0\xaf\x4a\x38\x36\x58\xe8\x15\xd9\x52\xa9\x35\x53\xd4\x72\x5b\x68\x6c\x55\x3f\xf9\xef\x61\xd3\xf4\xfd\x18\x0e\x5a\x0f\xfc\x7f\xc7\x13\x13\xdf\x4f\x9c\x61\xca\x12\x44\x84\xfe\x78\xe0\x2c\xe4\x68\x1c\x35\x65\xe6\x1d\x52\x19\xb9\x25\x16\x61\x69\x48\x2f\xf2\x40\xc8\x75\x51\x68\x05\x99\xaa\xca\xac\x0f\xd2\x26\xa1\x41\x89\x94\x15\xad\xf8\xad\x8d\x8d\x86\x76\x63\x57\xdb\x49\x51\x3a\x22\x34\x16\x44\x69\x74\xe5\xb0\x0f\xe8\x72\xf2\x21\x58\xcd\x99\xf3\x5d\x33\x94\x42\x55\x9f\xa0\xd0\x95\x72\x80\x49\xb3\xa1\x4a\x71\x05\xcb\x42\xa7\xc0\xaa\x85\x0d\xeb\x46\xd4\x8a\x4d\x98\xcd\x9b\x8f\x27\x4f\x60\x41\x74\x20\x96\x22\xd9\x76\x85\xfa\x2f\xfd\x11\x4a\xb8\xc8\x31\x93\xa1\x8b\x28\x28\x36\x86\x50\xb1\x02\x21\x68\xf5\x75\x45\x49\xcc\xd0\x86\x16\x5d\x55\x46\x4d\xa8\x02\x08\x53\x94\xe8\x70\x77\x8c\x29\x20\x5c\x3c\x7a\x4e\x51\x49\x27\xc2\xca\xa2\x39\x3c\xea\xbb\x07\x9b\xd1\xf1\xfb\x9b\xbf\xfb\xe6\xe2\x47\x4e\x92\x9a\x33\x19\x2e\xec\x3f\x3b\xe7\x31\x67\x50\xca\xb8\xfb\x3b\xc9\xcd\x7f\x62\x3b\xb9\x97\x3b\xf9\x87\x24\x78\xbe\x75\x6f\x3f\x21\xe6\xa8\xc9\xd7\x3a\xa7\x7b\x0d\x63\xa2\x6a\xb6\x54\xfa\x39\x2b\x4b\x6c\x01\x7b\x8d\x5b\x5a\xe4\xb1\xb3\xe1\xe3\xdb\xda\x65\x89\xf6\x70\x46\xb8\xe5\x19\xef\xb7\xa0\xbd\x91\x12\xfb\xdc\x8b\xdb\x72\xc2\xb8\x4a\x2a\xe5\xaa\xb8\x6b\xd1\x85\x55\x19\xba\x1c\x43\x83\xa5\xb6\xc2\x69\xb3\x21\x3c\x56\xe9\x56\xe6\x0e\x50\x7c\xcf\xd2\xdc\xc6\x3e\x87\x62\xa1\x5a\x28\x75\xed\x41\xe1\xb2\x66\x4d\x71\x77\x87\x79\xe1\x1e\xe6\x11\x7e\xd7\x9f\x4f\x2d\xdc\xbe\xbb\x85\x25\x6e\xfc\x04\xb1\x33\x4f\x99\x84\xca\xd0\x94\x46\x28\x77\x30\x7e\xa0\x5d\xd8\xd2\x48\x0f\xcd\xe2\x9e\x63\x1b\xa0\xda\xb6\xc7\xb3\x1f\xa2\xd3\x73\xc2\x6c\xbd\x58\x08\x2e\x3c\x15\x6e\x98\x44\x7d\x23\xd9\xbb\xa0\xf1\xca\x18\x54\x4e\x6e\x7a\x0d\x6f\x28\x85\xaa\x19\x00\x73\xd1\x1e\xaa\x3d\xe2\x13\xc7\xec\xd2\xc6\xbe\x80\x1f\x78\xa3\x6e\x4e\x93\xeb\xd7\xc3\xe9\x8e\x1d\x78\xbd\xa2\xfe\xef\x1c\xc3\xd3\xdf\x1b\x2a\x60\x9d\x41\xc7\xf3\xc0\x87\xa0\xe1\xa6\xdb\x66\x87\x6a\xd5\xf3\x0d\xf3\x5e\xb7\x8c\xb6\xdd\xd8\x77\x55\xa7\x81\x01\xd1\xaf\xcf\x68\x34\xac\x98\xac\x28\x55\xac\x90\x9e\x7f\x36\x7c\xb8\x6e\xc4\xb0\x30\xba\xf0\xc8\xb9\xc4\x0d\xb5\xb0\xe0\x43\x3d\x71\xd9\x0e\x82\xae\x5c\x59\xd1\x55\x40\x57\x32\xf5\xa7\x27\x94\x88\x86\x12\xef\xd8\xba\x94\x08\x53\xd3\xc7\xa9\x31\x9b\x42\x28\x26\x4f\x6a\xda\x33\x78\x37\xbf\x7b\x3d\xfc\x78\x77\x33\x19\xcf\xef\x3e\x0c\xa6\xe3\xbb\xc9\xf8\xee\x66\x30\xbe\x1d\x4e\x27\xef\x67\x77\xef\x67\x9e\x10\x2d\x98\xb4\x18\xd4\xc8\x4e\xf8\x1c\x2e\xec\xec\x0d\x04\xbb\x94\x5e\x2b\xa9\x59\xba\x9f\xd6\x75\x3a\x1f\x1d\x47\x35\xd4\x68\xdb\xc2\xf5\x4f\x80\x3c\xd7\x10\x1c\x8d\x6e\x82\x93\x38\x2b\xb3\x00\xfe\xb6\x35\x85\x72\x29\xdc\x95\x70\x3b\xbc\x97\x72\x70\x3a\x7c\xf9\xea\xfa\xe6\xe2\x62\xb7\xaa\x26\x3e\x0d\x3b\x0b\x48\x46\x48\x1b\x77\xf9\xb8\x45\x6b\x80\x20\xc5\x04\x7e\x25\xc2\x7a\xd5\x6c\xf8\x0d\xfe\xbc\x19\x70\x74\x2c\x6d\x72\xd7\x36\xa4\x90\xdb\x13\xb0\xbe\x7d\x04\x3b\x35\x0f\x7b\x62\xb0\xad\xf4\xab\xa3\x2f\x87\x79\xf7\x35\xd8\x07\xa3\xeb\xf1\x08\x12\xa1\x98\x21\x1a\xe0\x34\xc4\xba\x74\x31\x57\x22\x4e\x84\x7a\x9c\xa7\xa7\x9a\xc3\x3a\xdf\x40\x0c\x6b\xba\x22\xe6\x68\x6a\x82\x78\x3d\x1e\xed\x13\xdf\x17\xd1\x69\xb0\x1d\x9f\x0f\xa6\x2f\x07\x6f\xde\x5c\x05\x5c\x89\xb0\x94\x55\x26\x94\x6d\x09\x6f\xb8\x3a\xfa\xb2\xb7\xf7\x6b\xe4\xb2\xcf\xbb\x9d\x9e\xef\xbe\x9f\xbe\xb9\xda\xa6\x83\x75\xda\xb0\x0c\xa3\x4c\xeb\x4c\x22\x2b\x85\xbd\x77\x41\x6c\x7d\x18\x2b\x74\x6b\x6d\x96\xed\x81\xf1\x4e\xaa\x17\x58\x1f\xdb\x1e\xf0\xb5\xfe\x6c\x34\xfd\xba\x9f\x8e\x94\x8d\x61\x68\xd0\x99\x0d\x3c\x87\x30\x6c\x2a\x22\x76\x45\x49\xae\x22\x85\xa1\x91\x46\x82\xf6\x82\x62\x73\xd6\x7f\xfe\xc2\x56\xc5\xe1\xe2\xed\x7c\xb1\x4c\x85\x81\xb0\x3c\x70\xfb\x6e\xda\x31\x03\xe1\xf5\xe1\x64\xf8\xe9\xf3\xe2\x40\x1a\x45\x73\x90\xa6\x74\x41\x6f\x18\x4a\xa5\xea\x5b\x8b\x4f\xa1\xbf\xe8\xca\x28\x26\xb9\x93\x0f\x2e\x32\xe0\xdc\xa6\x41\x39\x82\x65\xa9\x33\xa2\x89\x1e\x11\x5a\x51\x71\xfd\xae\xd1\xd0\xba\xa8\x73\x3d\x79\xf7\x71\x4f\x64\xe8\x74\xe8\xdc\xa6\x25\x13\x8f\xf6\x7d\xef\xc8\xfa\x37\x29\x81\xca\xbf\xfc\x3c\x90\xf1\x8d\x5b\x21\x7b\xd4\x32\xa2\xc9\xc4\x54\x95\x4e\x11\x12\xad\x5d\x55\x92\xfa\x82\xef\x21\x74\x52\x65\x9f\x85\x94\x6c\x9f\x75\xda\x5c\xaf\xef\x92\x2a\x8b\x78\x26\xfe\x57\xa4\x57\x17\xa7\x3f\x9e\x9d\xf7\xbf\x91\xe7\x84\xb0\xfe\xe9\xa3\x3d\xfd\xed\xe0\x9a\xf4\xf1\x0f\x42\xcc\x5a\x91\x29\x02\xdf\xbd\x23\xf7\xde\x2b\x0e\x9d\x90\xc6\xc2\xda\x0a\x6d\x7c\x7e\xfe\xc3\xb3\xae\xff\x4d\x84\x13\x95\x0b\xfb\x17\x17\x17\x17\xfd\xf3\xe7\xcf\xff\xa4\x9c\xf3\xf3\x1f\xfb\xcf\x2f\x9e\x5d\x9c\x93\x23\x6b\xaf\x79\x6f\xf9\x6b\x02\x26\x55\x96\x09\x95\x3d\x6e\xa9\xe7\xf3\xcd\xe3\x21\x2a\x27\x0c\x42\x22\x35\x5f\x7a\xd6\xb1\x6d\x3e\x5e\x08\x24\x12\xb3\xfc\x7f\xea\x3c\xf0\x23\x21\xc5\xe0\xdf\x48\x80\x87\x9b\x6b\xa1\x4d\xc7\x8c\x6c\x0e\x71\x65\x4d\xcd\x04\x29\xf3\xe3\x7a\xbe\x7f\xda\x36\xd5\x88\x6b\xf5\x38\xd5\xa4\x05\x12\x5d\x2b\x38\x4a\xfd\x9b\x83\x43\x62\x53\x4d\x2c\xfd\xf3\x96\xf0\x9d\x4b\xa8\xf6\x1e\x75\x2c\x1c\xac\x85\x94\xc0\x73\xe4\xcb\xed\xb3\xde\xae\x60\x50\xad\x4e\xfe\x63\xcf\x37\x74\x75\xdd\x89\xde\x72\xa4\x56\x43\xfc\x24\x9c\x25\x6a\x37\x1b\xdd\x4e\xe7\x6f\x47\xe3\xef\xcf\xeb\xe6\x3f\x1b\xdd\xce\x87\xd3\xb7\x70\x5c\x13\x48\x83\x21\x7e\x42\x5e\x11\x7b\x11\xee\xe4\x5f\xac\x80\xb3\xfe\xe9\xd9\x8b\xe7\x3f\x74\x66\xf3\xc9\xbb\xd9\xe8\x76\x3c\x78\xb3\x77\xde\x37\x6a\x32\xd5\xdc\xd3\x0e\x9f\x31\x4d\x2e\xfc\x0a\x01\x2a\x67\x36\xa5\x16\xca\xc5\xbb\x9f\x41\x0f\x82\x7b\x71\x0c\xe0\xb7\x4e\x17\x3e\xa0\xa7\xbe\xbb\x80\x68\xe2\x15\xef\x46\x37\x67\x6d\x81\x93\xeb\x57\xcc\x08\x5d\xd9\xbd\x3b\x71\x4b\xf2\x9a\x18\x63\x0f\xd0\xf1\x88\xcc\x1e\xa8\x0d\xdd\x19\x33\x40\x69\x91\x40\x4d\x35\x8f\xca\xa9\xee\xf9\xf7\x3f\x69\x6b\xc1\x8c\x68\x9e\x75\x7b\x30\xd7\x8a\xef\x0c\xc7\xf3\xe9\xc7\x77\x93\xd1\x78\x4e\x16\xdd\x53\xfc\x9e\x55\x96\xc6\xe8\x62\x40\x5f\x61\x28\x75\x16\xd6\xa4\xff\x8a\x6b\x65\x35\x75\xe7\xdf\x3a\x7f\x0f\x00\x00\xff\xff\x34\xe6\x65\xfd\x11\x18\x00\x00")

func imagesNodeDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_imagesNodeDockerfile,
		"images/node/Dockerfile",
	)
}

func imagesNodeDockerfile() (*asset, error) {
	bytes, err := imagesNodeDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/node/Dockerfile", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesNodeDebugNodeService = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4d\x8f\xdb\x36\x10\x86\xef\xfc\x15\x6f\x2d\x20\x4e\x80\xb5\x95\xec\xa9\x48\xe1\x83\xb3\x76\x51\x23\x81\x0d\x58\xde\x2e\x82\xc5\xa2\xa0\xc8\xb1\xc4\x86\x22\x95\xe1\xd0\x1f\xff\xbe\x90\xbd\x41\xbb\x48\x0a\x44\xc7\xe1\xf3\x0e\x1f\xcd\xb0\xc0\x5d\xec\xcf\xec\x9a\x56\x70\xfb\xf6\xdd\xaf\xd8\xb5\x84\x8f\xb9\x26\x0e\x24\x94\x30\xcf\xd2\x46\x4e\x53\x55\xa8\x02\x9f\x9c\xa1\x90\xc8\x22\x07\x4b\x0c\x69\x09\xf3\x5e\x9b\x96\xbe\x9d\xdc\xe0\x4f\xe2\xe4\x62\xc0\xed\xf4\x2d\x5e\x0f\xc0\xe8\xf9\x68\xf4\xe6\x37\x55\xe0\x1c\x33\x3a\x7d\x46\x88\x82\x9c\x08\xd2\xba\x84\xbd\xf3\x04\x3a\x19\xea\x05\x2e\xc0\xc4\xae\xf7\x4e\x07\x43\x38\x3a\x69\x2f\xd7\x3c\x37\x99\xaa\x02\x9f\x9f\x5b\xc4\x5a\xb4\x0b\xd0\x30\xb1\x3f\x23\xee\xff\xcb\x41\xcb\x45\x78\xf8\x5a\x91\xfe\x7d\x59\x1e\x8f\xc7\xa9\xbe\xc8\x4e\x23\x37\xa5\xbf\x82\xa9\xfc\xb4\xba\x5b\xae\xab\xe5\xe4\x76\xfa\xf6\x12\xb9\x0f\x9e\x52\x02\xd3\xd7\xec\x98\x2c\xea\x33\x74\xdf\x7b\x67\x74\xed\x09\x5e\x1f\x11\x19\xba\x61\x22\x0b\x89\x83\xef\x91\x9d\xb8\xd0\xdc\x20\xc5\xbd\x1c\x35\x93\x2a\x60\x5d\x12\x76\x75\x96\x17\xc3\xfa\x66\xe7\xd2\x0b\x20\x06\xe8\x80\xd1\xbc\xc2\xaa\x1a\xe1\xc3\xbc\x5a\x55\x37\xaa\xc0\xc3\x6a\xf7\xc7\xe6\x7e\x87\x87\xf9\x76\x3b\x5f\xef\x56\xcb\x0a\x9b\x2d\xee\x36\xeb\xc5\x6a\xb7\xda\xac\x2b\x6c\x7e\xc7\x7c\xfd\x19\x1f\x57\xeb\xc5\x0d\xc8\x49\x4b\x0c\x3a\xf5\x3c\xf8\x47\x86\x1b\xc6\x48\x76\x98\x59\x45\xf4\x42\x60\x1f\xaf\x42\xa9\x27\xe3\xf6\xce\xc0\xeb\xd0\x64\xdd\x10\x9a\x78\x20\x0e\x2e\x34\xe8\x89\x3b\x97\x86\x65\x26\xe8\x60\x55\x01\xef\x3a\x27\x5a\x2e\x95\xef\x7e\x6a\xaa\x54\x71\x5d\x67\x22\x3e\x38\x43\xe0\x1c\x12\x52\xec\x08\x96\xea\xdc\x20\x49\xde\xef\x55\x81\xdd\x66\xb1\x79\x5d\x53\x90\x96\xc8\x5b\xe2\x37\xef\xc1\xd4\xc5\xc3\xf5\x35\xfc\xa2\x1e\xef\x83\x93\x27\xb5\xa0\x64\xd8\xf5\xc3\x75\xb3\xc5\xa5\x41\x88\x96\xd4\xf6\xba\x98\x34\xeb\xb2\x17\x37\xc9\x89\x78\x2a\x9a\x1b\x12\x35\xdf\x0b\xf1\x0f\xea\xea\xb1\xba\x2a\x3d\xa9\xdd\xb9\xa7\x59\x0c\x94\xda\x28\x6a\x4b\x9d\x76\xe1\x92\x5a\x9e\x9c\xcc\x84\x33\xa9\x62\x79\x22\x53\x89\x66\x99\x95\xb5\x0b\x65\xad\x53\x8b\x89\xc1\x38\x5b\x3a\x68\xdb\xc1\x85\x7d\xc4\xe4\x2b\xb4\xf7\x98\x68\x94\x96\x0e\x65\xb2\xfa\xdd\xf8\xe7\xa3\x3d\xc7\x9e\x58\xce\x98\x84\x7f\xf3\x98\x4c\xe8\xd4\x47\x96\xb1\xfa\xbf\x3e\x36\x9a\x2f\xc4\xc3\x60\xd1\x92\xf7\x71\x72\x8c\xec\x2d\x5e\xbd\x02\x99\x36\x62\xd4\x73\xb4\xd9\x08\x72\x76\x76\x34\x94\x8d\x16\x94\xe9\x9c\x4a\xe3\x75\x4a\xa5\xed\x5c\xe9\x6c\xf9\x8c\xfd\x35\x60\xdf\x87\x83\xee\xe8\x27\xc2\x03\x36\x56\xc5\x00\x96\x39\x71\xe9\xa3\xd1\xfe\xa2\xfb\x25\xd7\xa4\x6d\x37\x4d\xed\x58\x55\xa2\x83\xd5\x6c\x37\x59\xfa\x2c\xb3\xbf\x63\xe6\xa0\xbd\x52\x8f\xab\x90\x44\x7b\xff\xa4\x1e\x74\x10\xb2\x1f\xce\x3f\xd8\xdb\x3f\x01\x00\x00\xff\xff\x02\xbc\xc5\x40\x9c\x04\x00\x00")

func imagesNodeDebugNodeServiceBytes() ([]byte, error) {
	return bindataRead(
		_imagesNodeDebugNodeService,
		"images/node/debug-node.service",
	)
}

func imagesNodeDebugNodeService() (*asset, error) {
	bytes, err := imagesNodeDebugNodeServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/node/debug-node.service", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesNodeEntrypointMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x55\xd1\x72\xdb\x38\x12\x7c\x16\xbe\x62\xc2\xab\xd4\x49\x89\x4c\x5a\x72\xec\x38\xba\xe4\x41\x67\x3b\x89\x2a\x3e\x39\x65\x29\xe7\x4b\xe5\x5c\x09\x08\x0e\x49\x94\x41\x80\x0b\x80\xa2\xb9\x5b\xf9\xf7\xad\x01\x69\xaf\xbc\xbb\x0f\x96\x05\x12\xe8\xe9\xe9\x69\xb4\x92\x17\xec\xcc\xd4\x9d\x95\x45\xe9\x61\x7e\x38\x3b\x85\x6d\x89\xf0\xa9\x49\xd1\x6a\xf4\xe8\x60\xd9\xf8\xd2\x58\x17\x33\x76\x29\x05\x6a\x87\x19\x34\x3a\x43\x0b\xbe\x44\x58\xd6\x5c\x94\x08\xc3\x9b\x29\xfc\x17\xad\x93\x46\xc3\x3c\x3e\x84\x31\x6d\x88\x86\x57\xd1\xe4\x5f\xac\x33\x0d\x54\xbc\x03\x6d\x3c\x34\x0e\xc1\x97\xd2\x41\x2e\x15\x02\xde\x0b\xac\x3d\x48\x0d\xc2\x54\xb5\x92\x5c\x0b\x84\x56\xfa\x32\x14\x19\x20\x62\xf6\x75\x00\x30\xa9\xe7\x52\x03\x07\x61\xea\x0e\x4c\xbe\xbf\x0b\xb8\x67\x0c\x00\xa0\xf4\xbe\x5e\x24\x49\xdb\xb6\x31\x0f\x2c\x63\x63\x8b\x44\xf5\xbb\x5c\x72\xb9\x3a\xbb\x58\x6f\x2e\x0e\xe6\xf1\x21\x63\x5f\xb4\x42\xe7\xc0\xe2\x2f\x8d\xb4\x98\x41\xda\x01\xaf\x6b\x25\x05\x4f\x15\x82\xe2\x2d\x18\x0b\xbc\xb0\x88\x19\x78\x43\x3c\x5b\x2b\xbd\xd4\xc5\x14\x9c\xc9\x7d\xcb\x2d\xb2\x4c\x3a\x6f\x65\xda\xf8\x27\x02\x3d\xb0\x92\x0e\xf6\x37\x18\x0d\x5c\x43\xb4\xdc\xc0\x6a\x13\xc1\xbf\x97\x9b\xd5\x66\xca\x6e\x56\xdb\x8f\x57\x5f\xb6\x70\xb3\xbc\xbe\x5e\xae\xb7\xab\x8b\x0d\x5c\x5d\xc3\xd9\xd5\xfa\x7c\xb5\x5d\x5d\xad\x37\x70\xf5\x1e\x96\xeb\xaf\xf0\x69\xb5\x3e\x9f\x02\x4a\x5f\xa2\x05\xbc\xaf\x2d\x71\x37\x16\x24\x49\x87\x59\xcc\x36\x88\x4f\x8a\xe7\xa6\x27\xe3\x6a\x14\x32\x97\x02\x14\xd7\x45\xc3\x0b\x84\xc2\xec\xd0\x6a\xa9\x0b\xa8\xd1\x56\xd2\xd1\xf0\x1c\x70\x9d\x31\x25\x2b\xe9\xb9\x0f\xeb\xbf\xb4\x13\xb3\x17\x09\x63\xc9\x0b\xb6\xa5\x11\x52\x5d\xac\x50\x7b\x07\xa6\x21\x1a\x84\x8c\xda\xdb\xae\x36\x52\xfb\x18\x56\x7e\xc1\x0e\xa0\xe5\xd2\xbb\xc0\x65\xb3\xfa\xf0\x65\x73\x3d\x63\x07\x84\xa9\x01\xef\x51\x38\x18\x73\x5b\xec\xbe\xcd\x6e\xa7\xd0\x7f\x59\xdc\x4e\x01\xf5\xee\xdb\xe2\x76\xc2\xfa\x3a\x5c\x29\xd3\x3a\x68\x1c\xcd\xa0\x46\x9b\x1b\x5b\x81\x09\x2a\x70\xd1\x33\x95\x3a\xd0\x8c\xb4\xc9\x30\x02\x61\x34\x19\x05\x2d\xec\x24\x87\x1f\x99\x11\x77\x41\x31\x14\x3f\xd8\xf7\x14\x73\x63\xf1\x3b\xb4\x08\x25\xdf\x21\x41\x34\x5c\xa9\x0e\xa2\xd4\x18\x8f\x59\x14\x90\xa4\x96\x9e\xf4\x00\xdc\xa1\xed\x7c\x49\x52\xa1\x22\x9b\x29\xa3\x8b\xde\xa3\xd2\xc7\x8c\xdd\x20\x08\xae\xfb\x86\x1c\xea\xec\xa1\x49\xe2\x1a\x8c\x5e\x5b\x23\x68\x50\xb4\xb6\xb2\x28\xd0\x82\xf3\xdc\x92\x8b\x7a\xce\x3d\x81\x88\xfd\xa1\x1c\xb4\x84\xd6\x22\x70\x8b\x90\x19\x8d\x50\x5b\x62\x5d\xd1\x19\xae\x3b\xc2\xdc\x49\x9a\x19\x3d\x30\xfb\xbd\xc7\x8c\xad\xaf\xb6\x17\x8b\xbe\xf6\xfe\x90\x30\x03\xee\x80\x83\x93\xba\x50\xe4\x80\x59\xfc\xbf\x70\x0b\xa7\xd0\xb8\x1e\x47\x75\xbd\x5d\x7c\xa6\x64\x1a\xf7\xe2\x57\xfc\x0e\x1d\x48\x0f\xc8\x9d\x24\x37\x18\x48\x1b\xa9\x32\xa8\x8d\xf5\x3c\x55\xdd\x34\xc8\xd4\x8f\x89\x38\x6b\xba\x2c\x34\xed\xb6\xe4\x9e\x0e\x66\x06\x5d\x6f\x9c\x9a\x8b\x3b\xf2\x48\xc5\xa5\x66\x4c\x56\x04\x01\x63\x36\x8a\x8c\x8b\xc2\x67\xe2\x64\xa1\x49\x8b\x51\xe4\x3a\x27\xb8\x52\x11\x63\xa3\x48\x99\x22\x62\x13\xc6\x92\x04\x3a\x74\x7d\x6b\xae\x34\x8d\xca\x20\xed\x1d\x2f\xa0\xe2\xc2\x9a\x29\xa4\x8d\x27\x41\x94\xd4\xcd\x3d\xb9\x62\x98\x7d\x67\x9a\x7f\x5a\x6a\x3a\xa8\x6e\xa0\x40\xdf\xc3\x70\xdd\x95\xa6\x25\xe4\x21\x37\x2a\xae\x5f\x0f\x89\xa1\x9b\x7b\x5a\x1e\xd4\xbc\x40\x17\x5e\x0c\xfc\xe2\xd7\x71\xe9\x2b\xf5\x70\xca\x2d\x92\xa4\x90\xbe\x6c\xd2\x58\x98\x2a\xa9\x4c\xda\xf5\x1f\xa9\x32\x69\x72\x7c\x32\xcf\xf2\x53\x31\xcf\x4e\xf2\x57\xa7\x27\x87\x33\x71\x9a\xcd\xb2\xfc\xf5\xfc\xf8\xe4\xe8\xf4\xcd\xf1\xc9\x9b\x6c\x7e\x2c\x0e\xd3\x7c\x96\xd4\x77\xc5\x80\x3f\xfc\xfb\x1e\x28\xc4\x85\xf9\xc7\xe5\xec\x90\x09\xa3\x9d\x07\x27\x0b\xeb\x2b\xa9\xe1\x1d\x1c\xbd\x62\x2c\x6f\xb4\x08\x72\x8e\x27\xf0\x1b\x1b\x25\x09\x39\x65\x87\xda\xc3\xaf\xa6\x4a\x25\x3e\xb8\x0f\x1d\xcd\x9d\xc2\x95\xf2\x55\x29\xd2\xed\xf3\xea\x7c\x16\x06\xc5\xa1\x2d\xa5\xc2\x70\xfc\xa1\x9f\xbe\x74\x26\x31\xd6\xe8\xa9\xf7\x64\x9e\xd0\x45\xae\x65\xc6\x46\x83\x0a\xab\x42\x1b\x8b\xe3\x61\x52\xf1\x66\xf5\xe1\xec\xe3\xe5\xf9\x84\x05\xa0\xc2\xf2\xb4\xf7\xa5\x45\xae\xa2\xbd\x5c\xa0\xac\xaf\xc8\x33\xe1\xcf\x16\x0e\x72\x6b\xaa\x10\x21\xb4\x62\x23\x99\x83\x42\x3d\x36\x2e\x5e\xda\xc2\x4d\xe0\x2d\xcc\xa9\xb9\x91\x32\x45\xfc\x9e\x7b\xae\xc6\xd1\xda\x78\x40\x6d\x9a\xa2\xa4\x33\x4d\x1f\x42\xde\xec\x55\x79\x16\x4d\xd8\xe8\x27\x1b\x89\x2a\xeb\x93\x05\x16\xef\x60\x80\x0c\x69\xf3\xf8\x7d\x71\xdb\x33\xa6\xf6\xf6\x63\x0a\xc6\x86\x42\x43\x06\x43\x6d\x56\x1f\xae\xb7\xff\x59\xad\x5f\x1e\x51\x99\x8a\x7b\x51\x82\xeb\x9c\xc7\x2a\x9b\xb0\x91\x20\x70\xba\x29\x63\x51\x72\x4d\xd0\x9b\x20\xd1\x14\x66\x93\x47\xb9\xd6\xc6\xcb\xbc\x1b\x8b\x29\xec\x29\x46\x85\xf6\x1e\x84\x9d\xe3\x87\x21\xbf\x3c\x9a\x4c\x58\x68\xfb\xb3\x95\xda\x2b\x3d\x8e\x6e\x78\xf8\xfd\x79\xc2\x33\x8e\xe3\xa8\x2f\x43\x34\xde\x1e\x88\x20\x21\x2d\x9f\xbd\xfb\x73\xad\x47\x21\x03\x62\x3e\x8e\x2e\xee\x7b\x40\x9e\x7b\x4a\xa6\x40\x60\x01\xcf\x77\x74\x76\x38\x13\x4d\xe9\xf9\x84\x8d\x46\x16\x7d\x63\x35\xe9\x1a\x14\x7b\x4c\xf1\x3e\xec\xfe\x66\xd6\x53\xb8\x43\xac\x1f\xa2\x0e\xf5\x8e\x3d\x29\x7e\x8d\x02\xe5\x0e\x1f\x43\x73\x1a\xd0\xfa\x3b\x1a\x48\x3c\xdf\xfd\x5f\x47\x53\x78\x1c\x22\xb5\x39\x34\x74\x71\x8f\x62\xfc\xf8\x22\xcc\xf3\x42\xef\xa4\x35\x7a\x3c\x99\xb0\x9f\xec\xf7\x00\x00\x00\xff\xff\xdc\xae\xdc\x05\xe1\x08\x00\x00")

func imagesNodeEntrypointMainGoBytes() ([]byte, error) {
	return bindataRead(
		_imagesNodeEntrypointMainGo,
		"images/node/entrypoint/main.go",
	)
}

func imagesNodeEntrypointMainGo() (*asset, error) {
	bytes, err := imagesNodeEntrypointMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/node/entrypoint/main.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesNodeJournalctlToTtyService = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xc1\x6e\xe3\x36\x10\xbd\xf3\x2b\x06\xd1\xa5\x05\x2c\x29\xcd\xa9\x48\xe1\x83\x37\x49\x51\x77\x17\x36\x60\x39\x0d\x16\x41\x0e\x23\x6a\x2c\x4d\x97\x22\x55\xce\x68\x15\xfd\x7d\x41\xd9\x69\x13\x14\xd5\x49\x9c\x79\xf3\xf8\xe6\x3d\x66\x70\x17\x86\x39\x72\xdb\x29\xdc\x5c\xff\xf4\x33\x1c\x3b\x82\xcf\x63\x4d\xd1\x93\x92\xc0\x66\xd4\x2e\x44\x29\x4c\x66\x32\xf8\xc2\x96\xbc\x50\x03\xa3\x6f\x28\x82\x76\x04\x9b\x01\x6d\x47\x6f\x9d\x15\xfc\x41\x51\x38\x78\xb8\x29\xae\xe1\x87\x04\xb8\xba\xb4\xae\x7e\xfc\xc5\x64\x30\x87\x11\x7a\x9c\xc1\x07\x85\x51\x08\xb4\x63\x81\x13\x3b\x02\x7a\xb5\x34\x28\xb0\x07\x1b\xfa\xc1\x31\x7a\x4b\x30\xb1\x76\xcb\x35\x17\x92\xc2\x64\xf0\xf5\x42\x11\x6a\x45\xf6\x80\x60\xc3\x30\x43\x38\xbd\xc7\x01\xea\x22\x38\x7d\x9d\xea\x70\x5b\x96\xd3\x34\x15\xb8\x88\x2d\x42\x6c\x4b\x77\x06\x4a\xf9\x65\x7b\xf7\xb0\xab\x1e\xf2\x9b\xe2\x7a\x19\x79\xf4\x8e\x44\x20\xd2\x5f\x23\x47\x6a\xa0\x9e\x01\x87\xc1\xb1\xc5\xda\x11\x38\x9c\x20\x44\xc0\x36\x12\x35\xa0\x21\xe9\x9d\x22\x2b\xfb\x76\x05\x12\x4e\x3a\x61\x24\x93\x41\xc3\xa2\x91\xeb\x51\x3f\x98\xf5\xa6\x8e\xe5\x03\x20\x78\x40\x0f\x57\x9b\x0a\xb6\xd5\x15\x7c\xda\x54\xdb\x6a\x65\x32\x78\xda\x1e\x7f\xdb\x3f\x1e\xe1\x69\x73\x38\x6c\x76\xc7\xed\x43\x05\xfb\x03\xdc\xed\x77\xf7\xdb\xe3\x76\xbf\xab\x60\xff\x2b\x6c\x76\x5f\xe1\xf3\x76\x77\xbf\x02\x62\xed\x28\x02\xbd\x0e\x31\xe9\x0f\x11\x38\xd9\x48\x4d\xf2\xac\x22\xfa\x20\xe0\x14\xce\x82\x64\x20\xcb\x27\xb6\xe0\xd0\xb7\x23\xb6\x04\x6d\xf8\x4e\xd1\xb3\x6f\x61\xa0\xd8\xb3\xa4\x30\x05\xd0\x37\x26\x03\xc7\x3d\x2b\xea\x52\xf9\xcf\x52\x85\x31\x19\xb0\x97\xe1\xcd\xb5\x96\xb5\x1b\xeb\xc2\x86\xbe\x6c\xb8\x5f\x38\x6e\x4d\xb6\xc4\x21\xb7\x65\xf9\xae\x1d\x06\xf2\xa2\x68\xbf\x95\xc8\x51\x3a\x1e\xf2\x1e\x51\xca\xda\x85\xba\xec\x51\x94\x62\xc9\x3d\xb6\x24\x65\xaa\xe7\x11\xed\xb7\xdc\x06\xaf\x31\x38\x47\xb1\x14\x1b\x79\x50\x29\xff\x0c\x63\xf4\xe8\xac\xba\x5c\x43\xae\x3a\x17\x42\xf1\x3b\xdb\x94\xc7\x31\xbd\xb3\xcb\x11\xe2\xe8\x05\xfe\x45\xa7\x0c\x4f\xc1\xb9\x30\x41\x1f\x1a\x5a\xc1\x10\xd8\xa7\x5c\x50\x97\x05\x55\xe7\x14\xc7\xd4\xb1\xed\xc0\x91\x0a\x8c\x89\xeb\xec\xe8\x85\x06\x5c\x68\x25\x11\xa5\x5a\xd2\x86\xec\x29\x82\x68\x13\x46\x2d\xcc\xf3\xa3\x67\x7d\x31\xf7\x74\xd6\xca\xc1\xaf\x7f\x3f\x0f\x36\x69\x12\x44\x23\x61\x4f\xd1\x1c\xce\xcf\x4e\xd6\x32\x8b\x52\xdf\xe4\x17\xfe\xe6\x9f\x5d\x36\x27\xa5\xf8\xff\x6d\xf3\x5c\x9d\xff\x5e\xcc\x81\x44\x31\xea\x1a\xdd\x84\xb3\xbc\x1d\x2b\xb2\xeb\x6b\xf3\xf0\x4a\xb6\x5a\xba\x65\xcd\xfe\x9d\x75\x90\xe7\x17\x2f\xf2\xdc\x87\x5c\x91\x9d\xa9\x14\x7d\x83\xb1\xd9\x8f\x3a\x8c\xba\x56\x9d\x8d\x79\xde\xa6\xc8\x9c\x7b\x31\x4f\x98\xcc\xfa\x34\xaf\xfb\xd1\x29\xe7\xa3\x50\x2c\x14\x63\x4b\x6a\xfe\x0e\x00\x00\xff\xff\xbb\xd7\xb0\x75\x5d\x04\x00\x00")

func imagesNodeJournalctlToTtyServiceBytes() ([]byte, error) {
	return bindataRead(
		_imagesNodeJournalctlToTtyService,
		"images/node/journalctl-to-tty.service",
	)
}

func imagesNodeJournalctlToTtyService() (*asset, error) {
	bytes, err := imagesNodeJournalctlToTtyServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/node/journalctl-to-tty.service", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesNodeKubeadmSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4f\x6f\xd4\x30\x10\xc5\xef\xfe\x14\x43\x57\xaa\xe0\xe0\x18\x38\x56\x6d\xa5\x22\xb6\x1c\x40\x5a\xa9\x2d\x12\xb7\x6a\xd6\x9e\x24\x66\x1d\x8f\xe5\x99\xf4\x0f\xe2\xc3\xa3\xcd\x26\xa8\x2a\x05\x89\x4b\xec\x99\xcc\xfc\x5e\xde\xcb\xea\x95\xdb\xc6\xec\xa4\x37\x2b\xb8\xd9\x7c\xdc\xbc\xde\x52\xd6\x9e\x28\x05\xaa\x6f\x4e\x40\xfb\x28\x10\x05\xbe\x8f\xa2\x80\xa0\x24\x0a\xe2\x6b\x2c\x0a\x2d\x57\x88\x59\x14\x53\x8a\xb9\x83\xb6\xf2\x00\x63\x11\xad\x84\x83\x59\x81\xf0\x58\x3d\x49\xd3\x34\x50\x69\xe0\x3b\x9a\x58\xc6\x08\x29\x58\x06\xaa\x95\x1e\xa2\x1e\xca\x07\x63\x56\x0b\x0b\x76\xe3\x96\x30\x0c\xe0\xa6\x5b\x22\x9d\x6f\x5e\xd3\x33\x11\x2c\x6a\x3b\x52\x18\x4b\x40\x25\x38\x3e\x86\xa5\xb3\xb0\xec\xe3\xd4\xd2\x8a\x59\x0a\x57\xb5\xbd\x6a\x11\xf0\x63\x4d\x66\xff\x00\x2b\x30\xb5\x4e\x9c\x2b\xe8\x77\xd8\x91\x34\x3e\xf1\x18\x9a\x8e\xb9\x4b\xd4\x78\x1e\x1c\x16\x75\x81\xfd\xfe\xb4\x3b\x7a\x6c\xba\xd2\xc1\x4f\x98\x2b\xc0\x10\xc0\x1a\x8f\x0a\xa7\xa7\xeb\xcd\x25\x9c\x3b\xd2\x69\xd6\x2d\x11\xa4\x28\xda\x04\xb7\x37\x51\x33\xe9\xdc\x31\x81\xb6\x93\xf8\x89\xdb\x4f\x37\x4f\x5e\x47\x3e\x58\x3e\x94\xf6\x81\x72\xc4\x04\x03\xc6\x6c\xd6\x9b\xcb\x67\xbe\xcd\x0b\xa6\xad\xcd\x6c\xe7\xda\x56\xf2\x3c\x0c\x94\x83\xfc\x4e\x74\xc9\x78\xce\x75\x42\x0c\x58\x77\xd0\x73\x0a\x7f\x9d\x32\xe4\x7b\x86\xa3\xcf\x5f\x3f\xac\xbf\xac\x6f\x6e\xd7\xdf\x6e\xae\x2e\x6e\x2f\xae\x3e\x5d\x9f\x59\xdb\x62\x4c\x56\xee\xb1\x58\xce\x67\x2d\x26\xa1\x23\x38\x3f\x87\x29\x8b\x40\x2d\x8e\x49\xdd\xcc\x35\x46\x1e\x45\x69\xd8\xff\x50\xe9\xf9\x7e\xd1\x6b\x84\xea\x5d\xf4\x34\xcb\x20\xe2\xd1\x14\xeb\xc4\x38\xac\x84\xf9\x74\xcf\x56\x9a\xe0\xde\xbd\xb5\xf3\xf7\x36\x9e\x73\x3b\x43\xb6\x09\xfb\xff\xa1\xbc\x7f\x99\x82\x3f\x9e\x40\xfe\xb0\xb3\x82\x2d\xb3\x82\xf6\x04\x03\x8a\x52\x85\xcc\x81\xcc\x0a\xac\x8d\x5d\xe6\x4a\xb6\x54\x6a\x53\xec\x7a\xb5\x54\x2b\x57\x39\xbb\xbe\xc7\x62\x96\x7c\x63\x8e\xfa\x8f\x59\x4c\xc9\xfc\x0a\x00\x00\xff\xff\x48\xe6\x84\x5b\xa7\x03\x00\x00")

func imagesNodeKubeadmShBytes() ([]byte, error) {
	return bindataRead(
		_imagesNodeKubeadmSh,
		"images/node/kubeadm.sh",
	)
}

func imagesNodeKubeadmSh() (*asset, error) {
	bytes, err := imagesNodeKubeadmShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/node/kubeadm.sh", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
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
	"images/node/20-kubeadm.conf": imagesNode20KubeadmConf,
	"images/node/Dockerfile": imagesNodeDockerfile,
	"images/node/debug-node.service": imagesNodeDebugNodeService,
	"images/node/entrypoint/main.go": imagesNodeEntrypointMainGo,
	"images/node/journalctl-to-tty.service": imagesNodeJournalctlToTtyService,
	"images/node/kubeadm.sh": imagesNodeKubeadmSh,
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
	"images": &bintree{nil, map[string]*bintree{
		"node": &bintree{nil, map[string]*bintree{
			"20-kubeadm.conf": &bintree{imagesNode20KubeadmConf, map[string]*bintree{}},
			"Dockerfile": &bintree{imagesNodeDockerfile, map[string]*bintree{}},
			"debug-node.service": &bintree{imagesNodeDebugNodeService, map[string]*bintree{}},
			"entrypoint": &bintree{nil, map[string]*bintree{
				"main.go": &bintree{imagesNodeEntrypointMainGo, map[string]*bintree{}},
			}},
			"journalctl-to-tty.service": &bintree{imagesNodeJournalctlToTtyService, map[string]*bintree{}},
			"kubeadm.sh": &bintree{imagesNodeKubeadmSh, map[string]*bintree{}},
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

