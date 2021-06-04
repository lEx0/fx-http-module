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
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module модуль http сервера
//goland:noinspection GoUnusedGlobalVariable
var Module = fx.Options(
	// прокидываем зависимости в Fx
	fx.Provide(
		mux.NewRouter,
		NewServer,
	),
	// вешаем хуки на старт и стоп приложения
	fx.Invoke(func(lifecycle fx.Lifecycle, server *http.Server, logger *zap.Logger) {
		lifecycle.Append(fx.Hook{
			// при старте приложения, запускаем http сервер
			OnStart: func(_ context.Context) error {
				listener, err := net.Listen("tcp", server.Addr)
				if err != nil {
					return err
				}

				go func() {
					if err = server.Serve(listener); err != nil {
						logger.Sugar().Fatalw(
							"serve http server failed",
							"error", err,
						)
					}
				}()

				return nil
			},
			// при остановке приложения, останавливаем http сервер
			// с ожиданием закрытия всех соединений
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		})
	}),
)
