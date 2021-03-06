/*
Copyright 2017 wechat-go Authors. All Rights Reserved.
MIT License

Copyright (c) 2017

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package wxbot

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
	"time"
)

func Run() {
	// syncheck
	msg := make(chan []byte, 10000)
	for {
		select {
		case <-time.After(1 * time.Second):
			go producer(msg)
		case m := <-msg:
			go consumer(m)
		}
	}

}

func producer(msg chan []byte) {
	for _, v := range WxWebDefaultCommon.SyncSrvs {
		ret, sel, err := wxweb.SyncCheck(WxWebDefaultCommon, WxWebXcg, Cookies, v, SynKeyList)
		logs.Debug(v, ret, sel)
		if err != nil {
			logs.Error(err)
		}
		if ret == 0 {
			// check success
			if sel == 2 {
				// new message
				err := wxweb.WebWxSync(WxWebDefaultCommon, WxWebXcg, Cookies, msg, SynKeyList)
				if err != nil {
					logs.Error(err)
				}
			}
			break
		}
	}

}

func consumer(msg []byte) {
	logs.Debug("received", string(msg))
}
