package main

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/AnimusPEXUS/appplugsys"
)

var _ appplugsys.PluginSearcherI = &PluginSearcher{}

type PluginSearcher struct {
	builtins []*appplugsys.Plugin
	paths    []string
}

func NewPluginSearcher(
	builtins []*appplugsys.Plugin,
	paths []string,
) (*PluginSearcher, error) {

	self := new(PluginSearcher)

	self.builtins = builtins
	self.paths = paths

	return self, nil
}

func (self *PluginSearcher) FindAll() ([]*appplugsys.PluginSearcherSearchResultItem, error) {

	ret := make([]*appplugsys.PluginSearcherSearchResultItem, 0)

	for _, i := range self.builtins {
		ni := &appplugsys.PluginSearcherSearchResultItem{
			BasicPluginInfo: appplugsys.BasicPluginInfo{
				Name:    i.Name,
				BuiltIn: true,
				Sha512:  "",
			},
			Path:   "",
			Plugin: i,
		}

		ret = append(ret, ni)
	}

	for _, pth := range self.paths {
		files, err := ioutil.ReadDir(pth)
		if err != nil {
			continue
		}
		for _, i := range files {
			var str_sum string
			file_path := path.Join(pth, i.Name())
			{
				h := sha512.New()

				rf, err := os.Open(file_path)
				if err != nil {
					continue
				}
				defer func(f *os.File) { f.Close() }(rf)

				_, err = io.Copy(h, rf)
				if err != nil {
					continue
				}

				rf.Close()

				sum := h.Sum([]byte{})
				str_sum = hex.EncodeToString(sum)
				str_sum = strings.ToLower(str_sum)
			}

			ni := &appplugsys.PluginSearcherSearchResultItem{
				BasicPluginInfo: appplugsys.BasicPluginInfo{
					Name:    "",
					BuiltIn: false,
					Sha512:  str_sum,
				},
				Path:   file_path,
				Plugin: nil,
			}

			ret = append(ret, ni)
		}
	}

	return ret, nil
}

func (self *PluginSearcher) FindBuiltIn(name string) (*appplugsys.PluginSearcherSearchResultItem, error) {

	all, err := self.FindAll()
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		if i.BuiltIn && i.Name == name {
			return i, nil
		}
	}

	return nil, errors.New("not found")
}

func (self *PluginSearcher) FindBySha512(sha512 string) (*appplugsys.PluginSearcherSearchResultItem, error) {
	if len(sha512) != 128 {
		return nil, errors.New("invalid input sha512")
	}

	sha512 = strings.ToLower(sha512)

	all, err := self.FindAll()
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		if !i.BuiltIn && i.Sha512 == sha512 {
			return i, nil
		}
	}

	return nil, errors.New("not found")
}
