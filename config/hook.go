package config

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"text/template"
)

const (
	defaultTemplateName = "config"
)

func StringRenderTextTemplateHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.String {
			return data, nil
		}

		return runRaw(data.(string), map[string]string{})
	}
}

func runt(tpl, expect string) error {
	return runtv(tpl, expect, map[string]string{})
}

func runtv(tpl, expect string, vars interface{}) error {
	t := template.Must(template.New(defaultTemplateName).Funcs(sprig.TxtFuncMap()).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return err
	}
	if expect != b.String() {
		return fmt.Errorf("expected '%s', got '%s'", expect, b.String())
	}
	return nil
}

func runRaw(tpl string, vars interface{}) (string, error) {
	t := template.Must(template.New(defaultTemplateName).Funcs(sprig.TxtFuncMap()).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
