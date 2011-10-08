package main

import (
	"fmt"
	"reflect"
	"./curl/_obj/curl"
)


const (
	a = iota
	b string = "100"
)


func defer_recover() {
	println("ok")
	println("recover")
	recover()

}


func printArgs(args ... interface{}) {
	for _, v := range args {
		fmt.Printf("%v -> %T\n", v, v)
	}
}


const endl = "\n"

func main() {
	var c complex64 = 5+5i;
	defer defer_recover()
	ret := curl.EasyInit()
	defer ret.Cleanup()
	print("init =>", ret, " ", reflect.TypeOf(ret).String(), endl)
	ret = ret.Duphandle()
	defer ret.Cleanup()
	print("dup =>", ret, " ", reflect.TypeOf(ret).String(), endl)

	print("global init =>", curl.GlobalInit(curl.GLOBAL_ALL), endl)
	print("version =>", curl.Version(), endl)
	// debug
	print("set verbose =>", ret.Setopt(curl.OPT_VERBOSE, true), endl)
	fmt.Printf("set verbose => %s. \n", ret.Setopt(curl.OPT_VERBOSE, true))


	print("set header =>", ret.Setopt(curl.OPT_HEADER, true), endl)
	// auto calculate port
	// print("set port =>", ret.EasySetopt(curl.OPT_PORT, 6060), endl)
	// curl.GlobalCleanup()
	print("set timeout =>", ret.Setopt(curl.OPT_TIMEOUT, 20), endl)
	print("set post data =>", ret.Setopt(curl.OPT_POSTFIELDS, "name=100"), endl)

	print("set url =>", ret.Setopt(curl.OPT_URL, "http://www.google.com"), endl)
	print("set user_agent =>", ret.Setopt(curl.OPT_USERAGENT, "go-curl v0.0.1"), endl)
	// add to DNS cache
	print("set resolve =>", ret.Setopt(curl.OPT_RESOLVE, []string{"www.baidu.com:6543:127.0.0.1",}), endl)

	// ret.EasyReset()  clean seted
	code := ret.Perform()
	print("perfom =>", code, endl)


	println("================================")
	print("pause =>", ret.Pause(curl.PAUSE_ALL), endl)

	print("escape =>", ret.Escape("http://baidu.com/"), endl)
	print("unescape =>", ret.Unescape("http://baidu.com/-%00-%5c"), endl)
	print("unescape lenght =>", len(ret.Unescape("http://baidu.com/-%00-%5c")), endl)

	// print("version info data =>", curl.VersionInfo(1), endl)
	ver := curl.VersionInfo(1)
	fmt.Printf("VersionInfo: Age: %d, Version:%s, Host:%s, Features:%d, SslVer: %s, LibzV: %s, ssh: %s\n",
		ver.Age, ver.Version, ver.Host, ver.Features, ver.SslVersion, ver.LibzVersion, ver.LibsshVersion)

	print("Protocols:")
	for _, p := range ver.Protocols {
		print(p, ", ")
	}
	print(endl)
	println(curl.Getdate("20111002 15:05:58 +0800").String())
	ret.Getinfo(curl.INFO_EFFECTIVE_URL)
	ret.Getinfo(curl.INFO_RESPONSE_CODE)
	ret.Getinfo(curl.INFO_TOTAL_TIME)
	// ret.Getinfo(curl.INFO_SSL_ENGINES)

/*	mret := curl.MultiInit()
	mret.AddHandle(ret)			// works
	defer mret.Cleanup()
	if ok, handles := mret.Perform(); ok == curl.OK {
		fmt.Printf("ok=%s, handles=%d\n", ok, handles)
	} else {
		fmt.Printf("error calling multi\n")
	}
*/
	println("================================")
	//println(curl.GlobalInit(curl.GLOBAL_SSL))
	println(reflect.TypeOf(printArgs).String())

	fmt.Printf("Hello World; \n 你好世界\n")
	fmt.Printf("%s hah%da\n", b, a)
	fmt.Printf("Value is:%v\n", c)
	printArgs("hello", []int{1,2,3}, 123)
	panic("fuck")
	println("recover~")


}