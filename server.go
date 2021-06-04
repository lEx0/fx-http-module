// Copyright (c) 2021 Amangeldy Kadyl
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package http

import (
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"time"
)

// Params указываем полями в структуре нужные нам зависимости
type Params struct {
	fx.In

	Options Options
	Router  *mux.Router
}

type Options struct {
	Listen       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewServer(params Params) (*http.Server, error) {
	if params.Options.Listen == "" {
		return nil, errors.New("empty listen addr")
	}

	return &http.Server{
		Addr:         params.Options.Listen,
		Handler:      params.Router,
		ReadTimeout:  params.Options.ReadTimeout,
		WriteTimeout: params.Options.WriteTimeout,
		IdleTimeout:  params.Options.IdleTimeout,
	}, nil
}
