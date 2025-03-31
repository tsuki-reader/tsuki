// Code generated by 'yaegi extract github.com/tsuki-reader/nisshoku/providers'. DO NOT EDIT.

package extract

import (
	"github.com/tsuki-reader/nisshoku/providers"
	"reflect"
)

func init() {
	Symbols["github.com/tsuki-reader/nisshoku/providers/providers"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Comic": reflect.ValueOf(providers.Comic),
		"Manga": reflect.ValueOf(providers.Manga),

		// type definitions
		"Chapter":         reflect.ValueOf((*providers.Chapter)(nil)),
		"LibraryEntry":    reflect.ValueOf((*providers.LibraryEntry)(nil)),
		"Page":            reflect.ValueOf((*providers.Page)(nil)),
		"Provider":        reflect.ValueOf((*providers.Provider)(nil)),
		"ProviderContext": reflect.ValueOf((*providers.ProviderContext)(nil)),
		"ProviderResult":  reflect.ValueOf((*providers.ProviderResult)(nil)),
		"ProviderType":    reflect.ValueOf((*providers.ProviderType)(nil)),

		// interface wrapper definitions
		"_Provider": reflect.ValueOf((*_github_com_tsuki_reader_nisshoku_providers_Provider)(nil)),
	}
}

// _github_com_tsuki_reader_nisshoku_providers_Provider is an interface wrapper for Provider type
type _github_com_tsuki_reader_nisshoku_providers_Provider struct {
	IValue           interface{}
	WGetChapterPages func(id string) ([]providers.Page, error)
	WGetChapters     func(id string) ([]providers.Chapter, error)
	WImageHeaders    func() map[string]string
	WProviderType    func() providers.ProviderType
	WSearch          func(query string) ([]providers.ProviderResult, error)
}

func (W _github_com_tsuki_reader_nisshoku_providers_Provider) GetChapterPages(id string) ([]providers.Page, error) {
	return W.WGetChapterPages(id)
}
func (W _github_com_tsuki_reader_nisshoku_providers_Provider) GetChapters(id string) ([]providers.Chapter, error) {
	return W.WGetChapters(id)
}
func (W _github_com_tsuki_reader_nisshoku_providers_Provider) ImageHeaders() map[string]string {
	return W.WImageHeaders()
}
func (W _github_com_tsuki_reader_nisshoku_providers_Provider) ProviderType() providers.ProviderType {
	return W.WProviderType()
}
func (W _github_com_tsuki_reader_nisshoku_providers_Provider) Search(query string) ([]providers.ProviderResult, error) {
	return W.WSearch(query)
}
