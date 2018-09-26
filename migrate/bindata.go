// Code generated by go-bindata.
// sources:
// migrations/20170523183416_init.sql
// migrations/20171108110099_rbac.sql
// migrations/20171122172200_indexes.sql
// DO NOT EDIT!

package migrate

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

var _migrations20170523183416_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x7d\x4b\x73\xdc\x36\xb6\xff\x3a\xfe\x14\xa8\xd9\x74\x5c\xd3\x4e\x39\x93\xc9\xff\x7f\xab\xa6\x66\xa1\x58\x8a\xaf\x72\x15\x39\xb1\xe5\xa9\x64\x45\xa1\x49\x74\x37\x2c\x12\xe0\x00\xa0\x64\xe5\xd3\xdf\x22\x00\x92\x78\x13\x6c\xa9\xdb\x99\xeb\x4d\xa2\xe6\xef\xe0\x79\x5e\x38\x38\x00\x5e\xbd\x02\x6f\x2f\xae\x2f\xde\x9f\xdd\x5c\x9c\x83\x0f\xbf\x5e\x7d\x03\xce\xdf\x81\xeb\x77\x37\xe0\xe2\xfc\xf2\xe6\xc5\x8b\x57\xaf\xc0\x5f\x77\x94\x72\x04\x3e\xb6\xfd\x1f\x1f\x7e\xbd\x02\x98\x00\x8e\x4a\x81\x29\x01\xab\x8f\xed\x0a\x60\x0e\xd0\x67\x54\x76\x02\x55\xe0\x61\x8f\x08\x10\x7b\xcc\x41\x83\x77\x0c\x4a\x10\xe6\x00\xb6\x6d\x8d\x51\x25\xcb\x03\x67\x6f\xae\xde\x32\x48\x84\xfc\x00\x1a\xd8\xb6\x98\xec\x00\xdd\x02\x86\x38\xed\x58\x89\x80\xa0\x80\xd1\x1a\x71\xf0\xb0\xc7\xe5\x1e\xb4\x8c\xde\xe3\x0a\x01\x08\x5a\xc4\x1a\xcc\x79\x5f\xea\x96\x32\x00\x2d\x12\x28\x89\x5e\xbc\x79\x7f\x71\x76\x73\x01\x6e\xce\x7e\xb8\xba\x00\xb7\xb0\xac\x8b\x5d\x5f\xdb\x2d\xf8\xfa\xc5\x57\xb7\xb8\xba\x05\xc6\xbf\x7f\x9d\xbd\x7f\xf3\xdf\x67\xef\xbf\xfe\x7f\x7f\x7f\x29\x3b\x7d\xfd\xf1\xea\x0a\xfc\xf2\xfe\xf2\xe7\xb3\xf7\xbf\x83\xff\xb9\xf8\x7d\xfd\xe2\xab\xdb\x72\x8f\xca\x3b\xde\x35\x9a\x70\xc0\xf7\x9f\x86\xda\x0b\x55\x6c\xa8\x34\x09\xa3\xf5\x00\x89\x55\xda\xc3\xa6\xce\x49\xe4\xc5\xf5\xc7\x9f\xbf\x5e\x31\x04\xab\xd5\x5a\xfe\xe7\x81\x61\x81\x56\xeb\x15\xac\x1a\x4c\x56\x36\xad\xfc\xad\xe8\x38\x62\xb2\xa2\x58\x15\x65\xc7\x05\x6d\x34\x28\xde\x92\x92\x21\x28\x50\x55\x40\x21\x5b\xf2\xc3\xe5\xdb\xcb\xeb\x1b\xf0\xf1\xfa\xc3\xe5\xdb\xeb\x8b\x73\x0b\xda\xb5\x55\x02\xba\x7e\xf1\xd5\xe5\xf5\xf9\xc5\x6f\x60\x9c\x86\xc2\x18\xb3\x02\x93\x0a\x7d\x06\x5f\x5b\xe3\xf8\x32\x44\xa3\x06\x70\xc2\xeb\x01\x0d\x61\x8d\x2e\x8e\x78\xb3\xdb\x2f\x5f\x7c\xf5\x12\x5c\x5c\xbf\xbd\xbc\xbe\xf8\xe7\x25\x21\xf4\xfc\x07\x70\x7e\xf1\xe3\xd9\xc7\xab\x1b\x39\xb5\x1f\x2e\x6e\xfe\xd9\x89\xed\x7f\x35\x9b\xbf\x83\x37\xef\xae\xae\xce\x6e\x2e\x86\xbf\x8b\x8e\xe0\x92\x56\xa8\x28\xf1\x3f\x46\x4e\x7e\x3f\x30\xa0\x64\xe6\x0a\x6d\x31\x41\xd5\xc4\x96\xb8\x97\x06\x04\xf8\x23\x17\xa8\x09\xb0\xe6\x00\x3c\x06\x77\x76\x8c\xdc\x06\x4a\xfb\xdb\xf7\xdf\xdb\xb3\x5d\x21\x5e\x32\xdc\x0a\xcd\x78\x37\x17\xbf\xdd\xf4\x3f\x0b\x2c\x6a\x34\x15\x30\xfc\xdc\x76\x9b\x1a\x97\xe3\xef\x37\x97\xd7\xbf\x5f\x5e\xdf\x7c\xfd\xad\x5d\xe6\x1e\x57\x15\x22\x73\x28\xc9\xb5\x46\x15\x61\xd4\x71\xb8\x71\xe4\xb9\x8e\x91\x91\x51\xfa\x21\x7b\x6e\x06\xa1\xb5\x66\x8e\x9e\x6b\x01\x6f\x51\x89\xb7\xb8\x04\x02\x6e\x6a\x34\x32\xcc\xe6\x11\x40\x02\xe4\x78\x84\xb8\x84\xd6\x47\xe1\x10\x02\x1b\x63\x86\xc7\xd2\xbe\x7d\xfd\x3a\x8b\x43\xe2\x4a\x27\xad\x6b\x8e\x38\xa9\xbd\x62\xe8\x7b\x35\xce\xa8\xec\xe2\x4b\x1f\x74\x62\x35\xf1\x91\x23\x26\x39\xa1\x81\x2d\x07\xfd\x90\xf1\xd1\xca\x05\x26\x5c\x8e\xe9\xb1\x66\x7d\x9c\xb0\xb4\x39\xca\xb4\x5a\x99\x26\x25\xd7\x40\x1d\x85\x2d\xc6\xe1\x1c\x1a\x30\x09\xbc\x6e\xd0\xcb\x08\x3e\xcf\xf2\x4c\xf8\x93\xb2\x55\xdb\x2a\xc5\xb2\xc7\xbb\x3d\x62\xa0\x46\xf7\xa8\x06\x70\xc3\x05\x83\xca\x3d\x83\x8c\x76\xa4\x02\x10\xd4\x74\x87\x4b\x58\x83\x8e\x60\xd1\x7b\x5a\x7d\x41\x00\x92\x4a\x1b\x26\x8f\x05\xdb\xf6\x84\xea\x26\xd7\x20\xf5\x9d\xba\x9f\x0a\xf8\xe1\xdd\xbb\x2b\x9b\x61\x51\x4b\x0b\x5c\x71\x8d\xf8\xe9\xc3\xbb\xeb\xe7\x66\xa9\x4c\x6e\xd7\x9c\xd1\xb6\x69\x5d\xd4\xb6\x27\xe5\x97\x37\xb4\x69\xb0\xf6\xba\x77\x88\x20\x36\x5a\x21\xe5\x66\xab\x61\xdf\xa0\xfe\x3b\xc7\x64\x57\x23\xc5\x27\xa5\xa4\x73\x58\x44\xfd\x18\xe4\x92\x03\x39\xc5\xe5\x16\x3d\x9d\x21\x17\xc6\xd5\x1a\x7c\x0f\xdd\x16\xc4\xa0\x1b\x06\x49\xb9\x0f\x72\xb5\xe4\xc3\x61\x74\xff\xd2\x40\x2e\x10\xfb\x4b\x4f\xd3\x20\xce\xe1\xce\x64\x5d\x70\xf5\xee\xfa\xed\xc0\x97\x0d\x62\x3b\x34\x8c\x87\xc1\x9b\x43\x51\x5b\x58\x73\xd4\x03\xd1\xe7\xb2\xee\x2a\x64\x76\x29\x0c\x6c\x21\x43\x44\xc4\x44\x6f\x42\x58\xc3\xe3\x20\x7a\x4e\x76\x06\x25\xc5\xfa\xb0\x13\x7b\xca\x0a\xd3\x36\x44\x75\xb4\xec\xaa\x40\x6c\x5e\x9d\x53\x56\x61\x02\x6b\xb3\x1d\x97\xd7\x37\x01\x41\x4a\xd5\xc7\xd0\xb6\x10\x8f\x6d\x50\x75\xbc\xf6\xa1\x19\x3a\x4b\xcd\x9a\x80\x15\x14\x26\xe3\x68\x8d\xa1\xa4\x53\xf5\xb2\xd0\x5c\x68\x2c\x55\x14\x57\xbe\x74\x81\x7c\x0f\x47\x50\xcf\x8f\x1e\xc0\x1e\xe1\x11\xeb\x0c\xbc\x47\xe6\x8d\xf6\xa4\x27\xbc\x79\xf0\x89\xe7\xd4\x8b\xdf\x5d\x35\xd6\x46\x7f\xf5\xe0\x07\xa1\xd6\xc0\x6c\x8f\xa4\xb0\xce\x7a\xbd\x8f\xc5\x63\x5c\x73\xd5\x98\xdc\x19\x5e\x55\x4d\xcb\xde\xa1\x56\xed\xd4\x81\x82\x5e\x1c\x82\x2a\xac\x80\xba\x78\x5f\x97\xe5\x28\x31\x53\xcc\x52\xf2\x65\xaa\xa8\x18\x4b\xe6\x08\x9f\xa9\x15\x63\x98\x2d\xae\xd1\x60\x6b\x07\xcc\xff\x77\x7d\x7a\x53\x34\x53\x0d\xaf\xe9\xb4\xd8\xeb\x41\xe3\x40\x0c\xd3\xfa\x5a\x76\xcf\x80\xc5\x51\x9b\x1a\x92\x3b\x6d\x9f\xe3\xa8\x7e\x5e\x10\x11\x12\x17\x45\xd9\xa2\xa5\xa7\x30\x43\x08\x07\xe4\xbc\x23\xe8\x52\x64\x6b\x82\x91\x62\x98\x87\x91\x64\x9c\x98\x0c\x9a\xa1\xba\x9e\xc1\x0a\x3d\x59\x7e\x41\xeb\xb1\x19\x6b\xc5\x8a\xeb\x71\x62\x8f\x20\x89\x3f\xe2\x61\x2d\x9b\xe9\x3f\x6c\x25\x01\xe9\x7f\x18\xd6\xbd\x09\x6f\x42\x76\xfe\x98\x2e\x85\xae\x66\x4e\xc6\x17\x7a\x1f\xc7\x31\x9d\xa6\x10\x3b\x0d\xf8\xf6\xf5\xdf\x1c\x70\x0d\xc9\xae\xb3\x5d\x94\xe0\x62\x7e\x72\x6f\x3a\x72\x47\xe8\x03\xf9\x8b\x5a\x9b\x55\xb8\x77\xb6\xf9\x2d\x98\x97\xf4\x0a\xd5\x28\x1b\xcc\xf1\x1f\xae\x13\x12\x07\xc3\x0d\x26\x90\x3d\x9a\x78\xcb\xc7\xf7\xfd\xa4\x85\x3e\xce\x9c\xef\xe7\x8f\x92\xe5\x04\x86\xbd\xbc\x78\xeb\x62\xee\x5e\x9c\xc2\xd4\xb5\xf3\x63\xeb\xa3\xf3\x74\xea\x2c\xd8\x54\xd3\xb3\xe0\x7b\xc8\x30\x24\x25\xca\x2a\x99\x0b\x28\x3a\xbb\x64\x15\xe6\x86\x55\x85\xaa\xd5\x7a\xd5\xd0\x0a\x6f\xb1\xfc\x5f\xc9\x68\xa8\x5a\xb9\x72\xd9\x0b\x45\xb5\x80\x49\x34\x45\xb1\x65\x74\xd0\x0b\x83\xe7\x3e\x7c\x12\x74\x8a\x54\xea\x4f\xff\x29\xfe\x69\xaf\x26\x8a\x51\xaf\x39\xbe\x61\xd0\x3c\x29\x8a\xa5\x8e\xa1\xa4\x5a\xe0\x1d\x8e\xf8\x93\xb8\x88\x97\x9c\x77\x09\xcb\xa4\xfc\x43\x4a\x10\xa0\x0c\x34\x94\xf5\x50\xde\x21\xae\x76\x8d\xde\x24\x2c\x92\xc4\x85\x4d\xd2\x21\xf6\x28\x65\x8c\x72\x57\xac\xf9\x2a\xcb\x09\xf4\x65\xdb\xba\x18\x4e\x0e\x46\x86\x40\xb8\x6a\x39\xa5\x93\x5d\x39\xcb\x15\xb2\x94\x84\x99\x23\x9b\x2a\x2f\xc7\x97\x96\x3c\x2b\xfb\xad\x6a\x4f\x4a\xb6\x31\x3e\xb9\x62\x1d\x92\x69\x55\xce\x22\xa1\xd6\x24\x4b\xa5\x5a\x91\x2d\x10\xeb\x89\xc0\x6f\x9e\x3d\xf4\x09\x52\xdf\x93\x9e\xe6\x22\x41\x36\x4d\x83\xad\x51\xa6\xd9\x99\x25\x76\x95\xd1\x30\x61\xcf\xa5\x92\x82\x3a\x84\x77\x4d\x23\xfd\x9a\x83\xb5\x48\xa6\xa6\xc8\x95\x63\xd7\x61\xcd\x0c\xed\x3f\x59\x36\x5d\x3f\x33\xdf\xc9\x8c\x23\x7b\x3b\xc3\x8b\x72\x0f\xc9\x4e\x7a\x05\x09\xaf\x26\xa5\x47\x43\xea\xd3\x75\xaa\x07\x82\xef\xe7\x3d\xea\x25\x1a\xd0\x8d\x2e\x1a\xa1\x45\x3b\xc0\xa4\xd8\x68\x99\x5e\x18\x88\xf2\xc3\x58\x9a\xe0\xc0\x68\xd5\xd8\xc8\xa5\x9a\x68\x6a\x68\x54\x17\x3d\xab\xdb\xc0\x05\x78\x83\x88\x40\x4c\x79\x0d\x7d\x0d\x6a\x63\x84\x0b\x50\x21\x01\x71\x0d\x4a\x4a\x04\xc4\x04\x31\x20\xf6\x50\x80\x12\x12\xb0\x41\x80\xa1\x57\x1d\x47\x95\x0c\x2b\x35\x5d\x2d\x70\x5b\x23\x15\x7a\xf2\xa4\x9f\x8b\xa2\x94\x75\x1c\xc9\x81\xc8\x14\x4f\x67\x15\x29\x5d\xdc\xd4\x7e\xcb\xe8\x03\xe3\x0a\x11\xd1\xbb\xe3\xec\x16\x98\xbf\x77\x44\xc6\x25\xf9\xb0\x10\xfa\xf1\xea\xdd\x99\x5d\xe2\x3d\xac\x47\x8b\x09\x26\x27\x7f\xc3\x69\xdd\xc9\x1c\x96\x16\xb1\x7e\x68\xe0\x0e\x39\xfe\xbd\x2c\xdb\x50\x26\x8a\x92\xc3\x1a\xb2\xc7\xd5\x7a\xb5\xa7\x1d\xab\xfb\xff\xe9\x27\x87\xc1\x52\x50\xd6\xff\x01\x79\x07\xeb\xd5\x7a\x45\xc5\x1e\xb1\x50\x89\x65\xc7\x18\x22\x65\xbf\xb8\x94\x63\xf4\x9d\xa3\x99\xea\x9a\x96\x70\x1c\x80\x40\x8f\x9c\x5d\xa4\x25\xdb\x48\x2e\xf6\x39\xf9\x58\xb3\x40\xc8\xf5\xed\x19\x74\x4f\xeb\x0a\x93\x1d\x18\x58\x65\x0a\xc3\x28\x1e\xf7\x78\x56\xe3\x02\x0c\x9b\xc3\xad\x16\xa7\x79\x6c\x66\xed\xdf\xcd\x2c\xda\xac\xe1\xce\x1e\xeb\xc0\x86\x9d\xe5\x6f\xf5\xce\xd6\x33\x8e\xfe\x39\x14\x10\xbc\x65\xb4\xd3\xfb\xb0\x0f\xf0\xb1\x5f\x51\x50\xb6\x83\x04\xff\x81\xc0\x9e\x3e\xc8\xec\x9f\x8f\x97\x60\xd7\xa3\xb8\xdc\x72\xdd\xe2\x5a\x20\xc6\x41\xdf\x2a\x67\xf8\xfb\x9f\x0a\x09\x3d\x6a\x14\x2c\x7f\x69\xeb\xaa\x8e\x1c\xed\x61\xe6\x0a\xcd\xee\x52\x29\x0f\xac\x65\xf4\x13\x2a\x85\xda\xbc\x1d\xf6\x6d\xb5\x8d\xe1\xfe\x02\xd8\xdd\xe9\x05\xf1\xdd\xde\xa5\xa2\x7a\x54\x69\xbd\xf8\x5c\xd6\x1d\xc7\x94\xfc\xac\x73\x1e\x95\xe9\x31\x16\xaa\x7b\xc8\x01\xac\x6b\xc9\x35\x43\xf8\x08\xa8\x65\xb3\xc3\x2a\x68\x28\xab\xd0\x09\x94\x21\x8e\xc9\x61\x16\x3b\x24\x91\xe7\xe2\xc5\x1c\x3c\x7b\xc1\x99\xdc\x90\x9c\xdd\x5b\x4f\x6e\xad\x7b\xbd\x8f\x85\x1a\x66\x28\xe6\x56\x3d\x3e\xd5\x29\x37\xec\xa3\x51\x0d\xb9\xa5\x45\x54\x10\xc3\xe1\x8b\x44\xc0\xe2\x50\xe5\x91\xc8\x1d\x0a\x6c\x4c\x8f\xa2\x3c\xc7\x52\x4e\x7e\x61\x48\xb5\x6c\x68\xf5\x68\xf7\xc1\xdc\x77\xe7\xc2\x0d\xff\x2a\x67\x81\xb6\x88\xf4\x4e\x41\x4d\xb9\x17\x3e\x74\xe2\x74\x23\xcd\xa6\xdb\xad\xd6\xab\x2d\x82\xa2\x63\xbd\x7f\x82\xc8\x1e\x92\x12\x35\x88\x88\x88\x5b\xc1\x90\x74\x66\x0c\xad\xa7\x4a\x22\x94\x20\x99\xb0\xcb\x69\x7d\x2f\x03\x99\x1d\x19\xff\x48\x26\x3b\x2d\x56\x55\x21\x19\x92\x9d\x4e\xa8\xb3\x09\xb4\x79\x8c\xa6\xee\x75\xac\xbe\x0d\x73\x8e\x97\xab\xe3\x1b\x93\xfc\x30\x69\x7e\x94\x34\x3f\x48\x6a\xc5\x53\x96\x46\x45\xf2\xc3\x21\x46\x60\x21\xa1\x74\xfc\xfa\x87\xf5\xd4\x24\x26\xc1\x26\xad\x47\x41\x5b\x9b\x12\xe5\x96\x2c\x25\xa0\x98\xd8\x28\xd8\x55\x25\x26\x6b\x93\xdb\xd6\xc9\xfe\xcb\xbe\x9b\x25\x05\x1a\x2b\x87\x63\xed\xb4\x39\xd1\xd2\x71\x4c\x55\x93\x8d\x12\x43\x4d\x1e\x07\x7c\x3d\xb6\xde\x28\x3b\xdd\x7a\xa7\xa6\x51\x20\x0e\xa8\x73\x12\xa6\x27\xd4\x3f\x4d\xce\x01\x0d\x30\xa7\x2c\xbb\x05\xba\xe2\xf4\x04\x8e\x55\xe4\xce\xa0\x2a\x55\x73\xfc\x34\xaa\x6e\x79\x5a\x06\xcc\xc1\x9b\x6f\xdf\x54\x9c\x2e\xdf\x12\xbf\x70\x53\xcd\xc9\x19\x2b\x7d\xf6\x40\x81\x32\xc1\x6f\xd4\xee\x59\xde\xd6\xb7\xde\x6a\xd3\x79\x27\x51\x1b\x5d\x68\x5c\xc4\x56\x1f\x64\xaa\x7d\xcf\x3a\x23\x1a\xe8\x99\xf4\x39\x8b\x1e\xdd\xa8\x70\xad\xf5\x60\xac\x9d\xe5\x82\xdc\xd6\x33\x3c\xc1\x99\x3d\x5e\x85\x8f\xa7\x99\xbb\x46\x74\x91\x0d\x0d\x58\xc7\xa1\xc2\x14\xc6\x33\x90\xf9\xf6\x31\xdb\x3c\x66\x5b\xc7\x6c\xe3\x18\xb2\x8d\x8a\x09\xfd\xf0\xf8\x14\x1a\x8f\x50\xcc\x27\xd0\xd8\xf8\xa5\x56\x78\xa0\xcb\xb7\xc6\x26\x45\x86\x55\x0e\xb4\xeb\x89\xd6\xf9\x64\xae\x3f\xd0\xb5\x6a\xed\xa3\x03\x95\x5c\xef\x6a\x06\xb5\x8d\xa6\x38\x74\x95\x68\xab\x9a\x78\x52\xb7\x27\xee\xb6\xa4\x78\x9f\x6d\xe9\xcd\xdf\x28\x7c\xda\x3a\x35\x6f\xcd\x6b\xcb\x4f\x40\x78\x86\xb9\x5f\xca\xda\x03\x5d\x3e\x6b\x9b\x14\xc7\xdd\x50\x97\x8c\xf7\x41\x05\xe3\x15\xff\xa9\xc8\x3c\xfe\x43\x1d\xe6\xd4\x56\x0f\x8b\xbd\x6a\x9a\x3e\x75\xa0\xc3\xf7\xa1\xa0\x96\x36\xfa\xa9\x4d\xb0\xa7\x5b\x3e\xf3\x5f\xd2\x0a\xce\xd5\x24\x17\x89\x54\xc0\x5a\xa9\x44\xee\x11\x46\x19\x94\xa0\x87\xef\x5e\x17\x15\x7c\xf4\x69\x12\x54\xaa\x2e\xe5\xcc\xe4\xd7\xa5\xf0\x91\xea\xa2\x54\x88\x0b\xdc\x48\x71\x7b\xa0\xec\xae\x68\x28\x11\xfb\x91\xda\x0f\x79\x07\xe1\x63\x9d\x3e\xde\x5f\xd0\x1b\xff\x66\x14\x43\x00\xad\x7c\x0f\x4c\x99\xcc\xcd\x4d\x4c\xdf\xb7\xaf\x5f\xbf\x34\xc1\xc1\xe9\x76\xbc\x86\x40\x22\x52\xbc\x60\x05\x8e\x70\x51\x30\xa0\xe9\x87\x19\xbc\x82\x03\x99\x15\x4a\x17\xa4\xa2\x36\xe1\xc0\x43\xa2\xe9\x13\x78\xbe\x60\x3f\x9e\x12\x2c\x38\xe4\xdf\x14\x5b\x8c\xea\xca\x09\xc9\x4e\x41\x59\x81\x60\x13\x1b\xec\x11\x34\x04\x8b\x83\x12\x1b\x3c\xfe\xc0\x47\xb0\x55\xf8\x50\x60\x20\x83\x2a\x50\xab\x17\xaf\x0a\xd7\x6a\x27\x73\xb7\x0c\xdb\xc9\x74\x81\x92\x6b\xb8\x41\x75\xac\xd7\xe3\xc8\xd0\xb6\x90\x1a\x34\x84\xd3\xe7\x61\xc3\xc2\x8c\x79\x51\x23\xb8\x0d\x97\x9f\xa4\x6c\xa1\xd8\x47\x45\x2f\x9e\x63\x1a\xd8\x36\x98\x1f\xa9\x40\x6e\x9c\x4f\x94\x91\x25\x37\x5f\x93\xcb\x88\xc5\x3d\x66\xa2\xd3\x09\xee\x53\xc2\x5f\x8d\x20\x47\x45\xd5\x31\xe8\x49\x90\xf6\x3e\xbc\x60\xb3\x65\xc0\x16\x78\xca\x03\x85\xa1\x97\x46\x22\x53\x57\xc5\xe8\x2c\xa5\xe0\x54\x37\xa8\x8a\x18\xad\x25\xf7\xd6\x39\xfb\x49\x1b\x44\xdb\x6b\x0b\xe2\xd4\x66\x47\x40\xe3\xfd\xf5\xbc\xe7\x54\x5c\x41\x53\x8d\x92\x30\xc5\x98\x46\xd9\x88\xd1\x2c\x75\xbc\xcc\x24\x83\xf9\x15\x42\xa8\x96\xd8\xd8\x38\x61\x94\xcc\x91\x32\x4b\xf6\xfb\x6f\x97\xb9\x6c\x34\x8c\xd2\x82\x5c\x14\x2b\x7b\x9d\xc9\x5f\x39\x75\xd9\x7c\xbf\xa0\xd2\xb5\x2d\x1e\xcf\xe8\xdd\xfe\x84\xc8\x1d\x26\xfc\x87\x0e\xd7\x55\xef\xdd\x8a\x3d\x02\x1b\xf9\x07\x25\xc3\xc7\x6f\xc0\x4f\x74\xc3\xa7\x45\x95\xfc\xce\xbf\x71\xbc\xda\x4f\x0a\x5c\xc8\xaf\x29\xaf\xf6\x70\x77\x76\x66\xcf\x36\xa1\x0b\xad\xc6\x15\xa4\x6b\x36\x88\xdd\x06\xf4\xdb\xed\x27\xba\x49\xfb\x05\xee\x96\x47\x57\x0b\x17\x3d\xdc\x53\xd2\x7f\x2b\x74\x12\xd4\x6a\xbd\xe2\x5d\x59\x22\xce\x57\xeb\xd5\x16\xe2\x5a\xed\xaa\xc0\x0d\x65\x7e\x4a\x77\x60\xcb\x76\x66\x2d\xb8\xc5\x04\xf3\xbd\x4b\x12\x88\xd3\x04\x0e\x86\x3a\xc1\x9a\x50\x66\x58\x18\x15\x76\x58\x63\x81\x1f\x25\x34\xf6\x44\xe4\xab\x2b\x9b\x2e\x34\x9d\x63\x01\xc1\xb9\x8e\x96\x24\xe7\x7b\xa2\x55\xd3\x7f\x04\x09\xfb\x89\x6e\xd4\xea\xf1\x13\xdd\x8c\x97\x5c\x98\x32\x76\x4d\x59\x03\xeb\xfa\x11\x94\x94\x31\xc4\x5b\x4a\x2a\x9d\x8e\x3d\xc4\x36\x20\xa9\x72\x85\xf0\x13\xdd\x3c\x7f\x60\x23\x2f\xe0\x60\x87\x3f\xa2\x61\xc0\x9d\x3e\x4a\x7a\x9b\xc1\x5a\x79\x2c\xd5\xcf\xdd\x72\x86\xea\xa9\x02\x67\xd2\x9f\x93\x01\x30\x83\x2a\x2d\xe9\xc7\xde\x15\xf3\x73\x1c\xc6\x39\xed\x95\x6f\x47\xf0\xbf\x3b\xa4\xf3\x93\x80\x74\xde\x38\xd8\xaa\x5b\x0b\x4a\x46\xb9\x4a\x87\xf8\x84\xd9\xc8\x17\x6e\x78\xab\xff\x56\x98\xce\xdf\xf1\x36\xbf\x07\xdf\x32\x18\x2a\x9f\xc9\x9b\x89\x72\x06\x2f\xf7\xa8\x81\xa6\x93\x1c\x5d\x25\x6a\xa8\xea\xec\x6d\x4e\xa9\x7a\x5c\xfa\x26\xcf\x5f\x66\x30\x77\x9f\x81\x37\xd2\xc9\xdb\x0d\x7c\xf4\x29\x53\x27\x1c\x2e\xfc\x17\xac\x87\x70\xea\x12\x56\x04\x32\x8f\xf1\xa9\x0c\x59\xc8\x52\x8e\xb5\xcf\xe3\xf1\x4f\x94\x7d\x74\x33\xa6\x7f\xc1\xfd\x19\x67\x1d\x95\x9f\x1d\xbe\x94\x79\x54\x8e\xa8\x53\x9d\xc3\x18\x53\x33\x52\x8c\x65\x96\x74\x22\xf6\xba\x7c\x7f\xa6\xe3\x9f\x2a\x46\xda\xa2\x12\xc3\x1a\xff\x81\x2a\x70\x8f\x98\xbc\xa0\x8e\x6e\x01\xd4\x71\xfc\x2d\x65\x92\x24\xc4\x29\xc9\x84\x9d\xa7\x68\x2e\x9f\x55\x62\xc1\xd0\xa8\x16\x89\x05\x47\xd2\x07\x84\xfc\x30\x56\xf6\xd6\xa3\x8f\x97\x89\xa3\x9c\xe3\x1d\x41\x5e\xd3\xbd\x8d\xca\x70\xfc\x2f\x18\xfa\xf3\xab\x4d\x78\xbd\xc1\x10\x9a\x53\xec\x16\x7f\x2e\xf4\xd4\x5b\x01\xa9\x21\xd2\x13\xfa\x66\x7e\x47\xe4\x1e\x33\x4a\x1a\xf7\xe2\x8f\xf1\x40\x22\x6d\x5a\x4a\x54\x3a\x65\x28\xf1\x51\x86\x9b\xbc\xd2\xa7\xef\x55\x87\xe4\x51\x72\xe8\x38\xf0\x01\xbb\xd0\xd6\x90\x10\x54\x15\x5c\x40\x26\xd2\xf9\x46\x03\x14\x11\xdb\x19\x0f\x40\xef\x50\x3a\x8c\xeb\x46\x81\x62\x91\xc5\xa1\x43\x2a\x08\xe7\xcf\xe3\x38\xa0\x2d\x2e\x33\x98\x2b\x90\xa3\x1a\x82\x71\xc1\xa0\x40\x3b\x5c\x16\x36\x81\x1b\xfe\x25\x45\xcb\xe8\x8e\x21\xce\x8b\x92\x76\xe3\x5c\xba\x0b\x2f\x86\x68\x8b\x88\x05\x09\xe2\xcc\xf2\xcc\xa0\x55\x60\x2d\x77\x8f\x18\xde\x62\x95\xc5\x6e\x62\x03\x50\xb3\x54\x73\x92\x15\x34\x15\xec\x4b\x1e\x93\x5b\x10\x17\x54\x37\x66\xe5\x9f\x86\x35\x4e\x1b\x58\xff\xde\x5f\x9c\x05\x92\x07\x5e\xdb\xc6\x42\xeb\xbd\x8c\xc0\x9d\x01\xcf\x8d\x62\x79\x35\xe4\x06\xec\x0c\xc2\x3b\xf4\x38\xc2\x7b\x31\x09\x82\x34\x3f\x8f\xc0\x81\xbf\xc3\xad\x1f\x98\xd4\x8b\xde\x45\x08\x02\xfc\x6d\x24\xe0\xf8\xbc\x1f\x2c\x64\xc1\x7a\x64\x22\x9a\x0b\xc6\x19\x50\x83\x6b\xec\xbc\x83\x67\xde\x08\x8d\x99\x69\x7d\xce\xac\xa8\xe9\xc1\xe9\xd6\xc9\x14\x69\xdb\xd0\x25\x13\x5d\x02\x0a\x56\xaf\x7e\x5c\xb5\xea\x1a\x75\xa5\x50\x07\x7f\x71\x30\x2f\xc6\xe9\x78\x27\xcd\xb6\xff\xd2\xb3\x87\xcc\x31\xb7\xbe\x18\xc7\xe6\x81\xfb\x65\xa2\xb0\xbf\xe4\x6d\xbd\xe7\xad\xbd\x6d\x7f\x26\x86\xb2\x9d\x98\x18\xca\xbe\xa9\x2a\xb8\x0f\x10\xe4\x83\xc1\x3e\x8d\x17\xb3\xf4\x73\x10\x16\x8e\x89\xe6\x20\x39\x31\xe8\x13\xaa\xec\xb9\x9d\xdc\x40\x02\xdc\x8c\xb3\xfb\xc6\xc8\x7f\x33\xfc\xe4\xe1\x06\x8a\x84\x70\xc5\x13\xe2\x9e\x18\xc9\x19\x92\x9c\x2c\x0e\xf0\x04\xe3\x39\xd9\x29\x29\xe6\x76\x16\xda\xf3\x24\xbd\xcc\xca\x94\xcf\x4f\x43\x02\xd4\x38\x3a\xd6\x81\xd2\xa4\x9e\x5f\x90\x35\x16\x20\x3b\xc0\xb4\xe6\x67\x9d\x85\x7a\x78\x90\xb8\xa5\xb3\xc9\x8e\x23\x67\x57\xbd\x13\xaf\xa4\x4c\x87\x23\xa4\x5b\xaf\xb3\xbe\x76\xf8\x1e\x91\x21\xec\x10\x97\x23\x49\x72\xec\x44\xaf\xcc\x84\xc7\x5c\x66\x54\xab\x97\xd9\x78\x92\x09\x3e\xf9\x7a\x5f\x4d\xd1\x2f\x7a\x95\x39\x6c\x1f\x0d\xab\x4e\x15\x58\x8a\xcf\xca\x94\xca\x72\xbc\x93\x7d\x76\x4d\x56\xf4\xc6\x9b\xa7\x9c\xa3\x7d\xb8\xa4\xa4\x98\xcb\xe0\xcb\x3e\x4d\x18\x70\xb1\x87\x86\xba\xed\x76\x14\x4a\x78\xb7\x3c\x54\xd2\x17\x63\x0a\xb5\x9a\x02\x82\xc1\xf2\x6e\x60\x0c\xfd\x9b\x29\xbc\xca\x16\x6e\x1e\xe5\x89\xf1\x14\xb3\x28\xda\xe7\x97\xe2\xa4\x61\xca\xb3\x80\x6a\xc5\x38\x1c\xe7\x0b\x38\xa2\x88\x54\xc5\x74\xda\x2f\x94\x62\x6d\xa5\x5f\x1c\x7a\x14\xf0\x40\xbb\x37\xae\x7c\x17\xd9\x92\x91\x6a\xf9\x32\x52\x47\x03\x0e\x31\x42\x23\xf5\x71\xad\x50\x94\x0f\xd5\x2a\xe1\xf9\x6f\x82\x76\x02\x85\x99\xdb\x28\x51\x5d\x16\xb9\x30\xda\x51\x5f\x33\x1b\xba\xee\x67\xde\x6d\x04\xe4\x77\x63\xf0\xcb\xbd\x6e\x7a\xc9\x55\xd0\xce\x90\xe6\x1a\x3a\xef\x24\xd6\x49\x54\x5a\xc2\xc1\x38\x5a\x16\x79\xd8\xab\xf6\x26\xda\x8e\x21\x46\x77\x3a\xe0\x3d\x14\x90\x0d\x53\x1f\x77\x55\xa0\x40\x3b\xca\x1e\xc3\xcb\xed\xa3\xa8\xa0\xa1\x9f\x8b\x5c\xe0\xe5\x79\xe6\x16\xd9\xc9\x75\xc7\x54\xf1\x94\xa9\x7a\x8c\x0b\x50\xdc\xc8\x7c\xde\x92\x2d\x7f\xd7\x36\x57\xdf\x44\x2f\x0d\xcc\xe7\x07\xa3\x2b\x07\xb1\x86\x41\x3f\xab\x5b\x02\x34\xa7\xd2\x31\x09\x66\x51\xe1\xfc\x53\x9f\x0f\x78\x42\xfa\x6f\x24\x45\xd5\x73\x90\x03\x8e\x76\x02\x7d\x50\x96\x6d\x49\x9b\xb6\x46\x02\x15\x81\x0b\x51\xe7\x1d\xb4\x79\xbc\xed\xf1\x65\xe1\xe3\x99\xe4\x83\xb0\xec\x28\x8c\x9f\x00\x18\x4d\x38\xc1\x02\x0f\x87\x31\x62\x89\xde\x5b\x75\x37\x72\x08\xe2\x82\xf4\xa9\x51\x17\x3b\x80\xec\xea\x9c\xdd\x19\x77\x17\xc5\xac\x76\x76\x23\xc7\xaa\xd8\x45\xbb\x60\x79\x1b\x6a\x6e\xc9\x0c\x35\xf4\x3e\x06\xf7\x77\x7e\x8c\x0e\x72\xe7\xdc\xc9\x0c\x18\x93\x42\x0f\xa2\xaa\x23\x1e\x24\xb5\x65\xfa\x30\x8d\xa6\x69\xff\x0c\x9a\x49\x9f\x15\x39\x52\xfe\x86\xbb\x1d\x9d\x6b\xc5\xf2\x93\x8f\xb2\xed\x58\xcc\x73\x36\x2c\x9c\x76\x99\xa2\x1f\x8a\xc1\x45\xf3\x3e\x94\xb4\xa6\xec\x36\xf0\x21\x60\x96\x83\xc9\x04\xcf\x98\x88\x62\x4f\xed\x81\x0c\xaa\x68\xb3\xcd\xad\xc6\x9f\xdc\x9d\xef\x97\xb5\x46\xe0\x7e\x8b\xcb\x68\xf0\xa1\xff\xfd\x54\x21\x07\x8f\x19\x7b\xd8\xc4\xba\x33\x66\xd4\xfb\x8c\x1a\x88\xeb\xf9\xcd\xaf\xcc\x95\x41\x5e\x42\xe6\x91\xd6\x08\x72\xbc\x32\x83\x13\xf2\xe7\x85\xeb\x82\xd0\x2e\xea\x51\x17\x05\x0d\x2d\xef\x60\x8b\x0b\xf9\x0a\xe8\x70\xcd\xde\x97\x49\xd7\x45\x9f\x4d\xbf\x2d\x90\x1a\x91\x66\x33\x35\x8a\x81\xfe\x7c\x21\x3b\x35\xb4\xa4\x42\x6d\x4d\x1f\x13\xd7\x49\x3c\xc9\x5c\x85\xcc\x56\xd4\x4f\x8d\x8d\xbc\x39\x58\x6a\x16\x8a\x61\x32\x02\xd3\x10\x40\xa7\x8d\x80\x3d\xb1\xc6\xbf\xd9\x39\x36\xfe\xa5\xa7\x7b\x1a\xe4\x05\x12\x17\x20\x8e\x8e\xc4\x58\x54\x7c\xac\x96\x16\x9c\x2e\xf4\x68\x7c\x49\xd0\x03\x43\x35\x2e\x0b\x58\x23\x26\x78\x51\x52\xa2\xae\xfe\x3d\x46\x30\x31\x2f\x10\x77\xdb\xd2\x1a\x97\x8f\x29\xa6\xd3\x88\xb9\x30\x81\xc3\x6a\x81\x82\x9c\x13\x8f\x73\x4b\x42\xdf\x98\x11\xb8\xa9\x27\xbf\xdc\x0b\x3d\x36\x48\x30\xf3\x0d\x56\x97\x5e\x20\x66\x9e\xb7\xed\x57\x37\x3e\x5b\x47\xe7\x68\x01\x77\xc7\xcb\xb0\x46\x7b\xf2\xaa\xac\x39\x58\x50\x8e\x5f\xc6\xc9\x78\x17\x93\x52\x5e\x84\xfb\x67\x3b\xe4\x15\xd6\x78\x01\x6e\xa4\x2d\x22\xde\xf1\xaa\x10\xd0\xbf\x3c\x0e\x58\x49\x81\xc3\x48\x14\x2d\x43\x5b\xc4\x90\x7a\xf5\xc3\xe3\xbe\x7b\x4c\x6b\x53\xcd\x28\x4e\x74\x99\xd0\x80\xb9\xe9\xb2\x29\xb1\xf5\x1a\xe5\x08\xad\x3f\x68\x51\x26\x1b\xbb\x73\x38\xbf\x4f\x23\xf2\x24\x76\x77\x8b\xf9\x82\xdc\xae\xea\xfb\x8f\xe5\xf5\x4c\x16\x0d\x7b\x00\x33\xb7\xd3\x24\xaa\x0d\x5c\x23\x9b\x38\x50\x1e\x1e\xf3\x2f\xe4\x42\xba\x8d\x19\xc5\xf2\x78\xe7\xa8\xf2\xef\x8c\xf4\xe7\x3d\x30\xf8\xfe\x61\xfe\xb8\x6d\x37\x26\xde\x9f\xed\xd1\xe8\x0c\x28\x0f\xe2\xdf\x47\xe2\x41\x7c\x6d\x1b\x68\x0d\x22\x02\x0b\x79\x52\xbe\xea\x4a\x11\xac\x4a\x43\x0c\x57\x22\x06\x91\x17\x46\xab\x61\x8a\x57\x95\x76\x58\x34\x28\x35\x3a\xbe\x1e\x9e\xf7\xa0\x92\x21\xeb\x61\xb0\x13\x0e\xd9\x04\xca\x0a\xed\x44\x59\xf9\x09\x1a\x7e\x2a\xe3\x69\x2a\xde\x2b\x27\xac\xe3\x73\xfa\xe2\x8c\x9c\x91\xb5\xe7\x0c\xe9\xb2\xd2\x82\x25\x9d\x40\xed\xa4\x83\x02\x4f\xdf\x25\xcc\x7d\x26\xc1\xd5\x35\x19\x0b\x47\x4f\x48\xf6\x08\xd6\x62\x3f\x46\x8c\x8d\x7b\x47\x5a\xca\xc4\x90\x0f\xed\x7b\xf4\xee\x2c\x7d\xf1\xc0\xc2\xd8\x94\x9c\xc8\xc2\xa1\x1e\xc1\x9c\x3f\x90\x1f\x4a\x98\x8d\x21\x2c\x5c\xd3\xc5\x18\x80\xa1\x7b\xcc\x9d\x7b\x63\x7c\x3b\x22\x13\xb5\xe5\x19\x01\x0b\x14\x8c\xb8\x9b\xcf\x6a\xf0\xe1\xa1\x8d\x44\xe1\x02\x37\x88\x0b\xd8\xb4\xb7\xa9\x86\x3a\xfc\x74\x50\xe4\x22\x44\x1d\x8a\x21\x44\xc3\x0b\x0b\xca\x8a\x94\x73\x7c\xf6\x1e\x56\xd2\xff\x87\x59\x3b\xe4\xeb\xfa\x6b\x36\x79\x84\x76\x26\xc6\x2e\xb7\x81\x9d\xa2\x12\xfe\x58\xc1\x51\x99\x88\xb4\x7a\xc7\x91\x01\x00\xe7\xef\x3e\xf6\xf3\x13\xe7\x64\x35\x5f\x87\x70\xb1\xa6\x7c\x32\x07\x07\xca\xf9\x62\xdc\xcb\x11\xbb\x8f\xbd\x2e\xf4\x74\x1e\xce\xe3\xe4\x34\x3f\x47\xc3\xb2\x41\xae\x2e\xe5\xce\x72\xde\x4a\x2f\x9e\x56\xe1\x5b\xe5\xc8\xe1\xc1\x0c\x03\x6e\x43\x43\xa6\xdc\x68\xa4\x1b\xa6\xab\x21\x57\xb7\x7e\x30\x77\x6d\x18\xe8\xcf\x78\xb3\x52\xdb\x05\x85\xc2\xc1\x14\x5c\xd0\x1a\x8d\x46\xc4\xc7\x54\x98\xdf\x15\xd8\x3c\xa1\xe5\x63\x1a\xd4\x50\xfb\xc1\xdd\x18\xa6\xe8\xa6\x3c\x81\x29\xf6\xe2\x60\xe4\x5d\x96\xb7\x11\xcc\xb6\xab\x6b\xc4\x85\x6c\xd7\x6d\xa4\x2e\x13\x53\x6c\x19\x42\xb7\x53\x41\x8e\x18\x2a\xd6\x3f\xe9\x6e\xea\x2f\x8c\x96\x88\x73\x4c\x76\xe0\x82\x31\x2a\x37\x55\xe5\x53\x60\x82\xaa\xbc\x6f\x80\xfa\x9f\x39\xa8\x3a\xd6\x83\xda\x11\xef\xc8\xef\xf4\xa1\x90\x14\x47\xd8\x0c\xcb\xbe\x56\xd7\x0e\x55\x7b\xf2\x60\x3f\x4e\xf7\xf3\xc5\xf9\xe5\xc7\x9f\x3d\xd0\x16\x0a\x23\x95\x67\x26\x51\x36\xb9\x82\x73\x47\xe6\xa4\xf3\xfb\x1e\xb5\x34\xf7\x8a\xf7\x0a\x81\x5e\xb0\x39\x16\x94\x3d\x3a\xd3\xab\xee\xf9\x79\xfe\x7d\x8e\xdc\x6c\xe8\x44\x5e\xb3\xa3\x6e\x8d\xf7\xb8\x62\xf7\xb0\x0b\xd6\xc9\x6b\xd8\xb7\x94\xdd\xdd\x46\x74\x9d\x7f\x69\xbb\x73\xa8\xde\xc9\xec\x88\xa4\x6b\x3b\xd1\xb6\x05\xf7\xb8\xcf\xdf\x6d\x93\xda\x9b\x0f\x5c\xbd\x33\x73\x23\x65\xac\x34\xe7\xc6\x51\xeb\xca\x68\xf9\x66\x55\x2a\x59\x44\x02\xf2\x1d\x1b\x09\xcf\xb9\x41\x7a\x04\x9e\x6a\xbb\x5f\x56\xf8\xc4\x57\xa9\x9e\xba\xcf\x6f\x3f\x85\xfa\xa4\xcc\x8c\xbc\x3b\xc5\x8d\xb1\x9e\x1e\x98\x9a\x7b\x73\xd3\x81\x1f\xfb\xa0\xe1\x07\xbc\x23\xb0\xd6\x5a\xcd\x52\x74\x1c\x09\x40\xb7\x80\x4b\x00\x57\xef\x5d\x42\x86\x40\x09\xeb\xb2\xab\x7b\x51\x03\x1b\xb4\xc7\xa4\x92\x27\x9b\x78\x89\x88\x77\xcd\xbc\xa2\x3d\x74\xbe\x6d\xdd\xa6\xe0\xc9\xeb\x8b\xe2\x17\x3f\xc8\x25\x72\x47\x86\x67\xfa\xbd\xd5\x89\x95\x9a\x7a\xde\xf7\x20\x2e\xc2\xe3\xab\x72\x4f\xbb\x7c\x3e\xc9\x6b\x23\x37\xa8\x11\x4c\xea\x08\x0d\xc9\xd7\x12\x9a\x20\x47\x4f\x18\xd0\xd4\x75\x0c\x66\x2b\x87\x91\x0e\x36\x48\x36\x7d\x3d\xcd\x46\xec\xd1\x9a\xd9\x02\xe5\xb5\x35\x79\xa5\xae\xd5\xf4\x1e\x47\xaf\x71\x4a\x20\xfb\x77\xb7\x41\x73\x51\x83\x67\x38\xc9\x99\x7f\xaa\x72\xcc\x54\x34\x96\x5a\x4b\x2e\x74\x9a\x2b\xf6\x0e\x4d\x41\xff\x68\xb9\x81\x65\x5e\x3c\xe9\xd4\xcf\x0d\x0f\x2c\x87\x9c\x14\x87\x54\x4b\x03\x81\x84\xa4\x93\xe9\xce\xe3\x12\x71\x72\x49\x67\x32\x47\x8f\xcc\x87\x47\x3b\x8e\x75\x48\x32\x9d\xb7\x2c\xb0\x0f\x6b\x2d\x4c\xe8\x74\xc7\x3b\xfb\x0c\xd4\x73\xda\xcb\x47\x2e\x50\xa3\x2c\xa5\xb6\xd2\xd6\x92\xa0\x81\x6d\x6f\x2b\x31\xab\x5e\xb5\x90\x89\x47\x80\x89\x40\x3b\xa6\x1f\xca\xa0\x80\x12\x04\x28\x03\x0d\x65\xc6\x6a\x01\xfb\x96\x53\x56\xf3\xfc\x93\x78\x72\x1f\x68\xc9\x4a\x4f\x75\x3a\xc3\x47\xd2\xc0\x53\xce\xfb\x47\x1e\x7e\x51\xd9\x98\x75\xa8\xd3\xaa\x89\xbc\xb2\x76\x4c\xac\x1e\x5a\x34\xbc\x01\x2c\x59\xc1\x99\xee\x58\x92\xf5\x17\x30\x1d\x76\xd2\xb4\xa7\xb6\x23\x61\x36\x73\xbf\x22\xb2\x26\xf5\x4f\x64\xda\x6b\x4d\x10\x8a\x13\xc0\x4e\x50\x44\x18\xad\xa7\xd4\x37\x0f\x73\x8f\x39\xde\xd8\xaf\x8c\xf8\xe5\x58\x39\xde\x5e\xaf\x9c\x11\x68\x6b\xfa\x68\x5d\x2d\x38\x74\xcf\x7e\x11\xdc\xfa\xc4\xd0\xce\x7d\x3c\x60\xda\xe0\xe9\x55\x81\x75\x8b\x9f\x1d\x39\xe4\x85\xa0\xd1\xc7\xcc\x02\x4f\xa8\x8c\xd4\xd4\x7b\x5c\x25\xb2\x70\x06\x4f\x7e\xf3\x38\xb0\x96\x05\xd6\xf3\xcb\xa9\x7b\xdb\x02\xcb\x6a\xe0\xb9\xc3\x4a\xb6\xe5\x30\x48\x26\x9c\x6e\x35\x93\x2c\xe9\x40\x06\x4e\x0b\x6a\x81\x91\x0d\x63\x0e\xa6\x51\x4b\x88\x5e\xd5\xb8\x3e\xae\x36\x79\x43\xb9\x00\x6f\x10\x11\xde\x59\x0c\x86\x4a\xca\x2a\x7d\x49\x44\xdf\x54\x0e\xca\x1e\x1c\x7e\xa0\x5d\xe5\xfb\x53\x2e\x8a\x52\x16\xf6\x65\xb7\xcc\xdd\xfb\x3b\xe3\xc9\x1e\x63\x83\x43\x5c\xaf\x3c\xfe\x29\x48\x2f\x9f\x16\x1a\x3f\x18\x7c\xa4\xae\x9a\xe7\xb0\x86\xec\x71\xb5\x5e\xed\x69\xc7\xea\xfe\x7f\x4a\x4a\x04\x83\xa5\xa0\xac\xff\x03\xf2\x0e\xd6\xd3\xab\xbd\x53\x0d\x1d\x63\x88\x94\xbd\x53\x22\x6b\xff\x4e\xdd\xb9\x16\x7a\xfa\x3f\x24\x55\x7f\x8e\x17\xff\x25\x37\x5d\xd1\x1d\x26\xb1\xd7\xe1\x64\x28\xba\x77\x55\xa4\x85\xaa\x7b\x68\x90\x8b\xe4\x97\x2f\xe5\x3a\x46\xaf\x21\xb1\x50\x1b\x46\x1f\xa6\x4d\xf1\x51\xc5\xa6\x1e\x12\x1f\xe9\x8f\xe2\x7a\xcb\x66\x7f\xe1\xd0\x56\xde\xd0\x3d\x7b\x68\xcb\xec\x7a\xc6\xe9\x24\x0b\x7e\xdc\xd0\x56\x2f\x17\x7f\xdd\x51\xca\x11\x38\xa7\x0f\xa4\xff\xf3\xc3\xaf\x57\x80\xa3\x52\x3a\xe2\xab\xfe\xc7\x55\x2f\x2d\xe8\x33\x2a\x3b\x81\x2a\xf0\xb0\x47\xa4\x77\xdc\x39\x68\xf0\xe0\xaf\x63\x0e\x94\xeb\x01\x36\xb0\xbc\x7b\xf1\xe2\xfc\xfd\xbb\x5f\xf4\xb4\x5f\xfe\x08\x2e\x7e\xbb\xfc\x70\xf3\x01\xdc\xc2\xb2\x2e\x76\x0c\x12\x71\xfb\x8f\x04\x82\x21\x4e\x3b\x56\xa2\x34\xa8\x37\xe8\x29\x80\x3a\xa7\x95\x44\xb5\x6d\xf4\x5b\x49\x9b\x06\xc7\x9b\xa9\x3e\x17\xd2\x2f\xc3\xe2\x71\x0e\xb7\xc5\x89\x66\x68\x8c\xba\xc9\x7a\x06\x34\xbc\x03\x18\x87\x4d\xb6\x2d\x8a\xd1\x72\x12\x05\xf4\x7e\x8b\x4a\x8b\x8c\x42\xbc\x97\xfd\xa3\xc8\x74\xb7\xec\xab\x0b\xd3\xa8\x61\xb9\x9e\x46\xcd\x8d\x90\xfd\xf2\xcc\x1c\xea\x13\xdd\xc4\x31\xde\xd3\x09\xd9\x48\x7d\xa7\x7d\x12\x9f\x1e\xb8\xf0\xad\xaa\x39\xe8\x99\xb1\xf6\x6e\xc2\xcb\x00\x8e\x79\xbd\x59\x58\x7d\x37\x57\x06\x56\xea\xd6\x24\x6e\x8e\x27\x62\x17\xaa\x64\xe1\xf5\x9d\x1a\x79\x58\x95\xf2\x90\xc4\xca\xa5\x6b\x0c\x11\x3a\xe5\x39\x87\x35\xb2\x0b\x63\xd0\xf8\x51\xb2\x5c\x8a\xf1\x00\x4f\x2e\x81\x3e\x03\x91\x0b\x9f\x32\xe6\xe7\x29\x32\xc6\x26\x94\x7a\x39\x8b\xd5\xe1\xd2\x59\x9c\xce\x18\x8a\xe1\xbc\xcc\x84\x18\x50\xee\x71\xa7\x3e\xce\xaa\x55\xbd\x5f\x14\xfd\xec\x06\xdb\xe7\x81\x73\xa2\xa4\xe3\x6c\xb1\xcf\x49\xe6\xf6\xd6\x5d\x49\xa0\x72\xad\x93\x90\x69\x78\xfe\x37\x00\x00\xff\xff\x3a\xcb\x83\xd0\x82\xb1\x00\x00")

func migrations20170523183416_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20170523183416_initSql,
		"migrations/20170523183416_init.sql",
	)
}

func migrations20170523183416_initSql() (*asset, error) {
	bytes, err := migrations20170523183416_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20170523183416_init.sql", size: 45442, mode: os.FileMode(420), modTime: time.Unix(1537981572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20171108110099_rbacSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\xdf\x6b\xfa\x30\x14\xc5\xdf\xf3\x57\x5c\xf2\xa2\xf2\x4d\xe1\x5b\xdb\x6a\x8b\xec\x41\x30\x63\x82\x3f\x98\x3f\xb6\x47\x13\x93\xab\x86\xb5\x69\x69\x3a\xdd\x9f\x3f\x32\x71\x1b\xc8\xc0\x87\xed\x29\x87\x70\x0f\xf7\xdc\xf3\x09\x02\xf8\xb7\x2f\x4b\x87\xb0\xae\x48\x10\xc0\xf2\x71\x02\xc6\x82\x43\xd5\x98\xd2\x42\x6b\x5d\xb5\xc0\x38\xc0\x37\x54\xaf\x0d\x6a\x38\x1d\xd0\x42\x73\x30\x0e\x0a\xb3\xaf\xe5\xc7\x90\x71\x20\xab\x2a\x37\xa8\x09\x19\xcf\x96\x7c\xb1\x82\xf1\x6c\x35\x07\x21\x55\xbe\xa9\xcb\x1c\x05\xb4\x85\xd1\x82\x09\x2b\x0b\x14\x4c\x68\x74\xaa\x36\x95\x37\x0b\x26\x54\x8d\xb2\x41\xbd\x91\x8d\xe8\xc0\xd3\x70\xb2\xe6\x4b\x68\xd3\x70\x17\x26\x32\x56\x99\x8e\xfb\x59\xda\x8d\xfe\x53\x46\xa5\x2e\x8c\xa5\x8c\x0e\xfd\x6b\x5c\xe3\xd7\x1f\x11\xfc\x06\xca\xc2\x24\xea\x67\x69\x98\xf4\xbb\x69\xda\xeb\x0c\x7e\x2f\x48\x24\x31\x8c\x33\x99\x6a\x15\xc7\xf1\x56\x21\x65\xf4\xdc\x86\x39\x7a\xcd\x2f\xfa\xaf\x73\xc4\x49\x92\x84\xe1\x2e\xc5\xa8\xd7\xd5\xb8\xcd\x28\xa3\x85\xb4\x72\x8f\x35\x65\x74\x7a\x56\x3f\x64\x20\x5f\x94\x47\xe5\xc9\x5e\x38\x7f\x42\xf6\x9f\x37\x61\xae\xcb\x3c\x47\x0d\x5b\xa9\x5e\x08\x21\x23\x3e\xe1\x2b\x0e\xf7\x8b\xf9\xf4\xfb\x69\xcf\x0f\x7c\xc1\xc1\xdf\x07\x77\x70\xcd\x71\x70\x9b\xef\xaa\xf6\x1b\x7d\x57\x35\x0d\xc8\x7b\x00\x00\x00\xff\xff\x28\xa8\xe1\xcb\xe4\x02\x00\x00")

func migrations20171108110099_rbacSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20171108110099_rbacSql,
		"migrations/20171108110099_rbac.sql",
	)
}

func migrations20171108110099_rbacSql() (*asset, error) {
	bytes, err := migrations20171108110099_rbacSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20171108110099_rbac.sql", size: 740, mode: os.FileMode(420), modTime: time.Unix(1537981572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20171122172200_indexesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcd\x72\xe2\x3a\x10\x85\xd7\xf1\x53\x9c\x5d\x4c\x5d\x53\x15\xc8\x7f\x65\xe5\x5c\x7c\xef\x50\x43\x08\x63\xcc\x54\x58\x19\x61\x37\xa0\xc4\x48\x2e\x49\x0e\xe1\xed\xa7\x64\x93\x1f\x1c\x2d\x66\x47\x75\x7f\x3a\x7d\x50\x5b\xa7\xdb\xc5\x3f\x6b\x29\x35\x61\x56\x7a\xdd\x2e\xa6\xbf\x46\xe0\x02\x9a\x32\xc3\xa5\xc0\xe9\xac\x3c\x05\xd7\xa0\x37\xca\x2a\x43\x39\x76\x1b\x12\x30\x1b\xae\xb1\xe5\x6b\xc5\x6a\x88\x6b\xb0\xb2\x2c\x38\xe5\x9e\x95\x30\x1b\xc2\x84\x29\xc3\xeb\xe6\x92\x0a\xb9\x83\x20\xca\x35\x8c\x84\xa2\xad\x7c\xa5\x9a\x29\x15\xdf\x32\xb5\xc7\x0b\xed\xa1\xb9\xc8\x08\xdc\x9c\x6a\x08\x69\x50\x69\x5a\x55\x05\x98\xd8\xef\xd8\xde\x0b\x47\x49\x14\x23\x09\xef\x47\x11\x16\x9a\xaf\x05\x2b\x16\x18\xc4\x8f\x13\x4c\xe2\xe1\x43\x18\xcf\xf1\x33\x9a\x07\x08\x07\x83\xaf\x05\x9f\xe7\x01\x72\x66\x28\x80\x60\x5b\xea\xdc\x39\xcc\xed\x78\x51\x60\x49\x58\xee\x31\x8f\xc2\x18\x4c\xe4\x96\x10\xd0\xd5\xb2\x7c\xa7\x28\xb7\xfd\x66\xae\xdb\xcb\x24\x8c\x93\x61\x32\x7c\x1c\xe3\x7e\x8e\x38\x1c\xff\x1f\xf9\x56\xce\xb7\xd3\x3b\x1d\x78\xc0\x74\x76\x7f\x04\x59\x7f\xb5\xab\x56\x6f\x8a\x5e\x1f\xbe\x77\x72\xf2\x49\x97\x67\xf8\x1d\x8e\x66\xd1\x14\xa3\x68\x3a\x45\xf2\x23\x1c\xc3\xef\x9f\xf5\xce\x3a\xc1\x31\xd7\x73\x73\xbd\x36\xd7\x77\x73\xfd\x36\x77\xee\xe6\xce\xdb\xdc\x85\x9b\xbb\x68\x73\x97\x6e\xee\xb2\xcd\x5d\xb9\xb9\xab\x36\x77\xed\xe6\xae\xdb\xdc\x8d\x9b\xbb\x69\x73\xb7\x6e\xee\xf6\xdb\x3d\xbb\x17\xd2\xff\xbe\x10\xf7\x46\xfa\xdf\x36\x72\xeb\x18\xfd\x10\x3e\xd5\x35\xcf\x7e\xb8\xff\xc6\x51\x98\x44\x18\x8e\x07\xd1\x13\xb8\xd6\x15\xa5\xda\x30\x43\x69\x56\x69\x23\xb7\xa4\x52\x9e\xa7\xa5\x92\xcf\x94\x19\xfb\x33\x2b\xa4\xa6\x3c\x65\x26\x55\xb4\xb2\x05\xb3\x2f\x29\xe5\x22\xa7\x37\x3c\x8e\x1b\x05\xbf\x56\x08\xf0\x45\x22\xc0\xa7\x46\x80\x0f\x91\x00\x8d\x4a\x00\x2b\xd3\xb9\x3b\xb6\xf3\xcc\x15\x4b\x1b\x4f\xd9\x86\x89\x35\xa5\x85\x5c\x1f\x0a\x3c\x4f\x57\x9c\x8a\xfc\x73\xb4\x93\xf6\xdf\xe9\x00\x35\x7e\x78\xab\x99\x22\x66\x08\x4c\xa0\x39\xbe\x92\xaa\x7e\xbf\x6b\x12\xa4\x98\x4d\xa3\x57\xae\x4c\xc5\x8a\xe6\xd4\xd1\xdb\x3c\x5c\x52\xb5\xad\x23\xc6\x46\xc3\x7f\xb3\xd1\x28\x89\x9e\x92\x83\xed\x45\xf3\xbf\x1b\x7f\x3a\x5d\x55\x45\x91\xf2\xfc\x6d\xe1\x1f\x35\x52\x9e\xeb\xf4\x30\x65\x71\xb0\x75\xc8\xcb\x81\xdc\x89\xf7\xc4\xfc\x88\x4b\x5b\xfc\xab\xc0\x54\xb2\x28\x6c\xb0\xb0\xec\xc5\xf3\xfe\x04\x00\x00\xff\xff\xaf\xa4\x57\xce\x85\x05\x00\x00")

func migrations20171122172200_indexesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20171122172200_indexesSql,
		"migrations/20171122172200_indexes.sql",
	)
}

func migrations20171122172200_indexesSql() (*asset, error) {
	bytes, err := migrations20171122172200_indexesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20171122172200_indexes.sql", size: 1413, mode: os.FileMode(420), modTime: time.Unix(1537981553, 0)}
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
	"migrations/20170523183416_init.sql": migrations20170523183416_initSql,
	"migrations/20171108110099_rbac.sql": migrations20171108110099_rbacSql,
	"migrations/20171122172200_indexes.sql": migrations20171122172200_indexesSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"20170523183416_init.sql": &bintree{migrations20170523183416_initSql, map[string]*bintree{}},
		"20171108110099_rbac.sql": &bintree{migrations20171108110099_rbacSql, map[string]*bintree{}},
		"20171122172200_indexes.sql": &bintree{migrations20171122172200_indexesSql, map[string]*bintree{}},
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
